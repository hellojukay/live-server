package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
)

var (
	dir string
)

func init() {
	flag.StringVar(&dir, "d", "./", "serve directory")
	flag.Parse()
}
func main() {

	http.Handle("/", http.FileServer(http.Dir(dir)))
	log.Printf("runing on http://127.0.0.1:8080 and http://%s:8080\n", IP())

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}

}

func IP() string {
	conn, _ := net.Dial("udp", "8.8.8.8:80")
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP.String()
}
