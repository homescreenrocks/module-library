package module

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/homescreenrocks/homescreen/shared"
)

type Module struct {
	shared.Module
	RouteSetup RouteSetup
}

func (p *Module) Main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <core-url>\n", os.Args[0])
		os.Exit(2)
	}

	err := p.run(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to start module.\n")
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func (m *Module) run(coreUrl string) error {
	log.Printf("Starting module for homescreen using the module helper.")
	log.Printf("  ID:          %s", m.Metadata.ID)
	log.Printf("  Name:        %s", m.Metadata.Name)
	log.Printf("  Version:     %s", m.Metadata.Version)
	log.Printf("  Description: %s", m.Metadata.Description)

	engine, err := m.setupGin()
	if err != nil {
		return err
	}

	listner, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		return fmt.Errorf("Failed started HTTP listener: %v", err)
	}

	err = m.register(coreUrl, "http://"+listner.Addr().String())
	if err != nil {
		return fmt.Errorf("Failed server HTTP: %v", err)
	}

	err = http.Serve(listner, engine)
	if err != nil {
		return fmt.Errorf("Failed server HTTP: %v", err)
	}

	return nil
}

func (p *Module) setupGin() (*gin.Engine, error) {
	gin.SetMode(gin.ReleaseMode)
	m := gin.Default()

	if p.RouteSetup != nil {
		err := p.RouteSetup(m.Group("/module"))
		if err != nil {
			return nil, err
		}
	}

	//v1 := m.Group("/v1")

	return m, nil

}

func (p *Module) register(coreUrl string, moduleUrl string) error {
	req := &shared.Module{
		ModuleURL: moduleUrl,
		Metadata:  p.Metadata,
		Settings:  p.Settings,
	}

	data, err := json.Marshal(req)
	if err != nil {
		return err
	}

	buf := bytes.NewBuffer(data)

	resp, err := http.Post(
		fmt.Sprintf("%s/api/v1/modules/", coreUrl),
		"application/json", buf)
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		return fmt.Errorf("unexpected response code %d: %s", resp.StatusCode, string(data))
	}

	return nil
}
