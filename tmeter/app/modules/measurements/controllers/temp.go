package controllers

import (
	"encoding/json"
	"net/http"
	"tmeter/app/api"
	auth "tmeter/app/modules/auth/services"
	"tmeter/app/modules/auth/views"
	devices "tmeter/app/modules/devices/services"
	"tmeter/app/modules/measurements/services"
	"tmeter/lib/debug"
	"tmeter/lib/middlewares"
	"tmeter/lib/middlewares/wrappers"
	"tmeter/lib/router"
)

const otherErrorsStatusCode = 422

type DeviceTempController struct {
	Handlers            *router.Routes
	api                 *api.APIProtocol
	formatter           api.FormatterConfig
	auth                auth.AuthServiceInterface
	devicesService      devices.DevicesServiceInterface
	measurementsService services.MeasurementsServiceInterface
}

func NewDeviceTempController(
	protocol *api.APIProtocol,
	auth auth.AuthServiceInterface,
	devices devices.DevicesServiceInterface,
	measurements services.MeasurementsServiceInterface,
) *DeviceTempController {
	c := &DeviceTempController{
		Handlers:            router.NewRoutes(),
		api:                 protocol,
		auth:                auth,
		devicesService:      devices,
		measurementsService: measurements,
		formatter: api.FormatterConfig{
			DocumentView: views.NewAuthTokenResponseView(),
		},
	}

	c.Handlers.POST(
		"/temp?token",
		middlewares.LogMiddleware(
			wrappers.NewJsonErrorsWrapper().AsMiddleware()(
				c.handlerSaveTempMetrics,
			),
		),
		// BODY: [
		//   { time: string, value: float },
		// }
	)
	return c
}

type tempMetricPoint struct {
	Time  int64   `json:"time"`
	Value float32 `json:"value"`
}

type tokenSaveMetricsDTO = []tempMetricPoint

func (c *DeviceTempController) handlerSaveTempMetrics(resp http.ResponseWriter, req *http.Request) {
	// TODO: extract device uuid by middleware
	token := req.URL.Query().Get("token")
	uuid, err := c.auth.GetUUIDFromToken(token)

	debug.Printf("token %s", token)
	debug.Printf("uuid %s", *uuid)

	if err != nil {
		debug.Printf("Error: %s", err.Error())
		resp.WriteHeader(http.StatusUnauthorized)
		return
	}

	debug.Printf("device (uuid:%s) wants to save data", *uuid)
	d, err := c.devicesService.GetDeviceById(*uuid)
	if err != nil {
		debug.Printf("device (uuid:%s) is not found", *uuid)
		resp.WriteHeader(http.StatusUnauthorized)
		return
	}
	debug.Printf("device (uuid:%s) is found", d.UUID)

	var dto tokenSaveMetricsDTO
	if err := json.NewDecoder(req.Body).Decode(&dto); err != nil {
		http.Error(resp, err.Error(), otherErrorsStatusCode)
		return
	}

	debug.Printf("got metrics (count=%d)", len(dto))

	for _, record := range dto {
		if err := c.measurementsService.WriteTemperature(d.UUID, record.Time, record.Value); err != nil {
			http.Error(resp, err.Error(), otherErrorsStatusCode)
			return
		}
	}

	resp.WriteHeader(http.StatusCreated)
}
