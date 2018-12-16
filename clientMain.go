package main

import (
	"bufio"
	"fmt"
	"net"
)

const trackerUrl = "udp://tracker.coppersurfer.tk:6969/announce"

func updMessage() {
	p := make([]byte, 2048)
	conn, err := net.Dial("udp", "127.0.0.1:1234")
	if err != nil {
		fmt.Printf("Some error %v", err)
		return
	}
	fmt.Fprintf(conn, "Hi UDP Server, How are you doing?")
	_, err = bufio.NewReader(conn).Read(p)
	if err == nil {
		fmt.Printf("%s\n", p)
	} else {
		fmt.Printf("Some error %v\n", err)
	}
	conn.Close()
}
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
