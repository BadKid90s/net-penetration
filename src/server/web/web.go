package main

import (
	"encoding/json"
	"log"
	"net-penetration/define"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		values := request.URL.Query()
		marshal, err := json.Marshal(values)
		if err != nil {
			log.Printf("Marshal Error: %v", err)
		}
		writer.Write(marshal)
	})
	log.Println("本地服务已启动" + define.LocalServerAddress)
	http.ListenAndServe(define.LocalServerAddress, nil)
}
