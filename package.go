/*
! Package
* package adalah tempat yg bisa untuk mengorganisir kode program yg kita buat di go-lang
* dengan menggunakan package ,  kita bisa merapikan kode program yg kita buat
* package sebenarnya hanya direktori folder dalam operasi kita
*/

package main

import (
	"fmt"
	"s1/helper"
)

func main(){
	result := helper.SayHello("ambasuke")
	fmt.Println(result)
}
