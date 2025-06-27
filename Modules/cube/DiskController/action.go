package DiskController

import (
	"github.com/goccy/go-json"
	"github.com/ycyun/Cube-API/Modules/cube/DiskController/BLK"
	"github.com/ycyun/Cube-API/Modules/cube/DiskController/Disk"
	"github.com/ycyun/Cube-API/Modules/cube/DiskController/RBD"
	"github.com/ycyun/Cube-API/utils"
	"log/slog"
)

func (s *DiskController) GetDisks() []Disk.Interface {
	return s.Disks
}

func (s *DiskController) GetRBDs() []Disk.Interface {
	return s.RBDs
}

func (s *DiskController) GetBLKs() []*BLK.BLK {
	return s.BLKs
}

func (s *DiskController) UpdateRBDs() {
	tmplist := make([]Disk.Interface, 0)

	//Todo: get Pool list
	// ceph osd pool ls --format json
	PoolListCmd := []string{
		"ceph", "osd", "pool", "ls", "--format", "json",
	}
	PoolListReturn, err := utils.Execute(PoolListCmd)
	utils.HandleError(err)
	PoolList := make([]string, 0)
	err = json.Unmarshal(PoolListReturn, &PoolList)

	Images := make([]Disk.Interface, 0)
	for i, pool := range PoolList {
		slog.Info("Pool", "index", i, "name", pool)
		//Todo: get image list
		// rbd ls -l -p rbd --format json
		ImageListCommand := []string{
			"rbd", "ls", "-l", "-p", pool, "--format", "json",
		}
		ImageListReturn, err := utils.Execute(ImageListCommand)

		utils.HandleError(err)
		ImageList := make([]*RBD.RBD, 0)
		err = json.Unmarshal(ImageListReturn, &ImageList)
		utils.HandleError(err)

		for _, image := range ImageList {
			Logger := slog.New(utils.AbleJsonHandler)
			if image.Parent != nil {
				image.Type = "ChildImage"
			} else if image.Snapshot != nil {
				image.Type = "SnapshotImage"
			} else {
				image.Type = "Image"
			}
			image.Name = image.Image
			image.Pool = pool

			//if image.Name == "ccvm" {
			image.Update()
			Images = append(Images, image)
			Logger.Debug("Image in "+pool, "Image", image)
			//}
		}

	}
	utils.HandleError(err)
	tmplist = append(tmplist, Images...)

	//
	//for i := range s.RBDs {
	//	s.RBDs = append(s.RBDs[:i], s.RBDs[i+1:]...)
	//}
	//for i := range tmplist {
	//	s.RBDs = append(s.RBDs, tmplist[i])
	//}
	s.RBDs = tmplist
	//slog.Info("RBDs:", s.RBDs)

	/* 기존 목록에 신규정보 업데이트
	// 1. 목록 업데이트
	tmplist=[]
	PoolList = ceph osd pool ls?
	for pool in PoolList
	  tmplist.append(rbd -p pool ls -l)
	for disk in tmplist
	  if disk not in RBDs
	     RBDs.append(disk)
	for disk in RBDs
	  disk.update()

	// 2. 목록 교체
	tmplist=[]
	PoolList = ceph osd pool ls?
	for pool in PoolList
	  tmplist.append(rbd -p pool ls -l)
	for disk in tmplist
	  disk.update()
	RBDs = tmplist
	*/
}

func (s *DiskController) UpdateBLKs() {

}
