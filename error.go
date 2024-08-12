/*
! error interface
* golang memiliki interface yg digunakan sebagai kontrak untuk membuat error , nama interface nya adalah error
*/

/*
! membuat error
* untuk membuat error , kita tidak perlu manual
* golang sudah menyediakan library untuk membuat helper secara mudah , yg terdapat di package errors
*/

package main

import (
	"errors"
	"fmt"
)

func Pembagian(nilai int , pembagi int) (int,error){
	if pembagi == 0 {
		return 0 , errors.New("Pembagian Dengan Nol")
	} else {
		return nilai / pembagi , nil
	}
}

func main(){
	hasil, err := Pembagian(100,100)
	if err == nil {
		fmt.Println("Hasil" , hasil)
	} else {
		fmt.Println("Error", err.Error())
	}
}
