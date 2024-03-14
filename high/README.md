# purego-opencl

fork from BlackCL to work with OpenCL. These are highly opinionated OpenCL bindings for Go. It tries to make GPU computing easy, with some sugar abstraction, Go's concurency and channels. (without magic)

# still in develop !!!
this is still in develop, I translate CGO to [purego](https://github.com/ebitengine/purego), inspirate by low level wrapper for OpenCL in C (https://github.com/krrishnarraj/libopencl-stub), every helps, corrects are welcome ...
# pure
this is fork and simplify of https://github.com/Zyko0/go-opencl, but for most people too much, because it is wrapper 1:1 so no space to regular GO's error handling, also my fork of opencl(blackcl) is too much hight level of abstraction, so choose is on your preferencies ...

# examples

```go
package main

import (
	"fmt"
	opencl "github.com/opencl-pure/triple-opencl/high"
)

func main() {
	err := opencl.Init(2) //init with version of OpenCL
	if err != nil {
		panic("no possible call opencl functions")
	}
	//Do not create platforms/devices/contexts/queues/...
	//Just get the GPU
	d, err := opencl.GetDefaultDevice()
	if err != nil {
		panic("no opencl device")
	}
	defer d.Release()

	// has several kinds of device memory object: Bytes, Vector, Image
	//allocate buffer on the device (16 elems of float32)
	v, err := opencl.NewVector[float32](d, []float32{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15})
	if err != nil {
		panic("could not allocate buffer")
	}
	defer v.Release()

	//an complicated kernel
	const kernelSource = `
__kernel void addOne(__global float* data) {
	const int i = get_global_id (0);
	data[i] += 1;
}
`

	//Add program source to device, get kernel
	_, err = d.AddProgram(kernelSource)
	if err != nil {
		panic("could not run kernel")
	}
	k, err := d.Kernel("addOne")
	if err != nil {
		panic("could not run kernel")
	}
	//run kernel (global work size 16 and local work size 1)
	_, err = k.Global(16).Local(1).Run(false, nil, v)
	if err != nil {
		panic("could not run kernel")
	}

	//Get data from vector
	newData, err := v.Data()
	if err != nil {
		panic("could not get data from buffer")
	}

	//prints out [1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16]
	fmt.Println(newData)
}
```
