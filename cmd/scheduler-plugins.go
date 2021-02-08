package main

import (
	"math/rand"
	"os"
	"time"

	"github.com/morningfish/kube-scheduler-framework/pkg/plugins"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	command := plugins.Register()
	if err := command.Execute(); err != nil {
		os.Exit(1)
	}

}
