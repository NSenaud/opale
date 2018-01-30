package api

import (
	"errors"
	"log"

	"github.com/NSenaud/opale"
	"golang.org/x/net/context"
)

// Server represents the gRPC server
type Server struct{}

/*
 Cpu service
*/

func (s *Server) GetCpuInfo(ctx context.Context, in *StatusRequest) (*CpuInfo, error) {
	log.Println("StatusRequest for CpuInfo")

	cpu, cores := opale.GetCpu()
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
	log.Println("StatusRequest for AdvancedCpuInfo")

	cpu, cores := opale.GetCpu()
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
	log.Println("CoreStatusRequest for CpuInfo")

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
	log.Println("CoreStatusRequest for AdvancedCpuInfo")

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

func (s *Server) GetRamInfo(ctx context.Context, in *StatusRequest) (*RamInfo, error) {
	log.Println("StatusRequest for RamInfo")

	ram := opale.GetMem()
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
	log.Println("StatusRequest for AdvancedRamInfo")

	ram := opale.GetMem()
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
