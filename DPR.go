// ! defer , panic , recover

package main

import (
	"fmt"
)

// ! defer
// * defer function adalah function yg bisa kita jadwalkan untuk dieksekusi setelah menambahkan function selesai dieksekusi
// * defer function akan selalu diieksekusi walaupun terjadi erro di function yg dieksekusi

func logging(){
	fmt.Println("Selesai memanggil function")
}

func runApplication(){
	// ! artinya di dalam function ini akan memanggil function  logging diakhir function runApplication
	defer logging()
	fmt.Println("run aplication")
}

// ! panic
// * panic function adalah function yg bisa kita gunakan untuk menghentikan proggram
// * panic function biasanya dipanggil ketika terjadi panic pada saat program kita berjalan
// * saat panic function dipanggil proggram akan terhenti namun defer funtion tetap akan dieksekusi

func endApp(){
	fmt.Println("End App")
}

func runApp(error bool){
	defer endApp()
	if error{
		panic("ERROR")
	}
}

// ! recover
// * recover adalah function yg bisa kita gunakan untuk menangkap data panic
// * dengan recover proses panic akan terhenti sehingga program akan tetap berjalan

// todo , recover yg salah
func runAppFalseRecover(error bool){
	// ! ketika seperti ini kode tidak akan berjalan sampai bawah ketika proggram menemukan kata panic maka akan berhenti disitu 
	defer endApp()
	if error{
		panic("ERROR")
	}
	message := recover()
	fmt.Println("Terjadi error", message)
}
// todo , cara yg benar

func endAppRecover(){
	fmt.Println("End App")
	message := recover()
	fmt.Println("Terjadi error", message)
}

func runAppTrueRecover(error bool){
	// ! recover nya disimpan didalam defer
	defer endAppRecover()
	if error{
		panic("ups error")
	}
}

func main(){
	runApplication()
	// runApp(true)
	// runAppFalseRecover(true)
	runAppTrueRecover(true)
	fmt.Println("Ambasuke")
}

/*
!hi
*/
