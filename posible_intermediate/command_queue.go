package posible_intermediate

// #include "opencl.h"
import "C"
import (
	"errors"
	"github.com/opencl-pure/triple-opencl/constants"
	"github.com/opencl-pure/triple-opencl/pure"
	"unsafe"
)

type CommandQueue struct {
	commandQueue pure.CommandQueue
}

func createCommandQueue(context *Context, device *Device) (*CommandQueue, error) {
	var errInt pure.Status
	if pure.CreateCommandQueue == nil {
		return nil, pure.Uninitialized("CreateCommandQueue")
	}
	queue := pure.CreateCommandQueue(context.context, device.deviceID, 0, &errInt)
	if errInt != constants.CL_SUCCESS {
		return nil, pure.StatusToErr(errInt)
	}
	return &CommandQueue{queue}, nil
}

func (c *CommandQueue) EnqueueNDRangeKernel(kernel *Kernel, workDim uint32, globalWorkSize []uint64) error {
	gws := make([]pure.Size, 0, len(globalWorkSize))
	for _, u := range globalWorkSize {
		gws = append(gws, pure.Size(u))
	}
	if pure.EnqueueNDRangeKernel == nil {
		return pure.Uninitialized("EnqueueNDRangeKernel")
	}
	errInt := pure.EnqueueNDRangeKernel(c.commandQueue,
		kernel.kernel,
		uint(workDim),
		nil,
		gws,
		nil, 0, nil, nil)
	return pure.StatusToErr(errInt)
}

func (c CommandQueue) EnqueueReadBuffer(buffer Buffer, blockingRead bool, dataPtr interface{}) error {
	var br C.cl_bool
	if blockingRead {
		br = C.CL_TRUE
	} else {
		br = C.CL_FALSE
	}

	var ptr unsafe.Pointer
	var dataLen uint64
	switch p := dataPtr.(type) {
	case []float32:
		dataLen = uint64(len(p) * 4)
		ptr = unsafe.Pointer(&p[0])
	default:
		return errors.New("Unexpected type for dataPtr")
	}

	errInt := clError(C.clEnqueueReadBuffer(c.commandQueue,
		buffer.buffer,
		br,
		0,
		C.size_t(dataLen),
		ptr,
		0, nil, nil))
	return clErrorToError(errInt)
}

func (c CommandQueue) EnqueueWriteBuffer(buffer Buffer, blockingRead bool, dataPtr interface{}) error {
	var br C.cl_bool
	if blockingRead {
		br = C.CL_TRUE
	} else {
		br = C.CL_FALSE
	}

	var ptr unsafe.Pointer
	var dataLen uint64
	switch p := dataPtr.(type) {
	case []float32:
		dataLen = uint64(len(p) * 4)
		ptr = unsafe.Pointer(&p[0])
	default:
		return errors.New("Unexpected type for dataPtr")
	}

	errInt := clError(C.clEnqueueWriteBuffer(c.commandQueue,
		buffer.buffer,
		br,
		0,
		C.size_t(dataLen),
		ptr,
		0, nil, nil))
	return clErrorToError(errInt)
}

func (c CommandQueue) Release() {
	C.clReleaseCommandQueue(c.commandQueue)
}

func (c CommandQueue) Flush() {
	C.clFlush(c.commandQueue)
}

func (c CommandQueue) Finish() {
	C.clFinish(c.commandQueue)
}
