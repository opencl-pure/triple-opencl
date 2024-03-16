package middle

import (
	"github.com/opencl-pure/triple-opencl/v1/constants"
	"github.com/opencl-pure/triple-opencl/v1/pure"
	"strings"
)

type Platform struct {
	P pure.Platform
}

func GetPlatforms() ([]Platform, error) {
	numPlatforms := uint32(0)
	if pure.GetPlatformIDs == nil {
		return nil, pure.Uninitialized("GetPlatformIDs")
	}
	st := pure.GetPlatformIDs(0, nil, &numPlatforms)
	if st != constants.CL_SUCCESS {
		return nil, pure.StatusToErr(st)
	}

	platformIDs := make([]pure.Platform, numPlatforms)
	st = pure.GetPlatformIDs(numPlatforms, platformIDs, nil)
	if st != constants.CL_SUCCESS {
		return nil, pure.StatusToErr(st)
	}
	res := make([]Platform, numPlatforms)
	for i := uint32(0); i < numPlatforms; i++ {
		res[i] = Platform{P: platformIDs[i]}
	}
	return res, nil
}

func (p *Platform) getInfo(name pure.PlatformInfo) (string, error) {
	size := pure.Size(0)
	st := pure.GetPlatformInfo(p.P, name, pure.Size(0), nil, &size)
	if st != constants.CL_SUCCESS {
		return "", pure.StatusToErr(st)
	}

	info := make([]byte, size)
	st = pure.GetPlatformInfo(p.P, name, size, info, nil)
	if st != constants.CL_SUCCESS {
		return "", pure.StatusToErr(st)
	}
	return string(info), nil
}
func (p *Platform) GetProfile() (string, error) {
	return p.getInfo(constants.CL_PLATFORM_PROFILE)
}
func (p *Platform) GetVersion() (string, error) {
	return p.getInfo(constants.CL_PLATFORM_VERSION)
}
func (p *Platform) GetName() (string, error) {
	return p.getInfo(constants.CL_PLATFORM_NAME)
}
func (p *Platform) GetVendor() (string, error) {
	return p.getInfo(constants.CL_PLATFORM_VENDOR)
}
func (p *Platform) GetExtensions() ([]pure.Extension, error) {
	extensions, err := p.getInfo(constants.CL_PLATFORM_EXTENSIONS)
	if err != nil {
		return nil, err
	}
	return strings.Split(extensions, " "), nil
}

func (p *Platform) GetDevices(deviceType pure.DeviceType) ([]Device, error) {
	numDevices := uint32(0)
	if pure.GetDeviceIDs == nil {
		return nil, pure.Uninitialized("GetDeviceIDs")
	}
	st := pure.GetDeviceIDs(p.P, deviceType, 0, nil, &numDevices)
	if st != constants.CL_SUCCESS {
		return nil, pure.StatusToErr(st)
	}
	deviceIDs := make([]pure.Device, numDevices)
	st = pure.GetDeviceIDs(p.P, deviceType, numDevices, deviceIDs, nil)
	if st != constants.CL_SUCCESS {
		return nil, pure.StatusToErr(st)
	}
	res := make([]Device, numDevices)
	for i := uint32(0); i < numDevices; i++ {
		res[i] = Device{D: deviceIDs[i]}
	}
	return res, nil
}
