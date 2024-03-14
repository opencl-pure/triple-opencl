package middle

import (
	"errors"
	"github.com/opencl-pure/triple-opencl/pure"
)

type Kernel struct {
	K pure.Kernel
}

type KernelArg struct {
	KA pure.KernelArg
}

func NewKernelArg(arg interface{}) (*KernelArg, error) {
	switch val := arg.(type) {
	case uint8:
		return &KernelArg{KA: pure.NewKernelArg(&val)}, nil
	case int8:
		return &KernelArg{KA: pure.NewKernelArg(&val)}, nil
	case uint16:
		return &KernelArg{KA: pure.NewKernelArg(&val)}, nil
	case int16:
		return &KernelArg{KA: pure.NewKernelArg(&val)}, nil
	case uint32:
		return &KernelArg{KA: pure.NewKernelArg(&val)}, nil
	case int32:
		return &KernelArg{KA: pure.NewKernelArg(&val)}, nil
	case float32:
		return &KernelArg{KA: pure.NewKernelArg(&val)}, nil
	case uint64:
		return &KernelArg{KA: pure.NewKernelArg(&val)}, nil
	case int64:
		return &KernelArg{KA: pure.NewKernelArg(&val)}, nil
	case float64:
		return &KernelArg{KA: pure.NewKernelArg(&val)}, nil
	case *Buffer:
		return &KernelArg{KA: pure.NewKernelArg(&val.B)}, nil
	default:
		return nil, errors.New("Unsuported arg")
	}
}
func newKernelArgBuffer(b *Buffer) *KernelArg {
	return &KernelArg{KA: pure.NewKernelArg[pure.Buffer](&b.B)}
}

func (k *Kernel) SetArg(index uint, arg *KernelArg) error {
	return k.K.SetArg(index, arg.KA)
}

func (k *Kernel) Release() error {
	return k.K.Release()
}
