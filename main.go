package main

import (
	"fmt"
	"log/slog"

	"github.com/TravisRoad/travis-blog-service/global"
	"github.com/TravisRoad/travis-blog-service/setup"
)

func main() {
	setup.Setup()

	r := setup.InitRouter()
	if err := r.Run(fmt.Sprintf("0.0.0.0:%d", global.Config.Port)); err != nil {
		slog.Error("failed to run server on 0.0.0.0:%d: %v", global.Config.Port, err)
	}
}
