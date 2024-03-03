package pure

import "unsafe"

type ImageFormat struct {
	ChannelOrder ImageChannelOrder
	ChannelType  ImageChannelType
}

type BufferData struct {
	TypeSize uintptr
	DataSize uintptr
	Pointer  unsafe.Pointer
}

type ImageData struct {
	*BufferData
	Origin     [3]uint
	Region     [3]uint
	RowPitch   uint
	SlicePitch uint
}

type (
	// primitive types
	Size                   uint
	Status                 int32
	Program                uint
	ProgramBuildInfo       uint32
	DeviceType             uint32
	PlatformInfo           uint
	Context                uint
	Platform               uint
	Device                 uint
	DeviceInfo             uint32
	MemFlag                uint32
	ImageChannelOrder      uint32
	ImageChannelType       uint32
	CommandQueueProperties uint32
	CommandQueue           uint
	Buffer                 uint
	CommandQueueProperty   uint32
	MemInfo                uint32
	MapFlag                uint32
	Kernel                 uint
	Event                  uint

	// notify function types
	CreateContextNotifyFuncType func(errinfo, privateInfo []byte, cb Size, userData []byte)
	BuildProgramNotifyFuncType  func(program Program, userData []byte)

	// structs
	KernelArg struct {
		ptr  unsafe.Pointer
		size Size
	}

	// GL
	GLEnum          uint32
	GLInt           int32
	GLUint          uint32
	CLGLObjectType  uint32
	CLGLTextureInfo uint32
)
