// if expression
// if adalah salah satu kata kunci yg digunakan untuk percabangan
// Percabangan artinya kita bisa mengeksekusi kode program tertentu ketika suatu kondisi terpenuhi
// Hampir disemua bahasa pemrograman mendukung if expression

package main

import "fmt"

func main(){
	name := "Amba"

	if name =="amba"{
		fmt.Println("Hello amba")
	} else if name == "Amba"{
		fmt.Println("tetep amba jir cuman beda kapital")
	} else {
		fmt.Println(false)
	}

	// if dengan short statement 
	// if mendukung short statement sebelum kondisi
	// hal ini sangat cocok untuk membuat statement yg sederhana sebelum melakukan pengecekan terhadap kondisi

	// short statement nya adalah di baris pertama karena menyyingkat banyak hal 
	if length := len(name); length > 5 {
		fmt.Println("nama terlalu panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}


}
