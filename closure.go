package main

import "fmt"

// * closure
// * closure adalah kemampuan sebuah funtion berinteraksi dengan data data disekitarnya dalam scope yg sama
// * harap gunakan fitur closure ini dengan bijak saat kita membuat aplikasi karena terlalu banyak dapat membunuh proggram kita

func main(){
	counter := 0




// ! anggap saja ini kode yg sudah jadi dan panjang disuatu program







	increment := func(){
		fmt.Println("Increment")
		counter ++
	}








  //! dalam kasus ini karena sudah terlalu panjang kode nya lalu kita memanggil suatu function tetapi ada nilai / value yg berubah ketika kita tidak pernah merasa merubah sesuatu value tapi value itu berubah karena kita memanggil function maka itulah kemampuan closure pada function bisa mengakses data disekitarnya

	increment()
	increment()
	increment()

	fmt.Println(counter)
}
