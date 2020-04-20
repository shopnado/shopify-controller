package controller

import (
	"context"

	v1 "github.com/shopnado/shopify-controller/controller/generated/controllers/shopnado.xyz/v1"
	"github.com/shopnado/shopify-controller/controller/handlers/event"
)

func Register(
	ctx context.Context, events v1.EventController) {

	eventHandler := event.NewHandler(ctx)
	events.OnChange(ctx, "events-handler", eventHandler.OnChange)
	events.OnRemove(ctx, "events-handler", eventHandler.OnRemove)
}
