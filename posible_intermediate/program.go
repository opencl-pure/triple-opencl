package posible_intermediate

// #include "opencl.h"
import "C"
import (
	"github.com/opencl-pure/triple-opencl/constants"
	"github.com/opencl-pure/triple-opencl/pure"
	"unsafe"
)

type Program struct {
	program pure.Program
}

func createProgramWithSource(context *Context, programCode string) (*Program, error) {
	var errInt pure.Status
	if pure.CreateProgramWithSource == nil {
		return nil, pure.Uninitialized("CreateProgramWithSource")
	}
	program := pure.CreateProgramWithSource(context.context, 1, []string{programCode}, nil, &errInt)
	if errInt != constants.CL_SUCCESS {
		return nil, pure.StatusToErr(errInt)
	}
	return &Program{program}, nil
}

func (p *Program) Build(device *Device) (string, error) {
	return p.BuildOptions(device, nil)
}

func (p *Program) BuildOptions(device *Device, opts *ProgramBuildOptions) (string, error) {
	var err error
	if pure.BuildProgram == nil {
		return "", pure.Uninitialized("BuildProgram")
	}
	if pure.GetProgramBuildInfo == nil {
		return "", pure.Uninitialized("GetProgramBuildInfo")
	}
	st := pure.BuildProgram(
		p.program, 1, []pure.Device{device.deviceID}, []byte(opts.String()), nil, nil)
	if st != constants.CL_SUCCESS {
		err = pure.StatusToErr(st)
	}
	var logsSize pure.Size
	st = pure.GetProgramBuildInfo(
		p.program, device.deviceID, constants.CL_PROGRAM_BUILD_LOG, 0, nil, &logsSize,
	)
	if st != constants.CL_SUCCESS {
		return "", pure.StatusToErr(st)
	}

	var logs = make([]byte, logsSize)
	st = pure.GetProgramBuildInfo(
		p.program, device.deviceID, constants.CL_PROGRAM_BUILD_LOG, logsSize, unsafe.Pointer(&logs[0]), nil,
	)
	if st != constants.CL_SUCCESS {
		return "", pure.StatusToErr(st)
	}

	return string(logs), err
}

func (p *Program) Release() error {
	if pure.ReleaseProgram == nil {
		return pure.Uninitialized("ReleaseProgram")
	}
	return pure.StatusToErr(pure.ReleaseProgram(p.program))
}

func (p *Program) CreateKernel(kernelName string) (*Kernel, error) {
	return createKernel(p, kernelName)
}
