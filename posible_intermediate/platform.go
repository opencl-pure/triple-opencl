package posible_intermediate

// #include "opencl.h"
import "C"
import (
	"github.com/opencl-pure/triple-opencl/constants"
	"github.com/opencl-pure/triple-opencl/pure"
	"strings"
)

// PlatformInfo is a type of info that can be retrieved by Platform.GetInfo.
type PlatformInfo pure.PlatformInfo

// PlatformInfo constants.
const (
	PlatformProfile    PlatformInfo = PlatformInfo(constants.CL_PLATFORM_PROFILE)
	PlatformVersion                 = PlatformInfo(constants.CL_PLATFORM_VERSION)
	PlatformName                    = PlatformInfo(constants.CL_PLATFORM_NAME)
	PlatformVendor                  = PlatformInfo(constants.CL_PLATFORM_VENDOR)
	PlatformExtensions              = PlatformInfo(constants.CL_PLATFORM_EXTENSIONS)
)

// Platform is a structure for an OpenCL platform.
type Platform struct {
	platformID pure.Platform
	version    MajorMinor
}

// GetPlatforms returns a slice containing all platforms available.
func GetPlatforms() ([]Platform, error) {
	var platformCount = uint32(0)
	if pure.GetPlatformIDs == nil {
		return nil, pure.Uninitialized("GetPlatformIDs")
	}
	errInt := pure.GetPlatformIDs(0, nil, &platformCount)
	if errInt != constants.CL_SUCCESS {
		return nil, pure.StatusToErr(errInt)
	}
	platformIDs := make([]pure.Platform, platformCount)
	errInt = pure.GetPlatformIDs(platformCount, platformIDs, nil)
	if errInt != constants.CL_SUCCESS {
		return nil, pure.StatusToErr(errInt)
	}
	platforms := make([]Platform, len(platformIDs))
	for i, platformID := range platformIDs {
		platforms[i] = Platform{
			platformID: platformID,
		}
		if err := platforms[i].GetInfo(PlatformVersion, &platforms[i].version); err != nil {
			return nil, err
		}
	}
	return platforms, nil
}

// GetInfo retrieves the information specified by name and stores it in output.
// The output must correspond to the return type for that type of info:
//
// PlatformProfile *string
// PlatformVersion *string or *PlatformMajorMinor
// PlatformName *string
// PlatformVendor *string
// PlatformExtensions *[]string or *string
// PlatformICDSuffixKHR *string
//
// Note that if PlatformExtensions is retrieved with output being a *string,
// the extensions will be a space-separated list as specified by the OpenCL
// reference for clGetPlatformInfo.
func (p *Platform) GetInfo(name PlatformInfo, output interface{}) error {
	var size pure.Size
	errInt := pure.GetPlatformInfo(p.platformID, pure.PlatformInfo(name), 0, nil, &size)
	if errInt != constants.CL_SUCCESS {
		return pure.StatusToErr(errInt)
	}
	if size == 0 {
		outputStr, _ := output.(*string)
		*outputStr = ""
		return nil
	}
	info := make([]byte, size)
	errInt = pure.GetPlatformInfo(p.platformID, pure.PlatformInfo(name),
		size, info, nil)
	if errInt != constants.CL_SUCCESS {
		return pure.StatusToErr(errInt)
	}
	outputString := zeroTerminatedByteSliceToString(info)
	switch t := output.(type) {
	case *string:
		*t = outputString
	case *MajorMinor:
		if name != PlatformVersion {
			return UnexpectedType
		}

		ver, errVer := parseVersion(outputString)
		if errVer != nil {
			return errVer
		}
		*t = *ver
	case *[]string:
		if name != PlatformExtensions {
			return UnexpectedType
		}
		elems := strings.Split(outputString, " ")
		*t = elems
	default:
		return UnexpectedType
	}
	return nil
}

// GetDevices returns a slice of devices of type deviceType for a Platform. If there are
// no such devices it returns an empty slice.
func (p *Platform) GetDevices(deviceType DeviceType) ([]Device, error) {
	return getDevices(p, deviceType)
}

// GetVersion returns the platform OpenCL version.
func (p *Platform) GetVersion() MajorMinor {
	return p.version
}

// parseVersion is a helper function to parse an OpenCL version. The version format
// is given by the specification to be:
//
// OpenCL<space><major_version.minor_version><space><platform-specific information>
//
// The only part that concerns us here is the major/minor version combination.
func parseVersion(ver string) (*MajorMinor, error) {
	elems := strings.SplitN(ver, " ", 3)
	if len(elems) < 3 || elems[0] != "OpenCL" {
		return nil, ErrorParsingVersion
	}

	return ParseMajorMinor(elems[1])
}
