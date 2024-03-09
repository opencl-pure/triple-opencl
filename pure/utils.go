package pure

import (
	"fmt"
	"github.com/ebitengine/purego"
	"image"
	"unsafe"
)

type BufferType interface {
	~float32 | ~float64 | ~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

func GetBufferData[T BufferType](data []T) *BufferData {
	size := unsafe.Sizeof(data[0])
	return &BufferData{
		TypeSize: size,
		DataSize: uintptr(len(data)) * size,
		Pointer:  unsafe.Pointer(&data[0]),
	}
}

func GetImageBufferData(img image.RGBA) *ImageData {
	bounds := img.Bounds()
	return &ImageData{
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

func joinErr(err error, err2 error) error {
	if err == nil {
		return err2
	}
	if err2 == nil {
		return err
	}
	return fmt.Errorf("%s, %s", err.Error(), err2.Error())
}

func registerLibFuncWithoutPanic(fptr interface{}, handle uintptr, name string, error0 error) (e error) {
	e = error0
	defer func() {
		if r := recover(); r != nil {
			err, ok := r.(error)
			if ok {
				e = err
				return
			}
			e = joinErr(error0, err)
		}
	}()
	purego.RegisterLibFunc(fptr, handle, name)
	return
}
