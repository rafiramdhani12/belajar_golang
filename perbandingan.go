package main

import "fmt"

func main(){
	// operator yg menghasilkan nilai boolean true / false 
	// > lebih dari < kurang dari >= lebih dari sama dengan <= kurang dari samadengan == sama dengan != tidak sama dengan

	var name1 = "eko"
	var name2 = "amba"
	var name3 = "amba"

	var result bool = name1 == name2
	var result2  = name2 == name3

	fmt.Println(result)
	fmt.Println(result2)

	// && dan || atau ! kebalikan (tidak)

	var nilaiAkhir = 90
	var absensi = 81

	var lulusNilaiAkhir = nilaiAkhir > 80
	var lulusAbsensi = absensi > 80

	var lulus = lulusNilaiAkhir && lulusAbsensi

	fmt.Println(lulus)

}
