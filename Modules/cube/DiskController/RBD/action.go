package RBD

import (
	"github.com/goccy/go-json"
	"github.com/ycyun/Cube-API/Modules/cube/DiskController/Disk"
	"github.com/ycyun/Cube-API/utils"
	"log/slog"
)

func (disk *RBD) GetID() string {
	//TODO implement me
	return disk.ID
}

func (disk *RBD) SetID(s string) bool {
	//TODO implement me
	disk.ID = s
	return true
}

func (disk *RBD) GetDisk() Disk.Base {
	return disk.Base
}

func (disk *RBD) GetType() string {
	return disk.Type
}

func (disk *RBD) Initialize() error {
	return nil
}

func (disk *RBD) Update() {
	disk.UpdateRBD()
}

func (disk *RBD) UpdateRBD() {
	Logger := slog.New(utils.AbleTextHandler)

	//Todo: get image info
	// rbd info ccvm -p rbd --format json
	// {"name":"ccvm","id":"3e9b2359a76e","size":536870912000,"objects":128000,"order":22,"object_size":4194304,"snapshot_count":10,"block_name_prefix":"rbd_data.3e9b2359a76e","format":2,"features":["layering","exclusive-lock","object-map","fast-diff","deep-flatten"],"op_features":[],"flags":[],"create_timestamp":"Mon Jun 17 13:16:57 2024","access_timestamp":"Thu Jun 26 16:03:40 2025","modify_timestamp":"Thu Jun 26 16:02:51 2025"}
	ImageInfoCommand := []string{
		"rbd", "info", disk.Image, "-p", disk.Pool, "--format", "json",
	}
	ImageInfoReturn, err := utils.Execute(ImageInfoCommand)
	utils.HandleError(err)

	err = json.Unmarshal(ImageInfoReturn, &disk)
	utils.HandleError(err)
	disk.Path = disk.Pool + "/" + disk.Image

	//Todo: get image usage
	// rbd Usage ccvm -p rbd --format json
	// {"images":[{"name":"ccvm","id":"3e9b2359a76e","snapshot":"2025-06-17-01:00:01","snapshot_id":80704,"provisioned_size":536870912000,"used_size":476791701504},{"name":"ccvm","id":"3e9b2359a76e","snapshot":"2025-06-18-01:00:01","snapshot_id":80706,"provisioned_size":536870912000,"used_size":4513071104},{"name":"ccvm","id":"3e9b2359a76e","snapshot":"2025-06-19-01:00:01","snapshot_id":80708,"provisioned_size":536870912000,"used_size":6798966784},{"name":"ccvm","id":"3e9b2359a76e","snapshot":"2025-06-20-01:00:02","snapshot_id":80710,"provisioned_size":536870912000,"used_size":4810866688},{"name":"ccvm","id":"3e9b2359a76e","snapshot":"2025-06-21-01:00:01","snapshot_id":80712,"provisioned_size":536870912000,"used_size":5033164800},{"name":"ccvm","id":"3e9b2359a76e","snapshot":"2025-06-22-01:00:02","snapshot_id":80714,"provisioned_size":536870912000,"used_size":4664066048},{"name":"ccvm","id":"3e9b2359a76e","snapshot":"2025-06-23-01:00:01","snapshot_id":80716,"provisioned_size":536870912000,"used_size":5347737600},{"name":"ccvm","id":"3e9b2359a76e","snapshot":"2025-06-24-01:00:01","snapshot_id":80719,"provisioned_size":536870912000,"used_size":5091885056},{"name":"ccvm","id":"3e9b2359a76e","snapshot":"2025-06-25-01:00:01","snapshot_id":80722,"provisioned_size":536870912000,"used_size":6794772480},{"name":"ccvm","id":"3e9b2359a76e","snapshot":"2025-06-26-01:00:01","snapshot_id":80724,"provisioned_size":536870912000,"used_size":9755951104},{"name":"ccvm","id":"3e9b2359a76e","provisioned_size":536870912000,"used_size":6291456000}]}
	// {"images":[{"name":"0062b883-7598-4b73-a217-5e982cdc5ca0","id":"ab9f5e31463026","provisioned_size":5242880000,"used_size":2428502016}]}

	ImageUsageCommand := []string{
		"rbd", "du", disk.Image, "-p", disk.Pool, "--format", "json",
	}

	ImageUsage := make(map[string]interface{})
	ImageUsageReturn, err := utils.Execute(ImageUsageCommand)
	utils.HandleError(err)

	err = json.Unmarshal(ImageUsageReturn, &ImageUsage)
	if err != nil {
		Logger.Info("UsageReturn", "image", ImageUsageReturn)
		Logger.Info("UsageReturn", "image", string(ImageUsageReturn))
		Logger.Info("ImageUsage", "image", ImageUsage)
		utils.HandleError(err)
	}
	ImageUsageList := ImageUsage["images"].([]interface{})
	for _, _image := range ImageUsageList {
		image := _image.(map[string]interface{})

		//if image["name"] == "ccvm" {
		if disk.Type != "SnapshotImage" {
			if image["snapshot"] == nil {
				disk.UsedSize = image["used_size"].(float64)
				disk.ProvisionedSize = image["provisioned_size"].(float64)
				Logger.Debug("In ", "image", image["name"])
			}
		} else if disk.Type == "SnapshotImage" {
			if image["snapshot_id"] == disk.SnapshotId {
				disk.UsedSize = image["used_size"].(float64)
				disk.ProvisionedSize = image["provisioned_size"].(float64)
				Logger.Debug("In ", "image", image["name"])
			}
		}
		//Logger.Info("Image Usage "+image["name"].(string), "provisioned_size", image["provisioned_size"], "used_size", image["used_size"])
		//}
	}

}
