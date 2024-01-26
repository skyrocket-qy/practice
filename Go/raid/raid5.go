package main

import (
	"fmt"
	"math"
)

const RAID5_INIT_PARITY = byte(0)

type Raid5 struct {
	Disks     []*Disk
	BlockSize int
}

func NewRaid5(disks []*Disk) (Raid, error) {
	numDisks := len(disks)
	if numDisks < 3 {
		return nil, fmt.Errorf("RAID 5 requires at least 3 disks, got %d", numDisks)
	}

	raid := &Raid5{
		Disks:     disks,
		BlockSize: numDisks,
	}
	// init the parity
	maxBlock := math.Ceil(float64(len(disks[0].Data)) / float64(raid.BlockSize))
	for i := 0; i < int(maxBlock); i++ {
		parityDisk := raid.computeParityDisk(i)
		raid.Disks[parityDisk].Data[i] = RAID5_INIT_PARITY
	}

	return raid, nil
}

func (r *Raid5) WriteString(data string) error {
	for i := range data {
		blockIndex := r.computeBlockIndex(i)
		parityDisk := r.computeParityDisk(blockIndex)
		writeDisk := r.computeRWDisk(parityDisk, i)

		r.Disks[parityDisk].Data[blockIndex] ^= r.Disks[writeDisk].Data[blockIndex] ^ data[i]
		r.Disks[writeDisk].Data[blockIndex] = data[i]
	}
	return nil
}

func (r *Raid5) ClearDisk(diskIndex int) {
	r.Disks[diskIndex].Clear()
}

func (r *Raid5) ReadString(length int) (string, error) {
	result := make([]byte, length)

	for i := 0; i < length; i++ {
		blockIndex := r.computeBlockIndex(i)
		parityDisk := r.computeParityDisk(blockIndex)

		parity := RAID5_INIT_PARITY
		var checkParity byte
		for j, disk := range r.Disks {
			if j == parityDisk {
				checkParity = disk.Data[blockIndex]
				continue
			}
			parity ^= disk.Data[blockIndex]
		}
		if parity != checkParity {
			return "", fmt.Errorf("parity check failed. data may be corrupted")
		}

		readDisk := r.computeRWDisk(parityDisk, i)
		result[i] = r.Disks[readDisk].Data[blockIndex]
	}

	return string(result), nil
}

func (r *Raid5) computeBlockIndex(dataIndex int) int {
	return dataIndex / (r.BlockSize - 1)
}

func (r *Raid5) computeParityDisk(blockIndex int) int {
	return blockIndex % r.BlockSize
}

func (r *Raid5) computeRWDisk(parityDisk, dataIndex int) int {
	writeDisk := dataIndex % (r.BlockSize - 1)
	if writeDisk >= parityDisk {
		writeDisk = (writeDisk + 1) % r.BlockSize
	}
	return writeDisk
}
