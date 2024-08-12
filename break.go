// break dan continue
// break dan continue adalah kata kunci yg bisa digunakan dalam perulangan
// break digunakan untuk menghentikan seluruh perulangan
// continue dogunakan untuk menghentikan perulangan yg berjalan dan langsung melanjutkan ke perulangan selanjutnya

package main

import "fmt"

func main(){
	// break
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println("Perulangan ke", i)
	}
	fmt.Println("ini batas break continue")
	// continue
	for c := 0; c < 10; c++{
		if c%2 == 0 {
			continue
		}
		fmt.Println("perulangan ke",c)
		// hanya menampilkan perulangan yg ganjil jika mendapatkan yg genap program ini mengeksekusi dan melanjutkan hanya menampilkan yg ganjil sja 

	}
}
