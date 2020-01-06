package main

import "fmt"

func main() {
	tampil_mahasiswa()
}

func tampil_mahasiswa() {
	var mahasiswa = map[string]int{"Aldo": 182, "Yosep": 178}
	for key, val := range mahasiswa {
		fmt.Println(key, "\t:", val)
	}
}
