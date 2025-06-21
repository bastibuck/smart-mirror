package kvg

import (
	"github.com/go-chi/chi/v5"
	"smartmirror.server/env"
	"smartmirror.server/widgets"
)

type KVGWidget struct{}

var _ widgets.Widget = (*KVGWidget)(nil)

func (v *KVGWidget) SetupEnv() {
	env.ValidateEnvKeys(getEnvKeys())
}

func (v *KVGWidget) SetupRouter(router *chi.Mux) {
	router.HandleFunc("/transportation/departures", nextDeparturesHandler)
}

func NewKVGWidget() *KVGWidget {
	return &KVGWidget{}
}
