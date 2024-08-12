package main

import "fmt"

func main(){
	// konversi tipe data 1
	var nilai32 int32 = 32768
	var nilai64 int64 = int64(nilai32)

	var nilai16 int16 = int16(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai16)

	// hasilnya sama seperti awal hanya di int16 jadi -32768 karena melebihi kapasitas karena number overflow

}
