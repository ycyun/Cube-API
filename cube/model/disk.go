package model

import (
	"encoding/xml"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"github.com/ycyun/Cube-API/utils"
	"net/http"
	"os/exec"
	"reflect"
	"sync"
	"time"
)

type TypeBlockDevice struct {
	Blockdevices []struct {
		Name         string        `json:"name"`
		Kname        string        `json:"kname"`
		Path         string        `json:"path"`
		MajMin       string        `json:"maj:min"`
		Fsavail      interface{}   `json:"fsavail"`
		Fssize       interface{}   `json:"fssize"`
		Fstype       *string       `json:"fstype"`
		Fsused       interface{}   `json:"fsused"`
		Fsuse        interface{}   `json:"fsuse%"`
		Fsroots      []interface{} `json:"fsroots"`
		Fsver        *string       `json:"fsver"`
		Mountpoint   interface{}   `json:"mountpoint"`
		Mountpoints  []interface{} `json:"mountpoints"`
		Label        interface{}   `json:"label"`
		Uuid         *string       `json:"uuid"`
		Ptuuid       *string       `json:"ptuuid"`
		Pttype       *string       `json:"pttype"`
		Parttype     interface{}   `json:"parttype"`
		Parttypename interface{}   `json:"parttypename"`
		Partlabel    interface{}   `json:"partlabel"`
		Partuuid     interface{}   `json:"partuuid"`
		Partflags    interface{}   `json:"partflags"`
		Ra           int           `json:"ra"`
		Ro           bool          `json:"ro"`
		Rm           bool          `json:"rm"`
		Hotplug      bool          `json:"hotplug"`
		Model        *string       `json:"model"`
		Serial       *string       `json:"serial"`
		Size         string        `json:"size"`
		State        *string       `json:"state"`
		Owner        string        `json:"owner"`
		Group        string        `json:"group"`
		Mode         string        `json:"mode"`
		Alignment    int           `json:"alignment"`
		MinIo        int           `json:"min-io"`
		OptIo        int           `json:"opt-io"`
		PhySec       int           `json:"phy-sec"`
		LogSec       int           `json:"log-sec"`
		Rota         bool          `json:"rota"`
		Sched        string        `json:"sched"`
		RqSize       int           `json:"rq-size"`
		Type         string        `json:"type"`
		DiscAln      int           `json:"disc-aln"`
		DiscGran     string        `json:"disc-gran"`
		DiscMax      string        `json:"disc-max"`
		DiscZero     bool          `json:"disc-zero"`
		Wsame        string        `json:"wsame"`
		Wwn          *string       `json:"wwn"`
		Rand         bool          `json:"rand"`
		Pkname       interface{}   `json:"pkname"`
		Hctl         *string       `json:"hctl"`
		Tran         *string       `json:"tran"`
		Subsystems   string        `json:"subsystems"`
		Rev          *string       `json:"rev"`
		Vendor       *string       `json:"vendor"`
		Zoned        string        `json:"zoned"`
		Dax          bool          `json:"dax"`
		Children     []struct {
			Name         string      `json:"name"`
			Kname        string      `json:"kname"`
			Path         string      `json:"path"`
			MajMin       string      `json:"maj:min"`
			Fsavail      *string     `json:"fsavail"`
			Fssize       *string     `json:"fssize"`
			Fstype       *string     `json:"fstype"`
			Fsused       *string     `json:"fsused"`
			Fsuse        *string     `json:"fsuse%"`
			Fsroots      []string    `json:"fsroots"`
			Fsver        *string     `json:"fsver"`
			Mountpoint   *string     `json:"mountpoint"`
			Mountpoints  []string    `json:"mountpoints"`
			Label        interface{} `json:"label"`
			Uuid         *string     `json:"uuid"`
			Ptuuid       *string     `json:"ptuuid"`
			Pttype       *string     `json:"pttype"`
			Parttype     *string     `json:"parttype"`
			Parttypename *string     `json:"parttypename"`
			Partlabel    interface{} `json:"partlabel"`
			Partuuid     *string     `json:"partuuid"`
			Partflags    *string     `json:"partflags"`
			Ra           int         `json:"ra"`
			Ro           bool        `json:"ro"`
			Rm           bool        `json:"rm"`
			Hotplug      bool        `json:"hotplug"`
			Model        interface{} `json:"model"`
			Serial       interface{} `json:"serial"`
			Size         string      `json:"size"`
			State        *string     `json:"state"`
			Owner        string      `json:"owner"`
			Group        string      `json:"group"`
			Mode         string      `json:"mode"`
			Alignment    int         `json:"alignment"`
			MinIo        int         `json:"min-io"`
			OptIo        int         `json:"opt-io"`
			PhySec       int         `json:"phy-sec"`
			LogSec       int         `json:"log-sec"`
			Rota         bool        `json:"rota"`
			Sched        *string     `json:"sched"`
			RqSize       int         `json:"rq-size"`
			Type         string      `json:"type"`
			DiscAln      int         `json:"disc-aln"`
			DiscGran     string      `json:"disc-gran"`
			DiscMax      string      `json:"disc-max"`
			DiscZero     bool        `json:"disc-zero"`
			Wsame        string      `json:"wsame"`
			Wwn          *string     `json:"wwn"`
			Rand         bool        `json:"rand"`
			Pkname       string      `json:"pkname"`
			Hctl         interface{} `json:"hctl"`
			Tran         interface{} `json:"tran"`
			Subsystems   string      `json:"subsystems"`
			Rev          interface{} `json:"rev"`
			Vendor       interface{} `json:"vendor"`
			Zoned        string      `json:"zoned"`
			Dax          bool        `json:"dax"`
			Children     []struct {
				Name         string      `json:"name"`
				Kname        string      `json:"kname"`
				Path         string      `json:"path"`
				MajMin       string      `json:"maj:min"`
				Fsavail      *string     `json:"fsavail"`
				Fssize       *string     `json:"fssize"`
				Fstype       *string     `json:"fstype"`
				Fsused       *string     `json:"fsused"`
				Fsuse        *string     `json:"fsuse%"`
				Fsroots      []string    `json:"fsroots"`
				Fsver        *string     `json:"fsver"`
				Mountpoint   *string     `json:"mountpoint"`
				Mountpoints  []string    `json:"mountpoints"`
				Label        interface{} `json:"label"`
				Uuid         *string     `json:"uuid"`
				Ptuuid       *string     `json:"ptuuid"`
				Pttype       *string     `json:"pttype"`
				Parttype     interface{} `json:"parttype"`
				Parttypename interface{} `json:"parttypename"`
				Partlabel    interface{} `json:"partlabel"`
				Partuuid     interface{} `json:"partuuid"`
				Partflags    interface{} `json:"partflags"`
				Ra           int         `json:"ra"`
				Ro           bool        `json:"ro"`
				Rm           bool        `json:"rm"`
				Hotplug      bool        `json:"hotplug"`
				Model        interface{} `json:"model"`
				Serial       interface{} `json:"serial"`
				Size         string      `json:"size"`
				State        string      `json:"state"`
				Owner        string      `json:"owner"`
				Group        string      `json:"group"`
				Mode         string      `json:"mode"`
				Alignment    int         `json:"alignment"`
				MinIo        int         `json:"min-io"`
				OptIo        int         `json:"opt-io"`
				PhySec       int         `json:"phy-sec"`
				LogSec       int         `json:"log-sec"`
				Rota         bool        `json:"rota"`
				Sched        interface{} `json:"sched"`
				RqSize       int         `json:"rq-size"`
				Type         string      `json:"type"`
				DiscAln      int         `json:"disc-aln"`
				DiscGran     string      `json:"disc-gran"`
				DiscMax      string      `json:"disc-max"`
				DiscZero     bool        `json:"disc-zero"`
				Wsame        string      `json:"wsame"`
				Wwn          interface{} `json:"wwn"`
				Rand         bool        `json:"rand"`
				Pkname       string      `json:"pkname"`
				Hctl         interface{} `json:"hctl"`
				Tran         interface{} `json:"tran"`
				Subsystems   string      `json:"subsystems"`
				Rev          interface{} `json:"rev"`
				Vendor       interface{} `json:"vendor"`
				Zoned        string      `json:"zoned"`
				Dax          bool        `json:"dax"`
			} `json:"children,omitempty"`
		} `json:"children,omitempty"`
	} `json:"blockdevices"`
	RefreshTime time.Time
} // @name TypeBlockDevice

var lockBlockDevice sync.Once

var _BlockDevice *TypeBlockDevice

func Disk() *TypeBlockDevice {
	if _BlockDevice == nil {
		lockBlockDevice.Do(
			func() {
				if gin.IsDebugging() {
					fmt.Println("Creating ", reflect.TypeOf(_BlockDevice), " now.")
				}
				_BlockDevice = &TypeBlockDevice{}
			})
	} else {
		if gin.IsDebugging() {
			fmt.Println("get old ", reflect.TypeOf(_BlockDevice), " instance.")
		}
	}
	return _BlockDevice
}

// Get godoc
//
//	@Summary		Show List of Disk
//	@Description	Cube의 Disk목록을 보여줍니다.
//	@Tags			API, CUBE
//	@Accept			x-www-form-urlencoded
//	@Produce		json
//	@Success		200	{object}	TypeBlockDevice
//	@Failure		400	{object}	HTTP400BadRequest
//	@Failure		404	{object}	HTTP404NotFound
//	@Failure		500	{object}	HTTP500InternalServerError
//	@Router			/cube/disk [get]
func (d *TypeBlockDevice) Get(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, d)
}

func (d *TypeBlockDevice) Update() {
	if gin.Mode() == gin.ReleaseMode {
		var stdout []byte

		cmd := exec.Command("lsblk", "-J", "-O")
		stdout, _ = cmd.CombinedOutput()
		if err := json.Unmarshal(stdout, &_BlockDevice); err != nil && gin.IsDebugging() {
			utils.FancyHandleError(err)

		}
	} else {

		stdout := []byte("{\"blockdevices\":[{\"name\":\"sda\",\"kname\":\"sda\",\"path\":\"/dev/sda\",\"maj:min\":\"8:0\",\"fsavail\":null,\"fssize\":null,\"fstype\":\"LVM2_member\",\"fsused\":null,\"fsuse%\":null,\"fsroots\":[null],\"fsver\":\"LVM2 001\",\"mountpoint\":null,\"mountpoints\":[null],\"label\":null,\"uuid\":\"5ap1Kx-1Z47-RecG-0bbj-KCz8-TTc7-hO9Xsj\",\"ptuuid\":null,\"pttype\":null,\"parttype\":null,\"parttypename\":null,\"partlabel\":null,\"partuuid\":null,\"partflags\":null,\"ra\":4096,\"ro\":false,\"rm\":false,\"hotplug\":false,\"model\":\"XA3840LE10063\",\"serial\":\"HKT0127E\",\"size\":\"3.5T\",\"state\":\"running\",\"owner\":\"root\",\"group\":\"disk\",\"mode\":\"brw-rw----\",\"alignment\":0,\"min-io\":4096,\"opt-io\":0,\"phy-sec\":4096,\"log-sec\":512,\"rota\":false,\"sched\":\"mq-deadline\",\"rq-size\":256,\"type\":\"disk\",\"disc-aln\":0,\"disc-gran\":\"4K\",\"disc-max\":\"128M\",\"disc-zero\":false,\"wsame\":\"0B\",\"wwn\":\"0x5000c5003ea0103f\",\"rand\":false,\"pkname\":null,\"hctl\":\"0:0:6:0\",\"tran\":null,\"subsystems\":\"block:scsi:pci\",\"rev\":\"1121\",\"vendor\":\"ATA     \",\"zoned\":\"none\",\"dax\":false,\"children\":[{\"name\":\"ceph--dc770352--040b--48a7--822c--f9c18f03885a-osd--block--e38b7673--b131--40dc--9609--1e34e9216f5b\",\"kname\":\"dm-2\",\"path\":\"/dev/mapper/ceph--dc770352--040b--48a7--822c--f9c18f03885a-osd--block--e38b7673--b131--40dc--9609--1e34e9216f5b\",\"maj:min\":\"253:2\",\"fsavail\":null,\"fssize\":null,\"fstype\":null,\"fsused\":null,\"fsuse%\":null,\"fsroots\":[null],\"fsver\":null,\"mountpoint\":null,\"mountpoints\":[null],\"label\":null,\"uuid\":null,\"ptuuid\":null,\"pttype\":null,\"parttype\":null,\"parttypename\":null,\"partlabel\":null,\"partuuid\":null,\"partflags\":null,\"ra\":4096,\"ro\":false,\"rm\":false,\"hotplug\":false,\"model\":null,\"serial\":null,\"size\":\"3.5T\",\"state\":\"running\",\"owner\":\"root\",\"group\":\"disk\",\"mode\":\"brw-rw----\",\"alignment\":0,\"min-io\":4096,\"opt-io\":0,\"phy-sec\":4096,\"log-sec\":512,\"rota\":false,\"sched\":null,\"rq-size\":null,\"type\":\"lvm\",\"disc-aln\":0,\"disc-gran\":\"4K\",\"disc-max\":\"128M\",\"disc-zero\":false,\"wsame\":\"0B\",\"wwn\":null,\"rand\":false,\"pkname\":\"sda\",\"hctl\":null,\"tran\":null,\"subsystems\":\"block\",\"rev\":null,\"vendor\":null,\"zoned\":\"none\",\"dax\":false}]},{\"name\":\"sdb\",\"kname\":\"sdb\",\"path\":\"/dev/sdb\",\"maj:min\":\"8:16\",\"fsavail\":null,\"fssize\":null,\"fstype\":null,\"fsused\":null,\"fsuse%\":null,\"fsroots\":[null],\"fsver\":null,\"mountpoint\":null,\"mountpoints\":[null],\"label\":null,\"uuid\":null,\"ptuuid\":\"335318d1-33e6-49c6-b148-be25e85548f1\",\"pttype\":\"gpt\",\"parttype\":null,\"parttypename\":null,\"partlabel\":null,\"partuuid\":null,\"partflags\":null,\"ra\":4096,\"ro\":false,\"rm\":false,\"hotplug\":false,\"model\":\"RAID 930-8i-2GB\",\"serial\":\"00ec376d050a62d528409b5f07b26200\",\"size\":\"446.1G\",\"state\":\"running\",\"owner\":\"root\",\"group\":\"disk\",\"mode\":\"brw-rw----\",\"alignment\":0,\"min-io\":262144,\"opt-io\":262144,\"phy-sec\":512,\"log-sec\":512,\"rota\":false,\"sched\":\"mq-deadline\",\"rq-size\":256,\"type\":\"disk\",\"disc-aln\":0,\"disc-gran\":\"0B\",\"disc-max\":\"0B\",\"disc-zero\":false,\"wsame\":\"0B\",\"wwn\":\"0x600062b2075f9b4028d5620a056d37ec\",\"rand\":false,\"pkname\":null,\"hctl\":\"0:2:0:0\",\"tran\":null,\"subsystems\":\"block:scsi:pci\",\"rev\":\"5.15\",\"vendor\":\"Lenovo  \",\"zoned\":\"none\",\"dax\":false,\"children\":[{\"name\":\"sdb1\",\"kname\":\"sdb1\",\"path\":\"/dev/sdb1\",\"maj:min\":\"8:17\",\"fsavail\":\"591.8M\",\"fssize\":\"598.8M\",\"fstype\":\"vfat\",\"fsused\":\"7M\",\"fsuse%\":\"1%\",\"fsroots\":[\"/\"],\"fsver\":\"FAT32\",\"mountpoint\":\"/boot/efi\",\"mountpoints\":[\"/boot/efi\"],\"label\":null,\"uuid\":\"D64E-6846\",\"ptuuid\":\"335318d1-33e6-49c6-b148-be25e85548f1\",\"pttype\":\"gpt\",\"parttype\":\"c12a7328-f81f-11d2-ba4b-00a0c93ec93b\",\"parttypename\":\"EFI System\",\"partlabel\":\"EFI System Partition\",\"partuuid\":\"9c772974-1da5-4a18-9443-3e57885b6813\",\"partflags\":null,\"ra\":4096,\"ro\":false,\"rm\":false,\"hotplug\":false,\"model\":null,\"serial\":null,\"size\":\"600M\",\"state\":null,\"owner\":\"root\",\"group\":\"disk\",\"mode\":\"brw-rw----\",\"alignment\":0,\"min-io\":262144,\"opt-io\":262144,\"phy-sec\":512,\"log-sec\":512,\"rota\":false,\"sched\":\"mq-deadline\",\"rq-size\":256,\"type\":\"part\",\"disc-aln\":0,\"disc-gran\":\"0B\",\"disc-max\":\"0B\",\"disc-zero\":false,\"wsame\":\"0B\",\"wwn\":\"0x600062b2075f9b4028d5620a056d37ec\",\"rand\":false,\"pkname\":\"sdb\",\"hctl\":null,\"tran\":null,\"subsystems\":\"block:scsi:pci\",\"rev\":null,\"vendor\":null,\"zoned\":\"none\",\"dax\":false},{\"name\":\"sdb2\",\"kname\":\"sdb2\",\"path\":\"/dev/sdb2\",\"maj:min\":\"8:18\",\"fsavail\":\"702.8M\",\"fssize\":\"1006M\",\"fstype\":\"xfs\",\"fsused\":\"303.2M\",\"fsuse%\":\"30%\",\"fsroots\":[\"/\"],\"fsver\":null,\"mountpoint\":\"/boot\",\"mountpoints\":[\"/boot\"],\"label\":null,\"uuid\":\"34f222db-f7f9-41ee-b4f1-f7c95670efbf\",\"ptuuid\":\"335318d1-33e6-49c6-b148-be25e85548f1\",\"pttype\":\"gpt\",\"parttype\":\"0fc63daf-8483-4772-8e79-3d69d8477de4\",\"parttypename\":\"Linux filesystem\",\"partlabel\":null,\"partuuid\":\"21be1667-a4f8-40e5-967b-d900529ce47d\",\"partflags\":null,\"ra\":4096,\"ro\":false,\"rm\":false,\"hotplug\":false,\"model\":null,\"serial\":null,\"size\":\"1G\",\"state\":null,\"owner\":\"root\",\"group\":\"disk\",\"mode\":\"brw-rw----\",\"alignment\":0,\"min-io\":262144,\"opt-io\":262144,\"phy-sec\":512,\"log-sec\":512,\"rota\":false,\"sched\":\"mq-deadline\",\"rq-size\":256,\"type\":\"part\",\"disc-aln\":0,\"disc-gran\":\"0B\",\"disc-max\":\"0B\",\"disc-zero\":false,\"wsame\":\"0B\",\"wwn\":\"0x600062b2075f9b4028d5620a056d37ec\",\"rand\":false,\"pkname\":\"sdb\",\"hctl\":null,\"tran\":null,\"subsystems\":\"block:scsi:pci\",\"rev\":null,\"vendor\":null,\"zoned\":\"none\",\"dax\":false},{\"name\":\"sdb3\",\"kname\":\"sdb3\",\"path\":\"/dev/sdb3\",\"maj:min\":\"8:19\",\"fsavail\":null,\"fssize\":null,\"fstype\":\"LVM2_member\",\"fsused\":null,\"fsuse%\":null,\"fsroots\":[null],\"fsver\":\"LVM2 001\",\"mountpoint\":null,\"mountpoints\":[null],\"label\":null,\"uuid\":\"qncuJa-Nl3y-bGGg-hKQz-BOku-MWES-nFO0Bv\",\"ptuuid\":\"335318d1-33e6-49c6-b148-be25e85548f1\",\"pttype\":\"gpt\",\"parttype\":\"e6d6d379-f507-44c2-a23c-238f2a3df928\",\"parttypename\":\"Linux LVM\",\"partlabel\":null,\"partuuid\":\"c2a0409d-8052-45cd-b90e-99fc7d088957\",\"partflags\":null,\"ra\":4096,\"ro\":false,\"rm\":false,\"hotplug\":false,\"model\":null,\"serial\":null,\"size\":\"444.5G\",\"state\":null,\"owner\":\"root\",\"group\":\"disk\",\"mode\":\"brw-rw----\",\"alignment\":0,\"min-io\":262144,\"opt-io\":262144,\"phy-sec\":512,\"log-sec\":512,\"rota\":false,\"sched\":\"mq-deadline\",\"rq-size\":256,\"type\":\"part\",\"disc-aln\":0,\"disc-gran\":\"0B\",\"disc-max\":\"0B\",\"disc-zero\":false,\"wsame\":\"0B\",\"wwn\":\"0x600062b2075f9b4028d5620a056d37ec\",\"rand\":false,\"pkname\":\"sdb\",\"hctl\":null,\"tran\":null,\"subsystems\":\"block:scsi:pci\",\"rev\":null,\"vendor\":null,\"zoned\":\"none\",\"dax\":false,\"children\":[{\"name\":\"ablestack_ablecube13--1-root\",\"kname\":\"dm-0\",\"path\":\"/dev/mapper/ablestack_ablecube13--1-root\",\"maj:min\":\"253:0\",\"fsavail\":\"322.3G\",\"fssize\":\"380.3G\",\"fstype\":\"xfs\",\"fsused\":\"58G\",\"fsuse%\":\"15%\",\"fsroots\":[\"/\"],\"fsver\":null,\"mountpoint\":\"/\",\"mountpoints\":[\"/\"],\"label\":null,\"uuid\":\"a3d08e53-5c13-43e5-97e1-af007bb4ab0b\",\"ptuuid\":null,\"pttype\":null,\"parttype\":null,\"parttypename\":null,\"partlabel\":null,\"partuuid\":null,\"partflags\":null,\"ra\":4096,\"ro\":false,\"rm\":false,\"hotplug\":false,\"model\":null,\"serial\":null,\"size\":\"380.5G\",\"state\":\"running\",\"owner\":\"root\",\"group\":\"disk\",\"mode\":\"brw-rw----\",\"alignment\":0,\"min-io\":262144,\"opt-io\":262144,\"phy-sec\":512,\"log-sec\":512,\"rota\":false,\"sched\":null,\"rq-size\":null,\"type\":\"lvm\",\"disc-aln\":0,\"disc-gran\":\"0B\",\"disc-max\":\"0B\",\"disc-zero\":false,\"wsame\":\"0B\",\"wwn\":null,\"rand\":false,\"pkname\":\"sdb3\",\"hctl\":null,\"tran\":null,\"subsystems\":\"block\",\"rev\":null,\"vendor\":null,\"zoned\":\"none\",\"dax\":false},{\"name\":\"ablestack_ablecube13--1-swap\",\"kname\":\"dm-1\",\"path\":\"/dev/mapper/ablestack_ablecube13--1-swap\",\"maj:min\":\"253:1\",\"fsavail\":null,\"fssize\":null,\"fstype\":\"swap\",\"fsused\":null,\"fsuse%\":null,\"fsroots\":[null],\"fsver\":\"1\",\"mountpoint\":\"[SWAP]\",\"mountpoints\":[\"[SWAP]\"],\"label\":null,\"uuid\":\"e2286fe3-2f15-4cdf-8b77-7e9fa5d3f0b3\",\"ptuuid\":null,\"pttype\":null,\"parttype\":null,\"parttypename\":null,\"partlabel\":null,\"partuuid\":null,\"partflags\":null,\"ra\":4096,\"ro\":false,\"rm\":false,\"hotplug\":false,\"model\":null,\"serial\":null,\"size\":\"64G\",\"state\":\"running\",\"owner\":\"root\",\"group\":\"disk\",\"mode\":\"brw-rw----\",\"alignment\":0,\"min-io\":262144,\"opt-io\":262144,\"phy-sec\":512,\"log-sec\":512,\"rota\":false,\"sched\":null,\"rq-size\":null,\"type\":\"lvm\",\"disc-aln\":0,\"disc-gran\":\"0B\",\"disc-max\":\"0B\",\"disc-zero\":false,\"wsame\":\"0B\",\"wwn\":null,\"rand\":false,\"pkname\":\"sdb3\",\"hctl\":null,\"tran\":null,\"subsystems\":\"block\",\"rev\":null,\"vendor\":null,\"zoned\":\"none\",\"dax\":false}]}]},{\"name\":\"rbd0\",\"kname\":\"rbd0\",\"path\":\"/dev/rbd0\",\"maj:min\":\"252:0\",\"fsavail\":null,\"fssize\":null,\"fstype\":null,\"fsused\":null,\"fsuse%\":null,\"fsroots\":[null],\"fsver\":null,\"mountpoint\":null,\"mountpoints\":[null],\"label\":null,\"uuid\":null,\"ptuuid\":null,\"pttype\":null,\"parttype\":null,\"parttypename\":null,\"partlabel\":null,\"partuuid\":null,\"partflags\":null,\"ra\":128,\"ro\":false,\"rm\":false,\"hotplug\":false,\"model\":null,\"serial\":null,\"size\":\"4.9G\",\"state\":null,\"owner\":\"root\",\"group\":\"disk\",\"mode\":\"brw-rw----\",\"alignment\":0,\"min-io\":65536,\"opt-io\":65536,\"phy-sec\":512,\"log-sec\":512,\"rota\":false,\"sched\":\"none\",\"rq-size\":128,\"type\":\"disk\",\"disc-aln\":0,\"disc-gran\":\"64K\",\"disc-max\":\"4M\",\"disc-zero\":false,\"wsame\":\"0B\",\"wwn\":null,\"rand\":false,\"pkname\":null,\"hctl\":null,\"tran\":null,\"subsystems\":\"block:rbd\",\"rev\":null,\"vendor\":null,\"zoned\":\"none\",\"dax\":false,\"children\":[{\"name\":\"rbd0p1\",\"kname\":\"rbd0p1\",\"path\":\"/dev/rbd0p1\",\"maj:min\":\"252:1\",\"fsavail\":null,\"fssize\":null,\"fstype\":null,\"fsused\":null,\"fsuse%\":null,\"fsroots\":[null],\"fsver\":null,\"mountpoint\":null,\"mountpoints\":[null],\"label\":null,\"uuid\":null,\"ptuuid\":null,\"pttype\":null,\"parttype\":null,\"parttypename\":null,\"partlabel\":null,\"partuuid\":null,\"partflags\":null,\"ra\":128,\"ro\":false,\"rm\":false,\"hotplug\":false,\"model\":null,\"serial\":null,\"size\":\"380M\",\"state\":null,\"owner\":\"root\",\"group\":\"disk\",\"mode\":\"brw-rw----\",\"alignment\":0,\"min-io\":65536,\"opt-io\":65536,\"phy-sec\":512,\"log-sec\":512,\"rota\":false,\"sched\":\"none\",\"rq-size\":128,\"type\":\"part\",\"disc-aln\":0,\"disc-gran\":\"64K\",\"disc-max\":\"4M\",\"disc-zero\":false,\"wsame\":\"0B\",\"wwn\":null,\"rand\":false,\"pkname\":\"rbd0\",\"hctl\":null,\"tran\":null,\"subsystems\":\"block:rbd\",\"rev\":null,\"vendor\":null,\"zoned\":\"none\",\"dax\":false},{\"name\":\"rbd0p2\",\"kname\":\"rbd0p2\",\"path\":\"/dev/rbd0p2\",\"maj:min\":\"252:2\",\"fsavail\":null,\"fssize\":null,\"fstype\":null,\"fsused\":null,\"fsuse%\":null,\"fsroots\":[null],\"fsver\":null,\"mountpoint\":null,\"mountpoints\":[null],\"label\":null,\"uuid\":null,\"ptuuid\":null,\"pttype\":null,\"parttype\":null,\"parttypename\":null,\"partlabel\":null,\"partuuid\":null,\"partflags\":null,\"ra\":128,\"ro\":false,\"rm\":false,\"hotplug\":false,\"model\":null,\"serial\":null,\"size\":\"1K\",\"state\":null,\"owner\":\"root\",\"group\":\"disk\",\"mode\":\"brw-rw----\",\"alignment\":1024,\"min-io\":65536,\"opt-io\":65536,\"phy-sec\":512,\"log-sec\":512,\"rota\":false,\"sched\":\"none\",\"rq-size\":128,\"type\":\"part\",\"disc-aln\":1024,\"disc-gran\":\"64K\",\"disc-max\":\"4M\",\"disc-zero\":false,\"wsame\":\"0B\",\"wwn\":null,\"rand\":false,\"pkname\":\"rbd0\",\"hctl\":null,\"tran\":null,\"subsystems\":\"block:rbd\",\"rev\":null,\"vendor\":null,\"zoned\":\"none\",\"dax\":false},{\"name\":\"rbd0p5\",\"kname\":\"rbd0p5\",\"path\":\"/dev/rbd0p5\",\"maj:min\":\"252:5\",\"fsavail\":null,\"fssize\":null,\"fstype\":null,\"fsused\":null,\"fsuse%\":null,\"fsroots\":[null],\"fsver\":null,\"mountpoint\":null,\"mountpoints\":[null],\"label\":null,\"uuid\":null,\"ptuuid\":null,\"pttype\":null,\"parttype\":null,\"parttypename\":null,\"partlabel\":null,\"partuuid\":null,\"partflags\":null,\"ra\":128,\"ro\":false,\"rm\":false,\"hotplug\":false,\"model\":null,\"serial\":null,\"size\":\"487M\",\"state\":null,\"owner\":\"root\",\"group\":\"disk\",\"mode\":\"brw-rw----\",\"alignment\":0,\"min-io\":65536,\"opt-io\":65536,\"phy-sec\":512,\"log-sec\":512,\"rota\":false,\"sched\":\"none\",\"rq-size\":128,\"type\":\"part\",\"disc-aln\":0,\"disc-gran\":\"64K\",\"disc-max\":\"4M\",\"disc-zero\":false,\"wsame\":\"0B\",\"wwn\":null,\"rand\":false,\"pkname\":\"rbd0\",\"hctl\":null,\"tran\":null,\"subsystems\":\"block:rbd\",\"rev\":null,\"vendor\":null,\"zoned\":\"none\",\"dax\":false},{\"name\":\"rbd0p6\",\"kname\":\"rbd0p6\",\"path\":\"/dev/rbd0p6\",\"maj:min\":\"252:6\",\"fsavail\":null,\"fssize\":null,\"fstype\":null,\"fsused\":null,\"fsuse%\":null,\"fsroots\":[null],\"fsver\":null,\"mountpoint\":null,\"mountpoints\":[null],\"label\":null,\"uuid\":null,\"ptuuid\":null,\"pttype\":null,\"parttype\":null,\"parttypename\":null,\"partlabel\":null,\"partuuid\":null,\"partflags\":null,\"ra\":128,\"ro\":false,\"rm\":false,\"hotplug\":false,\"model\":null,\"serial\":null,\"size\":\"4G\",\"state\":null,\"owner\":\"root\",\"group\":\"disk\",\"mode\":\"brw-rw----\",\"alignment\":0,\"min-io\":65536,\"opt-io\":65536,\"phy-sec\":512,\"log-sec\":512,\"rota\":false,\"sched\":\"none\",\"rq-size\":128,\"type\":\"part\",\"disc-aln\":0,\"disc-gran\":\"64K\",\"disc-max\":\"4M\",\"disc-zero\":false,\"wsame\":\"0B\",\"wwn\":null,\"rand\":false,\"pkname\":\"rbd0\",\"hctl\":null,\"tran\":null,\"subsystems\":\"block:rbd\",\"rev\":null,\"vendor\":null,\"zoned\":\"none\",\"dax\":false}]}]}\n")

		if err := xml.Unmarshal(stdout, &_BlockDevice); err != nil {
			utils.FancyHandleError(err)

		}
	}
	_BlockDevice.RefreshTime = time.Now()

}

func Update() *TypeBlockDevice {
	if gin.Mode() == gin.ReleaseMode {
		var stdout []byte

		cmd := exec.Command("lsblk", "-J", "-O")
		stdout, _ = cmd.CombinedOutput()
		if err := json.Unmarshal(stdout, &_BlockDevice); err != nil && gin.IsDebugging() {
			utils.FancyHandleError(err)

		}
	} else {

		stdout := []byte("{\"blockdevices\":[{\"name\":\"sda\",\"kname\":\"sda\",\"path\":\"/dev/sda\",\"maj:min\":\"8:0\",\"fsavail\":null,\"fssize\":null,\"fstype\":\"LVM2_member\",\"fsused\":null,\"fsuse%\":null,\"fsroots\":[null],\"fsver\":\"LVM2 001\",\"mountpoint\":null,\"mountpoints\":[null],\"label\":null,\"uuid\":\"5ap1Kx-1Z47-RecG-0bbj-KCz8-TTc7-hO9Xsj\",\"ptuuid\":null,\"pttype\":null,\"parttype\":null,\"parttypename\":null,\"partlabel\":null,\"partuuid\":null,\"partflags\":null,\"ra\":4096,\"ro\":false,\"rm\":false,\"hotplug\":false,\"model\":\"XA3840LE10063\",\"serial\":\"HKT0127E\",\"size\":\"3.5T\",\"state\":\"running\",\"owner\":\"root\",\"group\":\"disk\",\"mode\":\"brw-rw----\",\"alignment\":0,\"min-io\":4096,\"opt-io\":0,\"phy-sec\":4096,\"log-sec\":512,\"rota\":false,\"sched\":\"mq-deadline\",\"rq-size\":256,\"type\":\"disk\",\"disc-aln\":0,\"disc-gran\":\"4K\",\"disc-max\":\"128M\",\"disc-zero\":false,\"wsame\":\"0B\",\"wwn\":\"0x5000c5003ea0103f\",\"rand\":false,\"pkname\":null,\"hctl\":\"0:0:6:0\",\"tran\":null,\"subsystems\":\"block:scsi:pci\",\"rev\":\"1121\",\"vendor\":\"ATA     \",\"zoned\":\"none\",\"dax\":false,\"children\":[{\"name\":\"ceph--dc770352--040b--48a7--822c--f9c18f03885a-osd--block--e38b7673--b131--40dc--9609--1e34e9216f5b\",\"kname\":\"dm-2\",\"path\":\"/dev/mapper/ceph--dc770352--040b--48a7--822c--f9c18f03885a-osd--block--e38b7673--b131--40dc--9609--1e34e9216f5b\",\"maj:min\":\"253:2\",\"fsavail\":null,\"fssize\":null,\"fstype\":null,\"fsused\":null,\"fsuse%\":null,\"fsroots\":[null],\"fsver\":null,\"mountpoint\":null,\"mountpoints\":[null],\"label\":null,\"uuid\":null,\"ptuuid\":null,\"pttype\":null,\"parttype\":null,\"parttypename\":null,\"partlabel\":null,\"partuuid\":null,\"partflags\":null,\"ra\":4096,\"ro\":false,\"rm\":false,\"hotplug\":false,\"model\":null,\"serial\":null,\"size\":\"3.5T\",\"state\":\"running\",\"owner\":\"root\",\"group\":\"disk\",\"mode\":\"brw-rw----\",\"alignment\":0,\"min-io\":4096,\"opt-io\":0,\"phy-sec\":4096,\"log-sec\":512,\"rota\":false,\"sched\":null,\"rq-size\":null,\"type\":\"lvm\",\"disc-aln\":0,\"disc-gran\":\"4K\",\"disc-max\":\"128M\",\"disc-zero\":false,\"wsame\":\"0B\",\"wwn\":null,\"rand\":false,\"pkname\":\"sda\",\"hctl\":null,\"tran\":null,\"subsystems\":\"block\",\"rev\":null,\"vendor\":null,\"zoned\":\"none\",\"dax\":false}]},{\"name\":\"sdb\",\"kname\":\"sdb\",\"path\":\"/dev/sdb\",\"maj:min\":\"8:16\",\"fsavail\":null,\"fssize\":null,\"fstype\":null,\"fsused\":null,\"fsuse%\":null,\"fsroots\":[null],\"fsver\":null,\"mountpoint\":null,\"mountpoints\":[null],\"label\":null,\"uuid\":null,\"ptuuid\":\"335318d1-33e6-49c6-b148-be25e85548f1\",\"pttype\":\"gpt\",\"parttype\":null,\"parttypename\":null,\"partlabel\":null,\"partuuid\":null,\"partflags\":null,\"ra\":4096,\"ro\":false,\"rm\":false,\"hotplug\":false,\"model\":\"RAID 930-8i-2GB\",\"serial\":\"00ec376d050a62d528409b5f07b26200\",\"size\":\"446.1G\",\"state\":\"running\",\"owner\":\"root\",\"group\":\"disk\",\"mode\":\"brw-rw----\",\"alignment\":0,\"min-io\":262144,\"opt-io\":262144,\"phy-sec\":512,\"log-sec\":512,\"rota\":false,\"sched\":\"mq-deadline\",\"rq-size\":256,\"type\":\"disk\",\"disc-aln\":0,\"disc-gran\":\"0B\",\"disc-max\":\"0B\",\"disc-zero\":false,\"wsame\":\"0B\",\"wwn\":\"0x600062b2075f9b4028d5620a056d37ec\",\"rand\":false,\"pkname\":null,\"hctl\":\"0:2:0:0\",\"tran\":null,\"subsystems\":\"block:scsi:pci\",\"rev\":\"5.15\",\"vendor\":\"Lenovo  \",\"zoned\":\"none\",\"dax\":false,\"children\":[{\"name\":\"sdb1\",\"kname\":\"sdb1\",\"path\":\"/dev/sdb1\",\"maj:min\":\"8:17\",\"fsavail\":\"591.8M\",\"fssize\":\"598.8M\",\"fstype\":\"vfat\",\"fsused\":\"7M\",\"fsuse%\":\"1%\",\"fsroots\":[\"/\"],\"fsver\":\"FAT32\",\"mountpoint\":\"/boot/efi\",\"mountpoints\":[\"/boot/efi\"],\"label\":null,\"uuid\":\"D64E-6846\",\"ptuuid\":\"335318d1-33e6-49c6-b148-be25e85548f1\",\"pttype\":\"gpt\",\"parttype\":\"c12a7328-f81f-11d2-ba4b-00a0c93ec93b\",\"parttypename\":\"EFI System\",\"partlabel\":\"EFI System Partition\",\"partuuid\":\"9c772974-1da5-4a18-9443-3e57885b6813\",\"partflags\":null,\"ra\":4096,\"ro\":false,\"rm\":false,\"hotplug\":false,\"model\":null,\"serial\":null,\"size\":\"600M\",\"state\":null,\"owner\":\"root\",\"group\":\"disk\",\"mode\":\"brw-rw----\",\"alignment\":0,\"min-io\":262144,\"opt-io\":262144,\"phy-sec\":512,\"log-sec\":512,\"rota\":false,\"sched\":\"mq-deadline\",\"rq-size\":256,\"type\":\"part\",\"disc-aln\":0,\"disc-gran\":\"0B\",\"disc-max\":\"0B\",\"disc-zero\":false,\"wsame\":\"0B\",\"wwn\":\"0x600062b2075f9b4028d5620a056d37ec\",\"rand\":false,\"pkname\":\"sdb\",\"hctl\":null,\"tran\":null,\"subsystems\":\"block:scsi:pci\",\"rev\":null,\"vendor\":null,\"zoned\":\"none\",\"dax\":false},{\"name\":\"sdb2\",\"kname\":\"sdb2\",\"path\":\"/dev/sdb2\",\"maj:min\":\"8:18\",\"fsavail\":\"702.8M\",\"fssize\":\"1006M\",\"fstype\":\"xfs\",\"fsused\":\"303.2M\",\"fsuse%\":\"30%\",\"fsroots\":[\"/\"],\"fsver\":null,\"mountpoint\":\"/boot\",\"mountpoints\":[\"/boot\"],\"label\":null,\"uuid\":\"34f222db-f7f9-41ee-b4f1-f7c95670efbf\",\"ptuuid\":\"335318d1-33e6-49c6-b148-be25e85548f1\",\"pttype\":\"gpt\",\"parttype\":\"0fc63daf-8483-4772-8e79-3d69d8477de4\",\"parttypename\":\"Linux filesystem\",\"partlabel\":null,\"partuuid\":\"21be1667-a4f8-40e5-967b-d900529ce47d\",\"partflags\":null,\"ra\":4096,\"ro\":false,\"rm\":false,\"hotplug\":false,\"model\":null,\"serial\":null,\"size\":\"1G\",\"state\":null,\"owner\":\"root\",\"group\":\"disk\",\"mode\":\"brw-rw----\",\"alignment\":0,\"min-io\":262144,\"opt-io\":262144,\"phy-sec\":512,\"log-sec\":512,\"rota\":false,\"sched\":\"mq-deadline\",\"rq-size\":256,\"type\":\"part\",\"disc-aln\":0,\"disc-gran\":\"0B\",\"disc-max\":\"0B\",\"disc-zero\":false,\"wsame\":\"0B\",\"wwn\":\"0x600062b2075f9b4028d5620a056d37ec\",\"rand\":false,\"pkname\":\"sdb\",\"hctl\":null,\"tran\":null,\"subsystems\":\"block:scsi:pci\",\"rev\":null,\"vendor\":null,\"zoned\":\"none\",\"dax\":false},{\"name\":\"sdb3\",\"kname\":\"sdb3\",\"path\":\"/dev/sdb3\",\"maj:min\":\"8:19\",\"fsavail\":null,\"fssize\":null,\"fstype\":\"LVM2_member\",\"fsused\":null,\"fsuse%\":null,\"fsroots\":[null],\"fsver\":\"LVM2 001\",\"mountpoint\":null,\"mountpoints\":[null],\"label\":null,\"uuid\":\"qncuJa-Nl3y-bGGg-hKQz-BOku-MWES-nFO0Bv\",\"ptuuid\":\"335318d1-33e6-49c6-b148-be25e85548f1\",\"pttype\":\"gpt\",\"parttype\":\"e6d6d379-f507-44c2-a23c-238f2a3df928\",\"parttypename\":\"Linux LVM\",\"partlabel\":null,\"partuuid\":\"c2a0409d-8052-45cd-b90e-99fc7d088957\",\"partflags\":null,\"ra\":4096,\"ro\":false,\"rm\":false,\"hotplug\":false,\"model\":null,\"serial\":null,\"size\":\"444.5G\",\"state\":null,\"owner\":\"root\",\"group\":\"disk\",\"mode\":\"brw-rw----\",\"alignment\":0,\"min-io\":262144,\"opt-io\":262144,\"phy-sec\":512,\"log-sec\":512,\"rota\":false,\"sched\":\"mq-deadline\",\"rq-size\":256,\"type\":\"part\",\"disc-aln\":0,\"disc-gran\":\"0B\",\"disc-max\":\"0B\",\"disc-zero\":false,\"wsame\":\"0B\",\"wwn\":\"0x600062b2075f9b4028d5620a056d37ec\",\"rand\":false,\"pkname\":\"sdb\",\"hctl\":null,\"tran\":null,\"subsystems\":\"block:scsi:pci\",\"rev\":null,\"vendor\":null,\"zoned\":\"none\",\"dax\":false,\"children\":[{\"name\":\"ablestack_ablecube13--1-root\",\"kname\":\"dm-0\",\"path\":\"/dev/mapper/ablestack_ablecube13--1-root\",\"maj:min\":\"253:0\",\"fsavail\":\"322.3G\",\"fssize\":\"380.3G\",\"fstype\":\"xfs\",\"fsused\":\"58G\",\"fsuse%\":\"15%\",\"fsroots\":[\"/\"],\"fsver\":null,\"mountpoint\":\"/\",\"mountpoints\":[\"/\"],\"label\":null,\"uuid\":\"a3d08e53-5c13-43e5-97e1-af007bb4ab0b\",\"ptuuid\":null,\"pttype\":null,\"parttype\":null,\"parttypename\":null,\"partlabel\":null,\"partuuid\":null,\"partflags\":null,\"ra\":4096,\"ro\":false,\"rm\":false,\"hotplug\":false,\"model\":null,\"serial\":null,\"size\":\"380.5G\",\"state\":\"running\",\"owner\":\"root\",\"group\":\"disk\",\"mode\":\"brw-rw----\",\"alignment\":0,\"min-io\":262144,\"opt-io\":262144,\"phy-sec\":512,\"log-sec\":512,\"rota\":false,\"sched\":null,\"rq-size\":null,\"type\":\"lvm\",\"disc-aln\":0,\"disc-gran\":\"0B\",\"disc-max\":\"0B\",\"disc-zero\":false,\"wsame\":\"0B\",\"wwn\":null,\"rand\":false,\"pkname\":\"sdb3\",\"hctl\":null,\"tran\":null,\"subsystems\":\"block\",\"rev\":null,\"vendor\":null,\"zoned\":\"none\",\"dax\":false},{\"name\":\"ablestack_ablecube13--1-swap\",\"kname\":\"dm-1\",\"path\":\"/dev/mapper/ablestack_ablecube13--1-swap\",\"maj:min\":\"253:1\",\"fsavail\":null,\"fssize\":null,\"fstype\":\"swap\",\"fsused\":null,\"fsuse%\":null,\"fsroots\":[null],\"fsver\":\"1\",\"mountpoint\":\"[SWAP]\",\"mountpoints\":[\"[SWAP]\"],\"label\":null,\"uuid\":\"e2286fe3-2f15-4cdf-8b77-7e9fa5d3f0b3\",\"ptuuid\":null,\"pttype\":null,\"parttype\":null,\"parttypename\":null,\"partlabel\":null,\"partuuid\":null,\"partflags\":null,\"ra\":4096,\"ro\":false,\"rm\":false,\"hotplug\":false,\"model\":null,\"serial\":null,\"size\":\"64G\",\"state\":\"running\",\"owner\":\"root\",\"group\":\"disk\",\"mode\":\"brw-rw----\",\"alignment\":0,\"min-io\":262144,\"opt-io\":262144,\"phy-sec\":512,\"log-sec\":512,\"rota\":false,\"sched\":null,\"rq-size\":null,\"type\":\"lvm\",\"disc-aln\":0,\"disc-gran\":\"0B\",\"disc-max\":\"0B\",\"disc-zero\":false,\"wsame\":\"0B\",\"wwn\":null,\"rand\":false,\"pkname\":\"sdb3\",\"hctl\":null,\"tran\":null,\"subsystems\":\"block\",\"rev\":null,\"vendor\":null,\"zoned\":\"none\",\"dax\":false}]}]},{\"name\":\"rbd0\",\"kname\":\"rbd0\",\"path\":\"/dev/rbd0\",\"maj:min\":\"252:0\",\"fsavail\":null,\"fssize\":null,\"fstype\":null,\"fsused\":null,\"fsuse%\":null,\"fsroots\":[null],\"fsver\":null,\"mountpoint\":null,\"mountpoints\":[null],\"label\":null,\"uuid\":null,\"ptuuid\":null,\"pttype\":null,\"parttype\":null,\"parttypename\":null,\"partlabel\":null,\"partuuid\":null,\"partflags\":null,\"ra\":128,\"ro\":false,\"rm\":false,\"hotplug\":false,\"model\":null,\"serial\":null,\"size\":\"4.9G\",\"state\":null,\"owner\":\"root\",\"group\":\"disk\",\"mode\":\"brw-rw----\",\"alignment\":0,\"min-io\":65536,\"opt-io\":65536,\"phy-sec\":512,\"log-sec\":512,\"rota\":false,\"sched\":\"none\",\"rq-size\":128,\"type\":\"disk\",\"disc-aln\":0,\"disc-gran\":\"64K\",\"disc-max\":\"4M\",\"disc-zero\":false,\"wsame\":\"0B\",\"wwn\":null,\"rand\":false,\"pkname\":null,\"hctl\":null,\"tran\":null,\"subsystems\":\"block:rbd\",\"rev\":null,\"vendor\":null,\"zoned\":\"none\",\"dax\":false,\"children\":[{\"name\":\"rbd0p1\",\"kname\":\"rbd0p1\",\"path\":\"/dev/rbd0p1\",\"maj:min\":\"252:1\",\"fsavail\":null,\"fssize\":null,\"fstype\":null,\"fsused\":null,\"fsuse%\":null,\"fsroots\":[null],\"fsver\":null,\"mountpoint\":null,\"mountpoints\":[null],\"label\":null,\"uuid\":null,\"ptuuid\":null,\"pttype\":null,\"parttype\":null,\"parttypename\":null,\"partlabel\":null,\"partuuid\":null,\"partflags\":null,\"ra\":128,\"ro\":false,\"rm\":false,\"hotplug\":false,\"model\":null,\"serial\":null,\"size\":\"380M\",\"state\":null,\"owner\":\"root\",\"group\":\"disk\",\"mode\":\"brw-rw----\",\"alignment\":0,\"min-io\":65536,\"opt-io\":65536,\"phy-sec\":512,\"log-sec\":512,\"rota\":false,\"sched\":\"none\",\"rq-size\":128,\"type\":\"part\",\"disc-aln\":0,\"disc-gran\":\"64K\",\"disc-max\":\"4M\",\"disc-zero\":false,\"wsame\":\"0B\",\"wwn\":null,\"rand\":false,\"pkname\":\"rbd0\",\"hctl\":null,\"tran\":null,\"subsystems\":\"block:rbd\",\"rev\":null,\"vendor\":null,\"zoned\":\"none\",\"dax\":false},{\"name\":\"rbd0p2\",\"kname\":\"rbd0p2\",\"path\":\"/dev/rbd0p2\",\"maj:min\":\"252:2\",\"fsavail\":null,\"fssize\":null,\"fstype\":null,\"fsused\":null,\"fsuse%\":null,\"fsroots\":[null],\"fsver\":null,\"mountpoint\":null,\"mountpoints\":[null],\"label\":null,\"uuid\":null,\"ptuuid\":null,\"pttype\":null,\"parttype\":null,\"parttypename\":null,\"partlabel\":null,\"partuuid\":null,\"partflags\":null,\"ra\":128,\"ro\":false,\"rm\":false,\"hotplug\":false,\"model\":null,\"serial\":null,\"size\":\"1K\",\"state\":null,\"owner\":\"root\",\"group\":\"disk\",\"mode\":\"brw-rw----\",\"alignment\":1024,\"min-io\":65536,\"opt-io\":65536,\"phy-sec\":512,\"log-sec\":512,\"rota\":false,\"sched\":\"none\",\"rq-size\":128,\"type\":\"part\",\"disc-aln\":1024,\"disc-gran\":\"64K\",\"disc-max\":\"4M\",\"disc-zero\":false,\"wsame\":\"0B\",\"wwn\":null,\"rand\":false,\"pkname\":\"rbd0\",\"hctl\":null,\"tran\":null,\"subsystems\":\"block:rbd\",\"rev\":null,\"vendor\":null,\"zoned\":\"none\",\"dax\":false},{\"name\":\"rbd0p5\",\"kname\":\"rbd0p5\",\"path\":\"/dev/rbd0p5\",\"maj:min\":\"252:5\",\"fsavail\":null,\"fssize\":null,\"fstype\":null,\"fsused\":null,\"fsuse%\":null,\"fsroots\":[null],\"fsver\":null,\"mountpoint\":null,\"mountpoints\":[null],\"label\":null,\"uuid\":null,\"ptuuid\":null,\"pttype\":null,\"parttype\":null,\"parttypename\":null,\"partlabel\":null,\"partuuid\":null,\"partflags\":null,\"ra\":128,\"ro\":false,\"rm\":false,\"hotplug\":false,\"model\":null,\"serial\":null,\"size\":\"487M\",\"state\":null,\"owner\":\"root\",\"group\":\"disk\",\"mode\":\"brw-rw----\",\"alignment\":0,\"min-io\":65536,\"opt-io\":65536,\"phy-sec\":512,\"log-sec\":512,\"rota\":false,\"sched\":\"none\",\"rq-size\":128,\"type\":\"part\",\"disc-aln\":0,\"disc-gran\":\"64K\",\"disc-max\":\"4M\",\"disc-zero\":false,\"wsame\":\"0B\",\"wwn\":null,\"rand\":false,\"pkname\":\"rbd0\",\"hctl\":null,\"tran\":null,\"subsystems\":\"block:rbd\",\"rev\":null,\"vendor\":null,\"zoned\":\"none\",\"dax\":false},{\"name\":\"rbd0p6\",\"kname\":\"rbd0p6\",\"path\":\"/dev/rbd0p6\",\"maj:min\":\"252:6\",\"fsavail\":null,\"fssize\":null,\"fstype\":null,\"fsused\":null,\"fsuse%\":null,\"fsroots\":[null],\"fsver\":null,\"mountpoint\":null,\"mountpoints\":[null],\"label\":null,\"uuid\":null,\"ptuuid\":null,\"pttype\":null,\"parttype\":null,\"parttypename\":null,\"partlabel\":null,\"partuuid\":null,\"partflags\":null,\"ra\":128,\"ro\":false,\"rm\":false,\"hotplug\":false,\"model\":null,\"serial\":null,\"size\":\"4G\",\"state\":null,\"owner\":\"root\",\"group\":\"disk\",\"mode\":\"brw-rw----\",\"alignment\":0,\"min-io\":65536,\"opt-io\":65536,\"phy-sec\":512,\"log-sec\":512,\"rota\":false,\"sched\":\"none\",\"rq-size\":128,\"type\":\"part\",\"disc-aln\":0,\"disc-gran\":\"64K\",\"disc-max\":\"4M\",\"disc-zero\":false,\"wsame\":\"0B\",\"wwn\":null,\"rand\":false,\"pkname\":\"rbd0\",\"hctl\":null,\"tran\":null,\"subsystems\":\"block:rbd\",\"rev\":null,\"vendor\":null,\"zoned\":\"none\",\"dax\":false}]}]}\n")

		if err := xml.Unmarshal(stdout, &_BlockDevice); err != nil {
			utils.FancyHandleError(err)

		}
	}
	_BlockDevice.RefreshTime = time.Now()
	return _BlockDevice
}

func test() {
	stdout := []byte("{\"blockdevices\":[{\"name\":\"sda\",\"kname\":\"sda\",\"path\":\"/dev/sda\",\"maj:min\":\"8:0\",\"fsavail\":null,\"fssize\":null,\"fstype\":\"LVM2_member\",\"fsused\":null,\"fsuse%\":null,\"fsroots\":[null],\"fsver\":\"LVM2 001\",\"mountpoint\":null,\"mountpoints\":[null],\"label\":null,\"uuid\":\"5ap1Kx-1Z47-RecG-0bbj-KCz8-TTc7-hO9Xsj\",\"ptuuid\":null,\"pttype\":null,\"parttype\":null,\"parttypename\":null,\"partlabel\":null,\"partuuid\":null,\"partflags\":null,\"ra\":4096,\"ro\":false,\"rm\":false,\"hotplug\":false,\"model\":\"XA3840LE10063\",\"serial\":\"HKT0127E\",\"size\":\"3.5T\",\"state\":\"running\",\"owner\":\"root\",\"group\":\"disk\",\"mode\":\"brw-rw----\",\"alignment\":0,\"min-io\":4096,\"opt-io\":0,\"phy-sec\":4096,\"log-sec\":512,\"rota\":false,\"sched\":\"mq-deadline\",\"rq-size\":256,\"type\":\"disk\",\"disc-aln\":0,\"disc-gran\":\"4K\",\"disc-max\":\"128M\",\"disc-zero\":false,\"wsame\":\"0B\",\"wwn\":\"0x5000c5003ea0103f\",\"rand\":false,\"pkname\":null,\"hctl\":\"0:0:6:0\",\"tran\":null,\"subsystems\":\"block:scsi:pci\",\"rev\":\"1121\",\"vendor\":\"ATA     \",\"zoned\":\"none\",\"dax\":false,\"children\":[{\"name\":\"ceph--dc770352--040b--48a7--822c--f9c18f03885a-osd--block--e38b7673--b131--40dc--9609--1e34e9216f5b\",\"kname\":\"dm-2\",\"path\":\"/dev/mapper/ceph--dc770352--040b--48a7--822c--f9c18f03885a-osd--block--e38b7673--b131--40dc--9609--1e34e9216f5b\",\"maj:min\":\"253:2\",\"fsavail\":null,\"fssize\":null,\"fstype\":null,\"fsused\":null,\"fsuse%\":null,\"fsroots\":[null],\"fsver\":null,\"mountpoint\":null,\"mountpoints\":[null],\"label\":null,\"uuid\":null,\"ptuuid\":null,\"pttype\":null,\"parttype\":null,\"parttypename\":null,\"partlabel\":null,\"partuuid\":null,\"partflags\":null,\"ra\":4096,\"ro\":false,\"rm\":false,\"hotplug\":false,\"model\":null,\"serial\":null,\"size\":\"3.5T\",\"state\":\"running\",\"owner\":\"root\",\"group\":\"disk\",\"mode\":\"brw-rw----\",\"alignment\":0,\"min-io\":4096,\"opt-io\":0,\"phy-sec\":4096,\"log-sec\":512,\"rota\":false,\"sched\":null,\"rq-size\":null,\"type\":\"lvm\",\"disc-aln\":0,\"disc-gran\":\"4K\",\"disc-max\":\"128M\",\"disc-zero\":false,\"wsame\":\"0B\",\"wwn\":null,\"rand\":false,\"pkname\":\"sda\",\"hctl\":null,\"tran\":null,\"subsystems\":\"block\",\"rev\":null,\"vendor\":null,\"zoned\":\"none\",\"dax\":false}]},{\"name\":\"sdb\",\"kname\":\"sdb\",\"path\":\"/dev/sdb\",\"maj:min\":\"8:16\",\"fsavail\":null,\"fssize\":null,\"fstype\":null,\"fsused\":null,\"fsuse%\":null,\"fsroots\":[null],\"fsver\":null,\"mountpoint\":null,\"mountpoints\":[null],\"label\":null,\"uuid\":null,\"ptuuid\":\"335318d1-33e6-49c6-b148-be25e85548f1\",\"pttype\":\"gpt\",\"parttype\":null,\"parttypename\":null,\"partlabel\":null,\"partuuid\":null,\"partflags\":null,\"ra\":4096,\"ro\":false,\"rm\":false,\"hotplug\":false,\"model\":\"RAID 930-8i-2GB\",\"serial\":\"00ec376d050a62d528409b5f07b26200\",\"size\":\"446.1G\",\"state\":\"running\",\"owner\":\"root\",\"group\":\"disk\",\"mode\":\"brw-rw----\",\"alignment\":0,\"min-io\":262144,\"opt-io\":262144,\"phy-sec\":512,\"log-sec\":512,\"rota\":false,\"sched\":\"mq-deadline\",\"rq-size\":256,\"type\":\"disk\",\"disc-aln\":0,\"disc-gran\":\"0B\",\"disc-max\":\"0B\",\"disc-zero\":false,\"wsame\":\"0B\",\"wwn\":\"0x600062b2075f9b4028d5620a056d37ec\",\"rand\":false,\"pkname\":null,\"hctl\":\"0:2:0:0\",\"tran\":null,\"subsystems\":\"block:scsi:pci\",\"rev\":\"5.15\",\"vendor\":\"Lenovo  \",\"zoned\":\"none\",\"dax\":false,\"children\":[{\"name\":\"sdb1\",\"kname\":\"sdb1\",\"path\":\"/dev/sdb1\",\"maj:min\":\"8:17\",\"fsavail\":\"591.8M\",\"fssize\":\"598.8M\",\"fstype\":\"vfat\",\"fsused\":\"7M\",\"fsuse%\":\"1%\",\"fsroots\":[\"/\"],\"fsver\":\"FAT32\",\"mountpoint\":\"/boot/efi\",\"mountpoints\":[\"/boot/efi\"],\"label\":null,\"uuid\":\"D64E-6846\",\"ptuuid\":\"335318d1-33e6-49c6-b148-be25e85548f1\",\"pttype\":\"gpt\",\"parttype\":\"c12a7328-f81f-11d2-ba4b-00a0c93ec93b\",\"parttypename\":\"EFI System\",\"partlabel\":\"EFI System Partition\",\"partuuid\":\"9c772974-1da5-4a18-9443-3e57885b6813\",\"partflags\":null,\"ra\":4096,\"ro\":false,\"rm\":false,\"hotplug\":false,\"model\":null,\"serial\":null,\"size\":\"600M\",\"state\":null,\"owner\":\"root\",\"group\":\"disk\",\"mode\":\"brw-rw----\",\"alignment\":0,\"min-io\":262144,\"opt-io\":262144,\"phy-sec\":512,\"log-sec\":512,\"rota\":false,\"sched\":\"mq-deadline\",\"rq-size\":256,\"type\":\"part\",\"disc-aln\":0,\"disc-gran\":\"0B\",\"disc-max\":\"0B\",\"disc-zero\":false,\"wsame\":\"0B\",\"wwn\":\"0x600062b2075f9b4028d5620a056d37ec\",\"rand\":false,\"pkname\":\"sdb\",\"hctl\":null,\"tran\":null,\"subsystems\":\"block:scsi:pci\",\"rev\":null,\"vendor\":null,\"zoned\":\"none\",\"dax\":false},{\"name\":\"sdb2\",\"kname\":\"sdb2\",\"path\":\"/dev/sdb2\",\"maj:min\":\"8:18\",\"fsavail\":\"702.8M\",\"fssize\":\"1006M\",\"fstype\":\"xfs\",\"fsused\":\"303.2M\",\"fsuse%\":\"30%\",\"fsroots\":[\"/\"],\"fsver\":null,\"mountpoint\":\"/boot\",\"mountpoints\":[\"/boot\"],\"label\":null,\"uuid\":\"34f222db-f7f9-41ee-b4f1-f7c95670efbf\",\"ptuuid\":\"335318d1-33e6-49c6-b148-be25e85548f1\",\"pttype\":\"gpt\",\"parttype\":\"0fc63daf-8483-4772-8e79-3d69d8477de4\",\"parttypename\":\"Linux filesystem\",\"partlabel\":null,\"partuuid\":\"21be1667-a4f8-40e5-967b-d900529ce47d\",\"partflags\":null,\"ra\":4096,\"ro\":false,\"rm\":false,\"hotplug\":false,\"model\":null,\"serial\":null,\"size\":\"1G\",\"state\":null,\"owner\":\"root\",\"group\":\"disk\",\"mode\":\"brw-rw----\",\"alignment\":0,\"min-io\":262144,\"opt-io\":262144,\"phy-sec\":512,\"log-sec\":512,\"rota\":false,\"sched\":\"mq-deadline\",\"rq-size\":256,\"type\":\"part\",\"disc-aln\":0,\"disc-gran\":\"0B\",\"disc-max\":\"0B\",\"disc-zero\":false,\"wsame\":\"0B\",\"wwn\":\"0x600062b2075f9b4028d5620a056d37ec\",\"rand\":false,\"pkname\":\"sdb\",\"hctl\":null,\"tran\":null,\"subsystems\":\"block:scsi:pci\",\"rev\":null,\"vendor\":null,\"zoned\":\"none\",\"dax\":false},{\"name\":\"sdb3\",\"kname\":\"sdb3\",\"path\":\"/dev/sdb3\",\"maj:min\":\"8:19\",\"fsavail\":null,\"fssize\":null,\"fstype\":\"LVM2_member\",\"fsused\":null,\"fsuse%\":null,\"fsroots\":[null],\"fsver\":\"LVM2 001\",\"mountpoint\":null,\"mountpoints\":[null],\"label\":null,\"uuid\":\"qncuJa-Nl3y-bGGg-hKQz-BOku-MWES-nFO0Bv\",\"ptuuid\":\"335318d1-33e6-49c6-b148-be25e85548f1\",\"pttype\":\"gpt\",\"parttype\":\"e6d6d379-f507-44c2-a23c-238f2a3df928\",\"parttypename\":\"Linux LVM\",\"partlabel\":null,\"partuuid\":\"c2a0409d-8052-45cd-b90e-99fc7d088957\",\"partflags\":null,\"ra\":4096,\"ro\":false,\"rm\":false,\"hotplug\":false,\"model\":null,\"serial\":null,\"size\":\"444.5G\",\"state\":null,\"owner\":\"root\",\"group\":\"disk\",\"mode\":\"brw-rw----\",\"alignment\":0,\"min-io\":262144,\"opt-io\":262144,\"phy-sec\":512,\"log-sec\":512,\"rota\":false,\"sched\":\"mq-deadline\",\"rq-size\":256,\"type\":\"part\",\"disc-aln\":0,\"disc-gran\":\"0B\",\"disc-max\":\"0B\",\"disc-zero\":false,\"wsame\":\"0B\",\"wwn\":\"0x600062b2075f9b4028d5620a056d37ec\",\"rand\":false,\"pkname\":\"sdb\",\"hctl\":null,\"tran\":null,\"subsystems\":\"block:scsi:pci\",\"rev\":null,\"vendor\":null,\"zoned\":\"none\",\"dax\":false,\"children\":[{\"name\":\"ablestack_ablecube13--1-root\",\"kname\":\"dm-0\",\"path\":\"/dev/mapper/ablestack_ablecube13--1-root\",\"maj:min\":\"253:0\",\"fsavail\":\"322.3G\",\"fssize\":\"380.3G\",\"fstype\":\"xfs\",\"fsused\":\"58G\",\"fsuse%\":\"15%\",\"fsroots\":[\"/\"],\"fsver\":null,\"mountpoint\":\"/\",\"mountpoints\":[\"/\"],\"label\":null,\"uuid\":\"a3d08e53-5c13-43e5-97e1-af007bb4ab0b\",\"ptuuid\":null,\"pttype\":null,\"parttype\":null,\"parttypename\":null,\"partlabel\":null,\"partuuid\":null,\"partflags\":null,\"ra\":4096,\"ro\":false,\"rm\":false,\"hotplug\":false,\"model\":null,\"serial\":null,\"size\":\"380.5G\",\"state\":\"running\",\"owner\":\"root\",\"group\":\"disk\",\"mode\":\"brw-rw----\",\"alignment\":0,\"min-io\":262144,\"opt-io\":262144,\"phy-sec\":512,\"log-sec\":512,\"rota\":false,\"sched\":null,\"rq-size\":null,\"type\":\"lvm\",\"disc-aln\":0,\"disc-gran\":\"0B\",\"disc-max\":\"0B\",\"disc-zero\":false,\"wsame\":\"0B\",\"wwn\":null,\"rand\":false,\"pkname\":\"sdb3\",\"hctl\":null,\"tran\":null,\"subsystems\":\"block\",\"rev\":null,\"vendor\":null,\"zoned\":\"none\",\"dax\":false},{\"name\":\"ablestack_ablecube13--1-swap\",\"kname\":\"dm-1\",\"path\":\"/dev/mapper/ablestack_ablecube13--1-swap\",\"maj:min\":\"253:1\",\"fsavail\":null,\"fssize\":null,\"fstype\":\"swap\",\"fsused\":null,\"fsuse%\":null,\"fsroots\":[null],\"fsver\":\"1\",\"mountpoint\":\"[SWAP]\",\"mountpoints\":[\"[SWAP]\"],\"label\":null,\"uuid\":\"e2286fe3-2f15-4cdf-8b77-7e9fa5d3f0b3\",\"ptuuid\":null,\"pttype\":null,\"parttype\":null,\"parttypename\":null,\"partlabel\":null,\"partuuid\":null,\"partflags\":null,\"ra\":4096,\"ro\":false,\"rm\":false,\"hotplug\":false,\"model\":null,\"serial\":null,\"size\":\"64G\",\"state\":\"running\",\"owner\":\"root\",\"group\":\"disk\",\"mode\":\"brw-rw----\",\"alignment\":0,\"min-io\":262144,\"opt-io\":262144,\"phy-sec\":512,\"log-sec\":512,\"rota\":false,\"sched\":null,\"rq-size\":null,\"type\":\"lvm\",\"disc-aln\":0,\"disc-gran\":\"0B\",\"disc-max\":\"0B\",\"disc-zero\":false,\"wsame\":\"0B\",\"wwn\":null,\"rand\":false,\"pkname\":\"sdb3\",\"hctl\":null,\"tran\":null,\"subsystems\":\"block\",\"rev\":null,\"vendor\":null,\"zoned\":\"none\",\"dax\":false}]}]},{\"name\":\"rbd0\",\"kname\":\"rbd0\",\"path\":\"/dev/rbd0\",\"maj:min\":\"252:0\",\"fsavail\":null,\"fssize\":null,\"fstype\":null,\"fsused\":null,\"fsuse%\":null,\"fsroots\":[null],\"fsver\":null,\"mountpoint\":null,\"mountpoints\":[null],\"label\":null,\"uuid\":null,\"ptuuid\":null,\"pttype\":null,\"parttype\":null,\"parttypename\":null,\"partlabel\":null,\"partuuid\":null,\"partflags\":null,\"ra\":128,\"ro\":false,\"rm\":false,\"hotplug\":false,\"model\":null,\"serial\":null,\"size\":\"4.9G\",\"state\":null,\"owner\":\"root\",\"group\":\"disk\",\"mode\":\"brw-rw----\",\"alignment\":0,\"min-io\":65536,\"opt-io\":65536,\"phy-sec\":512,\"log-sec\":512,\"rota\":false,\"sched\":\"none\",\"rq-size\":128,\"type\":\"disk\",\"disc-aln\":0,\"disc-gran\":\"64K\",\"disc-max\":\"4M\",\"disc-zero\":false,\"wsame\":\"0B\",\"wwn\":null,\"rand\":false,\"pkname\":null,\"hctl\":null,\"tran\":null,\"subsystems\":\"block:rbd\",\"rev\":null,\"vendor\":null,\"zoned\":\"none\",\"dax\":false,\"children\":[{\"name\":\"rbd0p1\",\"kname\":\"rbd0p1\",\"path\":\"/dev/rbd0p1\",\"maj:min\":\"252:1\",\"fsavail\":null,\"fssize\":null,\"fstype\":null,\"fsused\":null,\"fsuse%\":null,\"fsroots\":[null],\"fsver\":null,\"mountpoint\":null,\"mountpoints\":[null],\"label\":null,\"uuid\":null,\"ptuuid\":null,\"pttype\":null,\"parttype\":null,\"parttypename\":null,\"partlabel\":null,\"partuuid\":null,\"partflags\":null,\"ra\":128,\"ro\":false,\"rm\":false,\"hotplug\":false,\"model\":null,\"serial\":null,\"size\":\"380M\",\"state\":null,\"owner\":\"root\",\"group\":\"disk\",\"mode\":\"brw-rw----\",\"alignment\":0,\"min-io\":65536,\"opt-io\":65536,\"phy-sec\":512,\"log-sec\":512,\"rota\":false,\"sched\":\"none\",\"rq-size\":128,\"type\":\"part\",\"disc-aln\":0,\"disc-gran\":\"64K\",\"disc-max\":\"4M\",\"disc-zero\":false,\"wsame\":\"0B\",\"wwn\":null,\"rand\":false,\"pkname\":\"rbd0\",\"hctl\":null,\"tran\":null,\"subsystems\":\"block:rbd\",\"rev\":null,\"vendor\":null,\"zoned\":\"none\",\"dax\":false},{\"name\":\"rbd0p2\",\"kname\":\"rbd0p2\",\"path\":\"/dev/rbd0p2\",\"maj:min\":\"252:2\",\"fsavail\":null,\"fssize\":null,\"fstype\":null,\"fsused\":null,\"fsuse%\":null,\"fsroots\":[null],\"fsver\":null,\"mountpoint\":null,\"mountpoints\":[null],\"label\":null,\"uuid\":null,\"ptuuid\":null,\"pttype\":null,\"parttype\":null,\"parttypename\":null,\"partlabel\":null,\"partuuid\":null,\"partflags\":null,\"ra\":128,\"ro\":false,\"rm\":false,\"hotplug\":false,\"model\":null,\"serial\":null,\"size\":\"1K\",\"state\":null,\"owner\":\"root\",\"group\":\"disk\",\"mode\":\"brw-rw----\",\"alignment\":1024,\"min-io\":65536,\"opt-io\":65536,\"phy-sec\":512,\"log-sec\":512,\"rota\":false,\"sched\":\"none\",\"rq-size\":128,\"type\":\"part\",\"disc-aln\":1024,\"disc-gran\":\"64K\",\"disc-max\":\"4M\",\"disc-zero\":false,\"wsame\":\"0B\",\"wwn\":null,\"rand\":false,\"pkname\":\"rbd0\",\"hctl\":null,\"tran\":null,\"subsystems\":\"block:rbd\",\"rev\":null,\"vendor\":null,\"zoned\":\"none\",\"dax\":false},{\"name\":\"rbd0p5\",\"kname\":\"rbd0p5\",\"path\":\"/dev/rbd0p5\",\"maj:min\":\"252:5\",\"fsavail\":null,\"fssize\":null,\"fstype\":null,\"fsused\":null,\"fsuse%\":null,\"fsroots\":[null],\"fsver\":null,\"mountpoint\":null,\"mountpoints\":[null],\"label\":null,\"uuid\":null,\"ptuuid\":null,\"pttype\":null,\"parttype\":null,\"parttypename\":null,\"partlabel\":null,\"partuuid\":null,\"partflags\":null,\"ra\":128,\"ro\":false,\"rm\":false,\"hotplug\":false,\"model\":null,\"serial\":null,\"size\":\"487M\",\"state\":null,\"owner\":\"root\",\"group\":\"disk\",\"mode\":\"brw-rw----\",\"alignment\":0,\"min-io\":65536,\"opt-io\":65536,\"phy-sec\":512,\"log-sec\":512,\"rota\":false,\"sched\":\"none\",\"rq-size\":128,\"type\":\"part\",\"disc-aln\":0,\"disc-gran\":\"64K\",\"disc-max\":\"4M\",\"disc-zero\":false,\"wsame\":\"0B\",\"wwn\":null,\"rand\":false,\"pkname\":\"rbd0\",\"hctl\":null,\"tran\":null,\"subsystems\":\"block:rbd\",\"rev\":null,\"vendor\":null,\"zoned\":\"none\",\"dax\":false},{\"name\":\"rbd0p6\",\"kname\":\"rbd0p6\",\"path\":\"/dev/rbd0p6\",\"maj:min\":\"252:6\",\"fsavail\":null,\"fssize\":null,\"fstype\":null,\"fsused\":null,\"fsuse%\":null,\"fsroots\":[null],\"fsver\":null,\"mountpoint\":null,\"mountpoints\":[null],\"label\":null,\"uuid\":null,\"ptuuid\":null,\"pttype\":null,\"parttype\":null,\"parttypename\":null,\"partlabel\":null,\"partuuid\":null,\"partflags\":null,\"ra\":128,\"ro\":false,\"rm\":false,\"hotplug\":false,\"model\":null,\"serial\":null,\"size\":\"4G\",\"state\":null,\"owner\":\"root\",\"group\":\"disk\",\"mode\":\"brw-rw----\",\"alignment\":0,\"min-io\":65536,\"opt-io\":65536,\"phy-sec\":512,\"log-sec\":512,\"rota\":false,\"sched\":\"none\",\"rq-size\":128,\"type\":\"part\",\"disc-aln\":0,\"disc-gran\":\"64K\",\"disc-max\":\"4M\",\"disc-zero\":false,\"wsame\":\"0B\",\"wwn\":null,\"rand\":false,\"pkname\":\"rbd0\",\"hctl\":null,\"tran\":null,\"subsystems\":\"block:rbd\",\"rev\":null,\"vendor\":null,\"zoned\":\"none\",\"dax\":false}]}]}\n")
	fmt.Println(stdout)
}
