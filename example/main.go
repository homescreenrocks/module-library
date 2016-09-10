package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/homescreenrocks/plugin-library"
)

func main() {
	app := &App{}

	p := plugin.New(plugin.Metadata{
		ID:          "example-plugin",
		Name:        "Example Plugin",
		Version:     "v0.0",
		Description: "An example plugin.",
	})

	p.RouteSetup = app.RouteSetup
	p.Main()
}

type App struct {
}

func (app *App) RouteSetup(group *gin.RouterGroup) error {
	group.GET("/date", func(c *gin.Context) {
		c.JSON(http.StatusOK, time.Now().String())
	})

	return nil
}
