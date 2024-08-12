/*
! type assertions
* merupakan kemampuan merubah tipe data menjadi tipe data yg diinginkan
* fitur ini sering kali digunakan ketika bertemu dengan data interface kosong
*/
package main

import "fmt"

func random() interface{}{
	return true
}

/* 
! type assertion menggunakan switch
* saat salah menggunakan type assertions, maka bisa berakibat terjadi panic di aplikasi kita
* jika panic dan tidak terecover maka automatis proggram kita akan mati
* agar lebih aman sebaiknya kita menggunakan switch untuk melakukan type assertion
*/

 func main(){
	result := random()

	/* 
	resultString := result.(string)
	fmt.Println(resultString)
	 jika diatas string lalu dikonversi menjadi int akan panic jadi jgn asal konversi 
	resultInt := result.(int)  panic
	fmt.Println(resultInt)
	*/
	
	switch value := result.(type){
	case string:
		fmt.Println("String",value)
	case int:
		fmt.Println("Int",value)
	default:
		fmt.Println("Unknown",value)
	}
 }
 