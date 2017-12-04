package main

import (
	"net/http"
	"log"
	"flag"
	"html/template"
	"strconv"
)

var Ip string = "localhost"
var Port int64 = 8000

//保存保存客户端上传的ip地址
func saveIP(resp http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	log.Println("saveIP", req.Body, req.Host, req.Form)
	Ip = req.Form.Get("ip")
	protStr := req.Form.Get("port")
	Port, _ = strconv.ParseInt(protStr, 10, 64)
	log.Println("Port", Port)
}

//dev首页
func devIndex(resp http.ResponseWriter, req *http.Request) {
	log.Println("devIndex")
	t, err := template.ParseFiles("./tpl.html")
	if err != nil {
		resp.Write([]byte("error"))
		log.Print(err)
	}else {
		data := struct {
			Ip string
			Port int64
		}{
			Ip,
			Port,
		}
		t.Execute(resp, data)
	}
}

func registerHandler() {
	http.HandleFunc("/", devIndex)
	http.HandleFunc("/saveIp", saveIP)
}

func main() {
	log.Println("#开始启动服务器#")

	port := flag.String("port", "8000", "服务监听端口")
	flag.Parse()
	log.Println("prot=" + *port)

	registerHandler()

	err := http.ListenAndServe(":" + *port, nil)
	if(err != nil) {
		log.Println(err)
	}
	log.Println("服务结束")
}
