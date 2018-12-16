package main

import (
	"fmt"
	"time"
        "sync"
)



/* Waits for a get request, sees if the file exists, and returns a response accordingly. */
func tracker_server(recv chan get_request, exit chan int, who_has map[string]file_info,print_mutex *sync.Mutex) {
    ticker := time.NewTicker(5000 * time.Millisecond)
    for range ticker.C {
	select {
	    case req := <-recv:
                _,ok := who_has[req.info_hash]
                if ok {
	            req.port <- tracker_response{"", "success, file found", 5, 0, "*_*", 
                                                 who_has[req.info_hash].num_seeders, 
                                                 who_has[req.info_hash].num_leechers, 
                                                 who_has[req.info_hash].peer_list}
                } else {
                    req.port <- tracker_response{"File not found", "Request denied", 0,0,"",
                                                 0,0,nil}
                }
	    case <-exit:
		fmt.Println("Tracker server closing.")
		return
	    default:
	        fmt.Println("Tracker server standby.")
        }
    }
}
