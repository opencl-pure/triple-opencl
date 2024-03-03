package pure

import (
	"errors"
	"github.com/ebitengine/purego"
)

func Init(version int) (e error) {
	handle, err := loadLibrary()
	if err != nil {
		return err
	}
	defer func() {
		reco := recover()
		if reco == nil {
			return
		}
		err, ok := recover().(error)
		if ok {
			e = err
			return
		}
		e = errors.New("unknown error")
	}()
	// Platform
	purego.RegisterLibFunc(&GetPlatformIDs, handle, "clGetPlatformIDs")
	purego.RegisterLibFunc(&GetPlatformInfo, handle, "clGetPlatformInfo")
	// Device
	purego.RegisterLibFunc(&GetDeviceIDs, handle, "clGetDeviceIDs")
	purego.RegisterLibFunc(&GetDeviceInfo, handle, "clGetDeviceInfo")
	purego.RegisterLibFunc(&ReleaseDevice, handle, "clReleaseDevice")
	// Event
	purego.RegisterLibFunc(&ReleaseEvent, handle, "clReleaseEvent")
	purego.RegisterLibFunc(&WaitForEvents, handle, "clWaitForEvents")
	// Context
	purego.RegisterLibFunc(&CreateContext, handle, "clCreateContext")
	purego.RegisterLibFunc(&ReleaseContext, handle, "clReleaseContext")
	purego.RegisterLibFunc(&CreateProgramWithSource, handle, "clCreateProgramWithSource")
	purego.RegisterLibFunc(&CreateBuffer, handle, "clCreateBuffer")
	purego.RegisterLibFunc(&CreateImage2D, handle, "clCreateImage2D")
	// Command queue
	if version >= 2 {
		purego.RegisterLibFunc(&CreateCommandQueueWithProperties, handle, "clCreateCommandQueueWithProperties")
	} else {
		purego.RegisterLibFunc(&CreateCommandQueue, handle, "clCreateCommandQueue")
	}
	purego.RegisterLibFunc(&EnqueueBarrier, handle, "clEnqueueBarrier")
	purego.RegisterLibFunc(&EnqueueNDRangeKernel, handle, "clEnqueueNDRangeKernel")
	purego.RegisterLibFunc(&EnqueueReadBuffer, handle, "clEnqueueReadBuffer")
	//TODO: purego: broken too many arguments
	//purego.RegisterLibFunc(&EnqueueReadImage, handle, "clEnqueueReadImage")

	//purego.RegisterLibFunc(&EnqueueReadImage, handle, "clEnqueueWriteImage")

	purego.RegisterLibFunc(&EnqueueWriteBuffer, handle, "clEnqueueWriteBuffer")
	//TODO: purego: broken too many arguments
	//purego.RegisterLibFunc(&EnqueueMapImage, handle, "clEnqueueMapImage")
	//purego.RegisterLibFunc(&EnqueueMapBuffer, handle, "clEnqueueMapBuffer")

	purego.RegisterLibFunc(&EnqueueUnmapMemObject, handle, "clEnqueueUnmapMemObject")
	purego.RegisterLibFunc(&FinishCommandQueue, handle, "clFinish")
	purego.RegisterLibFunc(&FlushCommandQueue, handle, "clFlush")
	purego.RegisterLibFunc(&ReleaseCommandQueue, handle, "clReleaseCommandQueue")
	// Program
	purego.RegisterLibFunc(&BuildProgram, handle, "clBuildProgram")
	purego.RegisterLibFunc(&GetProgramBuildInfo, handle, "clGetProgramBuildInfo")
	purego.RegisterLibFunc(&GetProgramInfo, handle, "clGetProgramInfo")
	purego.RegisterLibFunc(&CreateKernel, handle, "clCreateKernel")
	purego.RegisterLibFunc(&ReleaseProgram, handle, "clReleaseProgram")
	// Kernel
	purego.RegisterLibFunc(&SetKernelArg, handle, "clSetKernelArg")
	purego.RegisterLibFunc(&ReleaseKernel, handle, "clReleaseKernel")
	// Buffer
	purego.RegisterLibFunc(&GetMemObjectInfo, handle, "clGetMemObjectInfo")
	purego.RegisterLibFunc(&ReleaseMemObject, handle, "clReleaseMemObject")

	return nil
}
func InitializeGLSharing() error {
	handle, err := loadLibrary()
	if err != nil {
		return err
	}
	// GL
	purego.RegisterLibFunc(&CreateFromGLTexture, handle, "clCreateFromGLTexture")
	purego.RegisterLibFunc(&EnqueueAcquireGLObjects, handle, "clEnqueueAcquireGLObjects")
	purego.RegisterLibFunc(&EnqueueReleaseGLObjects, handle, "clEnqueueReleaseGLObjects")
	purego.RegisterLibFunc(&GetGLObjectInfo, handle, "clGetGLObjectInfo")
	purego.RegisterLibFunc(&GetGLTextureInfo, handle, "clGetGLTextureInfo")

	return nil
}
