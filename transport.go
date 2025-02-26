package pcore

import (
	"net"
	"time"
)

// TCPConn implements a TCP connection
type TCPConn struct {
	options   TransportOptions
	conn      net.Conn
	connected bool
}

// NewTCPConn creates a new TCP connection
func NewTCPConn(address string, timeout time.Duration) *TCPConn {
	options := DefaultTransportOptions()
	options.Address = address
	options.Timeout = timeout

	return &TCPConn{
		options: options,
	}
}

// Open establishes a TCP connection
func (t *TCPConn) Open() error {
	dialer := net.Dialer{
		Timeout: t.options.Timeout,
	}

	conn, err := dialer.Dial("tcp", t.options.Address)
	if err != nil {
		return NewConnError("connect", t.options.Address, err)
	}

	t.conn = conn
	t.connected = true

	return nil
}

// Close closes the connection
func (t *TCPConn) Close() error {
	if t.conn != nil {
		err := t.conn.Close()
		t.connected = false
		t.conn = nil

		if err != nil {
			return NewConnError("disconnect", t.options.Address, err)
		}
	}

	return nil
}

// IsOpen returns connection status
func (t *TCPConn) IsOpen() bool {
	return t.connected
}

// Send sends data over the TCP connection
func (t *TCPConn) Send(data []byte) ([]byte, error) {
	return t.SendWithTimeout(data, t.options.Timeout)
}

// SendWithTimeout sends data with a timeout
func (t *TCPConn) SendWithTimeout(data []byte, timeout time.Duration) ([]byte, error) {
	if !t.connected || t.conn == nil {
		return nil, ErrConnClosed
	}

	// Set write deadline
	err := t.conn.SetWriteDeadline(time.Now().Add(timeout))
	if err != nil {
		return nil, NewConnError("set_deadline", t.options.Address, err)
	}

	// Write data
	_, err = t.conn.Write(data)
	if err != nil {
		return nil, NewConnError("write", t.options.Address, err)
	}

	// Set read deadline
	err = t.conn.SetReadDeadline(time.Now().Add(timeout))
	if err != nil {
		return nil, NewConnError("settt_deadline", t.options.Address, err)
	}

	// Read response
	buffer := make([]byte, t.options.MaxFrameSize)
	n, err := t.conn.Read(buffer)
	if err != nil {
		return nil, NewConnError("read", t.options.Address, err)
	}

	return buffer[:n], nil
}

// TCPProvider implements TransportProvider for TCP
type TCPProvider struct{}

// CreateConn creates a new TCP connection
func (p *TCPProvider) CreateConn(options TransportOptions) (Conn, error) {
	return NewTCPConn(options.Address, options.Timeout), nil
}
