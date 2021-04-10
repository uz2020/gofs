package main

import (
	"flag"
	"fmt"
	"gopkg.in/ini.v1"
	"log"
	"net"
	"strconv"
)

var port *int = flag.Int("p", 8080, "监听端口")
var conf *string = flag.String("c", "gofs.conf", "配置文件")

func handleConnection(conn net.Conn) {
	log.Println("处理请求:", conn)
}

func parseConf() {
	flag.Parse()

	if *conf != "" {
		cfg, err := ini.Load(*conf)
		if err != nil {
			log.Fatalln("加载配置文件失败: %v", err)
		}

		portStr := cfg.Section("server").Key("port").String()
		p, err := strconv.Atoi(portStr)
		if err != nil {
			log.Fatalln("端口配置错误: %v", err)
		}
		*port = p
	}
}

func main() {
	parseConf()

	log.Println("监听端口:", *port)

	addr := fmt.Sprintf(":%d", *port)

	ln, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalln("监听端口失败:", err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatalln("接受请求失败:", err)
		}
		go handleConnection(conn)
	}
}
