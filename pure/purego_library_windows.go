//go:build windows

package pure

import (
	"errors"
	"syscall"
	"unsafe"
)

func loadLibrary() (uintptr, error) {
	handle, err := syscall.LoadLibrary("opencl.dll")
	if err != nil {
		return 0, err
	}
	return uintptr(handle), err
}

func initUnsupported(handle syscall.Handle, errIn error) error {
	// purego unsupported functions
	dll := syscall.DLL{
		Name:   "opencl.dll",
		Handle: handle,
	}

	// Note: Functions with unsupported arguments requiring syscall loading
	readImg, err := dll.FindProc("clEnqueueReadImage")
	if err != nil {
		errIn = errors.Join(err)
	}
	mapImg, err := dll.FindProc("clEnqueueMapImage")
	if err != nil {
		errIn = errors.Join(err)
	}
	mapBuffer, err := dll.FindProc("clEnqueueMapBuffer")
	if err != nil {
		errIn = errors.Join(err)
	}
	writeImg, err := dll.FindProc("clEnqueueWriteImage")
	if err != nil {
		errIn = errors.Join(err)
	}
	EnqueueReadImage = func(queue CommandQueue, image Buffer, blockingRead bool, origin, region [3]Size, row_pitch, slice_pitch Size, ptr unsafe.Pointer, numEventsWaitList uint, eventWaitList []Event, event *Event) Status {
		block := uintptr(0)
		if blockingRead {
			block = 1
		}
		r1, _, _ := readImg.Call(
			uintptr(queue),
			uintptr(image),
			uintptr(block),
			uintptr(unsafe.Pointer(&origin[0])),
			uintptr(unsafe.Pointer(&region[0])),
			uintptr(row_pitch),
			uintptr(slice_pitch),
			uintptr(ptr),
			uintptr(numEventsWaitList),
			uintptr(0), // TODO: eventWaitList if non-nil
			uintptr(unsafe.Pointer(event)),
		)
		return Status(r1)
	}
	EnqueueMapImage = func(queue CommandQueue, image Buffer, blockingMap bool, mapFlags MapFlag, origin, region [3]Size, imageRowPitch, imageSlicePitch *Size, numEventsWaitList uint, eventWaitList []Event, event *Event, errCodeRet *Status) uintptr {
		block := uintptr(0)
		if blockingMap {
			block = 1
		}
		r1, _, _ := mapImg.Call(
			uintptr(queue),
			uintptr(image),
			uintptr(block),
			uintptr(mapFlags),
			uintptr(unsafe.Pointer(&origin[:][0])),
			uintptr(unsafe.Pointer(&region[:][0])),
			uintptr(unsafe.Pointer(imageRowPitch)),
			uintptr(unsafe.Pointer(imageSlicePitch)),
			uintptr(numEventsWaitList),
			uintptr(0), // TODO: eventWaitList if non-nil
			uintptr(unsafe.Pointer(event)),
			uintptr(unsafe.Pointer(errCodeRet)),
		)

		return r1
	}
	EnqueueMapBuffer = func(queue CommandQueue, buffer Buffer, blockingMap bool, mapFlags MapFlag, offset, size Size, numEventsWaitList uint, eventWaitList []Event, event *Event, errCodeRet *Status) uintptr {
		block := uintptr(0)
		if blockingMap {
			block = 1
		}
		r1, _, _ := mapBuffer.Call(
			uintptr(queue),
			uintptr(buffer),
			uintptr(block),
			uintptr(mapFlags),
			uintptr(offset),
			uintptr(size),
			uintptr(numEventsWaitList),
			uintptr(0), // TODO: eventWaitList if non-nil
			uintptr(unsafe.Pointer(event)),
			uintptr(unsafe.Pointer(errCodeRet)),
		)

		return r1
	}
	EnqueueWriteImage = func(queue CommandQueue, image Buffer, blockingRead bool, origin, region [3]Size, row_pitch, slice_pitch Size, ptr unsafe.Pointer, numEventsWaitList uint, eventWaitList []Event, event *Event) Status {
		block := uintptr(0)
		if blockingRead {
			block = 1
		}
		r1, _, _ := writeImg.Call(
			uintptr(queue),
			uintptr(image),
			uintptr(block),
			uintptr(unsafe.Pointer(&origin[0])),
			uintptr(unsafe.Pointer(&region[0])),
			uintptr(row_pitch),
			uintptr(slice_pitch),
			uintptr(ptr),
			uintptr(numEventsWaitList),
			uintptr(0), // TODO: eventWaitList if non-nil
			uintptr(unsafe.Pointer(event)),
		)
		return Status(r1)
	}
	return errIn
}
