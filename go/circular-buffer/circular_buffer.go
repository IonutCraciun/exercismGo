package circular

import (
	"errors"
)
// Buffer type
type Buffer struct {
	data []byte
	used int
	readIndex int
	writeIndex int
}

var (
	// ErrorEmpty t
	ErrorEmpty = errors.New("buffer is empty")
	// ErrorFull t
	ErrorFull  = errors.New("buffer is full")
)

// NewBuffer func
func NewBuffer(size int) *Buffer {
	return &Buffer{data: make([]byte, size), used: 0, readIndex: 0, writeIndex: 0}
}

// ReadByte func
func (b *Buffer) ReadByte() (byte, error) {
	if b.used == 0 {
		//panic("Its empty")
		return 0, ErrorEmpty
	}
	defer func() {
		if b.readIndex == len(b.data)-1 {
			b.readIndex = 0
		} else {
			b.readIndex++
		}
		b.used--
	}()
	return b.data[b.readIndex], nil
}

// WriteByte func
func (b *Buffer) WriteByte(c byte) error {
	if b.isFull() {
		// panic("ITS FULL")
		return ErrorFull
	}
	b.data[b.writeIndex] = c
	if b.writeIndex == len(b.data)-1 {
		b.writeIndex = 0
	} else {
		b.writeIndex++
	}
	b.used++
	return nil
}

// isfull func
func (b *Buffer) isFull() bool {
	if len(b.data) == b.used {
		return true
	}
	return false
}
// Overwrite func
func (b *Buffer) Overwrite(c byte) {
	b.data[b.writeIndex] = c
	if b.writeIndex == len(b.data)-1 {
		b.writeIndex = 0
		b.readIndex = 0
	} else {
		b.writeIndex++
		b.readIndex++
	}
	if !b.isFull() {
		b.used++
	}
}

// Reset func
func (b *Buffer) Reset() {
	b.data = make([]byte, b.used)
	b.used = 0
	b.readIndex = 0
	b.writeIndex = 0
}