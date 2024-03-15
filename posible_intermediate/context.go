package posible_intermediate

// #include "opencl.h"
import "C"
import (
	"github.com/opencl-pure/triple-opencl/constants"
	"github.com/opencl-pure/triple-opencl/pure"
)

type Context struct {
	context pure.Context
}

func createContext(device *Device) (*Context, error) {
	// TODO add more functionality. Super simple context creation right now
	var errInt pure.Status
	if pure.CreateContext == nil {
		return nil, pure.Uninitialized("CreateContext")
	}
	ctx := pure.CreateContext(nil, 1,
		[]pure.Device{device.deviceID}, nil, nil, &errInt)
	if errInt != constants.CL_SUCCESS {
		return nil, pure.StatusToErr(errInt)
	}
	return &Context{ctx}, nil
}

func (c *Context) CreateCommandQueue(device *Device) (*CommandQueue, error) {
	return createCommandQueue(c, device)
}

func (c *Context) CreateProgramWithSource(programCode string) (*Program, error) {
	return createProgramWithSource(c, programCode)
}

func (c *Context) CreateBuffer(memFlags []MemFlags, size uint64) (*Buffer, error) {
	return createBuffer(c, memFlags, size)
}

func (c Context) Release() error {
	if pure.ReleaseContext == nil {
		return pure.Uninitialized("ReleaseContext")
	}
	return pure.StatusToErr(pure.ReleaseContext(c.context))
}
