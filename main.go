package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		content, err := os.ReadFile("clock.html")
		if err != nil {
			http.Error(w, "无法读取 clock.html", http.StatusInternalServerError)
			log.Printf("读取文件失败: %v\n", err)
			return
		}

		w.Header().Set("Content-Type", "text/html; charset=utf-8")

		if _, err := w.Write(content); err != nil {
			log.Printf("写入响应失败: %v\n", err)
		}
	})

	log.Println("启动 HTTP 服务器，监听 http://127.0.0.1:4000")

	err := http.ListenAndServe(":4000", nil)
	if err != nil {
		log.Fatalf("HTTP 服务器启动失败: %v\n", err)
	}
}
