package database

/*
! package initialization
* saat kita membuat package , kita bisa membuat sebuah function yg akan diakses ketika package kita diakses
* ini sangat cocok ketika contoh nya , jika package kita berisi function-function untuk berkomunikasi dengan database , kita membuat function inisialisasi untuk membuka koleksi database
* untuk membuat function yg diakses secara automatis ketika package diakses , kita cukup membuat function dengan nama init
*/

var connection string 

func init(){
	connection = "MySQL"
}

func GetDatabase()string {
	return connection
}

/* 
! blank identifier
* kadang kita hanya ingin menjalankjan init function di package tanpa harus mengeksekusi salah 1 function yg ada di package
* secara default go-lang kan komplen ketika ada package yg di import namun tidak digunakan 
* untuk menangani hal tsb kita bisa menggunakan blank identfier (_) sebelum nama package ketika melakukan import
*/
