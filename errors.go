package pcore

import (
	"errors"
	"fmt"
)

// Common errors that may occur in protocol communications
var (
	ErrConnClosed      = errors.New("connection is closed")
	ErrConnTimeout     = errors.New("connection timed out")
	ErrInvalidMessage  = errors.New("invalid message format")
	ErrInvalidResponse = errors.New("invalid response")
	ErrProtocolError   = errors.New("protocol error")
	ErrFrameTooLarge   = errors.New("frame exceeds maximum allowed size")
	ErrFrameTooSmall   = errors.New("frame is too small to be valid")
	ErrInvalidChecksum = errors.New("invalid checksum")
	ErrDeviceError     = errors.New("device reported an error")
	ErrBufferTooSmall  = errors.New("buffer is too small")
)

// ProtocolError represents a protocol-specific error
type ProtocolError struct {
	Code    int
	Message string
	Inner   error
}

func (e *ProtocolError) Error() string {
	if e.Inner != nil {
		return fmt.Sprintf("protocol error %d: %s - %v", e.Code, e.Message, e.Inner)
	}
	return fmt.Sprintf("protocol error %d: %s", e.Code, e.Message)
}

// Unwrap returns the unwrapped error
func (e *ProtocolError) Unwrap() error {
	return e.Inner
}

// NewProtocolError creates a new protocol error
func NewProtocolError(code int, message string, inner error) *ProtocolError {
	return &ProtocolError{
		Code:    code,
		Message: message,
		Inner:   inner,
	}
}

// ConnError represents a connection-related error
type ConnError struct {
	Op    string
	Addr  string
	Inner error
}

// Error returns the error message
func (e *ConnError) Error() string {
	if e.Inner != nil {
		return fmt.Sprintf("connection error during %s to %s: %s", e.Op, e.Addr, e.Inner)
	}
	return fmt.Sprintf("connection error during %s to %s", e.Op, e.Addr)
}

// Unwrap returns the unwrapped error
func (e *ConnError) Unwrap() error {
	return e.Inner
}

// NewConnError creates a new connection error
func NewConnError(op, addr string, inner error) *ConnError {
	return &ConnError{
		Op:    op,
		Addr:  addr,
		Inner: inner,
	}
}
