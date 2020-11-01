package controllers

import (
	"net/http"
)

func (c *DevicesController) handlerGetOneDevice(writer http.ResponseWriter, req *http.Request) {
	uuid := req.URL.Query().Get("id")
	token := req.Header.Get("Authorization")

	d, status := c.getDeviceAndVerify(uuid, token)
	if status > 0 {
		writer.WriteHeader(status)
		return
	}

	if d != nil {
		c.sendItem(writer, d)
		writer.WriteHeader(201)
	} else {
		writer.WriteHeader(500)
	}
}
