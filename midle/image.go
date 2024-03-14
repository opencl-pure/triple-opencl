package midle

import (
	"github.com/opencl-pure/triple-opencl/pure"
	"unsafe"
)

type BufferData struct {
	TypeSize uintptr
	DataSize uintptr
	Pointer  unsafe.Pointer
}
type ImageChannelOrder uint32

type ImageChannelType pure.ImageChannelType

type ImageFormat struct {
	ChannelOrder ImageChannelOrder
	ChannelType  ImageChannelType
}

type ImageData struct {
	*BufferData
	Origin     [3]uint
	Region     [3]uint
	RowPitch   uint
	SlicePitch uint
}
