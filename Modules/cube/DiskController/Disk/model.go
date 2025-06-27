package Disk

type Base struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Path string `json:"path"`
	Type string `json:"type"`
	Size string `json:"size"`
	Used string `json:"used"`
	Free string `json:"free"`
}

func (disk *Base) GetDisk() *Base {
	return disk
}

func (disk *Base) GetType() string {
	return disk.Type
}

func (disk *Base) SetID(s string) {

}

func (disk *Base) GetID() string {
	return disk.ID
}

type List struct {
	Disks []*Base `json:"disks"`
}

type Block struct {
	*Base
}

type HBA struct {
	*Base
}

type iSCSI struct {
	*Base
}
