package main

import (
	"fmt"
	"time"
)

type peer struct {
	id   string
	port chan request //change this to whatever type of struct request the peer wants

}
type response struct {
	failure_reason  string
	warning_message string
	interval        string
	complete        int
	incomplete      int
	peers           []peer
}
type request struct {
	info_hash  string //file name
	peer_id    string
	port       chan response
	uploaded   int  //size of bytes sent(just do size of request struct).
	downloaded int  //size of bytes receieved(size of response struct).
	left       int  //bytes left to download
	compact    bool //compact mode or no
	no_peer_id bool
	event      string // can be "started", "completed", "stopped"
	ip         string
	numwant    int
	key        int //unique ID to prove identity
	trackerid  int
}
type client_info struct {
	req      request
	interval int
	t        int
}

func handle_request(req request, peers map[string][]peer) {
	req.port <- response{"success", "no warning", "5", 0, 0, peers[req.info_hash]}
}
func tracker_server(recv chan request, exit chan int, peers map[string][]peer) {
	ticker := time.NewTicker(5000 * time.Millisecond)
	t := 0
	clients := make(map[string]client_info)
	for range ticker.C {
		t += 5
		select {
		case req := <-recv:
			temp_client := client_info{req, 5, t}
			clients[req.peer_id] = temp_client
			handle_request(req, peers)
		case <-exit:
			fmt.Println("Tracker server closing. ")
			return
		default:
			fmt.Println("tracker server standby. ")
		}
	}

}
