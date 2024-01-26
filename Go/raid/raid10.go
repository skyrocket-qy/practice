package main

import "fmt"

type Raid10 struct {
	Disks  []*Disk
	Raid1s [][]*Disk
}

// TODO: refactor to directly use raid0 and raid1 implementations
func NewRaid10(disks []*Disk) (Raid, error) {
	numDisks := len(disks)
	if numDisks < 4 {
		return nil, fmt.Errorf("RAID 10 requires at least 4 disks, got %d", numDisks)
	} else if numDisks&1 == 1 {
		return nil, fmt.Errorf("RAID 10 requires even disks, got %d", numDisks)
	}

	pairs := numDisks / 2
	raid1Pairs := make([][]*Disk, pairs)
	for i := range raid1Pairs {
		raid1Pairs[i] = []*Disk{disks[2*i], disks[2*i+1]}
	}

	raid := &Raid10{
		Disks:  disks,
		Raid1s: raid1Pairs,
	}
	return raid, nil
}

func (r *Raid10) WriteString(data string) error {
	for i := range data {
		raid1Index := i % len(r.Raid1s)
		blockIndex := i / len(r.Raid1s)

		for _, disk := range r.Raid1s[raid1Index] {
			disk.Data[blockIndex] = data[i]
		}
	}
	return nil
}

func (r *Raid10) ClearDisk(diskIndex int) {
	r.Disks[diskIndex].Clear()
}

func (r *Raid10) ReadString(length int) (string, error) {
	result := make([]byte, length)

	for i := 0; i < length; i++ {
		raid1Index := i % len(r.Raid1s)
		blockIndex := i / len(r.Raid1s)
		result[i] = r.Raid1s[raid1Index][0].Data[blockIndex]
	}

	return string(result), nil
}
