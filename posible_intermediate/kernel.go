package posible_intermediate

import (
	"errors"
	"github.com/opencl-pure/triple-opencl/constants"
	"github.com/opencl-pure/triple-opencl/pure"
	"unsafe"
)

type Kernel struct {
	kernel pure.Kernel
}

func createKernel(program *Program, kernelName string) (*Kernel, error) {
	var errInt pure.Status
	if pure.CreateKernel == nil {
		return nil, pure.Uninitialized("CreateKernel")
	}
	kernel := pure.CreateKernel(program.program, kernelName, &errInt)
	if errInt != constants.CL_SUCCESS {
		return nil, pure.StatusToErr(errInt)
	}
	return &Kernel{kernel}, nil
}

func (k *Kernel) SetArg(argIndex uint32, argSize uint64, argValue interface{}) error {
	var argPtr unsafe.Pointer
	switch argValue.(type) {
	case *Buffer:
		argPtr = unsafe.Pointer(argValue.(*Buffer))
	default:
		return errors.New("unknown type for argValue")
	}
	errInt := pure.SetKernelArg(k.kernel, argIndex, pure.Size(argSize), argPtr)
	return pure.StatusToErr(errInt)
}

func (k *Kernel) Release() {
	pure.ReleaseKernel(k.kernel)
}
