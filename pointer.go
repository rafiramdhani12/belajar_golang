/*
! pass by value
* secara dafault di go lang semua variable itu di passing by value bukan bye refence
* artinya jika kita mengirim sebuah variable kedalam function , method atau variable lain, sebenarnya yg dikirim adalah duplikasi value nya
*/

package main

import "fmt"

type Address struct{
	City,Province,Country string
}

/* 
! pointer
* pointer adalah kemampuan membuat reference ke lokasi data di memory yg sama , tanpa menduplikasi data yg sudah ada
* sederhananya , dengan kemampuan pointer , kita bisa membuat pass by reference
? bagaimana cara membuat nya kita harus tau operator & 
* untuk membuat sebuah variable dengan nilai pointer ke variable yg lain , kita bisa menggunakan operator & diikuti dengan nama variable
*/

func main(){
	// address1 := Address{"subang","jawa barat","indonesia"}
	// address2 := address1 //! copy value

	// address2.City = "Bandung"

	// fmt.Println(address1) // ! tidak berubah
	// fmt.Println(address2) // ! berubah menjadi bandung

	address1 := Address{"subang","jawa barat","indonesia"}
	// ! address 2 saat ini isinya bukan value tapi reference ke addres 1 / pointer artinya jika kita melakukan perubahan di addres2 sebenarnya dia merubah addres 1
	address2 := &address1 //! pointer

	address2.City = "Bandung"

	fmt.Println(address1) // ! berubah

}
