package module

import "github.com/gin-gonic/gin"

type RouteSetup func(*gin.RouterGroup) error
