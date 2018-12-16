package main

import (
	"fmt"
	"time"
)

/* Staged for deletion. Replaced by tracker2.0
type peer struct {
	id   string
	port chan get_request //change this to whatever type of struct get_request the peer wants
}
type client_info struct {
	req      get_request
	t        int
}
func handle_request(req get_request, who_has map[string][]peer) {
	req.port <- tracker_response{"success", "", 5, 0, "*_*", num_seeders, num_leechers,
                                     who_has[req.info_hash]}
}
func tracker_server(recv chan get_request, exit chan int, peers map[string][]peer) {
	ticker := time.NewTicker(5000 * time.Millisecond)
	clients := make(map[string]client_info)
	for range ticker.C {
		select {
		case req := <-recv:
			temp_client := client_info{req, t}
			clients[req.peer_id] = temp_client
			handle_request(req, peers)
		case <-exit:
			fmt.Println("Tracker server closing.")
			return
		default:
			fmt.Println("Tracker server standby.")
		}
	}

}
*/

/* Waits for a get request, sees if the file exists, and returns a response accordingly. */
func tracker(recv chan get_request, exit chan int, who_has map[string]file_info) {
	ticker := time.NewTicker(5000 * time.Millisecond)
	for range ticker.C {
		select {
		case req := <-recv:
			_, ok := who_has[req.info_hash]
			if ok {
				req.port <- tracker_response{"", "success, file found", 5, 0, "*_*",
					who_has[req.info_hash].num_seeders,
					who_has[req.info_hash].num_leechers,
					who_has[req.info_hash].peer_list}
			} else {
				req.port <- tracker_response{"File not found", "Request denied", 0, 0, "",
					0, 0, nil}
			}
		case <-exit:
			fmt.Println("Tracker server closing.")
			return
		default:
			fmt.Println("Tracker server standby.")
		}
	}

}
