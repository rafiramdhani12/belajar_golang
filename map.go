// tipe data map
// pada array / slice untuk mengakses data kita menggunakan index number dimulai dari 0
// map adalah tipe data lain yg berisikan kumpulan data yg sama namun kita bisa menentukan jenis tipe data index yg akan kita gunakan
// sederhananya map adalah tipe data kumpulan key-value (kata-kunci-nilai) dimana kata kuncinya bersifat unik tidak boleh sama
// berbeda dengana array dan slice jumlah data yg kita masukkan kedalam map boleh sebanyak banyak nya asalkan kata kuncinya berbeda jika kita gunakan kata kunci sama maka secara automatis data sebelumnya akan diganti dengan data baru

package main

import "fmt"

func main(){
	person := map[string]string{
		"name" : "amba",
		"address" : "ngawi",
	}

	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person)

	// function map 
	// len(map) untuk mendapatkan jumlah data di map
	// map[key] untuk mengambil data di map dengan key
	// map[key] = value * mengubah data di map dengan key
	// make (map[TypeKey]TypeValue) membuat map baru
	// delete(map,key) untuk menghapus data di map dengan key

	book := make(map[string]string)
	book["title"] = "Buku Go-lang"
	book["author"] = "Amba"
	book["wrong"] = "Ups"
	
	fmt.Println(book)

	delete(book,"wrong")

	fmt.Println(book)
}
