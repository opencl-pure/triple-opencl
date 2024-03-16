package high

import (
	"errors"
	"github.com/opencl-pure/triple-opencl/constants"
	"github.com/opencl-pure/triple-opencl/pure"
	"runtime"
	"unsafe"
)

// Device the only needed entrence for the BlackCL
// represents the device on which memory can be allocated and kernels run
// it abstracts away all the complexity of contexts/platforms/queues
type Device struct {
	id       []pure.Device
	ctx      pure.Context
	queue    pure.CommandQueue
	programs []pure.Program // only one
	platform *Platform
}

// Release releases the device
func (d *Device) Release() error {
	var result error
	for _, p := range d.programs {
		if err := pure.StatusToErr(pure.ReleaseProgram(p)); err != nil {
			result = pure.ErrJoin(result, err)
		}
	}
	if err := pure.StatusToErr(pure.ReleaseCommandQueue(d.queue)); err != nil {
		result = pure.ErrJoin(result, err)
	}
	if err := pure.StatusToErr(pure.ReleaseContext(d.ctx)); err != nil {
		result = pure.ErrJoin(result, err)
	}
	return pure.ErrJoin(result, pure.StatusToErr(pure.ReleaseDevice(d.id[0])))
}

func (d *Device) GetInfoString(param pure.DeviceInfo) (string, error) {
	strC := make([]byte, 1024)
	var strN pure.Size
	err := pure.StatusToErr(pure.GetDeviceInfo(d.id[0], param, 1024, strC, &strN))
	if err != nil {
		return "", err
	}
	return string(strC[:int(strN)]), nil
}

func (d *Device) String() (string, error) {
	name, err := d.Name()
	vendor, err2 := d.Vendor()
	return name + " " + vendor, pure.ErrJoin(err, err2)
}

// Name device info - name
func (d *Device) Name() (string, error) {
	return d.GetInfoString(constants.CL_DEVICE_NAME)
}

// Vendor device info - vendor
func (d *Device) Vendor() (string, error) {
	return d.GetInfoString(constants.CL_DEVICE_VENDOR)
}

// Extensions device info - extensions
func (d *Device) Extensions() (string, error) {
	return d.GetInfoString(constants.CL_DEVICE_EXTENSIONS)
}

// OpenCLCVersion device info - OpenCL C Version
func (d *Device) OpenCLCVersion() (string, error) {
	return d.GetInfoString(constants.CL_DEVICE_OPENCL_C_VERSION)
}

// Profile device info - profile
func (d *Device) Profile() (string, error) {
	return d.GetInfoString(constants.CL_DEVICE_PROFILE)
}

// Version device info - version
func (d *Device) Version() (string, error) {
	return d.GetInfoString(constants.CL_DEVICE_VERSION)
}

// DriverVersion device info - driver version
func (d *Device) DriverVersion() (string, error) {
	return d.GetInfoString(constants.CL_DRIVER_VERSION)
}

func (d *Device) PlatformName() (string, error) {
	return d.platform.GetName()
}

func (d *Device) PlatformProfile() (string, error) {
	return d.platform.GetProfile()
}

func (d *Device) PlatformOpenCLCVersion() (string, error) {
	return d.platform.GetVersion()
}

func (d *Device) PlatformDriverVersion() (string, error) {
	return d.platform.GetVersion()
}

func (d *Device) PlatformVendor() (string, error) {
	return d.platform.GetVendor()
}

func (d *Device) PlatformExtensions() ([]pure.Extension, error) {
	return d.platform.GetExtensions()
}

// AddProgram copiles program source
func (d *Device) AddProgram(source string) (*Program, error) {
	defer runtime.KeepAlive(source)
	var ret pure.Status
	p := pure.CreateProgramWithSource(d.ctx, 1, []string{source}, nil, &ret)
	err := pure.StatusToErr(ret)
	if err != nil {
		panic(err)
	}
	ret = pure.BuildProgram(p, 1, d.id, []byte(""), nil, nil)
	if ret != constants.CL_SUCCESS {
		if ret == constants.CL_BUILD_PROGRAM_FAILURE {
			var n pure.Size
			pure.GetProgramBuildInfo(p, d.id[0], constants.CL_PROGRAM_BUILD_LOG, 0, nil, &n)
			log := make([]byte, int(n))
			pure.GetProgramBuildInfo(p, d.id[0], constants.CL_PROGRAM_BUILD_LOG, n, unsafe.Pointer(&log[0]), nil)
			return nil, errors.New(string(log))
		}
		return nil, pure.StatusToErr(ret)
	}
	d.programs = append(d.programs, p)
	return &Program{program: p}, nil
}
