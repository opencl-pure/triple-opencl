# pure
this is fork and simplify of https://github.com/Zyko0/go-opencl, big thank, 
this package provide low level wrapper to OpenCL,
that means it is 1:1 wrapper C:GO - no GO error handling
only GO types map OpenCL function and cl_types without cgo, powered by
https://github.com/ebitengine/purego and inspired by https://github.com/krrishnarraj/libopencl-stub, 
thank to both of them!
# warning
this is really low level wrapper I recommend to start with https://github.com/Zyko0/go-opencl, which still is low level, but add GO errors wrapper or
my package "high", but the package "high" is high level and not all functions are implemented.
# example

```go
package main

import (
	"errors"
	"github.com/opencl-pure/triple-opencl/constants"
	"github.com/opencl-pure/triple-opencl/pure"
	"log"
)

func main() {
	err := pure.Init(2) //init with version of OpenCL
	if err != nil {
		log.Println(err)
		return
	}
	numPlatforms := uint32(0)
	st := pure.GetPlatformIDs(0, nil, &numPlatforms)
	if st != constants.CL_SUCCESS {
		log.Println(errors.New("oops platform error"))
		return
	}

	platformIDs := make([]pure.Platform, numPlatforms)
	st = pure.GetPlatformIDs(numPlatforms, platformIDs, nil)
	if st != constants.CL_SUCCESS {
		log.Println(errors.New("oops none ...."))
		return
	}
	// ....

}    
```
