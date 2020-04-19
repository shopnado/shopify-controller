package api

var Topics = map[string]struct{}{
	"carts/create":                struct{}{},
	"carts/update":                struct{}{},
	"checkouts/create":            struct{}{},
	"checkouts/update":            struct{}{},
	"checkouts/delete":            struct{}{},
	"collections/create":          struct{}{},
	"collections/update":          struct{}{},
	"collections/delete":          struct{}{},
	"collection_listings/add":     struct{}{},
	"collection_listings/update":  struct{}{},
	"collection_listings/remove":  struct{}{},
	"customers/create":            struct{}{},
	"customers/disable":           struct{}{},
	"customers/enable":            struct{}{},
	"customers/update":            struct{}{},
	"customers/delete":            struct{}{},
	"customer_groups/create":      struct{}{},
	"customer_groups/update":      struct{}{},
	"customer_groups/delete":      struct{}{},
	"draft_orders/create":         struct{}{},
	"draft_orders/update":         struct{}{},
	"draft_orders/delete":         struct{}{},
	"fulfillments/create":         struct{}{},
	"fulfillments/update":         struct{}{},
	"fulfillment_events/create":   struct{}{},
	"fulfillment_events/delete":   struct{}{},
	"inventory_items/create":      struct{}{},
	"inventory_items/update":      struct{}{},
	"inventory_items/delete":      struct{}{},
	"inventory_levels/connect":    struct{}{},
	"inventory_levels/disconnect": struct{}{},
	"inventory_levels/update":     struct{}{},
	"locations/create":            struct{}{},
	"locations/update":            struct{}{},
	"locations/delete":            struct{}{},
	"orders/cancelled":            struct{}{},
	"orders/create":               struct{}{},
	"orders/fulfilled":            struct{}{},
	"orders/paid":                 struct{}{},
	"orders/partially_fulfilled":  struct{}{},
	"orders/updated":              struct{}{},
	"orders/delete":               struct{}{},
	"orders/edited":               struct{}{},
	"order_transactions/create":   struct{}{},
	"products/create":             struct{}{},
	"products/update":             struct{}{},
	"products/delete":             struct{}{},
	"product_listings/add":        struct{}{},
	"product_listings/update":     struct{}{},
	"product_listings/remove":     struct{}{},
	"refunds/create":              struct{}{},
	"app/uninstalled":             struct{}{},
	"shop/update":                 struct{}{},
	"locales/create":              struct{}{},
	"locales/update":              struct{}{},
	"tender_transactions/create":  struct{}{},
	"themes/create":               struct{}{},
	"themes/publish":              struct{}{},
	"themes/update":               struct{}{},
	"themes/delete":               struct{}{},
}

//func registerWebhooks(baseurl string) error {
//	webhooks, err := cli.Webhook.List(goshopify.ListOptions{})
//	if err != nil {
//		return err
//	}
//	existingTopics := map[string]struct{}{}
//	for _, wh := range webhooks {
//		existingTopics[wh.Topic] = struct{}{}
//	}
//
//	for topic, _ := range Topics {
//		logrus.Infof("adding webhook for topic: %s", topic)
//		if _, ok := existingTopics[topic]; ok {
//			logrus.Infof("skipping, existing webhook for topic: %s", topic)
//			continue // already exists, skip
//		}
//
//		themesUpdate, err := cli.Webhook.Create(goshopify.Webhook{
//			Topic:   topic,
//			Address: fmt.Sprintf("%s/%s", baseurl, topic),
//			Format:  "json",
//		})
//		if err != nil {
//			log.Println(err)
//		}
//
//		fmt.Printf("%+v\n", themesUpdate)
//	}
//
//	return nil
//}
