package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/meta-byte/photo-app-server/client"
	"github.com/meta-byte/photo-app-server/util"
)

type dataString struct {
	Data string `json:"Data"`
}

func main() {
	initServer()
}

func initServer() {
	client.InitS3()
	fmt.Println("s3 done")

	mux := route()
	fmt.Println(http.ListenAndServe(":8080", &mux))
}

func route() http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", HelloServer)
	mux.HandleFunc("/data", getData)
	mux.HandleFunc("/upload", upload)
	mux.HandleFunc("/download", download)
	return *mux
}

func HelloServer(writer http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(writer, "Hello, World!")
}

func getData(writer http.ResponseWriter, req *http.Request) {
	data := dataString{
		Data: "This is some random data.",
	}

	json.NewEncoder(writer).Encode(data)
}

func upload(writer http.ResponseWriter, req *http.Request) {
	path := "gun.jpeg"
	file, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	client.PutObj(client.Client, path, file)
}

func download(writer http.ResponseWriter, req *http.Request) {
	key := "gun.jpeg"
	path := util.IncrementDownload("Download") + ".jpg"
	file, err := os.Create(path)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	client.GetObj(client.Client, key, file)
}
