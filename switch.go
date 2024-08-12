// switch expression
// selain if untuk melakukan percabangan kita juga bisa menggunakan switch
// switch sangat sederhana dibandingkan switch
// biasanya switch digunakan untuk melakukan pengecekan ke kondisi dalam satu variable

package main

import "fmt"

func main(){
	nama := "ambasuke"

	switch nama{
	case nama:
		fmt.Println("hello ambasuke")
	case "joko":
		fmt.Println("hello joko")
	default:
		fmt.Println("Hi boleh kenalan ?")
	}

	// switch dengan short statement
	switch length := len(nama); length > 10 {
	case true:
		fmt.Println("Nama terlalu panjang")
	case false:
		fmt.Println("nama sudah benar")
	}
	// switch tanpa kondisi
	// kondisi di switch tidak wajib
	// jika kita tidak menggunakan kondisi di switch expression kita bisa menambahkan kondisi tersebut disetiap case nya 

	name := "ambasigma"
	length := len(name)
	switch{
	case length > 10:
		fmt.Println("Nama terlalu panjang")
	case length > 5:
		fmt.Println("nama lumayan panjang")
	default:
		fmt.Println("nama sudah benar")
	}
	// tidak disarankan jika seperti ini mending memakai if 
}
