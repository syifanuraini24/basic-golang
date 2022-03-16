package main

import "fmt"

// Disini kalian coba untuk membaca nilai dari slice yang diberikan.
// Lalu kalian akan menambahkan kata "Olleh" pada slice tersebut.
func main() {
	slice := []string{"Hello", "World"}
	slice = append(slice, "Olleh")
	fmt.Println(len(slice), cap(slice))
	fmt.Println(slice)
}

//read slice
