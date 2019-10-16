package cache

import "io"

// StreamClient : todo
type StreamClient interface {
	CasWithStream(string, io.Reader, int, uint32, int32) error

	GetValueWriter(string, string, uint32, int32, int, uint64) (io.WriteCloser, error)

	GetWithStream(string, func(io.Reader) error) error

	Touch(string, int32) error
}
