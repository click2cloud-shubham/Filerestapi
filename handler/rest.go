package handler

import (
	"encoding/json"
	"net/http"
	"todoproject/disk"
)

type RestHandler struct {
	disk disk.Disk
}

type Rest struct {
	NewFile         string `json:"new_file"`
	CurrentFileName string `json:"current_filename"`
	NewFileName     string `json:"new_filename"`
	Path            string `json:path`
}

// NewRestHandler is a constructor for RestHandler
func NewRestHandler(disk disk.Disk) RestHandler {
	return RestHandler{disk: disk}
}

func (rh RestHandler) CreateFile(r http.ResponseWriter, w *http.Request) {
	var ra Rest
	if err := json.NewDecoder(w.Body).Decode(&ra); err != nil {
		r.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := rh.disk.CreateFile(ra.NewFile); err != nil {
		r.WriteHeader(http.StatusInternalServerError)
		return
	} //list all the files

	r.WriteHeader(http.StatusCreated)
}

func (rh RestHandler) RenameFile(r http.ResponseWriter, w *http.Request) {
	var ra Rest
	if err := json.NewDecoder(w.Body).Decode(&ra); err != nil {
		r.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := rh.disk.RenameFile(ra.CurrentFileName, ra.NewFileName); err != nil {
		r.WriteHeader(http.StatusInternalServerError)
		return
	}

	r.WriteHeader(http.StatusOK)
}

func (rh RestHandler) DeleteFile(r http.ResponseWriter, w *http.Request) {
	var ra Rest
	if err := json.NewDecoder(w.Body).Decode(&ra); err != nil {
		r.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := rh.disk.DeleteFile(ra.CurrentFileName); err != nil {
		r.WriteHeader(http.StatusInternalServerError)
		return
	}

	r.WriteHeader(http.StatusOK)
}

func (rh RestHandler) ListWorkspace(r http.ResponseWriter, w *http.Request) {
	fileNames, err := rh.disk.ListWorkspace()
	if err != nil {
		r.WriteHeader(http.StatusInternalServerError)
		return
	}

	b, err := json.Marshal(fileNames)
	if err != nil {
		r.WriteHeader(http.StatusInternalServerError)
		return
	}
	r.WriteHeader(http.StatusOK)
	r.Write(b)
}
