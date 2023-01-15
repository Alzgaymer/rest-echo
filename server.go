package main

import (
	"context"
	"server/src/electronic"
	"server/src/service"
)

func main() {
	electronic.Start()
	service.New().Disconnect(context.Background())
}
