package whatever

import (
	"fmt"
	"swing-go/backend/wayland/proxies"

	"golang.org/x/sys/unix"
)

type BufferManager struct {
	Data []byte
	Fd   int
	Size int

	wlShmPool *proxies.WlShmPool
}

func NewBufferManager(wlShmPool *proxies.WlShmPool, size int) (*BufferManager, error) {
	fd, err := createFd(size)
	if err != nil {
		return nil, err
	}

	data, err := mmap(fd, size)
	if err != nil {
		_ = unix.Close(fd)
		return nil, err
	}

	return &BufferManager{
		Data:      data,
		Fd:        fd,
		Size:      size,
		wlShmPool: wlShmPool,
	}, nil
}

func (bm *BufferManager) Close() error {
	var firstErr error

	if bm.Data != nil {
		if err := unix.Munmap(bm.Data); err != nil {
			firstErr = err
		}

		bm.Data = nil
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
