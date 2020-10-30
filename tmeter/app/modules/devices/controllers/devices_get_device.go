package controllers

import (
	"net/http"
	"tmeter/app/modules/devices/services"
	"tmeter/lib/debug"
)

func (c *DevicesController) handlerGetOneDevice(writer http.ResponseWriter, request *http.Request) {
	id := request.URL.Query().Get("id")
	debug.Printf("get device (uuid=%s)", id)

	d, err := c.devicesService.GetDeviceById(id)
	if err != nil && err.Error() == services.ErrNoDevice {
		writer.WriteHeader(404)
	}
	if d != nil {
		debug.Printf("found: (uuid=%s; email=%s)", d.UUID, d.OwnerEmail)
		c.sendItem(writer, d)

		writer.WriteHeader(201)
	} else {
		writer.WriteHeader(500)
	}
}
