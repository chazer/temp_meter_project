package app

import (
	"tmeter/app/api"
	auth "tmeter/app/modules/auth/services"
	devicesRepo "tmeter/app/modules/devices/repositories"
	measureRepo "tmeter/app/modules/measurements/repositories"
	"tmeter/app/modules/measurements/services"
	usersRepo "tmeter/app/modules/users/repositories"
	"tmeter/lib/tokens"
)
import devicesServices "tmeter/app/modules/devices/services"

type AppFactoryInterface interface {
	GetDevicesRepository() devicesRepo.DevicesRepositoryInterface
	GetDevicesService() devicesServices.DevicesServiceInterface
	GetAPIProtocol() *api.APIProtocol
	GetAuthService() auth.AuthServiceInterface
	GetMeasurementsService() services.MeasurementsServiceInterface
}

type AppFactory struct {
	devicesRepository   devicesRepo.DevicesRepositoryInterface
	devicesService      *devicesServices.DevicesService
	apiProtocol         *api.APIProtocol
	authService         auth.AuthServiceInterface
	measurementsService services.MeasurementsServiceInterface
	tokenEncoder        tokens.TokenEncoderInterface
}

func NewAppFactory() *AppFactory {
	f := &AppFactory{}
	return f
}

func (f *AppFactory) GetDevicesService() devicesServices.DevicesServiceInterface {
	if f.devicesService == nil {
		config := &devicesServices.DevicesServiceConfig{}
		if config.Repository == nil {
			config.Repository = f.GetDevicesRepository()
		}
		f.devicesService = devicesServices.MakeDevicesService(config)
	}
	return f.devicesService
}

func (f *AppFactory) GetDevicesRepository() devicesRepo.DevicesRepositoryInterface {
	if f.devicesRepository == nil {
		f.devicesRepository = devicesRepo.MakeDevicesInmemoryRepository()
	}
	return f.devicesRepository
}

func (f *AppFactory) GetAPIProtocol() *api.APIProtocol {
	if f.apiProtocol == nil {
		f.apiProtocol = &api.APIProtocol{}
	}
	return f.apiProtocol
}

func (f *AppFactory) GetAuthService() auth.AuthServiceInterface {
	if f.authService == nil {
		f.authService = auth.NewAuthService(
			usersRepo.MakeUsersInmemoryRepository(),
			f.GetDevicesService(),
			f.GetTokenEncoder(),
		)
	}
	return f.authService
}

func (f *AppFactory) CreateDeviceLog() measureRepo.DeviceLogInterface {
	return &measureRepo.MeasurementsLog{}
}

func (f *AppFactory) GetMeasurementsService() services.MeasurementsServiceInterface {
	if f.measurementsService == nil {
		f.measurementsService = services.NewMeasurementsService(
			f.CreateDeviceLog,
			f.GetDevicesService(),
		)
	}
	return f.measurementsService
}

func (f *AppFactory) GetTokenEncoder() tokens.TokenEncoderInterface {
	if f.tokenEncoder == nil {
		f.tokenEncoder = tokens.NewBase64TokenEncoder()
	}
	return f.tokenEncoder
}
