package BLK

func (disk *BLK) GetID() string {
	//TODO implement me
	panic("implement me")
}

func (disk *BLK) SetID(s string) bool {
	//TODO implement me
	panic("implement me")
}

func (disk *BLK) GetDisk() *BLK {
	return disk
}

func (disk *BLK) GetType() string {
	return disk.Type
}

func (disk *BLK) Update() {
	disk.UpdateBLK()
}

func (disk *BLK) UpdateBLK() {

	/* 기존 목록에 신규정보 업데이트
	// 1. 목록 업데이트
	tmplist=[]
	poollist = ceph osd pool ls?
	for pool in poollist
	  tmplist.append(rbd -p pool ls -l)
	for disk in tmplist
	  if disk not in RBDs
	     RBDs.append(disk)
	for disk in RBDs
	  disk.update()

	// 2. 목록 교체
	tmplist=[]
	poollist = ceph osd pool ls?
	for pool in poollist
	  tmplist.append(rbd -p pool ls -l)
	for disk in tmplist
	  disk.update()
	RBDs = tmplist
	*/
}
