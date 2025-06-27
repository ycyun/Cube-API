package Disk

type Interface interface {
	GetID() string
	SetID(string) bool
	GetType() string
	Initialize() error
}
