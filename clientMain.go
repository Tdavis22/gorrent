package main

import (

)

/* Sends out a getrequest to the tracker, then using the peerlist, sends out a file request to the peers. */
func client_server(filename string, talk_to_tracker chan get_request) []byte {
	tracker_resp := make(chan tracker_response)
	talk_to_tracker <- get_request{filename, "0", tracker_resp, 0, 0, 0, true, "started", "", 30, "", ""}
	t_resp := <-tracker_resp //block waiting for tracker.
	peers := t_resp.peers
	peers_resp := make(chan other_message)
	for i := 0; i < len(peers); i++ {
		peer_connect := peers[i].port
		peer_connect <- other_message{i * 512, 2, nil, filename, peers_resp}
	}
	f := make([]byte, 2048)
	for data := range peers_resp {
		copy(f[data.offset:], data.payload)
	}
	return f
}
