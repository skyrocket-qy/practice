package main

type Raid interface {
	WriteString(data string) error
	ClearDisk(diskIndex int)
	ReadString(length int) (string, error)
}
