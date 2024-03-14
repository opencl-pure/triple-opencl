package posible_intermediate

// #include "opencl.h"
import "C"
import (
	"errors"
	"github.com/opencl-pure/triple-opencl/constants"
	"github.com/opencl-pure/triple-opencl/pure"
	"reflect"
	"strings"
	"unsafe"
)

// DeviceType is a type of OpenCL device.
type DeviceType pure.DeviceType

// DeviceType constants.
const (
	DeviceTypeDefault     DeviceType = constants.CL_DEVICE_TYPE_DEFAULT
	DeviceTypeCPU         DeviceType = constants.CL_DEVICE_TYPE_CPU
	DeviceTypeGPU         DeviceType = constants.CL_DEVICE_TYPE_GPU
	DeviceTypeAccelerator DeviceType = constants.CL_DEVICE_TYPE_ACCELERATOR
	DeviceTypeCustom      DeviceType = constants.CL_DEVICE_TYPE_CUSTOM
	DeviceTypeAll         DeviceType = constants.CL_DEVICE_TYPE_ALL
)

// DeviceInfo is a type of info that can be retrieved by Device.GetInfo.
type DeviceInfo pure.DeviceInfo

// DeviceInfo constants.
const (
	DeviceAddressBits       DeviceInfo = DeviceInfo(constants.CL_DEVICE_ADDRESS_BITS)
	DeviceAvailable                    = DeviceInfo(constants.CL_DEVICE_AVAILABLE)
	DeviceBuiltInKernels               = DeviceInfo(constants.CL_DEVICE_BUILT_IN_KERNELS)
	DeviceCompilerAvailable            = DeviceInfo(constants.CL_DEVICE_COMPILER_AVAILABLE)
	DeviceInfoType                     = DeviceInfo(constants.CL_DEVICE_TYPE)
	DeviceVendor                       = DeviceInfo(constants.CL_DEVICE_VENDOR)
	DriverVersion                      = DeviceInfo(constants.CL_DRIVER_VERSION)
)

var (
	deviceInfoTypes = map[DeviceInfo][]interface{}{
		DeviceAddressBits:       {uint32(0)},
		DeviceAvailable:         {false},
		DeviceBuiltInKernels:    {"", []string{}},
		DeviceCompilerAvailable: {false},
		DeviceInfoType:          {DeviceTypeDefault},
		DeviceVendor:            {""},
		DriverVersion:           {"", MajorMinor{}},
	}
)

// Device is a structure for an OpenCL device.
type Device struct {
	deviceID pure.Device
}

// getDevices returns a slice of devices of type deviceType for platform.
func getDevices(platform *Platform, deviceType DeviceType) ([]Device, error) {
	var deviceCount uint32
	if pure.GetDeviceIDs == nil {
		return nil, pure.Uninitialized("GetDeviceIDs")
	}
	errInt := pure.GetDeviceIDs(platform.platformID, pure.DeviceType(deviceType), 0, nil, &deviceCount)
	if errInt != constants.CL_SUCCESS {
		return nil, pure.StatusToErr(errInt)
	}

	deviceIDs := make([]pure.Device, deviceCount)
	errInt = pure.GetDeviceIDs(platform.platformID, pure.DeviceType(deviceType), deviceCount, deviceIDs, nil)
	if errInt != constants.CL_SUCCESS {
		return nil, pure.StatusToErr(errInt)
	}
	devices := make([]Device, len(deviceIDs))
	for i, deviceID := range deviceIDs {
		devices[i] = Device{deviceID}
	}
	return devices, nil
}

// CreateContext creates and returns an OpenCL context for a device. After use the
// context must be released.
func (d *Device) CreateContext() (*Context, error) {
	return createContext(d)
}

// GetInfo retrieves the information specified by name and stores it in output.
// The output must correspond to the return type for that type of info:
//
// DeviceAddressBits       *bool
// DeviceAvailable         *bool
// DeviceBuiltInKernels    *string or *[]string
// DeviceCompilerAvailable *bool
// DeviceInfoType          *DeviceType
// DriverVersion           *string or *MajorMinor
//
// Note that if DeviceBuiltInKernels is retrieved with output being a *string,
// the extensions will be a semicolon-separated list as specified by the OpenCL
// reference for clGetDeviceInfo.
func (d *Device) GetInfo(name DeviceInfo, output interface{}) error {
	// Output interface needs to be a pointer to the expected type
	v := reflect.ValueOf(output)
	if v.Kind() != reflect.Ptr {
		return errors.New("UnexpectedType")
	}

	elem := v.Elem()

	validType := false
	for _, curType := range deviceInfoTypes[name] {
		if elem.Type() == reflect.TypeOf(curType) {
			validType = true
			break
		}
	}

	if !validType {
		return errors.New("UnexpectedType")
	}

	var outputString string

	switch t := output.(type) {
	case *string:
		if err := d.getInfoStr(name, &outputString); err != nil {
			return err
		}
		*t = outputString

	case *[]string:
		if err := d.getInfoStr(name, &outputString); err != nil {
			return err
		}
		if name == DeviceBuiltInKernels {
			elems := strings.Split(outputString, ";")
			*t = elems
		}

	case *uint32, *bool, *DeviceType:
		return d.getInfoNum(name, output)

	case *MajorMinor:
		if err := d.getInfoStr(name, &outputString); err != nil {
			return err
		}
		if name == DriverVersion {
			ver, err := ParseMajorMinor(outputString)
			if err != nil {
				return err
			}

			*t = *ver
		}
	}

	return nil
}

// getInfoNum returns numeric device info. This includes all types of which the size is known
// beforehand, such as cl_uint or cl_bool.
func (d Device) getInfoNum(name DeviceInfo, output interface{}) error {
	var errInt pure.Status
	switch t := output.(type) {
	case *uint32:
		var u uint32
		errInt = pure.GetDeviceInfo(
			d.deviceID,
			pure.DeviceInfo(name),
			4,
			unsafe.Pointer(&u),
			nil,
		)
		*t = u

	case *bool:
		var u uint32
		errInt = clError(C.clGetDeviceInfo(
			d.deviceID,
			C.cl_device_info(name),
			4,
			unsafe.Pointer(&u),
			nil,
		))
		*t = u == C.CL_TRUE

	case *DeviceType:
		var u C.cl_device_type
		errInt = clError(C.clGetDeviceInfo(
			d.deviceID,
			C.cl_device_info(name),
			C.sizeof_cl_device_type,
			unsafe.Pointer(&u),
			nil,
		))
		*t = DeviceType(u)

	}
	if errInt != constants.CL_SUCCESS {
		return pure.StatusToErr(errInt)
	}
	return nil
}

// getInfoStr returns string device info. For strings, the size is not known beforehand and has to be queried
// before getting the info.
func (d Device) getInfoStr(name DeviceInfo, output interface{}) error {
	outputStr, _ := output.(*string)

	var size uint64
	errInt := clError(C.clGetDeviceInfo(
		d.deviceID,
		C.cl_device_info(name),
		0,
		nil,
		(*C.size_t)(&size),
	))
	if errInt != clSuccess {
		return clErrorToError(errInt)
	}

	if size == 0 {
		*outputStr = ""
		return nil
	}

	info := make([]byte, size)
	errInt = clError(C.clGetDeviceInfo(
		d.deviceID,
		C.cl_device_info(name),
		C.size_t(size),
		unsafe.Pointer(&info[0]),
		nil,
	))
	if errInt != clSuccess {
		return clErrorToError(errInt)
	}

	*outputStr = zeroTerminatedByteSliceToString(info)

	return nil
}
