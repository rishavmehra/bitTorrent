package torrentfile

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/jackpal/bencode-go"
	"github.com/rishavmehra/bitTorrent/peers"
)

/*

Request Format:
https://wiki.theory.org/BitTorrent_Tracker_Protocol
Retrieving peers from the tracker

This is how URL should should look like

http://some.tracker.com:999/announce
?info_hash=12345678901234567890
&peer_id=ABCDEFGHIJKLMNOPQRST
&ip=255.255.255.255  // this is optional read above mentioned doc for this
&port=6881
&downloaded=1234
&left=98765
&event=stopped // this is optional

Response Format:
interval
peers

*/

func (t *TorrentFile) buildTrackerURL(peerID [20]byte, port uint16) (string, error) {
	base, err := url.Parse(t.Announce)
	if err != nil {
		return "", err
	}

	params := url.Values{
		"info_hash": []string{string(t.InfoHash[:])},
		"peer_id":   []string{string(peerID[:])},
		"port":      []string{strconv.Itoa(int(port))},
		"uploaded":  []string{"0"},
		"download":  []string{"0"},
		"compact":   []string{"1"},
		"left":      []string{strconv.Itoa(t.Lenght)},
	}
	base.RawQuery = params.Encode()
	fmt.Println("Peers from the tracker URL: ", base.String())
	return base.String(), nil
}

type bencodeTrackerResp struct {
	Interval int    `bencode:"interval"`
	Peers    string `bencode:"peers"`
}

func (t *TorrentFile) requestPeers(peerID [20]byte, port uint16) ([]peers.Peer, error) {
	url, err := t.buildTrackerURL(peerID, port)
	if err != nil {
		return nil, err
	}

	c := &http.Client{Timeout: 15 * time.Second}
	resp, err := c.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	trackerResp := bencodeTrackerResp{}
	err = bencode.Unmarshal(resp.Body, &trackerResp)
	if err != nil {
		return nil, err
	}
	return peers.Unmarshal([]byte(trackerResp.Peers))
}
