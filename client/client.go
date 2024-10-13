package client

import (
	"net"
	"time"

	"github.com/rishavmehra/bitTorrent/peers"
)

func New(peer peers.Peer, peerID, infoHash [20]byte) {
	conn, err := net.DialTimeout("tcp", peer.String(), 5*time.Second)
	if err != nil {
		return nil, err
	}

	// need handshake
}
