package controllers

import (
	"encoding/json"
	"net/http"
	"tmeter/app/api"
	auth "tmeter/app/modules/auth/services"
	"tmeter/app/modules/auth/views"
	devices "tmeter/app/modules/devices/services"
	"tmeter/lib/debug"
	"tmeter/lib/middlewares"
	"tmeter/lib/middlewares/wrappers"
	"tmeter/lib/router"
)

type AuthDevicesController struct {
	Handlers       *router.Routes
	api            *api.APIProtocol
	formatter      api.FormatterConfig
	auth           *auth.AuthService
	devicesService devices.DevicesServiceInterface
}

func NewAuthDevicesController(
	protocol *api.APIProtocol,
	auth *auth.AuthService,
	devices devices.DevicesServiceInterface,
) *AuthDevicesController {
	c := &AuthDevicesController{
		Handlers:       router.NewRoutes(),
		api:            protocol,
		auth:           auth,
		devicesService: devices,
		formatter: api.FormatterConfig{
			DocumentView: views.NewAuthTokenResponseView(),
		},
	}

	c.Handlers.POST(
		"/token",
		// TODO: rewrite it
		middlewares.LogMiddleware(
			// TODO: Add ContentType guard middleware
			wrappers.NewJsonErrorsWrapper().AsMiddleware()(
				c.handlerCreateNewDeviceToken,
			),
		),
		// BODY: {
		//   device_name: string,
		//   user_email: string,
		// }
	)
	return c
}

type tokenRequestDTO1 struct {
	DeviceName string `json:"device_name"`
	UserEmail  string `json:"user_email"`
}

func (c *AuthDevicesController) handlerCreateNewDeviceToken(resp http.ResponseWriter, req *http.Request) {
	var dto tokenRequestDTO1
	err := json.NewDecoder(req.Body).Decode(&dto)
	if err != nil {
		http.Error(resp, err.Error(), http.StatusBadRequest)
		return
	}

	debug.Printf("create device (name=%s)", dto.DeviceName)

	d, err := c.devicesService.CreateDevice(dto.DeviceName, dto.UserEmail)
	if err != nil {
		http.Error(resp, err.Error(), http.StatusInternalServerError)
		return
	}

	debug.Printf("create token for device (uuid=%s; name=%s)", d.UUID, d.Name)

	token, err := c.auth.IssueTokenForDevice(d)
	if err != nil {
		http.Error(resp, err.Error(), http.StatusBadRequest)
		return
	}

	c.api.WriteEntityDocumentResponse(resp, &c.formatter, *token)
}
