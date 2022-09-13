package handlers

import (
	"github.com/res0lution/digital-house/config"
	"github.com/res0lution/digital-house/ent"
)

type Handler struct {
	Client *ent.Client
	Config *config.Config
}

func NewHandlers(client *ent.Client, config *config.Config) *Handler {
	return &Handler{
		Client: client,
		Config: config,
	}
}