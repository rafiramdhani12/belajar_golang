package helper

var version = "1.0.0" // ! tidak bisa diakses dari luar
var Application = "golang"

func sayHello(name string) string{ // ! tidak bisa diakses dari luar package
	return "hello " + name
}
func SayHello(name string) string{
	return "hello " + name
}

/* 
! acces modifier
* dibahasa pemroggraman lain,biasanya ada kata kunci yg bisa digunakan untuk menentukan access modifier terhadap suatu function / variable
* di golang untuk menentukan access modifier , cukup dengan nama function / variable
* jika nama nya diawalai dengan huruf besar , maka artinya bisa diakses dari package lain , jika dimulai dengan huruf kecil , artinya tidak bisa diakses dari package lain
*/

/* 
! import
* secara standar , file go-lang hanya bisa mengakses file golang lainya yg berada dalam package yg sama
* jika kita ingin mengakses file golang yg ada diluarb package , maka kita bisa menggunakan import
*/


