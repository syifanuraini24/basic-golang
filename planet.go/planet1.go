// package main

// import "fmt"

// func main() {
// 	for i := 1; i <= 50; i++ {
// 		for j := 1; j <= 50; i++ {
// 			if j%3 == 1 {
// 				continue
// 			}
// 			if j > 50 {
// 				break
// 			}
// 			fmt.Println("jumlah", j)
// 			// if i < 1 {
// 			// 	continue
// 			// }
// 			// if i > 50 {
// 			// 	break
// 			// }
// 			// fmt.Println("Hari ke:", i)
// 		}
// 	}
// }

// // - Jumlah penduduk di Planet Thanos sangat unik
// // - Di hari pertama, hanya ada 1 penduduk
// // - Di hari-hari berikutnya - setiap harinya - Dr Strange muncul dan menggandakan penduduknya menjadi 3x lipat hari sebelumnya
// // - Namun di hari kelipatan 3, Thanos selalu muncul dan menghilangkan 1/2 jumlah penduduk (pembulatan ke bawah)
// // - Di saat Thanos muncul, Dr Strange tidak berani muncul

// // Berapa jumlah penduduk Planet Thanos di hari ke 50?

// // Hint:
// // Hari ke-1: ada 1 penduduk
// // Hari ke-2: ada 3 penduduk = (1 * 3) -> Dr Strange muncul
// // Hari ke-3: ada 1 penduduk = (3 / 2) -> Thanos muncul
// // Hari ke-4: ada 3 penduduk = (1 * 3) -> Dr Strange muncul
// // Hari ke-5: ada 9 penduduk = (3 * 3) -> Dr Strange muncul
// // Hari ke-6: ada 4 penduduk = (9 / 2) -> Thanos muncul
