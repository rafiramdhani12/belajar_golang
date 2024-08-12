package main

import "fmt"

func main(){
	var name string

	// saat membuat variable wajib menentukan tipe variable nya , bisa juga tanpa tipe  
	// kata kunci var tidak wajib jika langsung menginisialisasi variable dengan :=

	umur:= 25
	fmt.Println("nama saya adalah amba saya berumur ", umur)

	// untuk yg ke 2 kalinya jika menggunakan := hanya = saja tidak perlu mengunakan : lagi karena dianggap membuat baru

	name = "ambatukam"
	fmt.Println(name)

	name = "ambasing"
	fmt.Println(name)
}
