package helloworld

import (
	"archive/zip"
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
)

// this is Google Cloud function syntax
func init() {
	functions.HTTP("HelloHTTP", uploadHandler)
}

// main is for local testing
func _main() {
	// For local testing, start HTTP server on :8080
	http.HandleFunc("/upload", uploadHandler)
	log.Println("http://localhost:8080/upload")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	callerIP := r.RemoteAddr
	if ip := r.Header.Get("X-Forwarded-For"); ip != "" {
		callerIP = ip
	}
	log.Printf("handling request from IP: %v", callerIP)

	if r.Method == http.MethodGet {
		uploadFormTmpl.Execute(w, nil)
		return
	}
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		http.Error(w, "error parsing form", http.StatusBadRequest)
		return
	}
	file, header, err := r.FormFile("logfile")
	if err != nil {
		http.Error(w, "error reading file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	var data []byte
	ext := filepath.Ext(header.Filename)
	if ext == ".zip" {
		data, err = extractFirstLogFileFromZip(file)
		if err != nil {
			http.Error(w, "error extracting zip: "+err.Error(), http.StatusBadRequest)
			return
		}
	} else {
		data, err = io.ReadAll(file)
		if err != nil {
			http.Error(w, "error reading file: "+err.Error(), http.StatusBadRequest)
			return
		}
	}

	ids := processLogData(data)
	result := struct {
		Result bool
		Count  int
		IDs    []string
	}{
		Result: true,
		Count:  len(ids),
		IDs:    ids,
	}
	sort.Strings(ids)
	err = uploadFormTmpl.Execute(w, result)
	if err != nil {
		http.Error(w, "error rendering template: "+err.Error(), http.StatusInternalServerError)
	}
}

func extractFirstLogFileFromZip(file multipart.File) ([]byte, error) {
	buf, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}
	zr, err := zip.NewReader(bytes.NewReader(buf), int64(len(buf)))
	if err != nil {
		return nil, err
	}
	for _, f := range zr.File {
		if !f.FileInfo().IsDir() && (filepath.Ext(f.Name) == ".txt" || filepath.Ext(f.Name) == ".ndjson") {
			rc, err := f.Open()
			if err != nil {
				return nil, err
			}
			defer rc.Close()
			return io.ReadAll(rc)
		}
	}
	return nil, fmt.Errorf("no .txt or .ndjson file found in zip")
}

func processLogData(data []byte) []string {
	lines := strings.Split(string(data), "\n")
	ruvIDs := make(map[string]bool)
	nLines := 0
	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			continue
		}
		nLines++
		var v BetterStackLogRow
		err := json.Unmarshal([]byte(line), &v)
		if err != nil {
			continue
		}
		if v.Message.RequestMethod != "DELETE" {
			continue
		}
		jwtTokenIdx := strings.Index(v.Message.RequestPath, "token=")
		if jwtTokenIdx == -1 {
			continue
		}
		token := v.Message.RequestPath[jwtTokenIdx+6:]
		if ampIdx := strings.Index(token, "&"); ampIdx != -1 {
			token = token[:ampIdx]
		}
		var tokenDelete TokenDeleteCase
		parts := strings.Split(token, ".")
		if len(parts) != 3 {
			continue
		}
		payloadBytes, err := base64.RawURLEncoding.DecodeString(parts[1])
		if err != nil {
			continue
		}
		err = json.Unmarshal(payloadBytes, &tokenDelete)
		if err != nil {
			continue
		}
		ruvIDs[tokenDelete.PassDelete.ExtId] = true
	}
	ids := make([]string, 0, len(ruvIDs))
	for id := range ruvIDs {
		ids = append(ids, id)
	}
	log.Printf("input data size %v bytes, %v lines: result %v IDs",
		len(data), nLines, len(ids))
	return ids
}

type BetterStackLogRow struct {
	App     string `json:"_app"` // "prd"
	Message struct {
		ClientAddr            string `json:"ClientAddr"`
		ClientHost            string `json:"ClientHost"`
		ClientPort            string `json:"ClientPort"`
		ClientUsername        string `json:"ClientUsername"`
		DownstreamContentSize int    `json:"DownstreamContentSize"`
		DownstreamStatus      int    `json:"DownstreamStatus"`
		Duration              int64  `json:"Duration"`
		OriginContentSize     int    `json:"OriginContentSize"`
		OriginDuration        int64  `json:"OriginDuration"`
		OriginStatus          int    `json:"OriginStatus"`
		Overhead              int    `json:"Overhead"`
		RequestAddr           string `json:"RequestAddr"`
		RequestContentSize    int    `json:"RequestContentSize"`
		RequestCount          int    `json:"RequestCount"`
		RequestHost           string `json:"RequestHost"`
		RequestMethod         string `json:"RequestMethod"`
		RequestPath           string `json:"RequestPath"`
		RequestPort           string `json:"RequestPort"`
		RequestProtocol       string `json:"RequestProtocol"`
		RequestScheme         string `json:"RequestScheme"`
		RetryAttempts         int    `json:"RetryAttempts"`
		RouterName            string `json:"RouterName"`
		ServiceAddr           string `json:"ServiceAddr"`
		ServiceName           string `json:"ServiceName"`
		ServiceURL            struct {
			Scheme      string      `json:"Scheme"`
			Opaque      string      `json:"Opaque"`
			User        interface{} `json:"User"`
			Host        string      `json:"Host"`
			Path        string      `json:"Path"`
			RawPath     string      `json:"RawPath"`
			OmitHost    bool        `json:"OmitHost"`
			ForceQuery  bool        `json:"ForceQuery"`
			RawQuery    string      `json:"RawQuery"`
			Fragment    string      `json:"Fragment"`
			RawFragment string      `json:"RawFragment"`
		} `json:"ServiceURL"`
		StartLocal     time.Time `json:"StartLocal"`
		StartUTC       time.Time `json:"StartUTC"`
		EntryPointName string    `json:"entryPointName"`
		Level          string    `json:"level"`
		Msg            string    `json:"msg"`
		Time           time.Time `json:"time"`
	} `json:"message"`
	SourceType string    `json:"source_type"`
	Stream     string    `json:"stream"`
	Timestamp  time.Time `json:"timestamp"`
}

type TokenDeleteCase struct {
	Exp        int    `json:"exp"`
	Jti        string `json:"jti"`
	PassDelete struct {
		ExtId     string `json:"extId"`
		ProjectId int    `json:"projectId"`
	} `json:"pass.delete"`
}

// uploadFormTmpl is the HTML for input file and result display
var uploadFormTmpl = template.Must(template.New("upload").Parse(`
<!DOCTYPE html>
<html>
<body>
<h2>Upload log file downloaded from BatterStack (.zip or .txt/.ndjson)</h2>
<form enctype="multipart/form-data" action="/upload" method="post">
	<input type="file" name="logfile" accept=".zip,.txt,.ndjson" required><br><br>
	<input type="submit" value="Parse to get externalIDs">
</form>
{{if .Result}}
<hr>
<h3>Result</h3>
<p>Count external IDs: {{.Count}}</p>
<ul>
	{{range .IDs}}
	<li>{{.}}</li>
	{{end}}
</ul>
{{end}}
</body>
</html>
`))
