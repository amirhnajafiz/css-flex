package handler

import (
	"net/http"

	"github.com/amirhnajafiz/natsbench-runbook/internal/nats"
)

type Handler struct {
	NatsCli nats.CLI
}

func (h Handler) Health(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
}
