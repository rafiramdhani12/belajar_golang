/*
! pointer di function
* saat kita membuat parameter di function secara default adalah pass by value , artinya data akan di copy lalu dikirim ke function tsb
* oleh karena itu jika kita mengubah data di dalam function, data yg aslinya tidak akan pernah berubah
* hal ini membuat variable menjadi aman , karena tidak akan bisa diubah
* namun kadang kita ingin membuat function yg bisa mengubah data asli parameter tsb
* untuk melakukan ini , kita juga bisa menggunakan pointer di function
* untuk menjadikan parameter sebagai pointer , kita bisa menggunakan operator * di parameter nya
*/

package main

import "fmt"

type Address struct {
	City,Province,Country string
}

func ChangeCountryToIndonesia(address *Address){
	address.Country = "indonesia"
}

func main(){
	address := Address{"subang" , "jawa barat" , "singapur"}
	ChangeCountryToIndonesia(&address)

	fmt.Println(address)
}
