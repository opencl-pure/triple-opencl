package high

import (
	"errors"
	"opencl-pure/opencl/constants"
	"opencl-pure/opencl/pure"
	"unsafe"
)

// Vector is a memory buffer on device that holds []float32
type Vector[T any] struct {
	buf        *buffer
	iSize, len int
}

// Length the length of the vector
func (v *Vector[T]) Length() int {
	return v.len
}

// Release releases the buffer on the device
func (v *Vector[T]) Release() error {
	return v.buf.Release()
}

// NewVector allocates new vector buffer with specified length
func NewVector[T pure.BufferType](d *Device, data []T) (*Vector[T], error) {
	iSize := int(unsafe.Sizeof(&data[0]))
	l := len(data)
	size := l * iSize
	buf, err := newBuffer(d, size)
	if err != nil {
		return nil, err
	}
	return &Vector[T]{buf: buf, iSize: iSize, len: l}, nil
}

// Copy copies the float32 data from host data to device buffer
// it's a non-blocking call, channel will return an error or nil if the data transfer is complete
func (v *Vector[T]) Copy(data []T) <-chan error {
	if v.Length() != len(data) {
		ch := make(chan error, 1)
		ch <- errors.New("vector length not equal to data length")
		return ch
	}
	return v.buf.copy(len(data)*v.iSize, unsafe.Pointer(&data[0]))
}

// Data gets float32 data from device, it's a blocking call
func (v *Vector[T]) Data() ([]T, error) {
	data := make([]T, int(v.buf.size)/v.iSize)
	err := pure.StatusToErr(pure.EnqueueReadBuffer(
		v.buf.device.queue,
		v.buf.memobj,
		constants.CL_TRUE,
		0,
		v.buf.size,
		unsafe.Pointer(&data[0]),
		0,
		nil,
		nil,
	))
	if err != nil {
		return nil, err
	}
	return data, nil
}

// Map applies an map kernel on all elements of the vector
func (v *Vector[T]) Map(k *Kernel, returnEvent bool, waitEvents []*Event) (*Event, error) {
	return k.Global(v.Length()).Local(1).Run(returnEvent, waitEvents, v)
}
