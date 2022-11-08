package network

import "bytes"


type NetAddr string

type RPC struct{
	// RPC is gonna be the messages that gonna be sent over the transport layer
	From NetAddr
	Payload[] byte
}
type Transport interface{
	

	// Return a channel for RPC. // Transport is a module on the server, and server needs have the access to the messages that are sent over the transport layer
	// We can do it with this consume method
	Consume() <-chan RPC
	// We need to make a connection with another Transport 
	Connect(Transport) error 
	SendMessage(NetAddr,[]byte) error
	Addr() NetAddr

}