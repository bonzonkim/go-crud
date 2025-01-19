package networks

import (
	"go-crud/services"

	"github.com/gin-gonic/gin"
)

type Network struct {
	engine		*gin.Engine
	serviec		*services.Service
}

func NewNetwork(service *services.Service) *Network {
	r := &Network {
		engine:		gin.New(),
	}
	newUserRouter(r, service.User)
	return r
}

func (n *Network) ServerStart(port string) {
	n.engine.Run(port)
}
