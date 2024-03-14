package middle

import (
	"github.com/opencl-pure/triple-opencl/constants"
	"github.com/opencl-pure/triple-opencl/pure"
	"unsafe"
)

type CommandQueue struct {
	C pure.CommandQueue
}

func (cq *CommandQueue) EnqueueBarrier() error {
	if pure.EnqueueBarrier == nil {
		return pure.Uninitialized("EnqueueBarrier")
	}
	return pure.StatusToErr(pure.EnqueueBarrier(cq.C))
}

func (cq *CommandQueue) EnqueueNDRangeKernel(kernel *Kernel, workDim uint, globalOffsets, globalWorkSizes, localWorkSizes []uint64) error {
	if pure.EnqueueNDRangeKernel == nil {
		return pure.Uninitialized("EnqueueNDRangeKernel")
	}
	var offsets, gsizes, lsizes []pure.Size
	if len(globalOffsets) > 0 {
		offsets = make([]pure.Size, len(globalOffsets))
		for i := range globalOffsets {
			offsets[i] = pure.Size(globalOffsets[i])
		}
	}
	if len(globalWorkSizes) > 0 {
		gsizes = make([]pure.Size, len(globalWorkSizes))
		for i := range globalWorkSizes {
			gsizes[i] = pure.Size(globalWorkSizes[i])
		}
	}
	if len(localWorkSizes) > 0 {
		lsizes = make([]pure.Size, len(localWorkSizes))
		for i := range localWorkSizes {
			lsizes[i] = pure.Size(localWorkSizes[i])
		}
	}
	return pure.StatusToErr(pure.EnqueueNDRangeKernel(
		cq.C, kernel.K, workDim, offsets, gsizes, lsizes, 0, nil, nil,
	))
}

func (cq *CommandQueue) EnqueueReadBuffer(buffer *Buffer, blockingRead bool, data *pure.BufferData) error {
	if pure.EnqueueReadBuffer == nil {
		return pure.Uninitialized("EnqueueReadBuffer")
	}
	return pure.StatusToErr(pure.EnqueueReadBuffer(
		cq.C, buffer.B, blockingRead, 0, pure.Size(data.DataSize), data.Pointer, 0, nil, nil,
	))
}

func (cq *CommandQueue) EnqueueReadImage(image Buffer, blockingRead bool, data *pure.ImageData) error {
	if pure.EnqueueReadImage == nil {
		return pure.Uninitialized("EnqueueReadImage")
	}
	origin := [3]pure.Size{pure.Size(data.Origin[0]), pure.Size(data.Origin[1]), pure.Size(data.Origin[2])}
	region := [3]pure.Size{pure.Size(data.Region[0]), pure.Size(data.Region[1]), pure.Size(data.Region[2])}
	return pure.StatusToErr(pure.EnqueueReadImage(
		cq.C, image.B, blockingRead, origin, region, 0, 0, data.Pointer, 0, nil, nil,
	))
}

func (cq *CommandQueue) EnqueueWriteBuffer(buffer Buffer, blockingWrite bool, data *pure.BufferData) error {
	if pure.EnqueueWriteBuffer == nil {
		return pure.Uninitialized("EnqueueWriteBuffer")
	}
	return pure.StatusToErr(pure.EnqueueWriteBuffer(
		cq.C, buffer.B, blockingWrite, 0, pure.Size(data.DataSize), data.Pointer, 0, nil, nil,
	))
}

func (cq *CommandQueue) EnqueueMapBuffer(buffer Buffer, blockingMap bool, flags []pure.MapFlag, data *pure.BufferData) error {
	if pure.EnqueueMapBuffer == nil {
		return pure.Uninitialized("EnqueueMapBuffer")
	}
	var st pure.Status
	mapFlags := pure.MapFlag(0)
	for _, f := range flags {
		mapFlags |= f
	}
	ptr := pure.EnqueueMapBuffer(
		cq.C, buffer.B, blockingMap, mapFlags, 0, pure.Size(data.DataSize), 0, nil, nil, &st,
	)
	if st != constants.CL_SUCCESS {
		return pure.StatusToErr(st)
	}
	data.Pointer = unsafe.Pointer(ptr)

	return nil
}

func (cq *CommandQueue) EnqueueMapImage(image Buffer, blockingMap bool, flags []pure.MapFlag, data *pure.ImageData) error {
	if pure.EnqueueMapImage == nil {
		return pure.Uninitialized("EnqueueMapImage")
	}
	var st pure.Status
	mapFlags := pure.MapFlag(0)
	for _, f := range flags {
		mapFlags |= f
	}
	origin := [3]pure.Size{pure.Size(data.Origin[0]), pure.Size(data.Origin[1]), pure.Size(data.Origin[2])}
	region := [3]pure.Size{pure.Size(data.Region[0]), pure.Size(data.Region[1]), pure.Size(data.Region[2])}
	rowpitch, slicepitch := (*pure.Size)(&data.RowPitch), (*pure.Size)(&data.SlicePitch)
	ptr := pure.EnqueueMapImage(
		cq.C, image.B, blockingMap, mapFlags, origin, region, rowpitch, slicepitch, 0, nil, nil, &st,
	)
	if st != constants.CL_SUCCESS {
		return pure.StatusToErr(st)
	}
	data.Pointer = unsafe.Pointer(ptr)
	return nil
}

func (cq *CommandQueue) EnqueueUnmapBuffer(buffer Buffer, data *pure.BufferData) error {
	if pure.EnqueueUnmapMemObject == nil {
		return pure.Uninitialized("EnqueueUnmapMemObject")
	}
	return pure.StatusToErr(pure.EnqueueUnmapMemObject(
		cq.C, buffer.B, data.Pointer, 0, nil, nil,
	))
}

func (cq *CommandQueue) Finish() error {
	if pure.FinishCommandQueue == nil {
		return pure.Uninitialized("FinishCommandQueue")
	}
	return pure.StatusToErr(pure.FinishCommandQueue(cq.C))
}

func (cq *CommandQueue) Flush() error {
	if pure.FlushCommandQueue == nil {
		return pure.Uninitialized("FlushCommandQueue")
	}
	return pure.StatusToErr(pure.FlushCommandQueue(cq.C))
}

func (cq *CommandQueue) Release() error {
	if pure.ReleaseCommandQueue == nil {
		return pure.Uninitialized("ReleaseCommandQueue")
	}
	return pure.StatusToErr(pure.ReleaseCommandQueue(cq.C))
}

// GL

func (cq CommandQueue) EnqueueAcquireGLObjects(objects []Buffer) error {
	if pure.EnqueueAcquireGLObjects == nil {
		return pure.Uninitialized("EnqueueAcquireGLObjects")
	}
	obsCL := make([]pure.Buffer, len(objects))
	for i := 0; i < len(obsCL); i++ {
		obsCL[i] = objects[i].B
	}
	return pure.StatusToErr(pure.EnqueueAcquireGLObjects(
		cq.C, uint32(len(objects)), unsafe.Pointer(&obsCL[0]), 0, nil, nil,
	))
}

func (cq CommandQueue) EnqueueReleaseGLObjects(objects []Buffer) error {
	if pure.EnqueueReleaseGLObjects == nil {
		return pure.Uninitialized("EnqueueReleaseGLObjects")
	}
	obsCL := make([]pure.Buffer, len(objects))
	for i := 0; i < len(obsCL); i++ {
		obsCL[i] = objects[i].B
	}
	return pure.StatusToErr(pure.EnqueueReleaseGLObjects(
		cq.C, uint32(len(objects)), unsafe.Pointer(&obsCL[0]), 0, nil, nil,
	))
}
