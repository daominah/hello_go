package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"net/http/httputil"
	"os"
	"time"
)

const cwd = "/home/tungdt/go/src/github.com/daominah/hello_go/http_client_server_form_data"
const listen = ":20892"

func main() {
	log.SetFlags(log.Lshortfile)

	go func() {
		err := http.ListenAndServe(listen, fileHandler())
		if err != nil {
			log.Fatal(err)
		}
	}()

	req := createPostFilesReq(map[string]string{
		"files0": cwd + "/files0.txt",
		"files1": cwd + "/files1.jpg",
	})
	_, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	_ = time.Sleep
	//time.Sleep(100 * time.Millisecond)
	select {}
}

// filePaths map form key (server read file by this key) to filePath
func createPostFilesReq(filePaths map[string]string) *http.Request {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	for key, filePath := range filePaths {
		file, err := os.Open(filePath)
		if err != nil {
			panic(err)
		}
		fw, err := writer.CreateFormFile(key, file.Name())
		if err != nil {
			panic(err)
		}
		_, err = io.Copy(fw, file)
		if err != nil {
			panic(err)
		}
		file.Close()
	}
	writer.Close()

	req, _ := http.NewRequest("POST",
		fmt.Sprintf("http://localhost%v/", listen), body)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	if !true {
		dumpedRequest, err := httputil.DumpRequestOut(req, true)
		log.Printf("err: %v. req:\n%v\n", err, string(dumpedRequest))
	}

	return req
}

func fileHandler() *http.ServeMux {
	h := http.NewServeMux()
	h.HandleFunc(
		"/", func(w http.ResponseWriter, r *http.Request) {
			reqStr, _ := httputil.DumpRequest(r, true)
			log.Printf("received request: %s\n", reqStr)
			maxMemory := 100 * 1024 * 1024 // 100MB
			r.ParseMultipartForm(int64(maxMemory))
			file, _, err := r.FormFile("files0")
			if err != nil {
				log.Println(err)
				return
			}
			defer file.Close()
			data, err := ioutil.ReadAll(file)
			if err != nil {
				log.Println(err)
				return
			}
			targetFile := cwd + "/savedFiles0.txt"
			err = ioutil.WriteFile(targetFile, data, 0644)
			log.Printf("wrote file %v: %v\n", targetFile, err)

			//# CORS spec: Your server will need to validate the origin header
			// and then you can echo the origin value in the response header
			w.Header().Set("Access-Control-Allow-Origin",
				"*")
			w.WriteHeader(200)
			w.Write([]byte("received file"))
		},
	)
	return h
}
