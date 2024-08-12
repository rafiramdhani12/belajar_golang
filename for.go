// for loops
//

package main

import "fmt"

func main(){
	counter := 1

	for counter <= 20 {
		fmt.Println("Perulangan ke " , counter)
		counter++
	}

	fmt.Println("selesai")

	// for dengan statement
	//  dalam for bisa menambahkan statement dimana terdapat 2 statement yg bisa ditambahkan di for
	// init statement yaitu statement sebelum for dieksekusi'
	// post statement yaitu statement yg akan selalu dieksekusi diakhir setiap pengulangan

	// A = variable , B = cek kondisi , C = post statement 
	for counter2 := 1 ; counter2 <=10; counter2++{
		fmt.Println("Perulangan ke " , counter2 )
	}

	fmt.Println("selesai")

	// for range 
	// for bisa digunakan untuk melakukan iterasi terhadap semua data collection
	// Data collection contohnya Array,Slice dan Map

	// ini adalah for biasa
	names := []string{"amba" , "fuad" , "rusdi"}
	for i := 0; i< len(names); i++{
		fmt.Println(names[i])
	}

	names2 := []string{"amba" , "fuad" , "rusdi"}
	// i = key , name = value
	for i, name := range names2{
		fmt.Println("index" , i, "=",name)
	}
	// jika tidak butuh key nya hanya value nya saja bisa diganti garis bawah
	for _, name := range names2{
		fmt.Println( "=",name)
	}

}
