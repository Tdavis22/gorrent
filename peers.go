package main

import (
	"fmt"
	"os"
	"time"
)

//PIECE = 1
//REQUEST = 2
//CANCEL = 3
// HAVE = 4
// INTERESTED = 5
// KEEP-ALIVE = 6
func peer_server(files []string, recv chan other_message, connect chan handshake) {
	file_pointers := make(map[string]*os.File)
	//initialize files
	ticker := time.NewTicker(5000 * time.Millisecond)
	for _, filename := range files {
		f, err := os.Open(filename)
		check(err)
		file_pointers[filename] = f
	}
	for range ticker.C {
		//doesnt work needs a default case
		select {
		//im just going to use length as offset for now..
		case msg := <-recv:
			if msg.message_id == 2 {
				//filename field has to be added to other_message struct
				//also gonna need a port to send this data back to?
				filename := "file1.txt"
				fp := file_pointers[filename]
				fp.Seek(int64(msg.length), 0)
				b1 := make([]byte, 512)
				numBytesRead, err := fp.Read(b1)
				check(err)
				print(numBytesRead)
				//send b1(512 bytes) back to client aka construct new other message
				ret := other_message{512, 1, b1}
				fmt.Printf("%+v", ret)
			}

		}
	}
}
