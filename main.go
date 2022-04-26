package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func init() {
	LoadMobileAddress()
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		static := "static" + r.URL.Path
		log.Printf("'%s', '%s'\n", r.URL.Path, static)
		http.ServeFile(w, r, static)
	})
	// 路由与视图函数绑定
	http.HandleFunc("/mobileAddress", mobileAddressHandler)

	// 启动服务,监听地址
	err := http.ListenAndServe(":80", nil)
	if err == nil {
		return
	}
	err = http.ListenAndServe(":8080", nil)
	if err == nil {
		return
	}
	err = http.ListenAndServe(":9999", nil)
	if err != nil {
		log.Fatal(err)
	}

}

func mobileAddressHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	phone := r.FormValue("phone")
	log.Println("phone number", phone)

	address := QueryMobile(phone)
	data, err := json.Marshal(address)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, `{"error":"error"}`)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	_, err = io.WriteString(w, string(data))
	if err != nil {
		log.Println("reponse err ", err)
	}
}
