package main

import "fmt"

type Raid1 struct {
	Disks []*Disk
}

func NewRaid1(disks []*Disk) (Raid, error) {
	numDisks := len(disks)
	if numDisks < 2 {
		return nil, fmt.Errorf("RAID 1 requires at least 2 disks, got %d", numDisks)
	}

	raid := &Raid1{
		Disks: disks,
	}
	return raid, nil
}

func (r *Raid1) WriteString(data string) error {
	for i, char := range data {
		for _, disk := range r.Disks {
			disk.Data[i] = byte(char)
		}
	}
	return nil
}

func (r *Raid1) ClearDisk(diskIndex int) {
	for i := range r.Disks[diskIndex].Data {
		r.Disks[diskIndex].Data[i] = 0
	}
}

func (r *Raid1) ReadString(length int) (string, error) {
	result := make([]byte, length)

	for i := 0; i < length; i++ {
		result[i] = r.Disks[0].Data[i]
	}
	return string(result), nil
}
