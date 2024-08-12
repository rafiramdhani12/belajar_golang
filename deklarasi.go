package main

import "fmt"

func main(){
	// tipe deklarasi adalah kemampuan untuk membuat ulang tipe data baru dari tipe data yg sudah ada
	// deklarasi biasanya untuk membuat alias terhadap tipe data yg sudah ada dengan tujuan agar lebih mudah dimengerti 

	type NoKTP string

	var ktpAmba NoKTP = "69696969"
	
	fmt.Println(ktpAmba)
	fmt.Println(NoKTP("2222222"))

}
