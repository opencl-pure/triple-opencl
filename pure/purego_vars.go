package pure

import "unsafe"

// functions
var (

	// Platform
	GetPlatformIDs  func(numEntries uint32, platforms []Platform, numPlatforms *uint32) Status                                                 = nil
	GetPlatformInfo func(platform Platform, platformInfo PlatformInfo, paramValueSize Size, paramValue []byte, paramValueSizeRet *Size) Status = nil
	// Device
	GetDeviceIDs  func(platform Platform, deviceType DeviceType, numEntries uint32, devices []Device, numDevices *uint32) Status     = nil
	GetDeviceInfo func(device Device, deviceInfo DeviceInfo, paramValueSize Size, paramValue []byte, paramValueSizeRet *Size) Status = nil
	ReleaseDevice func(id Device) Status                                                                                             = nil

	// Event
	ReleaseEvent  func(event Event) Status                        = nil
	WaitForEvents func(numEvents uint32, eventList *Event) Status = nil

	// Context
	CreateContext           func(properties unsafe.Pointer, numDevices uint32, devices []Device, pfnNotify *CreateContextNotifyFuncType, userData []byte, errCodeRet *Status) Context     = nil
	ReleaseContext          func(ctx Context) Status                                                                                                                                      = nil
	CreateProgramWithSource func(ctx Context, count Size, strings []string, lengths []Size, errCodeRet *Status) Program                                                                   = nil
	CreateBuffer            func(ctx Context, memFlags MemFlag, size Size, hostPtr unsafe.Pointer, errCodeRet *Status) Buffer                                                             = nil
	CreateImage2D           func(ctx Context, memFlags MemFlag, imageFormat *ImageFormat, imageWidth, imageHeight, imageRowPitch Size, hostPtr unsafe.Pointer, errCodeRet *Status) Buffer = nil
	// Command queue
	CreateCommandQueue               func(context Context, device Device, properties CommandQueueProperty, errCodeRet *Status) CommandQueue                                                                            = nil
	CreateCommandQueueWithProperties func(context Context, device Device, properties CommandQueueProperty, errCodeRet *Status) CommandQueue                                                                            = nil
	EnqueueBarrier                   func(queue CommandQueue) Status                                                                                                                                                   = nil
	EnqueueNDRangeKernel             func(queue CommandQueue, kernel Kernel, workDim uint, globalWorkOffset, globalWorkSize, localWorkSize []Size, numEventsWaitList uint, eventWaitList []Event, event *Event) Status = nil
	EnqueueReadBuffer                func(queue CommandQueue, buffer Buffer, blockingRead bool, offset, cb Size, ptr unsafe.Pointer, numEventsWaitList uint, eventWaitList []Event, event *Event) Status               = nil

	EnqueueReadImage  func(queue CommandQueue, image Buffer, blockingRead bool, origin, region [3]Size, row_pitch, slice_pitch Size, ptr unsafe.Pointer, numEventsWaitList uint, eventWaitList []Event, event *Event) Status = nil
	EnqueueWriteImage func(queue CommandQueue, image Buffer, blockingRead bool, origin, region [3]Size, row_pitch, slice_pitch Size, ptr unsafe.Pointer, numEventsWaitList uint, eventWaitList []Event, event *Event) Status = nil

	EnqueueWriteBuffer    func(queue CommandQueue, buffer Buffer, blockingWrite bool, offset, cb Size, ptr unsafe.Pointer, numEventsWaitList uint, eventWaitList []Event, event *Event) Status                                                              = nil
	EnqueueMapBuffer      func(queue CommandQueue, buffer Buffer, blockingMap bool, mapFlags MapFlag, offset, size Size, numEventsWaitList uint, eventWaitList []Event, event *Event, errCodeRet *Status) uintptr                                           = nil
	EnqueueUnmapMemObject func(queue CommandQueue, buffer Buffer, mappedPtr unsafe.Pointer, numEventsWaitList uint, eventWaitList []Event, event *Event) Status                                                                                             = nil
	EnqueueMapImage       func(queue CommandQueue, image Buffer, blockingMap bool, mapFlags MapFlag, origin, region [3]Size, imageRowPitch, imageSlicePitch *Size, numEventsWaitList uint, eventWaitList []Event, event *Event, errCodeRet *Status) uintptr = nil
	FinishCommandQueue    func(queue CommandQueue) Status                                                                                                                                                                                                   = nil
	FlushCommandQueue     func(queue CommandQueue) Status                                                                                                                                                                                                   = nil
	ReleaseCommandQueue   func(queue CommandQueue) Status                                                                                                                                                                                                   = nil
	// Program
	BuildProgram        func(program Program, numDevices uint32, devices []Device, options []byte, pfnNotify *BuildProgramNotifyFuncType, userData []byte) Status = nil
	GetProgramBuildInfo func(program Program, device Device, info ProgramBuildInfo, paramSize Size, paramValue unsafe.Pointer, paramSizeRet *Size) Status         = nil
	CreateKernel        func(program Program, kernelName string, errCodeRet *Status) Kernel                                                                       = nil
	ReleaseProgram      func(program Program) Status                                                                                                              = nil
	GetProgramInfo      func(program Program, info ProgramBuildInfo, size Size, pointer unsafe.Pointer, t *Size) Status                                           = nil
	// Kernel
	SetKernelArg  func(kernel Kernel, argIndex uint32, argSize Size, argValue unsafe.Pointer) Status = nil
	ReleaseKernel func(kernel Kernel) Status                                                         = nil
	// Buffer
	GetMemObjectInfo func(buffer Buffer, memInfo MemInfo, paramValueSize Size, paramValue unsafe.Pointer, paramValueSizeRet *Size) Status = nil
	ReleaseMemObject func(buffer Buffer) Status                                                                                           = nil

	// GL
	CreateFromGLTexture     func(ctx Context, memFlags MemFlag, textureTarget GLEnum, mipLevel GLInt, texture GLUint, errCodeRet *Status) Buffer                           = nil
	GetGLObjectInfo         func(memObj Buffer, objectType *CLGLObjectType, objectName *GLUint) Status                                                                     = nil
	GetGLTextureInfo        func(memObj Buffer, paramName CLGLTextureInfo, paramValueSize Size, paramValue unsafe.Pointer, paramValueSizeRet *Size) Status                 = nil
	EnqueueAcquireGLObjects func(queue CommandQueue, numObjects uint32, memObjects unsafe.Pointer, numEventsInWaitList uint32, eventWaitList []Event, event *Event) Status = nil
	EnqueueReleaseGLObjects func(queue CommandQueue, numObjects uint32, memObjects unsafe.Pointer, numEventsInWaitList uint32, eventWaitList []Event, event *Event) Status = nil
)
