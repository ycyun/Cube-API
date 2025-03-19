package task

import "log"

// 클러스터 상태 업데이트
func UpdateClusterStatus() {
	log.Println("Updating cluster status...")
	//err := cluster.CheckClusterHealth()
	//if err != nil {
	//	log.Println("Error updating cluster status:", err)
	//}
}

// 스토리지 상태 동기화
func SyncStorageStatus() {
	log.Println("Syncing storage status...")
	//err := glue.CheckStoragePools()
	//if err != nil {
	//	log.Println("Error syncing storage status:", err)
	//}
}
