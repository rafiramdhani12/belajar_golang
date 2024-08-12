package main

import "fmt"

func main(){
	// array
	// len(array) untuk mendapatkan panjang array , array[index] mendapatkan data posisi index , array[index]=value mengubah data di posisi index
	// yg dibawh ini tidak bisa [...] harus langsung ditentukan
	var names[3]string
	names[0] = "John"
	names[1] = "Mary"
	names[2] = "Jane"
	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	// di go juga bisa membuat array secara langsung saat deklarasi variable

	// jika tidak ingin menentukan secara pasti maka bisa menggunakan [...] yg berawal [3]

	var values=[...]int{
		11,25,37,
	}

	// jika kita tulis 3 tapi yg ke 3 tidak diisi tidak apa apa karena ada nilai default 0 / empty string

	fmt.Println(values)
	fmt.Println(len(values))
	values[0] = 100
	
	fmt.Println(values)
}
