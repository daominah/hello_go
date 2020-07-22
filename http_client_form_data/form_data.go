package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"net/http/httputil"
	"os"
	"time"
)

const cwd = "/home/tungdt/go/src/github.com/daominah/hello_go/http_request_form_data"

func main() {
	log.SetFlags(log.Lshortfile)

	go func() {
		err := http.ListenAndServe(":20891", fileHandler())
		if err != nil {
			log.Fatal(err)
		}
	}()

	req := createPostFilesReq(map[string]string{
		"file1": cwd + "/file1.jpg",
		"file2": cwd + "/file2.txt",
	})
	_, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	time.Sleep(100 * time.Millisecond)
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

	req, _ := http.NewRequest("POST", "http://localhost:20891/", body)
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
			maxMemory := 100 * 1024 * 1024 // 100MB
			r.ParseMultipartForm(int64(maxMemory))
			file, _, err := r.FormFile("file2")
			if err != nil {
				log.Fatal(err)
			}
			defer file.Close()
			data, err := ioutil.ReadAll(file)
			if err != nil {
				log.Fatal(err)
			}
			targetFile := cwd + "/savedFile2.txt"
			err = ioutil.WriteFile(targetFile, data, 0644)
			log.Printf("wrote file %v: %v\n", targetFile, err)
		},
	)
	return h
}
