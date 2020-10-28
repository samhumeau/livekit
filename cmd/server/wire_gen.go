// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
	"github.com/livekit/livekit-server/pkg/config"
	"github.com/livekit/livekit-server/pkg/node"
	"github.com/livekit/livekit-server/pkg/service"
)

// Injectors from wire.go:

func InitializeServer(conf *config.Config) (*LivekitServer, error) {
	nodeNode, err := node.NewLocalNode(conf)
	if err != nil {
		return nil, err
	}
	roomService, err := service.NewRoomService(conf, nodeNode)
	if err != nil {
		return nil, err
	}
	rtcService := service.NewRTCService()
	livekitServer, err := NewLivekitServer(conf, roomService, rtcService)
	if err != nil {
		return nil, err
	}
	return livekitServer, nil
}