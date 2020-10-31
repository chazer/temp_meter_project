package controllers

import (
	"net/http"
	"tmeter/lib/debug"
)

func (c *DevicesController) handlerGetLog(resp http.ResponseWriter, req *http.Request) {
	uuid := req.URL.Query().Get("id")
	token := req.URL.Query().Get("token")

	email, err := c.auth.GetEmailFromToken(token)
	if err != nil {
		resp.WriteHeader(http.StatusUnauthorized)
		return
	}

	debug.Printf("user (email:%s) wants get device (uuid:%s)", email, uuid)

	d, err := c.devicesService.GetDeviceById(uuid)
	if err != nil {
		debug.Printf("device (uuid:%s) is not found", uuid)
		resp.WriteHeader(http.StatusNotFound)
		return
	}

	if d.OwnerEmail != *email {
		debug.Printf("user (email:%s) is not owner of device (uuid:%s)", *email, uuid)
		resp.WriteHeader(http.StatusForbidden)
		return
	}

	log := c.measurements.GetDeviceLog(d.UUID).ToSlice()

	var is = make([]interface{}, len(log))
	for i, d := range log {
		is[i] = d
	}
	c.api.SendSlice(resp, &c.logFormatter, is)
}
