package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/homescreenrocks/homescreen/shared"
	"github.com/homescreenrocks/module-library"
)

func main() {
	m := new(module.Module)

	m.Metadata = shared.ModuleMetadata{
		ID:          "hue-control",
		Name:        "Philips Hue Home Automation Control Module",
		Version:     "v0.1",
		Description: "This module is used for controlling your Philips Hue light bulbs and accessory.",
	}

	m.Settings = shared.ModuleSettings{
		shared.ModuleSetting{
			Name:      "API-Point",
			Default:   "http://192.168.2.100/api",
			Type:      "string",
			Mandatory: true,
		},
		shared.ModuleSetting{
			Name:      "Username",
			Default:   "4JOBPVSlXXUuKHZYALuJ5o-HaGxwdTcT73cVCLpu",
			Type:      "string",
			Mandatory: true,
		},
	}
	m.RouteSetup = func(group *gin.RouterGroup) error {
		group.GET("/date", func(c *gin.Context) {
			c.JSON(http.StatusOK, time.Now().String())
		})

		return nil
	}

	m.Main()
}