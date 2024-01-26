package main

import (
	"fmt"
	"os"
	"strconv"
)

const (
	NUM_DISKS  int    = 4
	CAPACITY   int    = 50
	CLEAR_DISK int    = 1
	TEXT       string = "Hello World!"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run . <raid number>")
		return
	}
	raidNumber, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Invalid number of raid:", err)
		return
	}

	disks := make([]*Disk, NUM_DISKS)
	for i := 0; i < NUM_DISKS; i++ {
		disks[i] = NewDisk(i, CAPACITY)
	}

	var raid Raid
	switch raidNumber {
	case 0:
		raid, err = NewRaid0(disks)
	case 1:
		raid, err = NewRaid1(disks)
	case 5:
		raid, err = NewRaid5(disks)
	case 6:
		raid, err = NewRaid6(disks)
	case 10:
		raid, err = NewRaid10(disks)
	default:
		panic("Invalid raid number")
	}
	if err != nil {
		panic(err.Error())
	}

	raid.WriteString(TEXT)
	readDataBeforeClear, err := raid.ReadString(len(TEXT))
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("Before clear disk %d, read data: %s\n", CLEAR_DISK, readDataBeforeClear)
	raid.ClearDisk(CLEAR_DISK)

	readDataAfterClear, err := raid.ReadString(len(TEXT))
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("After clear disk %d, read data: %s\n", CLEAR_DISK, readDataAfterClear)
}
