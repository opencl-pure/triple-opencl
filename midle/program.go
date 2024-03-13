package midle

import (
	"github.com/opencl-pure/triple-opencl/constants"
	"opencl-pure/opencl/pure"
	"strings"
	"unsafe"
)

type Program struct {
	P pure.Program
}

type programBuildInfo uint32

type Version string

const (
	Version1_0 Version = "CL1.0"
	Version1_1 Version = "CL1.1"
	Version1_2 Version = "CL1.2"
	Version2_0 Version = "CL2.0"
	Version3_0 Version = "CL3.0"
)

type ProgramBuildOptions struct {
	// Preprocessor options
	Warnings          bool
	Macros            map[string]string
	DirectoryIncludes []string
	Version           Version
	// Math intrinsics options
	SinglePrecisionConstant bool
	MadEnable               bool
	NoSignedZeros           bool
	FastRelaxedMaths        bool
	// Extensions
	NvidiaVerbose bool
}

func (po *ProgramBuildOptions) String() string {
	if po == nil {
		return ""
	}

	var sb strings.Builder

	// Preprocessor
	if po.Warnings {
		sb.WriteString("-w")
		sb.WriteRune(' ')
	}
	if po.Version != "" {
		sb.WriteString("-cl-std=" + string(po.Version))
		sb.WriteRune(' ')
	}
	// Math intrinsics
	if po.SinglePrecisionConstant {
		sb.WriteString("-cl-single-precision-constant")
		sb.WriteRune(' ')
	}
	if po.MadEnable {
		sb.WriteString("-cl-mad-enable")
		sb.WriteRune(' ')
	}
	if po.NoSignedZeros {
		sb.WriteString("-cl-no-signed-zeros")
		sb.WriteRune(' ')
	}
	if po.FastRelaxedMaths {
		sb.WriteString("-cl-fast-relaxed-math")
		sb.WriteRune(' ')
	}
	// Extensions
	if po.NvidiaVerbose {
		sb.WriteString("-cl-nv-verbose")
		sb.WriteRune(' ')
	}

	return sb.String()
}

func (p *Program) Build(device *Device, opts *ProgramBuildOptions) (string, error) {
	var err error
	if pure.BuildProgram == nil {
		return "", pure.Uninitialized("BuildProgram")
	}
	if pure.GetProgramBuildInfo == nil {
		return "", pure.Uninitialized("GetProgramBuildInfo")
	}
	st := pure.BuildProgram(
		p.P, 1, []pure.Device{device.D}, []byte(opts.String()), nil, nil,
	)
	if st != constants.CL_SUCCESS {
		err = pure.StatusToErr(st)
	}

	var logsSize pure.Size
	st = pure.GetProgramBuildInfo(
		p.P, device.D, constants.CL_PROGRAM_BUILD_LOG, 0, nil, &logsSize,
	)
	if st != constants.CL_SUCCESS {
		return "", pure.StatusToErr(st)
	}

	var logs = make([]byte, logsSize)
	st = pure.GetProgramBuildInfo(
		p.P, device.D, constants.CL_PROGRAM_BUILD_LOG, logsSize, unsafe.Pointer(&logs[0]), nil,
	)
	if st != constants.CL_SUCCESS {
		return "", pure.StatusToErr(st)
	}

	return string(logs), err
}

func (p *Program) CreateKernel(name string) (*Kernel, error) {
	var st pure.Status
	if pure.CreateKernel == nil {
		return nil, pure.Uninitialized("CreateKernel")
	}
	kernel := pure.CreateKernel(p.P, name, &st)
	if st != constants.CL_SUCCESS {
		return nil, pure.StatusToErr(st)
	}
	return &Kernel{K: kernel}, nil
}

func (p *Program) Release() error {
	if pure.ReleaseProgram == nil {
		return pure.Uninitialized("ReleaseProgram")
	}
	return pure.StatusToErr(pure.ReleaseProgram(p.P))
}
