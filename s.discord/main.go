package main

import (
	"context"

	"swallowtail/libraries/mariana"
	"swallowtail/s.discord/client"
	"swallowtail/s.discord/dao"
	"swallowtail/s.discord/handler"
	discordproto "swallowtail/s.discord/proto"
)

const (
	svcName = "s.discord"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Init dao
	if err := dao.Init(ctx, svcName); err != nil {
		panic(err)
	}

	// Init Client
	if err := client.Init(ctx); err != nil {
		panic(err)
	}

	// Init gRPC server
	s := mariana.Init(svcName)
	discordproto.RegisterDiscordServer(s.Grpc(), &handler.DiscordService{})
	s.Run(ctx)
}
