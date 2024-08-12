/*
! Operator *
* saat kita mengubah variable pointer , maka yg berubah hanya variable tsb
* semua variable yg mengacu ke data yg sama tidak akan berubah
* jika kita ingin mengubah sluruh variable yg mengacu ke data tsb kita bisa menggunakan operator *
*/

package main

import "fmt"

type Address struct{
	City,Province,Country string
}

func main(){
	address1 := Address{"subang" , "jawa barat" , "indonesia"}
	address2 := &address1 //pointer

	address2.City = "bandung"

	*address2 = Address{"jakarta" , "dki jakarta" , "indonesia"}

	fmt.Println(address1) // address 1 berubah
	fmt.Println(address2)
}
