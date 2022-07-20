package main

import (
	"fmt"
	"go-qr-app/qrgen"
	"image/png"
	"log"
	"net/http"
	"strconv"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	msg := r.URL.Query().Get("message")
	// Printlnなどとは違い、HTTPのレスポンスやファイルなど標準出力以外のものにも文字列を書き込める
	fmt.Fprintf(w, "Hello %s", msg)
}

func generateQRCode(w http.ResponseWriter, r *http.Request) {
	param := r.URL.Query()
	url := param.Get("url")
	if url == "" {
		http.Error(w, "missing url", http.StatusBadRequest)
		return
	}
	widthStr := param.Get("width")
	width, err := strconv.Atoi(widthStr)
	if err != nil {
		http.Error(w, "invalid width", http.StatusBadRequest)
		return
	}
	heightStr := param.Get("height")
	height, err := strconv.Atoi(heightStr)
	if err != nil {
		http.Error(w, "invalid height", http.StatusBadRequest)
		return
	}

	if width <= 45 {
		http.Error(w, "width must be more than 45px", http.StatusBadRequest)
		return
	}

	if height <= 45 {
		http.Error(w, "height must be more than 45px", http.StatusBadRequest)
		return
	}

	img, err := qrgen.GenQRCode(url, width, height)
	if err != nil {
		errStr := fmt.Sprintf("failed to generate QR code%v", err)
		http.Error(w, errStr, http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "image/png")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	png.Encode(w, img)
}

func main() {
	// ルーティング
	// 第１引数：パス
	// 第２引数：呼ばれる関数
	http.HandleFunc("/", sayHello)
	http.HandleFunc("/generate", generateQRCode)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("failed to create server: %v", err)
	}
}
