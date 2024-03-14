package midle

import (
	"github.com/opencl-pure/triple-opencl/constants"
	"github.com/opencl-pure/triple-opencl/pure"
	"strings"
)

type DeviceType uint32

type Device struct {
	D pure.Device
}

type deviceInfo uint32

func (d *Device) getInfo(name pure.DeviceInfo) (string, error) {
	size := pure.Size(0)
	st := pure.GetDeviceInfo(d.D, name, pure.Size(0), nil, &size)
	if st != constants.CL_SUCCESS {
		return "", pure.StatusToErr(st)
	}
	info := make([]byte, size)
	st = pure.GetDeviceInfo(d.D, name, size, info, nil)
	if st != constants.CL_SUCCESS {
		return "", pure.StatusToErr(st)
	}
	return string(info), nil
}

func (d *Device) GetExtensions() ([]pure.Extension, error) {
	extensions, err := d.getInfo(constants.CL_DEVICE_EXTENSIONS)
	if err != nil {
		return nil, err
	}
	return strings.Split(extensions, " "), nil
}
