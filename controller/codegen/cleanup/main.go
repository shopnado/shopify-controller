package main

import (
	"github.com/rancher/wrangler/pkg/cleanup"
	"github.com/sirupsen/logrus"
)

func main() {
	if err := cleanup.Cleanup("controller/apis"); err != nil {
		logrus.Fatal(err)
	}
}
