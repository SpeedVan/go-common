// +build linux

package file

import (
	"os"
	"syscall"
)

// Lock todo
type Lock struct {
	dir string
	fd  int
	// +build linux
	ft *syscall.Flock_t
}

// New todo
func New(dir string) (*Lock, error) {
	fd, err := syscall.Open(dir, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0777)
	if err != nil {
		return nil, err
	}
	return &Lock{
		dir: dir,
		fd:  fd,
		ft: &syscall.Flock_t{
			Type: syscall.F_WRLCK,
		},
	}, nil
}

// Lock todo
func (s *Lock) Lock() error {
	return syscall.FcntlFlock(uintptr(s.fd), syscall.F_SETLK, s.ft)
}

// Unlock todo
func (s *Lock) Unlock() error {

	err := syscall.FcntlFlock(uintptr(s.fd), syscall.F_UNLCK, s.ft)
	if err != nil {
		return err
	}
	syscall.Close(s.fd)
	return syscall.Unlink(s.dir)
	// return syscall.Flock(int(s.fd), syscall.LOCK_UN)
}
