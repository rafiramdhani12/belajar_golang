package main

import "fmt"

func main(){
	var(
		firstName = "amba"
		lastName = "tukam"
	)

	// di go jika variable tidak digunakan automatis error 

	fmt.Println(firstName,lastName)

	// saat membuat variable constant langsung inisialisasi data
	// const  tidak akan error meski tidak digunakan

	const nama = "amba"

	fmt.Println(nama)

	// di go juga bisa membuat multiple const caranya sama seperti diatas tinggal ganti var nya aja

}
