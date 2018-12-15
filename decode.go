package main

import (
	"bytes"
	"github.com/jackpal/bencode-go"
	"io/ioutil"
)

type torrentMetaData struct {
	Announce string
	Encoding string
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func getMetaData(f string) torrentMetaData {
	dat, err := ioutil.ReadFile(f)
	check(err)
	r := bytes.NewReader(dat)
	torrent_obj := torrentMetaData{}
	err = bencode.Unmarshal(r, &torrent_obj)
	check(err)
	return torrent_obj
}

/*func main(){
	dat, err := ioutil.ReadFile("puppy.torrent")
	if err != nil{
		print("ohno")
		panic(err)
	}
	r := bytes.NewReader(dat)
	data, error := bencode.Decode(r)
	if error != nil{
		print("ohno")
		panic(error)
	}
	print(data)
	r.Seek(0, 0)
	torrent_obj := torrentMetaData{}
	error1 := bencode.Unmarshal(r, &torrent_obj)
	if error1 != nil{
		print("ohno")
		panic(error1)
	}
	fmt.Printf("%+v\n", torrent_obj)
}
*/
