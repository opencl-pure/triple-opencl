package posible_intermediate

// #include "opencl.h"
import "C"
import (
	"github.com/opencl-pure/triple-opencl/constants"
	"github.com/opencl-pure/triple-opencl/pure"
	"unsafe"
)

type MemFlags uint64 // pure.Memflag diferent

const (
	MemReadWrite MemFlags = constants.CL_MEM_READ_WRITE
	MemWriteOnly MemFlags = constants.CL_MEM_WRITE_ONLY
	MemReadOnly  MemFlags = constants.CL_MEM_READ_ONLY
	// ...
)

type Buffer struct {
	buffer pure.Buffer
}

func createBuffer(context *Context, flags []MemFlags, size uint64) (*Buffer, error) {
	// AND together all flags
	flagBitField := uint64(0)
	for _, flag := range flags {
		flagBitField &= uint64(flag)
	}
	var errInt pure.Status
	if pure.CreateBuffer == nil {
		return nil, pure.Uninitialized("CreateBuffer")
	}
	buffer := pure.CreateBuffer(context.context, pure.MemFlag(flagBitField), pure.Size(size), nil, &errInt)
	if errInt != constants.CL_SUCCESS {
		return nil, pure.StatusToErr(errInt)
	}
	return &Buffer{buffer}, nil
}

func (b *Buffer) Size() uint64 {
	return uint64(unsafe.Sizeof(b.buffer))
}

func (b Buffer) Release() error {
	if pure.ReleaseMemObject == nil {
		return pure.Uninitialized("ReleaseMemObject")
	}
	return pure.StatusToErr(pure.ReleaseMemObject(b.buffer))
}
