package pure

func Init(version int) error {
	handle, err := loadLibrary()
	if err != nil {
		return err
	}
	// Platform
	err = registerLibFuncWithoutPanic(&GetPlatformIDs, handle, "clGetPlatformIDs", nil)
	err = registerLibFuncWithoutPanic(&GetPlatformInfo, handle, "clGetPlatformInfo", err)
	// Device
	err = registerLibFuncWithoutPanic(&GetDeviceIDs, handle, "clGetDeviceIDs", err)
	err = registerLibFuncWithoutPanic(&GetDeviceInfo, handle, "clGetDeviceInfo", err)
	err = registerLibFuncWithoutPanic(&ReleaseDevice, handle, "clReleaseDevice", err)
	// Event
	err = registerLibFuncWithoutPanic(&ReleaseEvent, handle, "clReleaseEvent", err)
	err = registerLibFuncWithoutPanic(&WaitForEvents, handle, "clWaitForEvents", err)
	// Context
	err = registerLibFuncWithoutPanic(&CreateContext, handle, "clCreateContext", err)
	err = registerLibFuncWithoutPanic(&ReleaseContext, handle, "clReleaseContext", err)
	err = registerLibFuncWithoutPanic(&CreateProgramWithSource, handle, "clCreateProgramWithSource", err)
	err = registerLibFuncWithoutPanic(&CreateBuffer, handle, "clCreateBuffer", err)
	err = registerLibFuncWithoutPanic(&CreateImage2D, handle, "clCreateImage2D", err)
	// Command queue
	if version >= 2 {
		err = registerLibFuncWithoutPanic(&CreateCommandQueueWithProperties, handle, "clCreateCommandQueueWithProperties", err)
	} else {
		err = registerLibFuncWithoutPanic(&CreateCommandQueue, handle, "clCreateCommandQueue", err)
	}
	err = registerLibFuncWithoutPanic(&EnqueueBarrier, handle, "clEnqueueBarrier", err)
	err = registerLibFuncWithoutPanic(&EnqueueNDRangeKernel, handle, "clEnqueueNDRangeKernel", err)
	err = registerLibFuncWithoutPanic(&EnqueueReadBuffer, handle, "clEnqueueReadBuffer", err)
	//TODO: purego: broken too many arguments
	err = registerLibFuncWithoutPanic(&EnqueueReadImage, handle, "clEnqueueReadImage", err)
	err = registerLibFuncWithoutPanic(&EnqueueWriteImage, handle, "clEnqueueWriteImage", err)

	err = registerLibFuncWithoutPanic(&EnqueueWriteBuffer, handle, "clEnqueueWriteBuffer", err)
	//TODO: purego: broken too many arguments
	err = registerLibFuncWithoutPanic(&EnqueueMapImage, handle, "clEnqueueMapImage", err)
	err = registerLibFuncWithoutPanic(&EnqueueMapBuffer, handle, "clEnqueueMapBuffer", err)

	err = registerLibFuncWithoutPanic(&EnqueueUnmapMemObject, handle, "clEnqueueUnmapMemObject", err)
	err = registerLibFuncWithoutPanic(&FinishCommandQueue, handle, "clFinish", err)
	err = registerLibFuncWithoutPanic(&FlushCommandQueue, handle, "clFlush", err)
	err = registerLibFuncWithoutPanic(&ReleaseCommandQueue, handle, "clReleaseCommandQueue", err)
	// Program
	err = registerLibFuncWithoutPanic(&BuildProgram, handle, "clBuildProgram", err)
	err = registerLibFuncWithoutPanic(&GetProgramBuildInfo, handle, "clGetProgramBuildInfo", err)
	err = registerLibFuncWithoutPanic(&GetProgramInfo, handle, "clGetProgramInfo", err)
	err = registerLibFuncWithoutPanic(&CreateKernel, handle, "clCreateKernel", err)
	err = registerLibFuncWithoutPanic(&ReleaseProgram, handle, "clReleaseProgram", err)
	// Kernel
	err = registerLibFuncWithoutPanic(&SetKernelArg, handle, "clSetKernelArg", err)
	err = registerLibFuncWithoutPanic(&ReleaseKernel, handle, "clReleaseKernel", err)
	// Buffer
	err = registerLibFuncWithoutPanic(&GetMemObjectInfo, handle, "clGetMemObjectInfo", err)
	err = registerLibFuncWithoutPanic(&ReleaseMemObject, handle, "clReleaseMemObject", err)
	return nil
}

func InitializeGLSharing() error {
	handle, err := loadLibrary()
	if err != nil {
		return err
	}
	// GL
	err = registerLibFuncWithoutPanic(&CreateFromGLTexture, handle, "clCreateFromGLTexture", err)
	err = registerLibFuncWithoutPanic(&EnqueueAcquireGLObjects, handle, "clEnqueueAcquireGLObjects", err)
	err = registerLibFuncWithoutPanic(&EnqueueReleaseGLObjects, handle, "clEnqueueReleaseGLObjects", err)
	err = registerLibFuncWithoutPanic(&GetGLObjectInfo, handle, "clGetGLObjectInfo", err)
	err = registerLibFuncWithoutPanic(&GetGLTextureInfo, handle, "clGetGLTextureInfo", err)

	return err
}
