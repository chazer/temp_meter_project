package controllers

import (
	"encoding/json"
	"net/http"
)

import "tmeter/lib/debug"

type RegisterDeviceDTO struct {
	DeviceName string `json:"device_name"`
	ForEmail   string `json:"for_email"`
}

func (c *DevicesController) handlerCreateNewDevice(resp http.ResponseWriter, req *http.Request) {
	var dto RegisterDeviceDTO
	err := json.NewDecoder(req.Body).Decode(&dto)
	if err != nil {
		http.Error(resp, err.Error(), http.StatusBadRequest)
		return
	}

	debug.Printf("try create device (name=%s0 for email %s", dto.DeviceName, dto.ForEmail)

	if len(dto.DeviceName) == 0 {
		http.Error(resp, "Empty value in device_name", http.StatusBadRequest)
		return
	}

	if len(dto.ForEmail) == 0 {
		http.Error(resp, "Empty value in for_email", http.StatusBadRequest)
		return
	}

	d, err := c.devicesService.CreateDevice(dto.ForEmail)
	if err != nil {
		http.Error(resp, err.Error(), http.StatusInternalServerError)
		return
	}

	debug.Printf("device created (UUID:%s) for User(%s)", d.UUID, d.UserEmail)

	c.api.WriteEntityDocumentResponse(resp, &c.formatter, d)
}
