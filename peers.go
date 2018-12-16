package main

import (
	"os"
	"time"
        "fmt"
)

//PIECE = 1
//REQUEST = 2


func check(e error) {
	if e != nil {
		panic(e)
	}
}

/* Waits for a request message from a client, then sends them their chunk of the data. */
func peer_server(files []string, recv chan other_message) {
	file_pointers := make(map[string]*os.File)
	//initialize files
	ticker := time.NewTicker(5000 * time.Millisecond)
	for _, filename := range files {
		f, err := os.Open(filename)
		check(err)
		file_pointers[filename] = f
	}
	for range ticker.C {
		select {
		// Use length as offset
		case msg := <-recv:
			if msg.message_id == 2 {
				filename := msg.file_name
				fp := file_pointers[filename]
				fp.Seek(int64(msg.offset), 0)
				b1 := make([]byte, 512)
				_, err := fp.Read(b1)
				check(err)
				//send b1(512 bytes) back to client aka construct new other message
				ret := other_message{msg.offset, 1, b1, filename, nil}
				msg.port <- ret
			}
                default:
                        fmt.Printf("")
		}
	}
}
