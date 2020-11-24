package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"net"
	"os"
	"strconv"
	"strings"

	utils "./utils"
	quic "github.com/lucas-clemente/quic-go"
)

func main() {

	// tcp connection
	servAddr := "localhost:8002"
	tcpAddr, err := net.ResolveTCPAddr("tcp", servAddr)
	if err != nil {
		println("ResolveTCPAddr failed:", err.Error())
		os.Exit(1)
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		println("Dial failed:", err.Error())
		os.Exit(1)
	}

	defer conn.Close()

	// quic config
	const addr = "localhost:4242"
	quicConfig := &quic.Config{
		CreatePaths: true,
	}

	sess, err := quic.DialAddr(addr, &tls.Config{InsecureSkipVerify: true}, quicConfig)
	utils.HandleError(err)

	fmt.Println("session created: ", sess.RemoteAddr())

	stream, err := sess.OpenStream()
	utils.HandleError(err)
	defer stream.Close()
	defer stream.Close()

	fmt.Println("stream created...")
	fmt.Println("Client connected")
	counter := 0

	for {

		reply_size := make([]byte, 20)
		// _, err = conn.Read(reply_size)
		_, err = io.ReadFull(conn, reply_size)

		if reply_size == nil {
			break
		}

		size, _ := strconv.ParseInt(strings.Trim(string(reply_size), ":"), 10, 64)

		if size == 0 {
			stream.Write(reply_size)
			break
		}

		// if size == 0 {
		// 	// break
		// 	time.Sleep(10 * time.Millisecond)
		// 	continue
		// }
		print("frame size: ", size)

		reply := make([]byte, size)
		println("waiting for server's reply...")
		// _, err = conn.Read(reply)
		_, err = io.ReadFull(conn, reply)
		// println(reply)
		// break
		if err != nil {
			println("read to server failed:", err.Error())
			os.Exit(1)
		}

		println("reply from server: ", len((reply)), "bytes")

		// f, err := os.Create("sample/img.jpg")
		// if err != nil {
		// 	panic(err)
		// }
		// f.Write(reply)
		// f.Close()
		// os.Exit(1)

		// now send this via mpquic

		// f, err := os.Create("sample/img" + strconv.Itoa(counter) + "_server.jpg")
		// if err != nil {
		// 	panic(err)
		// }
		counter += 1
		// f.Write(reply)
		// f.Close()

		stream.Write(reply_size)
		stream.Write(reply)

	}
}
