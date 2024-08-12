/*
! nil
* biasanya di dalam bahasa pemrograman lain objek yg belum diinisialisasi maka secara automatis nilainya dalah null / nil
* berbeda dengan go lang di go lang saaat kita membuat variable dengan tipe data tertentu maka secara automatis akan dibuatkan default value nya
* namun di golang ada data nil yaitu data kosong
* nil sendiri hanya bisa digunakan di beberapa tipe data , seperti interface , funtion , map , slice , pointer , channel
*/

package main

import "fmt"

func newMap(name string) map[string]string{
	if name == ""{
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main(){
	data := newMap("")
	if data == nil{
		fmt.Println("data kosong")
	} else {
		fmt.Println(data["name"])
	}
}
