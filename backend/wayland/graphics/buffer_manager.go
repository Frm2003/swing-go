package graphics

import (
	"fmt"
	"swing-go/backend/wayland/proxies"

	"golang.org/x/sys/unix"
)

type AllocatorPool func(int, int) (*proxies.WlShmPool, error)

type AllocatorBuffer func(*Driver) (*Buffer, error)

type BufferManager struct {
	allocBuf AllocatorBuffer
	Buffers  []*Buffer

	Pixels []byte
	Fd     int

	Size   int
	Stride int

	WlShmPool *proxies.WlShmPool
}

func NewBufferManager(pool AllocatorPool, allocBuf AllocatorBuffer, width, height int) (*BufferManager, error) {
	stride := width * 4
	size := stride * height

	fd, err := createFd(size)

	if err != nil {
		return nil, err
	}

	data, err := mmap(fd, size)

	if err != nil {
		defer func() {
			_ = unix.Close(fd)
			_ = unix.Munmap(data)
		}()
		return nil, err
	}

	wlShmPool, err := pool(fd, size)

	if err != nil {
		return nil, err
	}

	return &BufferManager{
		allocBuf: allocBuf,
		Buffers:  make([]*Buffer, 0),

		Fd:     fd,
		Pixels: data,

		Size:   size,
		Stride: stride,

		WlShmPool: wlShmPool,
	}, nil
}

func (bm *BufferManager) CreateBuffer(d *Driver) (*Buffer, error) {
	buf, err := bm.allocBuf(d)

	if err != nil {
		return nil, err
	}

	bm.Buffers = append(bm.Buffers, buf)

	return buf, nil
}

func (bm *BufferManager) ClearBlack(buffer *Buffer) {
	start := buffer.Offset
	end := start + bm.Size

	for i := start; i < end; i++ {
		bm.Pixels[i] = 0
	}
}

func (bm *BufferManager) Close() error {
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
