package high

import (
	"errors"
	"github.com/opencl-pure/triple-opencl/constants"
	"github.com/opencl-pure/triple-opencl/pure"
	"image"
	"log"
	"unsafe"
)

// ImageType type of the image enum
type ImageType int

// available image types
const (
	ImageTypeGray = ImageType(constants.CL_INTENSITY)
	ImageTypeRGBA = ImageType(constants.CL_RGBA)
)

// Image memory buffer on the device with image data
type Image struct {
	buf       *buffer
	imageType ImageType
	bounds    image.Rectangle
	format    *pure.ImageFormat
}

// Bounds returns the image size
func (img *Image) Bounds() image.Rectangle {
	return img.bounds
}

// Release releases the buffer on the device
func (img *Image) Release() error {
	return img.buf.Release()
}

// NewImage2D NewImage allocates an image buffer
func (d *Device) NewImage2D(imageType ImageType, bounds image.Rectangle) (*Image, error) {
	return d.newImage2D(imageType, bounds, 0, nil)
}

// NewImageFromImage2D NewImageFromImage creates new Image and copies data from image.Image
func (d *Device) NewImageFromImage2D(img image.Image) (*Image, error) {
	data := imgData(img)
	var rowPitch int
	var imageType ImageType
	if g, ok := img.(*image.Gray); ok {
		imageType = ImageTypeGray
		rowPitch = g.Stride
	} else {
		imageType = ImageTypeRGBA
	}
	return d.newImage2D(imageType, img.Bounds(), rowPitch, unsafe.Pointer(&data[0]))
}

func (d *Device) newImage2D(imageType ImageType, bounds image.Rectangle, rowPitch int, data unsafe.Pointer) (*Image, error) {
	var format = &pure.ImageFormat{ChannelOrder: 0, ChannelType: 0}
	switch imageType {
	case ImageTypeGray:
		format.ChannelOrder = constants.CL_INTENSITY
		format.ChannelType = constants.CL_UNORM_INT8
	case ImageTypeRGBA:
		format.ChannelOrder = constants.CL_RGBA
		format.ChannelType = constants.CL_UNORM_INT8
	}
	flags := pure.MemFlag(constants.CL_MEM_READ_WRITE)
	if data != nil {
		flags = pure.MemFlag(constants.CL_MEM_READ_WRITE | constants.CL_MEM_COPY_HOST_PTR)
	}
	var ret pure.Status
	clBuffer := pure.CreateImage2D(
		d.ctx,
		flags,
		format,
		pure.Size(bounds.Dx()),
		pure.Size(bounds.Dy()),
		pure.Size(rowPitch),
		data,
		&ret)
	err := pure.StatusToErr(ret)
	if err != nil {
		return nil, err
	}
	if clBuffer == pure.Buffer(0) {
		return nil, ErrUnknown
	}
	size := bounds.Dx() * bounds.Dy()
	if imageType == ImageTypeRGBA {
		size *= 4
	}
	return &Image{
		buf: &buffer{
			memobj: clBuffer,
			size:   pure.Size(size),
			device: d,
		},
		bounds:    bounds,
		imageType: imageType,
		format:    format,
	}, nil
}

// Copy writes the image data to the buffer
func (img *Image) Copy(i image.Image) <-chan error {
	if !img.bounds.Eq(i.Bounds()) {
		ch := make(chan error, 1)
		ch <- errors.New("image bounds not equal")
		return ch
	}
	return img.copy(imgData(i))
}

func imgData(i image.Image) []byte {
	switch m := i.(type) {
	case *image.Gray:
		return m.Pix
	case *image.RGBA:
		return m.Pix
	}

	b := i.Bounds()
	w := b.Dx()
	h := b.Dy()
	data := make([]byte, w*h*4)
	dataOffset := 0
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			c := i.At(x+b.Min.X, y+b.Min.Y)
			r, g, b, a := c.RGBA()
			data[dataOffset] = uint8(r >> 8)
			data[dataOffset+1] = uint8(g >> 8)
			data[dataOffset+2] = uint8(b >> 8)
			data[dataOffset+3] = uint8(a >> 8)
			dataOffset += 4
		}
	}
	return data
}

func (img *Image) copy(data []byte) <-chan error {
	ch := make(chan error, 1)
	cOrigin := [3]pure.Size{0, 0, 0}
	cRegion := [3]pure.Size{pure.Size(img.bounds.Dx()), pure.Size(img.bounds.Dy()), 1}
	var event pure.Event
	err := pure.StatusToErr(pure.EnqueueWriteImage(
		img.buf.device.queue,
		img.buf.memobj,
		constants.CL_FALSE,
		cOrigin,
		cRegion,
		0,
		0,
		unsafe.Pointer(&data[0]),
		0,
		nil,
		&event,
	))
	if err != nil {
		ch <- err
		return ch
	}
	go func() {
		defer func() {
			if err2 := pure.StatusToErr(pure.ReleaseEvent(event)); err2 != nil {
				log.Println(err2)
			}
		}()
		ch <- pure.StatusToErr(pure.WaitForEvents(1, &event))
	}()
	return ch
}

// Data gets data from an image buffer and returns an image.Image
func (img *Image) Data() (image.Image, error) {
	data := make([]byte, img.buf.size)
	cOrigin := [3]pure.Size{0, 0, 0}
	cRegion := [3]pure.Size{pure.Size(img.bounds.Dx()), pure.Size(img.bounds.Dy()), 1}
	err := pure.StatusToErr(pure.EnqueueReadImage(
		img.buf.device.queue,
		img.buf.memobj,
		constants.CL_TRUE,
		cOrigin,
		cRegion,
		0,
		0,
		unsafe.Pointer(&data[0]),
		0,
		nil,
		nil,
	))
	if err != nil {
		return nil, errors.New("cannot get buffer data: " + err.Error())
	}
	switch img.imageType {
	case ImageTypeRGBA:
		img := image.NewRGBA(img.bounds)
		img.Pix = data
		return img, nil
	case ImageTypeGray:
		img := image.NewGray(img.bounds)
		img.Pix = data
		return img, nil
	}
	return nil, errors.New("cannot get image data from the buffer, not an image buffer")
}

// TODO General image
/*

// Image memory buffer on the device with image data
type Image struct {
	buf       *buffer
	imageType ImageType
	bounds    image.Rectangle
	format    *pure.ImageFormat
	desc      *C.cl_image_desc
}

//NewImage allocates an image buffer
func (d *Device) NewImage(imageType ImageType, bounds image.Rectangle) (*Image, error) {
	return d.newImage(imageType, bounds, 0, nil)
}

//NewImageFromImage creates new Image and copies data from image.Image
func (d *Device) NewImageFromImage(img image.Image) (*Image, error) {
	data := imgData(img)
	var rowPitch int
	var imageType ImageType
	if g, ok := img.(*image.Gray); ok {
		imageType = ImageTypeGray
		rowPitch = g.Stride
	} else {
		imageType = ImageTypeRGBA
	}
	return d.newImage(imageType, img.Bounds(), rowPitch, unsafe.Pointer(&data[0]))
}

func (d *Device) newImage(imageType ImageType, bounds image.Rectangle, rowPitch int, data unsafe.Pointer) (*Image, error) {
	var format C.cl_image_format
	switch imageType {
	case ImageTypeGray:
		format.image_channel_order = C.CL_INTENSITY
		format.image_channel_data_type = C.CL_UNORM_INT8
	case ImageTypeRGBA:
		format.image_channel_order = C.CL_RGBA
		format.image_channel_data_type = C.CL_UNORM_INT8
	}
	desc := C.create_image_desc(
		C.CL_MEM_OBJECT_IMAGE2D,
		C.size_t(bounds.Dx()),
		C.size_t(bounds.Dy()),
		0,
		0,
		C.size_t(rowPitch),
		0,
		0,
		0,
		nil)
	flags := C.cl_mem_flags(C.CL_MEM_READ_WRITE)
	if data != nil {
		flags = C.CL_MEM_READ_WRITE | C.CL_MEM_COPY_HOST_PTR
	}
	var ret C.cl_int
	clBuffer := C.clCreateImage(d.ctx, flags, &format, desc, data, &ret)
	err := toErr(ret)
	if err != nil {
		return nil, err
	}
	if clBuffer == nil {
		return nil, ErrUnknown
	}
	size := bounds.Dx() * bounds.Dy()
	if imageType == ImageTypeRGBA {
		size *= 4
	}
	return &Image{
		buf: &buffer{
			memobj: clBuffer,
			size:   size,
			device: d,
		},
		bounds:    bounds,
		imageType: imageType,
		format:    format,
		desc:      desc,
	}, nil
}
*/
/*
cl_image_desc* create_image_desc (
	cl_mem_object_type image_type,
	size_t image_width,
	size_t image_height,
	size_t image_depth,
	size_t image_array_size,
	size_t image_row_pitch,
	size_t image_slice_pitch,
	cl_uint num_mip_levels,
	cl_uint num_samples,
	cl_mem buffer
) {
	cl_image_desc *desc = malloc(sizeof(cl_image_desc));
	desc->image_type = image_type;
	desc->image_width = image_width;
	desc->image_height = image_height;
	desc->image_row_pitch = image_row_pitch;
	desc->image_slice_pitch = image_slice_pitch;
	desc->num_mip_levels = num_mip_levels;
	desc->num_samples = num_samples;
	desc->buffer = buffer;
	return desc;
}
*/
// import "C"
