package api

import (
	"errors"

	"github.com/NSenaud/opale/sensors/cpu"
	"github.com/NSenaud/opale/sensors/ram"
	"github.com/Sirupsen/logrus"
	"golang.org/x/net/context"
)

var log = logrus.New()

// Server represents the gRPC server
type Server struct{}

/*
 Cpu service
*/
func (s *Server) GetCpuUsedPercent(ctx context.Context, in *StatusRequest) (*UsedPercent, error) {
	log.Debug("StatusRequest for CPU UsedPercent")

	cpu := cpu.Last()
	// TODO Check for error instead
	if cpu != nil {
		return &UsedPercent{
			Value: cpu.UsedPercent,
		}, nil
	}

	return nil, errors.New("Can't get CPU infos!")
}

func (s *Server) GetCpuInfo(ctx context.Context, in *StatusRequest) (*CpuInfo, error) {
	log.Debug("StatusRequest for CpuInfo")

	cpu, cores := cpu.New()
	// TODO Check for error instead
	if cpu != nil && cores != nil {
		return &CpuInfo{
			Available:   true,
			CpuType:     CpuType_COMBINED,
			UsedPercent: cpu.UsedPercent,
			Frequency:   (*cores)[0].Mhz,
		}, nil
	}

	return nil, errors.New("Can't get CPU infos!")
}

func (s *Server) GetAdvancedCpuInfo(ctx context.Context, in *StatusRequest) (*AdvancedCpuInfo, error) {
	log.Debug("StatusRequest for AdvancedCpuInfo")

	cpu, cores := cpu.New()
	// TODO Check for error instead
	if cpu != nil && cores != nil {
		return &AdvancedCpuInfo{
			Available:   true,
			CpuType:     CpuType_COMBINED,
			UsedPercent: cpu.UsedPercent,
			Frequency:   (*cores)[0].Mhz,
			Microcode:   cpu.Microcode,
		}, nil
	}

	return nil, errors.New("Can't get CPU infos!")
}

func (s *Server) GetCoreInfo(ctx context.Context, in *CoreStatusRequest) (*CpuInfo, error) {
	log.Debug("CoreStatusRequest for CpuInfo")

	switch in.Type {
	case CpuType_COMBINED:
		return s.GetCpuInfo(ctx, &StatusRequest{})
	case CpuType_LOGICAL_CORE:
		return nil, errors.New("Not yet supported!")
	case CpuType_PHYSICAL_CORE:
		return nil, errors.New("Not yet supported!")
	default:
		return nil, errors.New("Unknown parameter!")
	}
}

func (s *Server) GetAdvancedCoreInfo(ctx context.Context, in *CoreStatusRequest) (*AdvancedCpuInfo, error) {
	log.Debug("CoreStatusRequest for AdvancedCpuInfo")

	switch in.Type {
	case CpuType_COMBINED:
		return s.GetAdvancedCpuInfo(ctx, &StatusRequest{})
	case CpuType_LOGICAL_CORE:
		return nil, errors.New("Not yet supported!")
	case CpuType_PHYSICAL_CORE:
		return nil, errors.New("Not yet supported!")
	default:
		return nil, errors.New("Unknown parameter!")
	}
}

/*
 Ram service.
*/

func (s *Server) GetRamUsedPercent(ctx context.Context, in *StatusRequest) (*UsedPercent, error) {
	log.Debug("StatusRequest for RAM UsedPercent")

	ram := ram.Last()
	// TODO Check for error instead
	if ram != nil {
		return &UsedPercent{
			Value: ram.UsedPercent,
		}, nil
	}

	return nil, errors.New("Can't get RAM infos!")
}

func (s *Server) GetRamInfo(ctx context.Context, in *StatusRequest) (*RamInfo, error) {
	log.Debug("StatusRequest for RamInfo")

	ram := ram.Last()
	// TODO Check for error instead
	if ram != nil {
		return &RamInfo{
			UsedPercent: ram.UsedPercent,
			Total:       ram.Total,
			Used:        ram.Used,
		}, nil
	}

	return nil, errors.New("Can't get RAM infos!")
}

func (s *Server) GetAdvancedRamInfo(ctx context.Context, in *StatusRequest) (*AdvancedRamInfo, error) {
	log.Debug("StatusRequest for AdvancedRamInfo")

	ram := ram.Last()
	// TODO Check for error instead
	if ram != nil {
		return &AdvancedRamInfo{
			UsedPercent: ram.UsedPercent,
			Total:       ram.Total,
			Used:        ram.Used,
			Cached:      ram.Cached,
		}, nil
	}

	return nil, errors.New("Can't get RAM infos!")
}
