package high

import "C"
import (
	"errors"
	"opencl-pure/opencl/constants"
	"opencl-pure/opencl/pure"
)

var (
	//ErrUnknown Generally an unexpected result from an OpenCL function (e.g. CL_SUCCESS but null pointer)
	ErrUnknown = errors.New("cl: unknown error")
)

// GetDefaultDevice ...
func GetDefaultDevice() (*Device, error) {
	id := make([]pure.Device, 1)
	err := pure.StatusToErr(
		pure.GetDeviceIDs(0, pure.DeviceType(constants.CL_DEVICE_TYPE_DEFAULT), 1, id, nil))
	if err != nil {
		return nil, err
	}
	return newDevice(id)
}

// GetDevices returns all devices of all platforms with specified type
func GetDevices(deviceType pure.DeviceType) ([]*Device, error) {
	platformIds, err := getPlatforms()
	if err != nil {
		return nil, err
	}
	var devices []*Device
	for _, p := range platformIds {
		var n uint32
		err = pure.StatusToErr(pure.GetDeviceIDs(p, deviceType, 0, nil, &n))
		if err != nil {
			return nil, err
		}
		deviceIds := make([]pure.Device, int(n))
		err = pure.StatusToErr(pure.GetDeviceIDs(p, deviceType, n, deviceIds, nil))
		if err != nil {
			return nil, err
		}
		for _, d := range deviceIds {
			device, err := newDevice([]pure.Device{d})
			if err != nil {
				return nil, err
			}
			devices = append(devices, device)
		}
	}
	return devices, nil
}

func getPlatforms() ([]pure.Platform, error) {
	var n uint32
	if err := pure.StatusToErr(pure.GetPlatformIDs(0, nil, &n)); err != nil {
		return nil, err
	}
	if n == 0 {
		return nil, errors.New("unknown error no available platform")
	}
	platformIds := make([]pure.Platform, int(n))
	if err := pure.StatusToErr(pure.GetPlatformIDs(n, platformIds, nil)); err != nil {
		return nil, err
	}
	return platformIds, nil
}

func newDevice(id []pure.Device) (*Device, error) {
	d := &Device{id: id}
	var ret pure.Status
	d.ctx = pure.CreateContext(nil, 1, id, nil, nil, &ret)
	err := pure.StatusToErr(ret)
	if err != nil {
		return nil, err
	}
	if d.ctx == pure.Context(0) {
		return nil, ErrUnknown
	}
	if pure.CreateCommandQueueWithProperties != nil {
		d.queue = pure.CreateCommandQueueWithProperties(d.ctx, d.id[0], 0, &ret)
	} else {
		d.queue = pure.CreateCommandQueue(d.ctx, d.id[0], 0, &ret)
	}
	if err = pure.StatusToErr(ret); err != nil {
		return nil, err
	}
	return d, nil
}

func Init(version int) (e error) {
	return pure.Init(version)
}
