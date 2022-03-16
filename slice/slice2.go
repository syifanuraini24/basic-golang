package main

import "fmt"

// Disini kalian coba untuk membaca nilai dari slice yang diberikan.
// Lalu kalian akan menambahkan kata "Olleh" pada slice tersebut.
func main() {
	var slice = []string{"Hello", "World"}

	slice = append(slice,"Olleh")

	fmt.Println(len(slice),cap(slice))
	fmt.Println(slice)
	// TODO: answer here
}