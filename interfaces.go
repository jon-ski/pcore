package pcore

import (
	"encoding"
	"io"
	"time"
)

type Frame interface {
	encoding.BinaryMarshaler
	encoding.BinaryUnmarshaler

	// Size returns the size in bytes of the encoded frame. It calls
	// `MarshalBinary()` and if it returns an error, Size will return
	// -1.
	Size() int

	// Type returns the frame type/function identifier
	Type() int
}

type Protocol interface {
	// Encode converts a protocol frame to raw bytes
	Encode(frame Frame) ([]byte, error)

	// Decode converts raw bytes to a protocol frame
	Decode(data []byte) (Frame, error)

	// Validate validates a response against a request
	Validate(request, response Frame) error
}

// Conn represents a basic connection to a device or system
type Conn interface {
	io.ReadWriter

	// Open establishes the connection
	Open() error

	// Close terminates the connection
	Close() error

	// IsOpen returns connection status
	IsOpen() bool

	// Send sends data and returns a response
	Send(data []byte) ([]byte, error)

	// SendWithTimeout sends data with a timeout
	SendWithTimeout(data []byte, timeout time.Duration) ([]byte, error)
}

type StreamConn interface {
	Conn

	// Subscribe registers a callback for data reception
	Subscribe(handler func(data []byte) error)

	// Unsubscribe removes subscription
	Unsubscribe() error
}

// Client defines a base client interface
type Client interface {
	// Connect establishes connection to the target
	Connect() error

	// Disconnect closes the connection
	Disconnect() error

	// IsConnected returns connection status
	IsConnected() bool
}

// TransportOptions defines common transport configuration
type TransportOptions struct {
	// Address is the target address (e.g., IP:port, COM port)
	Address string

	// Timeout is the default I/O timeout
	Timeout time.Duration

	// RetryCount is the number of retries on failure
	RetryCount int

	// RetryDelay is the delay between retries
	RetryDelay time.Duration

	// MaxFrameSize is the maximum allowed frame size
	MaxFrameSize int
}

func DefaultTransportOptions() TransportOptions {
	return TransportOptions{
		Timeout:      time.Second * 5,
		RetryCount:   0,
		RetryDelay:   time.Millisecond * 100,
		MaxFrameSize: 1024,
	}
}

// TransportProvider creates transport-specific connections
type TransportProvider interface {
	// CreateConn creates a new connection
	CreateConn(options TransportOptions) (Conn, error)
}
