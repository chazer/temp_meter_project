package controllers

import (
	"net/http"
	"tmeter/app/api"
	apiViews "tmeter/app/api/views"
	auth "tmeter/app/modules/auth/services"
	"tmeter/app/modules/devices/entities"
	"tmeter/app/modules/devices/views"
	measurements "tmeter/app/modules/measurements/services"
	"tmeter/lib/middlewares"
	"tmeter/lib/middlewares/wrappers"
)

import "tmeter/lib/router"
import "tmeter/app/modules/devices/services"

type DevicesController struct {
	api            *api.APIProtocol
	Handlers       *router.Routes
	devicesService services.DevicesServiceInterface
	formatter      api.FormatterConfig
	logFormatter   api.FormatterConfig
	auth           auth.AuthServiceInterface
	measurements   measurements.MeasurementsServiceInterface
}

func NewDevicesController(
	devices services.DevicesServiceInterface,
	protocol *api.APIProtocol,
	auth auth.AuthServiceInterface,
	measurements measurements.MeasurementsServiceInterface,
) *DevicesController {
	c := &DevicesController{
		api:            protocol,
		devicesService: devices,
		auth:           auth,
		measurements:   measurements,
		Handlers:       router.NewRoutes(),
		formatter: api.FormatterConfig{
			DocumentView: views.NewDeviceApiView(),
			PageApiView:  apiViews.NewPageApiView(views.NewDeviceApiView()),
		},
		logFormatter: api.FormatterConfig{
			DocumentView: views.NewDeviceLogPointApiView(),
			PageApiView:  apiViews.NewPageApiView(views.NewDeviceLogPointApiView()),
		},
	}

	// TODO: rewrite it
	//c.Handlers.POST("",
	//	middlewares.LogMiddleware(
	//		wrappers.NewJsonErrorsWrapper().AsMiddleware()(
	//			middlewares.AuthMiddleware(
	//				c.handlerCreateNewDevice))))
	c.Handlers.GET("/?token",
		middlewares.LogMiddleware(
			wrappers.NewJsonErrorsWrapper().AsMiddleware()(
				middlewares.AuthMiddleware(
					c.handlerGetMyDevices))))
	c.Handlers.GET("/byId?id",
		middlewares.LogMiddleware(
			wrappers.NewJsonErrorsWrapper().AsMiddleware()(
				middlewares.AuthMiddleware(
					c.handlerGetOneDevice))))
	c.Handlers.GET("/byId/log?id",
		middlewares.LogMiddleware(
			wrappers.NewJsonErrorsWrapper().AsMiddleware()(
				middlewares.AuthMiddleware(
					c.handlerGetLog))))

	return c
}

func (c *DevicesController) sendItem(writer http.ResponseWriter, d *entities.Device) {
	c.api.WriteEntityDocumentResponse(writer, &c.formatter, d)
}

func (c *DevicesController) sendSlice(writer http.ResponseWriter, d []*entities.Device) {
	var is = make([]interface{}, len(d))
	for i, d := range d {
		is[i] = d
	}
	c.api.SendSlice(writer, &c.formatter, is)
}
