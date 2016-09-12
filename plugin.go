package plugin

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type Plugin struct {
	Metadata   Metadata
	Settings   map[string]Setting
	RouteSetup RouteSetup
}

func (p *Plugin) Main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <core-url>\n", os.Args[0])
		os.Exit(2)
	}

	err := p.run(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to start plugin.\n")
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func (p *Plugin) run(coreUrl string) error {
	log.Printf("Starting plugin for homescreen using the plugin helper.")
	log.Printf("  ID:          %s", p.Metadata.ID)
	log.Printf("  Name:        %s", p.Metadata.Name)
	log.Printf("  Version:     %s", p.Metadata.Version)
	log.Printf("  Description: %s", p.Metadata.Description)

	engine, err := p.setupGin()
	if err != nil {
		return err
	}

	listner, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		return fmt.Errorf("Failed started HTTP listener: %v", err)
	}

	err = p.register(coreUrl, listner.Addr().String())
	if err != nil {
		return fmt.Errorf("Failed server HTTP: %v", err)
	}

	err = http.Serve(listner, engine)
	if err != nil {
		return fmt.Errorf("Failed server HTTP: %v", err)
	}

	return nil
}

func (p *Plugin) setupGin() (*gin.Engine, error) {
	gin.SetMode(gin.ReleaseMode)
	m := gin.Default()

	if p.RouteSetup != nil {
		err := p.RouteSetup(m.Group("/plugin"))
		if err != nil {
			return nil, err
		}
	}

	v1 := m.Group("/v1")
	v1.GET("/metadata", func(c *gin.Context) {
		c.JSON(http.StatusOK, p.Metadata)
	})

	return m, nil

}

func (p *Plugin) register(coreUrl string, pluginUrl string) error {
	req := &RegisterRequest{
		PluginUrl: pluginUrl,
	}

	data, err := json.Marshal(req)
	if err != nil {
		log.Fatal(err)
	}

	buf := bytes.NewBuffer(data)

	resp, err := http.Post(
		fmt.Sprintf("%s/api/v1/modules/register", coreUrl),
		"application/json", buf)
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		return fmt.Errorf("unexpected response code %d", resp.StatusCode)
	}

	return nil
}
