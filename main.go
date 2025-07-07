package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

func main() {
	port := flag.String("port", "4000", "监听端口号")
	flag.Parse()

	content, err := os.ReadFile("clock.html")
	if err != nil {
		log.Fatalf("读取文件失败: %v\n", err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		if _, err := w.Write(content); err != nil {
			log.Printf("写入响应失败: %v\n", err)
		}
	})

	addr := ":" + *port
	log.Printf("启动 HTTP 服务器，监听 http://127.0.0.1%s\n", addr)
	err = http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatalf("HTTP 服务器启动失败: %v\n", err)
	}
}
