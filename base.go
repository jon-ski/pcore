package pcore

import (
	"encoding/binary"
)

// BaseFrame implements the basic Frame interface
type BaseFrame struct {
	data      []byte
	frameType uint8
}

// NewBaseFrame creates a new base frame
func NewBaseFrame(data []byte, frameType uint8) *BaseFrame {
	return &BaseFrame{
		data:      data,
		frameType: frameType,
	}
}

// MarshalBinary returns the data portion of `BaseFrame` and nil.
func (f *BaseFrame) MarshalBinary() (data []byte, err error) {
	return f.data, nil
}

// Size returns the size in bytes of data
func (f *BaseFrame) Size() int {
	return len(f.data)
}

// Type returns the frame type
func (f *BaseFrame) Type() uint8 {
	return f.frameType
}

// Utility functions for common encoding/decoding operations

// ReadUint8 reads a uint8 from the buffer at the given offset
func ReadUint8(data []byte, offset int) (uint8, error) {
	if offset >= len(data) {
		return 0, ErrBufferTooSmall
	}
	return data[offset], nil
}

// ReadUint16 reads a uint16 from the buffer at the given offset (big endian)
func ReadUint16(data []byte) (uint16, error) {
	if len(data) < 2 {
		return 0, ErrBufferTooSmall
	}
	return binary.BigEndian.Uint16(data), nil
}

// ReadUint32 reads a uint32 from the buffer at the given offset (big endian)
func ReadUint32(data []byte) (uint32, error) {
	if len(data) < 4 {
		return 0, ErrBufferTooSmall
	}
	return binary.BigEndian.Uint32(data), nil
}

func ReadUint64(data []byte) (uint64, error) {
	if len(data) < 8 {
		return 0, ErrBufferTooSmall
	}
	return binary.BigEndian.Uint64(data), nil
}
