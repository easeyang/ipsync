package main

import (
	"fmt"
	"net"
	"strings"
	"net/http"
	"flag"
	"time"
	"log"
)

func getLocalIp() string {
	conn, err := net.Dial("udp", "www.easeyang.com:80")
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	defer conn.Close()
	return strings.Split(conn.LocalAddr().String(), ":")[0]
}

func uploadIp(host string) {
	ip := getLocalIp()
	if ip != "" {
		resp, err := http.Get(host + "/saveIp?ip=" + ip + "&port=9005")
		if err != nil {
			log.Print(err)
			return
		}
		defer resp.Body.Close()
		//body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Print(err)
		}
		//fmt.Println(string(body))
	}
}

func main() {
	addr := flag.String("addr", "http://www.easeyang.com", "server 服务器地址")
	//addr := flag.String("addr", "http://localhost:8000", "server 服务器地址")
	flag.Parse()
	fmt.Println(addr)

	t := time.Tick(time.Second*60)
	for {
		uploadIp(*addr)
		<-t
		fmt.Println("time")
	}
}