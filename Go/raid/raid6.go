package main

import (
	"fmt"

	"github.com/klauspost/reedsolomon"
)

type Raid6 struct {
	Disks        []*Disk
	Encoder      reedsolomon.Encoder
	NumDataDisks int
}

const NUM_PARITY_DISKS = 2

func NewRaid6(disks []*Disk) (Raid, error) {
	numDisks := len(disks)
	if numDisks < 4 {
		return nil, fmt.Errorf("RAID 6 requires at least 4 disks, got %d", numDisks)
	}

	numDataDisks := numDisks - NUM_PARITY_DISKS

	encoder, err := reedsolomon.New(numDataDisks, NUM_PARITY_DISKS)
	if err != nil {
		return nil, fmt.Errorf("error creating encoder: %s", err.Error())
	}

	raid := &Raid6{
		Disks:        disks,
		Encoder:      encoder,
		NumDataDisks: numDataDisks,
	}

	return raid, nil
}

func (r *Raid6) WriteString(data string) error {
	for i := range data {
		blockIndex := r.computeBlockIndex(i)
		diskIndex := r.computeRWDisk(i)
		r.Disks[diskIndex].Data[blockIndex] = data[i]
	}

	return r.encodeParity()
}

func (r *Raid6) ClearDisk(diskIndex int) {
	r.Disks[diskIndex].Clear()

	// suppose the data loss has been detected
	fmt.Printf("detect disk %d data loss, reconstruct the data\n", diskIndex)
	r.Disks[diskIndex].Data = nil
	r.reconstruct(diskIndex)
}

func (r *Raid6) ReadString(length int) (string, error) {
	result := make([]byte, length)

	for i := 0; i < length; i++ {
		blockIndex := r.computeBlockIndex(i)
		diskIndex := r.computeRWDisk(i)
		result[i] = r.Disks[diskIndex].Data[blockIndex]
	}

	return string(result), nil
}

func (r *Raid6) computeBlockIndex(dataIndex int) int {
	return dataIndex / r.NumDataDisks
}

func (r *Raid6) computeRWDisk(dataIndex int) int {
	return dataIndex % r.NumDataDisks
}

func (r *Raid6) encodeParity() error {
	data := make([][]byte, r.NumDataDisks+NUM_PARITY_DISKS)
	for i := range data {
		data[i] = r.Disks[i].Data
	}
	err := r.Encoder.Encode(data)
	return err
}

func (r *Raid6) reconstruct(diskIndex int) {
	data := make([][]byte, r.NumDataDisks+NUM_PARITY_DISKS)
	for i := range data {
		data[i] = r.Disks[i].Data
	}

	r.Encoder.Reconstruct(data)
	r.Disks[diskIndex].Data = data[diskIndex]
}
