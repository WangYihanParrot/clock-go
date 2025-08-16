package main

import (
	_ "embed"
	"flag"
	"log"
	"net/http"
)

//go:embed clock.html
var clockHTML []byte

func main() {
	port := flag.String("port", "4000", "监听端口号")
	flag.Parse()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		if _, err := w.Write(clockHTML); err != nil {
			log.Printf("写入响应失败: %v\n", err)
		}
	})

	addr := ":" + *port
	log.Printf("启动 HTTP 服务器，监听 http://127.0.0.1%s\n", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("HTTP 服务器启动失败: %v\n", err)
	}
}
