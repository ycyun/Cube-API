package DiskController

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/goccy/go-json"
	"github.com/ycyun/Cube-API/Modules/cube/DiskController/BLK"
	"github.com/ycyun/Cube-API/Modules/cube/DiskController/Disk"
	"github.com/ycyun/Cube-API/Modules/cube/DiskController/RBD"
	"github.com/ycyun/Cube-API/utils"
)

// GetDisks는 컨트롤러에 저장된 모든 디스크(RBD, BLK 등)의 목록을 반환합니다.
// 현재 구현에서는 이 함수가 직접 사용되기보다는 GetRBDs, GetBLKs가 각각 사용됩니다.
func (s *DiskController) GetDisks() []Disk.Interface {
	return s.Disks
}

// GetRBDs는 컨트롤러에 저장된 RBD(RADOS Block Device) 디스크 목록을 반환합니다.
func (s *DiskController) GetRBDs() []Disk.Interface {
	return s.RBDs
}

// GetBLKs는 컨트롤러에 저장된 로컬 블록 디바이스(하드디스크, SSD 등) 목록을 반환합니다.
func (s *DiskController) GetBLKs() []Disk.Interface {
	return s.BLKs
}

// UpdateRBDs는 Ceph 클러스터의 모든 RBD 이미지 정보를 조회하여 s.RBDs 필드를 최신 상태로 업데이트합니다.
func (s *DiskController) UpdateRBDs() (err error) {
	startTime := time.Now()
	defer func() {
		// 3. time.Since()를 사용하면 더 간결합니다.
		slog.Debug(
			"UpdateRBDs finished",
			"duration", time.Since(startTime),
			"RBDs_found", len(s.RBDs),
		)
		// 4. (고급) 에러가 발생했을 때 추가적인 로그를 남길 수도 있습니다.
		if err != nil {
			slog.Error("UpdateRBDs failed with an error", "error", err)
		}
	}()

	Logger := slog.New(utils.AbleHandler)
	slog.Debug("Start Get Pool List", "start", startTime)

	// Ceph OSD 풀 목록을 가져옵니다.
	// `ceph osd pool ls --format json` 명령어를 실행합니다.
	PoolListCmd := []string{
		"ceph", "osd", "pool", "ls", "--format", "json",
	}
	PoolListReturn, err := utils.Execute(PoolListCmd)
	utils.HandleError(err) // 에러 발생 시 처리

	// JSON으로 반환된 풀 목록을 문자열 슬라이스로 언마샬링합니다.
	PoolList := make([]string, 0)
	err = json.Unmarshal(PoolListReturn, &PoolList)
	utils.HandleError(err)
	slog.Debug("Pool", "num", len(PoolList), "PoolList", PoolList)

	// 모든 풀에서 찾은 RBD 이미지를 저장할 슬라이스
	Images := make([]Disk.Interface, 0)
	// 각 풀을 순회하며 이미지 목록을 가져옵니다.
	for i, pool := range PoolList {
		slog.Debug("Pool", "index", i, "name", pool)

		// 특정 풀에 속한 이미지 목록을 가져옵니다.
		// `rbd ls -l -p <pool_name> --format json` 명령어를 실행합니다.
		ImageListCommand := []string{
			"rbd", "ls", "-l", "-p", pool, "--format", "json",
		}
		ImageListReturn, err := utils.Execute(ImageListCommand)
		utils.HandleError(err)

		// JSON으로 반환된 이미지 목록을 RBD 구조체 슬라이스로 언마샬링합니다.
		ImageList := make([]*RBD.RBD, 0)
		err = json.Unmarshal(ImageListReturn, &ImageList)
		utils.HandleError(err)

		// 각 이미지의 상세 정보를 처리합니다.
		for j, image := range ImageList {
			// 이미지 타입을 결정합니다.
			if image.Parent != nil {
				// 부모 정보가 있으면 자식 이미지(clone)입니다.
				image.Type = "ChildImage"
				image.Name = image.Image

			} else if image.Snapshot != nil {
				// 스냅샷 정보가 있으면 스냅샷 이미지입니다.
				image.Type = "SnapshotImage"
				image.Name = image.Image + "@" + *image.Snapshot
				Logger.Debug("snapshotImage", "image", image.Image, "snapshot", image.Snapshot)

			} else {
				// 둘 다 없으면 일반 이미지입니다.
				image.Type = "Image"
				image.Name = image.Image
			}
			// 이미지의 소속 풀 정보를 설정합니다.
			image.Pool = pool

			// `rbd info`와 같은 명령어로 이미지의 상세 정보를 업데이트합니다.
			image.Update()
			// 처리된 이미지를 전체 목록에 추가합니다.
			Images = append(Images, image)
			Logger.Debug("Image in "+pool, "index", fmt.Sprintf("%d/%d", j, len(ImageList)), "Image", image.Path)
		}
	}

	// 최종적으로 수집된 이미지 목록을 컨트롤러의 RBDs 필드에 저장합니다.
	s.RBDs = Images
	return nil

}

// UpdateBLKs 는 시스템의 블록 디바이스 정보를 조회하여 s.BLKs 필드를 최신 상태로 업데이트합니다.
func (s *DiskController) UpdateBLKs() (err error) {
	start_time := time.Now()

	defer func() {
		// 3. time.Since()를 사용하면 더 간결합니다.
		slog.Debug(
			"UpdateBLKs finished",
			"duration", time.Since(start_time),
			"BLKs_found", len(s.BLKs),
		)
		// 4. (고급) 에러가 발생했을 때 추가적인 로그를 남길 수도 있습니다.
		if err != nil {
			slog.Error("UpdateBLKs failed with an error", "error", err)
		}
	}()

	slog.Debug("Start Get Block List")
	// 시스템의 블록 디바이스 목록을 가져옵니다.
	// `lsblk -J --exclude 252` 명령어를 실행합니다.
	// -J: JSON 형식으로 출력
	// --exclude 252: major number 252번(rbd-nbd에 의해 매핑된 디바이스)을 제외하여 중복 조회를 방지합니다.
	DiskListCmd := []string{
		"lsblk", "-J", "--exclude", "252",
	}
	DiskListReturn, err := utils.Execute(DiskListCmd)
	utils.HandleError(err)

	var lsblkData BLK.LsblkOutput
	err = json.Unmarshal(DiskListReturn, &lsblkData)
	utils.HandleError(err)

	// lsblk의 JSON 출력은 {"blockdevices": [...]} 형태이므로, 먼저 map으로 파싱합니다.
	_DiskList := make(map[string]interface{})
	err = json.Unmarshal(DiskListReturn, &_DiskList)
	utils.HandleError(err)
	// "blockdevices" 키에 해당하는 디바이스 목록을 추출합니다.
	mDiskList := _DiskList["blockdevices"]

	// 추출한 디바이스 목록 부분을 다시 JSON으로 마샬링합니다.
	jDiskList, err := json.Marshal(mDiskList)
	utils.HandleError(err)

	// 최종적으로 BLK 구조체 슬라이스로 언마샬링합니다.
	DiskList := make([]*BLK.BLK, 0)
	err = json.Unmarshal(jDiskList, &DiskList)
	utils.HandleError(err)

	// 인터페이스 슬라이스로 변환하여 저장합니다.
	slog.Debug("Disks", "num", len(lsblkData.BlockDevices), "DiskList", lsblkData.BlockDevices)

	blockDevices := make([]Disk.Interface, 0, len(lsblkData.BlockDevices))
	for _, blkDevice := range lsblkData.BlockDevices {
		blockDevices = append(blockDevices, blkDevice)
	}
	s.BLKs = blockDevices
	return nil
}
