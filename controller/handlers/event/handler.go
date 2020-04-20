package event

import (
	"context"

	v1 "github.com/shopnado/shopify-controller/controller/apis/shopnado.xyz/v1"
)

type handler struct {
	ctx context.Context
}

func NewHandler(ctx context.Context) *handler {
	return &handler{
		ctx: ctx,
	}
}

func (h *handler) OnChange(key string, obj *v1.Event) (*v1.Event, error) {
	return obj, nil
}

func (h *handler) OnRemove(key string, obj *v1.Event) (*v1.Event, error) {
	return obj, nil
}
