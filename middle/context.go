package middle

import (
	"github.com/opencl-pure/triple-opencl/constants"
	"github.com/opencl-pure/triple-opencl/pure"
	"unsafe"
)

type contextProperty pure.ContextProperty

type ContextProperties pure.ContextProperties

func (cp *ContextProperties) compile() []contextProperty {
	if cp == nil {
		return []contextProperty{0}
	}
	var properties []contextProperty
	if cp.Platform != nil {
		properties = append(properties, constants.CL_CONTEXT_PROPERTIES, contextProperty(*cp.Platform))
	}
	if cp.InteropUserSync != nil {
		b := contextProperty(0)
		if *cp.InteropUserSync {
			b = 1
		}
		properties = append(properties, constants.CL_CONTEXT_INTEROP_USER_SYNC, b)
	}
	if cp.GLContextKHR != nil {
		properties = append(properties, constants.CL_GL_CONTEXT_KHR, contextProperty(*cp.GLContextKHR))
	}
	if cp.WGL_HDC_KHR != nil {
		properties = append(properties, constants.CL_WGL_HDC_KHR, contextProperty(*cp.WGL_HDC_KHR))
	}
	// End of list should be marked as an extra zero
	return append(properties, 0)
}

type Context struct {
	C pure.Context
}

// TODO: make properties into a struct instead of weird map<uint32>

func (d *Device) CreateContext(properties *ContextProperties) (*Context, error) {
	var st pure.Status
	flattened := properties.compile()
	ctx := pure.CreateContext(unsafe.Pointer(&flattened[0]), 1, []pure.Device{d.D}, nil, nil, &st)
	if st != constants.CL_SUCCESS {
		return nil, pure.StatusToErr(st)
	}
	return &Context{ctx}, nil
}

func (c *Context) CreateCommandQueue(device *Device) (*CommandQueue, error) {
	var st pure.Status
	queue := pure.CreateCommandQueue(c.C, device.D, 0, &st)
	if st != constants.CL_SUCCESS {
		return nil, pure.StatusToErr(st)
	}
	return &CommandQueue{C: queue}, nil
}

type CommandQueueProperty pure.CommandQueueProperty

func (c *Context) CreateCommandQueueWithProperties(device Device, properties []pure.CommandQueueProperty) (*CommandQueue, error) {
	var st pure.Status
	property := pure.CommandQueueProperty(0)
	for _, p := range properties {
		property |= p
	}
	queue := pure.CreateCommandQueueWithProperties(c.C, device.D, property, &st)
	if st != constants.CL_SUCCESS {
		return nil, pure.StatusToErr(st)
	}
	return &CommandQueue{queue}, nil
}

func (c *Context) Release() error {
	return pure.StatusToErr(pure.ReleaseContext(c.C))
}

func (c *Context) CreateProgram(source string) (*Program, error) {
	var st pure.Status
	program := pure.CreateProgramWithSource(c.C, 1, []string{source}, []pure.Size{pure.Size(len(source))}, &st)
	if st != constants.CL_SUCCESS {
		return nil, pure.StatusToErr(st)
	}
	return &Program{program}, nil
}

func (c Context) CreateBuffer(flags []pure.MemFlag, size uint) (*Buffer, error) {
	var st pure.Status
	memFlags := pure.MemFlag(0)
	for _, f := range flags {
		memFlags |= f
	}
	buffer := pure.CreateBuffer(c.C, memFlags, pure.Size(size), nil, &st)
	if st != constants.CL_SUCCESS {
		return nil, pure.StatusToErr(st)
	}
	return &Buffer{B: buffer}, nil
}

func (c Context) CreateImage2D(flags []pure.MemFlag, format pure.ImageFormat, width, height, rowPitch int, data unsafe.Pointer) (*Image, error) {
	var st pure.Status
	memFlags := pure.MemFlag(0)
	for _, f := range flags {
		memFlags |= f
	}
	w, h, r := pure.Size(width), pure.Size(height), pure.Size(rowPitch)
	buffer := pure.CreateImage2D(c.C, memFlags, &format, w, h, r, data, &st)
	if st != constants.CL_SUCCESS {
		return nil, pure.StatusToErr(st)
	}
	return &Image{B: buffer}, nil
}

// GL

func (c *Context) CreateFromGLTexture(flags []pure.MemFlag, target pure.GLEnum, texture pure.GLUint) (*Buffer, error) {
	var st pure.Status
	memFlags := pure.MemFlag(0)
	for _, f := range flags {
		memFlags |= f
	}
	buffer := pure.CreateFromGLTexture(c.C, memFlags, target, 0, texture, &st)
	if st != constants.CL_SUCCESS {
		return nil, pure.StatusToErr(st)
	}
	return &Buffer{B: buffer}, nil
}
