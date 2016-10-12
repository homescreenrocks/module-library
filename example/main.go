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
		ID:          "example-module",
		Name:        "Example Module",
		Version:     "v0.0",
		Description: "An example module.",
	}

	m.Settings = shared.ModuleSettings{
		shared.ModuleSetting{
			Name:      "city",
			Default:   "Berlin",
			Type:      "string",
			Mandatory: false,
		},
		shared.ModuleSetting{
			Name:        "autopoll",
			Default:     true,
			Type:        "bool",
			Mandatory:   false,
			Description: "enables or disables something",
		},
		shared.ModuleSetting{
			Name:        "pollinterval",
			Default:     10,
			Type:        "number",
			Mandatory:   true,
			Description: "describes the interval on how often the modul should refresh the informations on the screen (in seconds)",
		},
		shared.ModuleSetting{
			Name:      "date",
			Default:   "2016-01-01",
			Type:      "date",
			Mandatory: false,
		},
		shared.ModuleSetting{
			Name:      "time",
			Default:   "15:00",
			Type:      "time",
			Mandatory: false,
		},
		shared.ModuleSetting{
			Name:           "difficulty",
			Default:        "easy",
			Type:           "string",
			Mandatory:      false,
			PossibleValues: []string{"easy", "medium", "hard"},
		},
	}
	m.RouteSetup = func(group *gin.RouterGroup) error {
		group.GET("/date", func(c *gin.Context) {
			c.JSON(http.StatusOK, time.Now().String())
		})

		templatesGroup := group.Group("/templates")
		templatesGroup.StaticFile("/main.html", "html/main.html")

		jsGroup := group.Group("/js")
		jsGroup.StaticFile("/ExampleModuleController.js", "js/ExampleModuleController.js")

		return nil
	}

	m.Main()
}
