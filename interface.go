/*
! interface
* interface adalah tipe data abstract dia tidak memiliki implementasi langsung
* sebuah interface berisikan defenisi defenisi method
* biasanya interface digunakan sebagai kontrak
*/

// todo, function yg menempel bisa dibilang method contoh getname bisa dibilang method

package main

import "fmt"

type Animal struct {
	Name string
}

type HasName interface{
	GetName() string
}

type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

func (animal Animal) GetName() string {
	return animal.Name
}

func SayHello(hasname HasName){
	fmt.Println("Hello", hasname.GetName())
}

/* 
! implementasi interface
* setiap tipe data yg sesuai dengan kontrak interface secara automatis dianggap sebagai interface tersebut
* sehingga kita tidak perlu mengimplementasikan interface secara manual
* hal ini agak berbeda dengan bahasa pemroggraman lain yg ketika memmbuat interface kita harus menyebutkan secara eksplisit akan menggunakan interface mana
*/

/* 
! interface kosong
* Golang bukanlah sebuah bahasa yg berorientasi objek
* biasanya dalam pemrograman berorinteasi objek ada satu data parent di puncak yg bisa dianggap sebagai semua implementasio data yg ada di bahasa pemrograman tsb
* contoh di java ada java.lang.Object
* untuk menangani kasus seperti ini , golang kita bisa menggunakan interface kosong
! interface kosong adalah interface yg tidak memiliki deklarasi method satupun hal ini membuiat secara automatis semua tipe data akan menjadi implementasi nya
* interface kosong juga memiliki typef alias any
*/

func Ups() any{
	// ! bisa juga diisi oleh any / interface{}
	// return 1
	// return true
	return "Ups"
}

/* 
! contoh interface kosong di go
* fmt.Prinln(a...interface{})
* panic(v.interface{})
* recover() interface{}
* dll
! interface kosong artinya bisa dikirim data apa pun karena smeua tipe data di golang mengikuti kontrak data interface yg kosong
*/

func main(){
	person := Person{Name: "fuad"}
	SayHello(person)

	animal := Animal{Name: "kucing"}
	SayHello(animal)

	kosong := Ups()
	fmt.Println(kosong)
}
