package controllers

import (
	"net/http"
	devices "tmeter/app/modules/devices/services"
	"tmeter/lib/debug"
)

func (c *DevicesController) handlerGetMyDevices(writer http.ResponseWriter, request *http.Request) {
	token := request.Header.Get("Authorization")
	debug.Printf("token: %s", token)

	// TODO: extract user by middleware
	email, err := c.auth.GetEmailFromToken(token)
	if err != nil {
		debug.Printf("Error: %s", err.Error())
		writer.WriteHeader(http.StatusUnauthorized)
		return
	}

	d, err := c.devicesService.GetDevicesByEmail(*email)

	if err != nil && err.Error() == devices.ErrNoDevice {
		writer.WriteHeader(404)
	}
	if d != nil {
		c.sendSlice(writer, d)

		writer.WriteHeader(201)
	} else {
		writer.WriteHeader(500)
	}
}
