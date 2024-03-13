package midle

import (
	"github.com/opencl-pure/triple-opencl/constants"
	"github.com/opencl-pure/triple-opencl/pure"
	"unsafe"
)

type Buffer struct {
	B pure.Buffer
}

func (b *Buffer) getInfo(name pure.MemInfo) (uint, error) {
	if pure.GetMemObjectInfo == nil {
		return 0, pure.Uninitialized("GetMemObjectInfo")
	}
	info := uint(0)
	st := pure.GetMemObjectInfo(b.B, name, pure.Size(unsafe.Sizeof(info)), unsafe.Pointer(&info), nil)
	if st != constants.CL_SUCCESS {
		return 0, pure.StatusToErr(st)
	}
	return info, nil
}

func (b *Buffer) Size() (uint, error) {
	return b.getInfo(pure.MemInfo(constants.CL_MEM_SIZE))
}

func (b *Buffer) Release() error {
	if pure.ReleaseMemObject == nil {
		return pure.Uninitialized("ReleaseMemObject")
	}
	return pure.StatusToErr(pure.ReleaseMemObject(b.B))
}

// GL

func (b *Buffer) GetGLObjectInfo() (pure.CLGLObjectType, error) {
	if pure.GetGLObjectInfo == nil {
		return 0, pure.Uninitialized("GetGLObjectInfo")
	}
	var objectType pure.CLGLObjectType
	st := pure.GetGLObjectInfo(b.B, &objectType, nil)
	if st != constants.CL_SUCCESS {
		return 0, pure.StatusToErr(st)
	}
	return objectType, nil
}

func (b *Buffer) GetGLTextureInfo(info pure.CLGLTextureInfo) (uint32, error) {
	if pure.GetGLTextureInfo == nil {
		return 0, pure.Uninitialized("GetGLTextureInfo")
	}
	var results = []uint32{0}
	st := pure.GetGLTextureInfo(
		b.B, info, pure.Size(unsafe.Sizeof(&results[0])),
		unsafe.Pointer(&results[0]), nil,
	)
	if st != constants.CL_SUCCESS {
		return 0, pure.StatusToErr(st)
	}
	return results[0], nil
}
