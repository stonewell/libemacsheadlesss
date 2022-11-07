package main

/*
   #include "headless.h"
*/
import "C"

import (
	ehs "github.com/stonewell/emacsheadless/pkg/service"
	pb "github.com/stonewell/emacsheadless/proto"
)

func main() {
}

type CServerConfig struct {
	ehs.ServerConfig

	sc *C.ServerConfig
}

type CClientCallbacks struct {
	cc *C.ClientCallbacks
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

//export CursorTo
func CursorTo(terminalId int32, vpos int32, hpos int32) {
}

//export RawCursorTo
func RawCursorTo(terminalId int32, row int32, col int32) {
}

//export ClearToEnd
func ClearToEnd(terminal_id int32) {
}

//export ClearFrame
func ClearFrame(terminal_id int32) {
}

//export ClearEndOfLine
func ClearEndOfLine(terminal_id int32, first_unused_hpos int32) {
}

//export InsDelLines
func InsDelLines(terminal_id int32, vpos int32, n int32) {
}

//export InsertGlyphs
func InsertGlyphs(terminal_id int32, start string, len int32) {
}

//export WriteGlyphs
func WriteGlyphs(terminal_id int32, data string, len int32) {
}

//export DeleteGlyphs
func DeleteGlyphs(terminal_id int32, n int32) {
}

//export RingBell
func RingBell(terminal_id int32) {
}

//export ResetTerminalModes
func ResetTerminalModes(terminal_id int32) {
}

//export SetTerminalModes
func SetTerminalModes(terminal_id int32) {
}

//export UpdateEnd
func UpdateEnd(terminal_id int32) {
}

//export SetTerminalWindow
func SetTerminalWindow(terminal_id int32, size int32) {
}

//export ReadAvailInput
func ReadAvailInput(terminal_id int32) {
}

//export DeleteFrame
func DeleteFrame(terminal_id int32) {
}

//export DeleteTerminal
func DeleteTerminal(terminal_id int32) {
}
