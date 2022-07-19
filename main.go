package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
)

var (
	dir  string
	port = 8080
)

func init() {
	flag.StringVar(&dir, "d", "./", "Serve directory")
	flag.IntVar(&port, "port", port, "Specify alternate bind port")
	flag.Parse()
}
func main() {

	http.Handle("/", http.FileServer(http.Dir(dir)))
	log.Printf("runing on http://127.0.0.1:%d and http://%s:%d\n", port, GetLocalIP(), port)

	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		fmt.Println(err)
	}

}

// GetLocalIP get local ipv4 address
func GetLocalIP() string {
	conn, _ := net.Dial("udp", "8.8.8.8:80")
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP.String()
}
