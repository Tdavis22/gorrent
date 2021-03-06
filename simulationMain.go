/* This is the main driver program for our Gorrent simulator.
 * Final Project CPE469.
 *
 * * BitTorrent Protocol Actors:
 *
 * Tracker - Webserver that holds connection data. Modeled as a goroutine.
 *           HTTP requests and other messages are modeled with channels.
 *
 * Client  - The user downloading a torrent. Initializes handshakes with other entities.
 *
 * Peer    - Theoretially equivalent to a client. Modeled here as the users who source the files.
 *
 *
 * * The simulation runs as follows:
 *
 * 1. Client performs a handshake with the tracker.
 * 2. Client asks tracker for a list of peers to download a file from.
 * 3. Tracker accepts or denys the request. Sends list upon accept.
 * 4. Client performs a handshake with the peers from the list.
 * 5. Client indicates interest in a file from the peers.
 * 6. Client requests pieces of the file from the peers.
 * 7. Peers send the pieces to the client.
 *
 */

package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

/* Starts the BitTorrent protocol simulation. All simulation parameters should be set here. */
func startSimulation() {

	var num_peers int = 4
	print_mutex := &sync.Mutex{}

	/* Initialize tracker, client, peers and channels. */

	/* For tracker requests. */
	c_2_t := make(chan get_request)
	stop := make(chan int)

	/* For peer communications. */
	c_2_p0 := make(chan other_message)
	c_2_p1 := make(chan other_message)
	c_2_p2 := make(chan other_message)
	c_2_p3 := make(chan other_message)

	/* Group up the peers and give the tracker some files to manage. */
	peer_com := []peer{peer{"0", c_2_p0}, peer{"1", c_2_p1},
		peer{"2", c_2_p2}, peer{"3", c_2_p3}}
	file_data := init_file_data(peer_com)

	/* Start the goroutines */
	go tracker_server(c_2_t, stop, file_data, print_mutex)

	go client_server("test1", c_2_t)

	str := make([]string, 1)
	str[0] = "test1"

	peer_server(str, c_2_p0)
	peer_server(str, c_2_p1)
	peer_server(str, c_2_p2)
	peer_server(str, c_2_p3)

	wg.Add(2 + num_peers)
}

/* Populates the map with one entry: test1. */
func init_file_data(peer_list []peer) map[string]file_info {
	file_data := make(map[string]file_info)

	torrent := file_info{"test1", peer_list, 0, 0}
	file_data["test1"] = torrent

	return file_data
}

/* Launches the sim. */
func main() {

	startSimulation()
	wg.Wait()

	fmt.Printf("Simulation Complete.\n")
}
