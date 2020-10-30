package app

import (
	"tmeter/app/api"
	auth "tmeter/app/modules/auth/services"
	"tmeter/app/modules/devices/repositories"
	repositories2 "tmeter/app/modules/measurements/repositories"
	"tmeter/app/modules/measurements/services"
	repositories3 "tmeter/app/modules/users/repositories"
)
import devicesServices "tmeter/app/modules/devices/services"

type AppFactoryInterface interface {
	GetDevicesRepository() repositories.DevicesRepositoryInterface
	GetDevicesService() devicesServices.DevicesServiceInterface
	GetAPIProtocol() *api.APIProtocol
	GetAuthService() auth.AuthServiceInterface
	GetMeasurementsService() services.MeasurementsServiceInterface
}

type AppFactory struct {
	devicesRepository   repositories.DevicesRepositoryInterface
	devicesService      *devicesServices.DevicesService
	apiProtocol         *api.APIProtocol
	authService         auth.AuthServiceInterface
	measurementsService services.MeasurementsServiceInterface
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

func (f *AppFactory) GetDevicesRepository() repositories.DevicesRepositoryInterface {
	if f.devicesRepository == nil {
		f.devicesRepository = repositories.MakeDevicesInmemoryRepository()
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
			repositories3.MakeUsersInmemoryRepository(),
		)
	}
	return f.authService
}

func (f *AppFactory) GetMeasurementsService() services.MeasurementsServiceInterface {
	if f.measurementsService == nil {
		f.measurementsService = services.NewMeasurementsService(
			&repositories2.FloatLog{},
		)
	}
	return f.measurementsService
}
