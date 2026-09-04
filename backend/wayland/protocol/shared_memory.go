package protocol

import (
	"fmt"

	"golang.org/x/sys/unix"
)

type SharedMemory struct {
	Pixels []byte
	Fd     int
	Size   int
}

func NewSharedMemory(size int) (*SharedMemory, error) {
	fd, err := createFd(size)
	if err != nil {
		return nil, err
	}

	data, err := mmap(fd, size)
	if err != nil {
		unix.Close(fd)
		return nil, err
	}

	return &SharedMemory{
		Fd:     fd,
		Pixels: data,
		Size:   size,
	}, nil
}

func (bm *SharedMemory) Close() error {
	var firstErr error

	if bm.Pixels != nil {
		if err := unix.Munmap(bm.Pixels); err != nil {
			firstErr = err
		}

		bm.Pixels = nil
	}

	if bm.Fd >= 0 {
		if err := unix.Close(bm.Fd); err != nil && firstErr == nil {
			firstErr = err
		}

		bm.Fd = -1
	}

	return firstErr
}

func createFd(size int) (int, error) {
	if size <= 0 {
		return -1, fmt.Errorf("invalid shm size: %d", size)
	}

	fd, err := unix.MemfdCreate("wayland-shm", unix.MFD_CLOEXEC)
	if err != nil {
		return -1, fmt.Errorf("memfd_create: %w", err)
	}

	return fd, nil
}

func mmap(fd, size int) ([]byte, error) {
	if err := unix.Ftruncate(fd, int64(size)); err != nil {
		return nil, fmt.Errorf("ftruncate: %w", err)
	}

	data, err := unix.Mmap(
		fd,
		0,
		size,
		unix.PROT_READ|unix.PROT_WRITE,
		unix.MAP_SHARED,
	)

	if err != nil {
		return nil, fmt.Errorf("mmap: %w", err)
	}

	return data, nil
}
