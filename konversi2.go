package main

import "fmt"

func main(){
	// konversi tipe data 2
	var name = "Amba"
	var e = name[0]
	var eString = string(e)

	// jika hanya print var e hanya akan keluar byte 

	fmt.Println(name)
	fmt.Println(eString)
	fmt.Println(e)

}
