package midle

import (
	"opencl-pure/opencl/pure"
)

type Kernel struct {
	K pure.Kernel
}

type KernelArg struct {
	KA pure.KernelArg
}

func NewKernelArg[T any](arg *T) *KernelArg {
	return &KernelArg{KA: pure.NewKernelArg(arg)}
}

func (k *Kernel) SetArg(index uint, arg KernelArg) error {
	return k.K.SetArg(index, arg.KA)
}

func (k Kernel) Release() error {
	return k.K.Release()
}
