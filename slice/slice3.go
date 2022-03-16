package main

import "fmt"

// Lanjutan nomor 2
// Sehabis menambahkan "Olleh" pada slice tersebut coba ubah nilai "World" menjadi "Marcus"
// dan "Olleh" menjadi "Aurelius"
func main() {
	slice := []string{"Hello", "World"}

	slice = append(slice,"Olleh")

	fmt.Println(len(slice),cap(slice))
	fmt.Println(slice)
	// Dibawah ini adalah jawaban nomor 3 silahkan kalian isi
	// TODO: answer here

	slice[1] = "Marcus"
	slice[2] = "Aurilius"
	fmt.Println(slice)
}