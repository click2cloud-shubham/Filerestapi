
package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"todoproject/disk"
	"todoproject/handler"
)

func main() {
	workspace := "C:\\mytemporaryfiles"
	d, err := disk.NewDisk(workspace)
	if err != nil {
		log.Fatalln(err)
	}

	rh := handler.NewRestHandler(d)

	r := mux.NewRouter()
	r.HandleFunc("/create", rh.CreateFile).Methods(http.MethodPost)
	r.HandleFunc("/rename", rh.RenameFile).Methods(http.MethodPut)
	r.HandleFunc("/delete", rh.DeleteFile).Methods(http.MethodDelete)
	r.HandleFunc("/listfiles", rh.ListWorkspace).Methods(http.MethodGet)

	http.ListenAndServe(":8008", r)
}
