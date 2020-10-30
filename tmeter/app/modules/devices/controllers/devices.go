package controllers

import (
	"net/http"
	"tmeter/app/api"
	apiViews "tmeter/app/api/views"
	"tmeter/app/modules/devices/entities"
	"tmeter/app/modules/devices/views"
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
}

func NewDevicesController(
	devices services.DevicesServiceInterface,
	protocol *api.APIProtocol,
) *DevicesController {
	c := &DevicesController{
		api:            protocol,
		devicesService: devices,
		Handlers:       router.NewRoutes(),
		formatter: api.FormatterConfig{
			DocumentView: views.NewDeviceApiView(),
			PageApiView:  apiViews.NewPageApiView(views.NewDeviceApiView()),
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
