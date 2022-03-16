package main

import "fmt"

func main() {
	result := 0
	for i := 1; i <= 50; i++ {
		if i*3 == i { // skip odd numbers
			continue
		}
	}
	fmt.Println(result) // 6 (2+4)
}

// Hari ke-1: ada 1 penduduk
// Hari ke-2: ada 3 penduduk = (1 * 3) -> Dr Strange muncul
// Hari ke-3: ada 1 penduduk = (3 / 2) -> Thanos muncul
// Hari ke-4: ada 3 penduduk = (1 * 3) -> Dr Strange muncul
// Hari ke-5: ada 9 penduduk = (3 * 3) -> Dr Strange muncul
// Hari ke-6: ada 4 penduduk = (9 / 2) -> Thanos muncul
