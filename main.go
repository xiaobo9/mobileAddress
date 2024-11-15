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
		log.Printf("'%s','%s', '%s'\n", remoteAddr(r), r.URL.Path, static)
		http.ServeFile(w, r, static)
	})

	http.HandleFunc("/mobileAddress", mobileAddressHandler)

	ports := []string{"8080", "9999"}
	var err error
	for _, v := range ports {
		err = http.ListenAndServe(":"+v, nil)
		if err != nil {
			log.Printf("'%s' 监听失败\b", v)
		}
	}
	log.Fatal(err)
}

func mobileAddressHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	phone := r.FormValue("phone")
	log.Printf("'%s','%s', '%s'\n", remoteAddr(r), r.URL.Path, phone)

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
		log.Println("response err ", err)
	}
}

// proxy_set_header X-Real-IP $remote_addr;
// proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
func remoteAddr(r *http.Request) string {
	header := r.Header
	remoteAddr := header.Get("X-Forwarded-For")
	if remoteAddr != "" {
		return remoteAddr
	}
	return r.RemoteAddr
}
