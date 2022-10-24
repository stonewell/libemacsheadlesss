package main

/*
   #include "headless.h"
*/
import "C"

import (
	ehs "github.com/stonewell/emacsheadless/pkg/service"
	pb "github.com/stonewell/emacsheadless/proto"
)

type CServerConfig struct {
	ehs.ServerConfig

	sc *C.ServerConfig
}

type CClientCallbacks struct {
	cc * C.ClientCallbacks
}

//export StartServer
func StartServer(sc *C.ServerConfig) {
	base_sc := ehs.ServerConfig{
		Port: uint(sc.port),
		Addr: C.GoString(sc.addr),
	}

	ehs.StartServer(&CServerConfig{
		base_sc,
		sc,
	})
}

//export RegisterClientCallback
func RegisterClientCallback(clientId int32, callback *C.ClientCallbacks) {
	ehs.G_ServImpl.ClientChannel <- ehs.ClientOP{
		Op: pb.CmdType_Cmd_ClientInfo,
		Client: ehs.Client{
			ClientId: clientId,
			Callback: &CClientCallbacks{
				callback,
			},
		},
	}
}

//export DisconnectClient
func DisconnectClient(clientId int32) {
	ehs.G_ServImpl.ClientChannel <- ehs.ClientOP{
		Op: pb.CmdType_Cmd_ClientDisconnect,
		Client: ehs.Client{
			ClientId: clientId,
		},
	}
}

func (self *CServerConfig) OnNewClient(clientId uint) {
	if self.sc != nil && self.sc.new_client_callback != nil {
		C.bridge_new_client(self.sc.new_client_callback, C.uint(clientId))
	}
}

func (self *CServerConfig) OnClientDisconnect(clientId uint) {
	if self.sc != nil && self.sc.client_disconnect_callback != nil {
		C.bridge_client_disconnect(self.sc.client_disconnect_callback, C.uint(clientId))
	}
}

func (self *CClientCallbacks) KeyboardInput(keycode uint) {
	if self.cc != nil && self.cc.input != nil {
		C.bridge_keyboard_input(self.cc.input, C.uint(keycode))
	}
}

func main() {
}
