package peers

import (
	"encoding/binary"
	"fmt"
	"net"
)

/*
check the Response Formate, We are building this now
https://wiki.theory.org/BitTorrent_Tracker_Protocol
*/
type Peer struct {
	IP   net.IP
	Port uint16
}

func Unmarshal(peersBin []byte) ([]Peer, error) {
	const peerSize = 6
	numPeers := len(peersBin) / 6
	if len(peersBin)%peerSize != 0 {
		err := fmt.Errorf("received wrong peers")
		return nil, err
	}
	peers := make([]Peer, numPeers)
	for i := 0; i < numPeers; i++ {
		offset := i * peerSize
		peers[i].IP = net.IP(peersBin[offset : offset+4])
		peers[i].Port = binary.BigEndian.Uint16(peersBin[offset+4 : offset+6])
	}
	return peers, nil
}
