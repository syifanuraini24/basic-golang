package main

import "fmt"

// Buatlah map dengan key "Nama Provinsi" pada pulau Kalimantan dan value nya adalah "Ibu Kota" provinsi tersebut
// Output satu semua key dan value yang ada di map tersebut
func main() {
	// TODO: answer here
	var pulau map[string]string
	pulau = map[string]string{}

	pulau["Kalimantan Barat"] = "Pontianak"
	fmt.Println("Kalimantan Barat", pulau["Kalimantan Barat"])
}
