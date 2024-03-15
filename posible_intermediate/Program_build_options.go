package posible_intermediate

import (
	"github.com/opencl-pure/triple-opencl/pure"
	"strings"
)

type ProgramBuildOptions struct {
	// Preprocessor options
	Warnings          bool
	Macros            map[string]string
	DirectoryIncludes []string
	Version           pure.Version
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
