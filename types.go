package plugin

import "github.com/gin-gonic/gin"

type Metadata struct {
	ID          string
	Name        string
	Version     string
	Description string
}

type RouteSetup func(*gin.RouterGroup) error

type RegisterRequest struct {
	PluginUrl string `json:"plugin-url"`
}

type Setting struct {
	Default        string
	Type           string
	Mandatory      bool
	Description    string
	PossibleValues []string
	Value          interface{}
}
