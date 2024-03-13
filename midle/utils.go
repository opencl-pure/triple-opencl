package midle

import (
	"github.com/opencl-pure/triple-opencl/pure"
	"image"
	"unsafe"
)

type BufferType interface {
	~float32 | ~float64 | ~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

func GetBufferData[T BufferType](data []T) *pure.BufferData {
	size := unsafe.Sizeof(data[0])
	return &pure.BufferData{
		TypeSize: size,
		DataSize: uintptr(len(data)) * size,
		Pointer:  unsafe.Pointer(&data[0]),
	}
}

func GetImageBufferData(img image.RGBA) *pure.ImageData {
	bounds := img.Bounds()
	return &pure.ImageData{
		BufferData: GetBufferData(img.Pix),
		Origin: [3]uint{
			uint(bounds.Min.X), uint(bounds.Min.Y), 0,
		},
		Region: [3]uint{
			uint(bounds.Dx()), uint(bounds.Dy()), 1,
		},
		RowPitch:   0,
		SlicePitch: 0,
	}
}
