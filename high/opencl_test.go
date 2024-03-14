package high

import (
	"fmt"
	"github.com/opencl-pure/triple-opencl/constants"
	"github.com/opencl-pure/triple-opencl/pure"
	"image"
	"image/color"
	"image/jpeg"
	_ "image/png"
	"os"
	"testing"
)

func TestGetDevices(t *testing.T) {
	err := Init(pure.Version2_0)
	if err != nil {
		t.Fatal(err)
	}
	ds, err := GetDevices(constants.CL_DEVICE_TYPE_ALL)
	if err != nil {
		t.Fatal(err)
	}
	for _, d := range ds {
		t.Log(d.Name())
		t.Log(d.Profile())
		t.Log(d.OpenCLCVersion())
		t.Log(d.DriverVersion())
		t.Log(d.Extensions())
		t.Log(d.Vendor())
		err = d.Release()
		if err != nil {
			t.Fatal(err)
		}
	}
	d, err := GetDefaultDevice()
	if err != nil {
		t.Fatal(err)
	}
	if d == nil {
		t.Fatal("device is nil")
	}
	fmt.Println(d)
	err = d.Release()
	if err != nil {
		t.Fatal(err)
	}
}

func TestBytes(t *testing.T) {
	err := Init(pure.Version2_0)
	if err != nil {
		t.Fatal(err)
	}
	d, err := GetDefaultDevice()
	if err != nil {
		t.Fatal(err)
	}
	defer d.Release()
	b, err := d.NewBytes(16)
	if err != nil {
		t.Fatal(err)
	}
	data := []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	err = <-b.Copy(data)
	if err != nil {
		t.Fatal(err)
	}
	retrievedData, err := b.Data()
	if err != nil {
		t.Fatal(err)
	}
	if len(retrievedData) != len(data) {
		t.Fatal("data not same length")
	}
	for i := 0; i < 16; i++ {
		if data[i] != retrievedData[i] {
			t.Fatal("retrieved data not equal to sended data")
		}
	}
	err = b.Release()
	if err != nil {
		t.Fatal(err)
	}
}

func TestVector(t *testing.T) {
	err := Init(pure.Version2_0)
	if err != nil {
		t.Fatal(err)
	}
	d, err := GetDefaultDevice()
	if err != nil {
		t.Fatal(err)
	}
	defer d.Release()
	v, err := NewVector(d, []float32{0, 1, 2, 3, 0, 5, 6, 7, 8, 9, 0, 1, 12, 13, 4, 15})
	if err != nil {
		t.Fatal(err)
	}
	data := []float32{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	err = <-v.Copy(data)
	if err != nil {
		t.Fatal(err)
	}
	retrievedData, err := v.Data()
	if err != nil {
		t.Fatal(err)
	}
	if len(retrievedData) != len(data) {
		t.Fatal("data not same length")
	}
	for i := 0; i < 16; i++ {
		if data[i] != retrievedData[i] {
			t.Fatal("retrieved data not equal to sended data")
		}
	}
	err = v.Release()
	if err != nil {
		t.Fatal(err)
	}
}

const testKernel = `
__kernel void testKernel(__global float* data) {
	const int i = get_global_id (0);
	data[i] += 1;
}
__kernel void testByteKernel(__global char* data) {
	const int i = get_global_id (0);
	data[i] += 1;
}
`

func TestBadProgram(t *testing.T) {
	err := Init(pure.Version2_0)
	if err != nil {
		t.Fatal(err)
	}
	d, err := GetDefaultDevice()
	if err != nil {
		t.Fatal(err)
	}
	defer d.Release()
	defer recoverAddProgram(t)
	_, err = d.AddProgram("meh")
	if err != nil {
		panic(err)
	}
}

func recoverAddProgram(t *testing.T) {
	if err := recover(); err == nil {
		t.Fatal("not correct program compiled without error")
	}
}

func TestBadKernel(t *testing.T) {
	err := Init(pure.Version2_0)
	if err != nil {
		t.Fatal(err)
	}
	d, err := GetDefaultDevice()
	if err != nil {
		t.Fatal(err)
	}
	defer d.Release()
	defer recoverKernel(t)
	_, err = d.AddProgram(testKernel)
	if err != nil {
		panic(err)
	}
	_, err = d.Kernel("meh")
	if err != nil {
		panic(err)
	}
}

func recoverKernel(t *testing.T) {
	if err := recover(); err == nil {
		t.Fatal("getting nonexisting kernel")
	}
}

func TestKernel(t *testing.T) {
	err := Init(pure.Version2_0)
	if err != nil {
		t.Fatal(err)
	}
	d, err := GetDefaultDevice()
	if err != nil {
		t.Fatal(err)
	}
	d.AddProgram(testKernel)
	k, err := d.Kernel("testKernel")
	if err != nil {
		t.Fatal(err)
	}
	v, err := NewVector(d, []float32{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15})
	if err != nil {
		t.Fatal(err)
	}
	defer v.Release()
	data := []float32{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	err = <-v.Copy(data)
	if err != nil {
		t.Fatal(err)
	}
	_, err = k.Global(16).Local(1).Run(false, nil, v)
	if err != nil {
		t.Fatal(err)
	}
	receivedData, err := v.Data()
	if err != nil {
		t.Fatal(err)
	}
	for i := 0; i < 16; i++ {
		if data[i]+1 != receivedData[i] {
			t.Error("receivedData not equal to data")
		}
	}
	_, err = v.Map(k, false, nil)
	if err != nil {
		t.Fatal(err)
	}
	receivedData, err = v.Data()
	if err != nil {
		t.Fatal(err)
	}
	for i := 0; i < 16; i++ {
		if data[i]+2 != receivedData[i] {
			t.Error("receivedData not equal to data")
		}
	}
	err = d.Release()
	if err != nil {
		t.Fatal(err)
	}
}

const invertColorKernel = `
__constant sampler_t sampler = CLK_NORMALIZED_COORDS_FALSE | CLK_ADDRESS_CLAMP_TO_EDGE | CLK_FILTER_NEAREST;

__kernel void invert(__read_only image2d_t src, __write_only image2d_t dest) {
	const int2 pos = {get_global_id(0), get_global_id(1)};
	float4 pixel = read_imagef(src, sampler, pos);
	pixel.x = 1 - pixel.x;
	pixel.y = 1 - pixel.y;
	pixel.z = 1 - pixel.z;
	write_imagef(dest, pos, pixel);
}`

func readImage(d *Device, path string) (*Image, error) {
	imgFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	img, _, err := image.Decode(imgFile)
	if err != nil {
		return nil, err
	}
	i, err := d.NewImageFromImage2D(img)
	if err != nil {
		return nil, err
	}
	return i, nil
}

func writeImage(img *Image, path string) error {
	receivedImg, err := img.Data()
	if err != nil {
		return err
	}
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	return jpeg.Encode(f, receivedImg, nil)
}

func TestImage(t *testing.T) {
	err := Init(pure.Version2_0)
	if err != nil {
		t.Fatal(err)
	}
	d, err := GetDefaultDevice()
	if err != nil {
		t.Fatal(err)
	}
	defer d.Release()
	img, err := readImage(d, "test_data/opencl.jpg")
	if err != nil {
		t.Fatal(err)
	}
	defer img.Release()
	_, err = d.AddProgram(invertColorKernel)
	if err != nil {
		t.Fatal(err)
	}
	k, err := d.Kernel("invert")
	if err != nil {
		t.Fatal(err)
	}
	invertedImg, err := d.NewImage2D(ImageTypeRGBA, img.Bounds())
	if err != nil {
		t.Fatal(err)
	}
	_, err = k.Global(img.Bounds().Dx(), img.Bounds().Dy()).Local(1, 1).Run(false, nil, img, invertedImg)
	if err != nil {
		t.Fatal(err)
	}
	err = writeImage(invertedImg, "/tmp/test.jpg")
	if err != nil {
		t.Fatal(err)
	}
	grayImg, err := readImage(d, "test_data/gopher.jpg")
	if err != nil {
		t.Fatal(err)
	}
	defer grayImg.Release()
	invertedGrayImg, err := d.NewImage2D(ImageTypeGray, grayImg.Bounds())
	if err != nil {
		t.Fatal(err)
	}
	_, err = k.Global(grayImg.Bounds().Dx(), grayImg.Bounds().Dy()).Local(1, 1).Run(false, nil, grayImg, invertedGrayImg)
	if err != nil {
		t.Fatal(err)
	}
	err = writeImage(invertedGrayImg, "/tmp/test_gray.jpg")
	if err != nil {
		t.Fatal(err)
	}
}

const gaussianBlurKernel = `
__constant sampler_t sampler = CLK_NORMALIZED_COORDS_FALSE | CLK_ADDRESS_CLAMP_TO_EDGE | CLK_FILTER_NEAREST;
 
__kernel void gaussian_blur(
        __read_only image2d_t image,
        __constant float * mask,
        __global float * blurredImage,
        __private int maskSize
    ) {
 
    const int2 pos = {get_global_id(0), get_global_id(1)};
 
    // Collect neighbor values and multiply with Gaussian
    float sum = 0.0f;
    for(int a = -maskSize; a < maskSize+1; a++) {
        for(int b = -maskSize; b < maskSize+1; b++) {
            sum += mask[a+maskSize+(b+maskSize)*(maskSize*2+1)]
                *read_imagef(image, sampler, pos + (int2)(a,b)).x;
        }
    }
 
    blurredImage[pos.x+pos.y*get_global_size(0)] = sum;
}`

func TestGaussianBlur(t *testing.T) {
	err := Init(pure.Version2_0)
	if err != nil {
		t.Fatal(err)
	}
	d, err := GetDefaultDevice()
	if err != nil {
		t.Fatal(err)
	}
	defer d.Release()
}

func invertCPU(input image.Image) image.Image {
	bounds := input.Bounds()
	newImg := image.NewRGBA(bounds)

	for x := 0; x < bounds.Max.X; x++ {
		for y := 0; y < bounds.Max.Y; y++ {
			//r, g, b, a := input.At(x, y).(color.YCbCr).RGBA()
			r, g, b, a := input.At(x, y).(color.RGBA).RGBA()
			newImg.Set(x, y, color.RGBA{
				R: uint8(0xff - r),
				G: uint8(0xff - g),
				B: uint8(0xff - b),
				A: uint8(a),
			})
		}
	}
	return newImg
}

func invertGPU(d *Device, kernel *Kernel, img *Image) (image.Image, error) {
	invertedImg, err := d.NewImage2D(ImageTypeRGBA, img.Bounds())
	if err != nil {
		return nil, err
	}
	defer invertedImg.Release()
	_, err = kernel.Global(img.Bounds().Dx(), img.Bounds().Dy()).Local(1, 1).Run(false, nil, img, invertedImg)
	if err != nil {
		return nil, err
	}
	return invertedImg.Data()
}

func BenchmarkInvertCPU(b *testing.B) {
	imgFile, err := os.Open("/home/micro/Downloads/wallpaper.png")
	if err != nil {
		b.Fatal(err)
	}
	img, _, err := image.Decode(imgFile)
	if err != nil {
		b.Fatal(err)
	}
	for n := 0; n < b.N; n++ {
		invertCPU(img)
	}
}

func BenchmarkInvertGPU(b *testing.B) {
	err := Init(pure.Version2_0)
	if err != nil {
		b.Fatal(err)
	}
	d, err := GetDefaultDevice()
	if err != nil {
		b.Fatal(err)
	}
	defer d.Release()
	_, err = d.AddProgram(invertColorKernel)
	if err != nil {
		b.Fatal(err)
	}
	k, err := d.Kernel("invert")
	if err != nil {
		b.Fatal(err)
	}
	imgFile, err := os.Open("/home/micro/Downloads/wallpaper.png")
	if err != nil {
		b.Fatal(err)
	}
	input, _, err := image.Decode(imgFile)
	if err != nil {
		b.Fatal(err)
	}
	img, err := d.NewImageFromImage2D(input)
	if err != nil {
		b.Fatal(err)
	}
	defer img.Release()
	for n := 0; n < b.N; n++ {
		_, err := invertGPU(d, k, img)
		if err != nil {
			b.Fatal(err)
		}
	}
}
