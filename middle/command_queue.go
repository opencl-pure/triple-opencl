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
	return pure.StatusToErr(pure.EnqueueBarrier(cq.C))
}

func (cq *CommandQueue) EnqueueNDRangeKernel(kernel *Kernel, workDim uint, globalOffsets, globalWorkSizes, localWorkSizes []uint64) error {
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
	return pure.StatusToErr(pure.EnqueueReadBuffer(
		cq.C, buffer.B, blockingRead, 0, pure.Size(data.DataSize), data.Pointer, 0, nil, nil,
	))
}

func (cq *CommandQueue) EnqueueReadImage(image *Image, blockingRead bool, data *pure.ImageData) error {
	origin := [3]pure.Size{pure.Size(data.Origin[0]), pure.Size(data.Origin[1]), pure.Size(data.Origin[2])}
	region := [3]pure.Size{pure.Size(data.Region[0]), pure.Size(data.Region[1]), pure.Size(data.Region[2])}
	return pure.StatusToErr(pure.EnqueueReadImage(
		cq.C, image.B, blockingRead, origin, region, 0, 0, data.Pointer, 0, nil, nil,
	))
}

func (cq *CommandQueue) EnqueueWriteImage(image *Image, blockingRead bool, data *pure.ImageData) error {
	origin := [3]pure.Size{pure.Size(data.Origin[0]), pure.Size(data.Origin[1]), pure.Size(data.Origin[2])}
	region := [3]pure.Size{pure.Size(data.Region[0]), pure.Size(data.Region[1]), pure.Size(data.Region[2])}
	return pure.StatusToErr(pure.EnqueueWriteImage(
		cq.C, image.B, blockingRead, origin, region, 0, 0, data.Pointer, 0, nil, nil,
	))
}

func (cq *CommandQueue) EnqueueWriteBuffer(buffer *Buffer, blockingWrite bool, data *pure.BufferData) error {
	return pure.StatusToErr(pure.EnqueueWriteBuffer(
		cq.C, buffer.B, blockingWrite, 0, pure.Size(data.DataSize), data.Pointer, 0, nil, nil,
	))
}

func (cq *CommandQueue) EnqueueMapBuffer(buffer Buffer, blockingMap bool, flags []pure.MapFlag, data *pure.BufferData) error {
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
	return pure.StatusToErr(pure.EnqueueUnmapMemObject(
		cq.C, buffer.B, data.Pointer, 0, nil, nil,
	))
}

func (cq *CommandQueue) Finish() error {
	return pure.StatusToErr(pure.FinishCommandQueue(cq.C))
}

func (cq *CommandQueue) Flush() error {
	return pure.StatusToErr(pure.FlushCommandQueue(cq.C))
}

func (cq *CommandQueue) Release() error {
	return pure.StatusToErr(pure.ReleaseCommandQueue(cq.C))
}

// GL

func (cq *CommandQueue) EnqueueAcquireGLObjects(objects []Buffer) error {
	obsCL := make([]pure.Buffer, len(objects))
	for i := 0; i < len(obsCL); i++ {
		obsCL[i] = objects[i].B
	}
	return pure.StatusToErr(pure.EnqueueAcquireGLObjects(
		cq.C, uint32(len(objects)), unsafe.Pointer(&obsCL[0]), 0, nil, nil,
	))
}

func (cq *CommandQueue) EnqueueReleaseGLObjects(objects []Buffer) error {
	obsCL := make([]pure.Buffer, len(objects))
	for i := 0; i < len(obsCL); i++ {
		obsCL[i] = objects[i].B
	}
	return pure.StatusToErr(pure.EnqueueReleaseGLObjects(
		cq.C, uint32(len(objects)), unsafe.Pointer(&obsCL[0]), 0, nil, nil,
	))
}
