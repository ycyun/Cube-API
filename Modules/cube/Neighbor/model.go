package Neighbor

type TypeNeighbor struct {
	IP       string `json:"ip"`
	HostName string `json:"hostname"`
}

type TypeNeighbors struct {
	Neighbors []TypeNeighbor `json:"neighbors"`
}
