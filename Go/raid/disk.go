package main

type Disk struct {
	ID   int
	Data []byte
}

func NewDisk(id, capacity int) *Disk {
	return &Disk{
		ID:   id,
		Data: make([]byte, capacity),
	}
}

func (d *Disk) Clear() {
	for i := range d.Data {
		d.Data[i] = 0
	}
}
