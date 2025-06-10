package main

import (
	"context"
	"fmt"
	"log"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humafiber"
	"github.com/gofiber/fiber/v2"
	//_ "github.com/danielgtaylor/huma/v2/formats/cbor"
)

type DeviceOutput struct {
	Body struct {
		Message string `json:"message" example:"Device Type, Device Name, Device ID, Device IP Address, Current State, Active Deployments" doc:"Device deployed display message"`
	}
}

func main() {
	router := fiber.New()
	api := humafiber.New(router, huma.DefaultConfig("Autobox API", "0.0.1"))

	// register /device/{deviceID} handler
	huma.Get(api, "/device/{deviceId}", func(ctx context.Context, input *struct {
		DeviceId string `path:"deviceId" maxLength:"16" example:"device123" doc:"The ID of the device"`
	}) (*DeviceOutput, error) {
		resp := &DeviceOutput{}
		resp.Body.Message = fmt.Sprintf("Device ID: %s", input.DeviceId)
		return resp, nil
	})
	log.Fatal(router.Listen(":8888"))

}
