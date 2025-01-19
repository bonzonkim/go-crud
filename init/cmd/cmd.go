package cmd

import (
	"go-crud/configs"
	"go-crud/networks"
	"go-crud/repositories"
	"go-crud/services"
)

type Cmd struct {
	config		*config.Config
	network		*networks.Network 
	repository	*repositories.Repository
	service		*services.Service
}

func NewCmd(configFilePath string) *Cmd {
	c := &Cmd{
		config:		config.NewConfig(configFilePath),
	}

	c.repository = repositories.NewRepository()
	c.service = services.NewService(c.repository)
	c.network = networks.NewNetwork(c.service)

	c.network.ServerStart(c.config.Server.Port)

	return c
}
