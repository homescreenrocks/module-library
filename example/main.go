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

	m.Metadata = &shared.ModuleMetadata{
		ID:          "example-module",
		Name:        "Example Module",
		Version:     "v0.0",
		Description: "An example module.",
	}

	m.Settings = &shared.ModuleSettings{
		shared.ModuleSetting{
			Name:      "timezone",
			Default:   "Berlin",
			Type:      "string",
			Mandatory: false,
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
