package network

import (
	"fmt"
	"sync"
)

type LocalTransport struct{
	addr NetAddr
	// Channel of RPC
	consumeCh chan RPC 
	// Transport needs to maintain and connect to peers, so we need to track them.
	lock sync.RWMutex // We need to lock this. Read Write Mutex
	peers map[NetAddr]*LocalTransport
}

func NewLocalTransport()* LocalTransport{
	return &LocalTransport{
		addr: 		addr,
		consumeCh: 	make(chan RPC, 1024),
		peers: 		make(map[NetAddr]*LocalTransport),
	}
}

func(t *LocalTransport) Consume() <- chan RPC{
	return t.consumeCh
}

func(t *LocalTransport) Connnect(tr Transport) error{
	t.lock.Lock()
	defer t.lock.Unlock()

	t.peers[tr.Addr()] = tr

	return nil
}

func (t *LocalTransport) SendMessage(to NetAddr, payload []byte) error{
	t.lock.RLocker()
	defer t.lock.RUnlock()

	peer, ok := t.peers[to]

	if !ok {
		return fmt.Errorf("%s: couldn't send a message to %s", &t.addr, to)
	}
	peer.consumeCh <-RPC{
		From: t.addr,
		Payload: payload,
	}
	
	return nil
}

func (t *LocalTransport) Addr() NetAddr{
	return t.addr
}