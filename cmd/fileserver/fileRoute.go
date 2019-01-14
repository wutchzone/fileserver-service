package main

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/wutchzone/api-response"
	"github.com/wutchzone/fileserver-service/pkg/file-server"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
)

// HandleGetAll files
func HandleGetAll(w http.ResponseWriter, r *http.Request) {
	f, err := readDirectory(FileFolder, r)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_ = response.CreateResponse(w, response.ResponseError, nil)
	} else {
		_ = response.CreateResponse(w, response.ResponseOK, f)
	}
}

// HandleGetOne file
func HandleGetOne(w http.ResponseWriter, r *http.Request) {
	n := chi.URLParam(r, "name")
	f, err := ioutil.ReadFile(path.Join(FileFolder, n))

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	} else {
		w.Header().Set("Content-Type", "image/png")
		w.Write(f)
	}
}

func readDirectory(path string, r *http.Request) ([]localfile.File, error) {
	// Read DIR
	d, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}

	// Loop through local files and that can be send to the user
	var files []localfile.File
	for _, item := range d {
		files = append(files, *localfile.NewFile(
			item.Name(),
			fmt.Sprintf("%s/api/%s", BaseURL, url.PathEscape(item.Name())),
			"",
		))
	}

	return files, nil
}
