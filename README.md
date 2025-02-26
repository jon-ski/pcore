# pcore

pcore is a lightweight, flexible framework for building industrial protocol clients and servers. It provides a minimalist set of interfaces and utilities to simplify the development of industrial communication protocols like Modbus, S7, OPC UA, and more.

## Features

- Protocol-agnostic design with clear separation of concerns
- Simple, composable interfaces for different communication patterns
- Transport-agnostic design supporting TCP, UDP, serial, and more
- Extensible error handling with protocol-specific errors
- Context support for timeout and cancellation
- Clean, idiomatic Go API

## Installation

```bash
go get github.com/jon-ski/pcore
```

## Core Concepts

pcore is built around a few key abstractions:

- **Frame**: The basic unit of data in a protocol
- **Protocol**: Handles encoding/decoding of frames
- **Conn**: Manages the connection to a device
- **Client**: High-level API for protocol-specific operations

The design is intentionally minimal to provide flexibility while avoiding unnecessary abstraction.

## Extending pcore

The framework is designed to be easily extended:

1. Create your own Frame implementation for your protocol
2. Implement the Protocol interface for encoding/decoding
3. Create a client that uses these components

## Adding a New Protocol

Here's how you might implement a simple S7 protocol client:

```go
package s7

import (
    "github.com/jon-ski/pcore"
)

// S7Frame implements pcore.Frame for S7 protocol
type S7Frame struct {
    pcore.BaseFrame
    // S7-specific fields
}

// S7Protocol implements pcore.Protocol for S7
type S7Protocol struct {
    // Configuration fields
}

// Implement Encode, Decode, and Validate methods

// S7Client provides S7-specific functionality
type S7Client struct {
    conn     pcore.Conn
    protocol *S7Protocol
}

// Implement S7-specific methods like ReadDB, WriteDB, etc.
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the LICENSE file for details.
