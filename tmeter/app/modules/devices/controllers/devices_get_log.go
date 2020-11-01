package controllers

import (
	"net/http"
	"tmeter/app/modules/devices/entities"
	"tmeter/lib/debug"
)

func (c *DevicesController) getDeviceAndVerify(uuid string, token string) (*entities.Device, int) {
	email, err := c.auth.GetEmailFromToken(token)
	if err != nil {
		return nil, http.StatusUnauthorized
	}

	debug.Printf("user (email:%s) wants get device (uuid:%s)", *email, uuid)

	d, err := c.devicesService.GetDeviceById(uuid)
	if err != nil {
		debug.Printf("device (uuid:%s) is not found", uuid)
		return d, http.StatusNotFound
	}

	debug.Printf("found device (uuid:%d)", uuid)

	if d.OwnerEmail != *email {
		debug.Printf("user (email:%s) is not owner of device (uuid:%s)", *email, uuid)
		return d, http.StatusForbidden
	}

	return d, 0
}

func (c *DevicesController) handlerGetLog(resp http.ResponseWriter, req *http.Request) {
	uuid := req.URL.Query().Get("id")
	token := req.Header.Get("Authorization")

	d, status := c.getDeviceAndVerify(uuid, token)
	if status > 0 {
		resp.WriteHeader(status)
		return
	}

	log := c.measurements.GetDeviceLog(d.UUID).ToSlice()

	var is = make([]interface{}, len(log))
	for i, d := range log {
		is[i] = d
	}
	c.api.SendSlice(resp, &c.logFormatter, is)
}
