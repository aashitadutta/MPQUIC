package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	utils "./utils"
	quic "github.com/lucas-clemente/quic-go"
)

const addr = "0.0.0.0:4242"

func main() {

	quicConfig := &quic.Config{
		CreatePaths: true,
	}

	fmt.Println("Attaching to: ", addr)
	listener, err := quic.ListenAddr(addr, utils.GenerateTLSConfig(), quicConfig)
	utils.HandleError(err)

	fmt.Println("Server started! Waiting for streams from client...")

	sess, err := listener.Accept()
	utils.HandleError(err)

	fmt.Println("session created: ", sess.RemoteAddr())

	stream, err := sess.AcceptStream()
	utils.HandleError(err)

	defer stream.Close()
	defer stream.Close()

	fmt.Println("stream created: ", stream.StreamID())

	counter := 0
	for {
		// reply := make([]byte, 64000)

		reply_size := make([]byte, 20)
		// _, err = stream.Read(reply_size)
		_, err = io.ReadFull(stream, reply_size)

		size, _ := strconv.ParseInt(strings.Trim(string(reply_size), ":"), 10, 64)

		if size == 0 {
			break
		}

		println("frame size: ", size)

		reply := make([]byte, size)

		// stream.Read(reply)
		_, err = io.ReadFull(stream, reply)

		f, err := os.Create("sample/img" + strconv.Itoa(counter) + ".jpg")
		counter += 1
		// if counter == 100 {
		// 	break
		// }
		if err != nil {
			panic(err)
		}
		f.Write(reply)
		f.Close()
	}
	os.Exit(1)

}
