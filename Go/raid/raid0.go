package main

import "fmt"

type Raid0 struct {
	Disks []*Disk
}

func NewRaid0(disks []*Disk) (Raid, error) {
	numDisks := len(disks)
	if numDisks < 2 {
		return nil, fmt.Errorf("RAID 0 requires at least 2 disks, got %d", numDisks)
	}

	raid := &Raid0{
		Disks: disks,
	}
	return raid, nil
}

func (r *Raid0) WriteString(data string) error {
	for i, char := range data {
		diskIndex := i % len(r.Disks)
		blockIndex := i / len(r.Disks)
		r.Disks[diskIndex].Data[blockIndex] = byte(char)
	}
	return nil
}

func (r *Raid0) ClearDisk(diskIndex int) {
	for i := range r.Disks[diskIndex].Data {
		r.Disks[diskIndex].Data[i] = 0
	}
}

func (r *Raid0) ReadString(length int) (string, error) {
	result := make([]byte, length)

	for i := 0; i < length; i++ {
		diskIndex := i % len(r.Disks)
		blockIndex := i / len(r.Disks)
		result[i] = r.Disks[diskIndex].Data[blockIndex]
	}

	return string(result), nil
}
