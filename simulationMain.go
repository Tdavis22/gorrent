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

include(
    "sync"
    "fmt"
)


var wg sync.WaitGroup


/* Starts the BitTorrent protocol simulation. All simulation parameters should be set here. */
func startSimulation() {

    var num_peers int
    print_mutex := &sync.Mutex{}



    /* Initialize tracker, client, peers and channels. */
    //TODO wg.Add()

    /* Choose a file for each client to download. */
    //TODO e.g. init_client(file1.txt)

    /* Start */
    //TODO go tracker()
    
    //TODO go client()
    
    for i:= 0; i < num_peers; i++ {
        //TODO go peer()
    }

}












/* Launches the sim. */
func main() {
    
    startSimulation()
    wg.Wait()

    fmt.Printf("Simulation Complete.\n")
}
