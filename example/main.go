package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/homescreenrocks/plugin-library"
)

func main() {
	p := &plugin.Plugin{
		Metadata: plugin.Metadata{
			ID:          "example-plugin",
			Name:        "Example Plugin",
			Version:     "v0.0",
			Description: "An example plugin.",
		},
		Settings: map[string]plugin.Setting{
			"tz": plugin.Setting{
				Default:   "",
				Type:      "string",
				Mandatory: false,
				Value:     "Berlin",
			},
		},
		RouteSetup: func(group *gin.RouterGroup) error {
			group.GET("/date", func(c *gin.Context) {
				c.JSON(http.StatusOK, time.Now().String())
			})

			return nil
		},
	}

	p.Main()
}
