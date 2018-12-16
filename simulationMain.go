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
    "sync"
    "fmt"
)


var wg sync.WaitGroup


/* Starts the BitTorrent protocol simulation. All simulation parameters should be set here. */
func startSimulation() {

    var num_peers int
    print_mutex := &sync.Mutex{}

    

    /* Initialize tracker, client, peers and channels. */
    hs := make(chan handshake)
    c_2_t := make(chan get_request)
    t_2_c := make(chan tracker_response)

    file_data := init_file_data()

    //TODO wg.Add()

    /* Choose a file for each client to download. */
    //TODO e.g. init_client(file1.txt)

    /* Start */
    go tracker(c_2_t,t_2_c,file_data,print_mutex)

    
    //TODO go client()
    
    for i:= 0; i < num_peers; i++ {
        //TODO go peer()
    }

}



/* Populates the map with one entry: test1. */
func init_file_data() map[string]file_info {
    file_data := make(map[string]file_info)

    var peer1 peer
    var peer_list []peer
    peer_list = append(peer_list, peer1)

    torrent := file_info{"test1",peer_list,0,0}
    file_data["test1"] = torrent

    return file_data
}




/* Launches the sim. */
func main() {
    
    startSimulation()
    wg.Wait()

    fmt.Printf("Simulation Complete.\n")
}
