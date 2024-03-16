package main

import (
	"embed"
	_ "embed"
	"fmt"
	"github.com/opencl-pure/triple-opencl/high"
	"github.com/opencl-pure/triple-opencl/pure"
)

//go:embed addOne.cl
var kernel embed.FS

func main() {
	err := high.Init(pure.Version2_0) //init with version of OpenCL
	if err != nil {
		panic("no possible call high functions")
	}
	//Do not create platforms/devices/contexts/queues/...
	//Just get the GPU
	d, err := high.GetDefaultDevice()
	if err != nil {
		panic("no high device")
	}
	defer d.Release()

	// has several kinds of device memory object: Bytes, Vector, Image
	//allocate buffer on the device (16 elems of float32)
	v, err := d.NewVector([]float32{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15})
	if err != nil {
		panic("could not allocate buffer")
	}
	err = <-v.Reset([]float32{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}) // usefull at repeat call
	if err != nil {
		panic(err)
	}
	defer v.Release()
	b, err := kernel.ReadFile("addOne.cl")
	if err != nil {
		panic(err)
	}
	//Add program source to device, get kernel
	_, err = d.AddProgram(string(b))
	if err != nil {
		panic(err)
	}
	k, err := d.Kernel("addOne")
	if err != nil {
		panic("could not run kernel")
	}
	//run kernel (global work size 16 and local work size 1)
	event, err := k.Global(16).Local(1).Run(nil, v)
	if err != nil {
		panic("could not run kernel")
	}
	defer event.Release()
	//Get data from vector
	newData, err := v.Data()
	if err != nil {
		panic("could not get data from buffer")
	}
	newDataSlice := newData.Interface().([]float32)
	//prints out [1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16]
	fmt.Println(newDataSlice)
	n, err := v.DataArray()
	if err != nil {
		panic(err)
	}
	fmt.Println(*(n.Interface().(*[16]float32)))
}
