package main

import (
	"fmt"
	"os"
	"strings"
)

type Student struct {
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

type Students []Student

var students Students

func init() {
	students = Students{
		Student{
			Nama:      "Thomas",
			Alamat:    "Kota A",
			Pekerjaan: "A",
			Alasan:    "Alasan Thomas",
		},
		Student{
			Nama:      "Fauzan",
			Alamat:    "Kota B",
			Pekerjaan: "A",
			Alasan:    "Alasan Fauzan",
		},
		Student{
			Nama:      "Indra",
			Alamat:    "Kota C",
			Pekerjaan: "A",
			Alasan:    "Alasan Indra",
		},
	}
}

func main() {

	//os.Args to get the param from terminal

	//validate if os.Args < 2 , then this is invalid format
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run biodata.go <name>")
		return
	}

	//declare closure
	var sByName = searchByName()

	//initiate using closure function
	siswa := sByName(os.Args[1]) //using args slice index number 1 (name)
	fmt.Println(siswa)

}

// define function closure
func searchByName() func(name string) string {

	return func(name string) string {
		var murid = "Nama : " + name + " tidak ditemukan"
		for _, val := range students {
			if name == val.Nama {
				muridEl := []string{"Nama: " + val.Nama, "Alamat: " + val.Alamat, "Pekerjaan: " + val.Pekerjaan, "Alasan: " + val.Alasan}
				murid = strings.Join(muridEl, "\n")
				break
			}
		}

		return murid
	}
}
