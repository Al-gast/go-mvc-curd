package main

import (
	"fmt"
	"go-curd/controllers/pasiencontroller"
	"net/http"
)

func main() {

	http.HandleFunc("/", pasiencontroller.Index)
	http.HandleFunc("/pasien", pasiencontroller.Index)
	http.HandleFunc("/pasien/index", pasiencontroller.Index)
	http.HandleFunc("/pasien/add", pasiencontroller.CreatePasien)
	http.HandleFunc("/pasien/edit", pasiencontroller.UpdatePasien)
	http.HandleFunc("/pasien/delete", pasiencontroller.DeletePasien)

	fmt.Println("server running localhost:3000")
	http.ListenAndServe(":3000", nil)

}
