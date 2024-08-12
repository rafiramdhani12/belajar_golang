/*
shift alt a
! struct
* struct adalah sebuah template data yg digunakan untuk menggabungkan nol / lebih tipe data lainya dalam 1 kesatuan
* struct biasanya representasi data dalam program aplikasi yg kita buat
* data struct  disimpan dalam field
* sederhananya struct adalah kumpulan dari field
*/

package main

import "fmt"

/*

* struct tidak bisa langsung digunakan
* namun kita bisa membuat data/object dari struct yg kita buat

 */

type Customer struct{
	// bisa di gabung seperti ini bisa jg 1 1
	Name , Address string
	Age int
}

/* 
! struct literal
* sebelumnya kita sudah membuat data dari struct namun sebenarnya ada banyak cara yg bisa kita gunakan untuk membuat data dari struct
*/

/* 
! struct method 
* struct bisa digunakan sebagai parameter untuk function
*/

// A adalah parameter B.struct C nama function
func (customer Customer) sayHello(name string){
	fmt.Println("Halo", name ,"saya adalah", customer.Name)
}

func main(){
	var amba Customer
	amba.Name = "ambasuke"
	amba.Address = "ngawi"
	amba.Age = 30

	fmt.Println(amba)
	fmt.Println(amba.Name)

	// ! struct literal

	fuad := Customer{
		Name: "fuad",
		Address: "ngawi",
		Age: 30,
	}
	fmt.Println(fuad)

	rusdi := Customer{"Rusdi","ngawi",30}
	fmt.Println(rusdi)

	jerome := Customer{Name:"jerome"}
	jerome.sayHello("rusdi")
	rusdi.sayHello("jerome")

}
