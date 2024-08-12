// function

package main

import "fmt"

func sayHello(){
	fmt.Println("Hello, World!")
}

//* function parameter
//* disaat membuat function terkadang kita membutuhkan data dari luar / paramaeter
//* kita bisa menambahkan parameter di funtion lebih dari 1
//* parameter tidaklah wajib jadi kita bisa membuat function tanpa parameter seperti sebelumnya yg sudah kita buat
//* namun jika kita menambahkan parameter di function maka ketika memanggil function tersebut kita wajib memasukkan data ke parameter nya

func sayHelloTo(firstName string, lastName string){
	// jika sudah menaruh 2 parameter / lebih harus memanggil semuanya tidak boleh salah 1
	fmt.Println("Hello" , firstName , lastName)
}

//* function return value
//* function bisa mengembalikan data
//* untuk memberitahukan bahwa fucntion mengembalikan data , kita harus mrnuliskan tipe data kembalian dari function tsb,
//* jika function tersebut kita deklarasikan dengan tipe data pengembalian maka wajib di dalam function nya kita harus mengembalikan data
//* untuk mengembalikan data dari function kita bisa menggunakan kata kunci return diikuti dengan datanya 

func getHello(name string) string {
	return "Hello " + name
}

//* returning multiple values 
//* function bisa mengembalikan lebih dari 1 data
//* untuk memberitahu jika function mengembalikan multiple value, kita harus menulis semua tipe data return value nya di function

func getFullName()(string,string){
	return "perrel" , "Brown"
}

//* return value wajib ditangkap semua value nya
//* jika kita ingin menghiraukan return value tersebut kita bisa menggunakan tanda _(garis bawah)

// named return value 
func getCompleteName()(firstName,middleName,lastName string){
	firstName = "perrel"
	middleName = "amba"
	lastName = "brown"

	return firstName ,middleName,lastName
}

//* variadic function 
//* parameter yg berada di posisi terakhir memiliki kemampuan dijadikan sebuah varargs(variable argumens)
//* varargs adalah parameter yang bisa diisi dengan banyak nilai
//* apa bedanya dengan parameter biasa dengan tipe data array
//* - jika parameter tipe array kita wajib membuat array terlebih dahulu sebelum mengirimkan ke function 
//* - jika parameter tipe varar*gs kita tidak perlu membuat array terlebih dahulu jika lebih 1 cukup gunakan tanda koma

func sumAll(numbers ...int)int{
	// ... adalah varargs number dibawah dianggap seperti slice  
	total := 0
	// di contoh ini index tidak dipakai makanya kosong
	for _, number := range numbers{
		total += number
	}
	return total
}
//* bagaimana jika kita sudah punya slice nya dan variable kita variadic argumen ? kita bisa mengubah slice menjadi variadic argumen

//* function value 
//* function adalah first class citizen (function adalah tipe data di golang)
//* function juga merupakan tipe data dan bisa disimpan di dalam variable
func getGoodBye(name string) string {
	return "Good Bye " + name
}

//* function as parameter 
//* function tidak hanya bisa kita simpan didalam variable sebagai value
//* namun juga bisa kita gunakan sebagai parameter untuk function lain

//* nama tipe nya string filter tipe nya function artinya filter adalah sebnuah function didalam kode ini jika ingin mengambil filter maka harus dengan parameter string

//* function type declaration 
//* kadang jika function terlalu panjang agak ribet untuk menuliskanya di dalam parameter
//* type declarations juga bisa digunakan untuk membuat alias function sehingga akan mempermudah kita menggunakan function sebagai parameter

type filter func(string) string

func sayHelloWithFilter(name string,filter filter){
	filteredName := filter(name)
	fmt.Println("Hello " , filteredName)
}

// * func sayHelloWithFilter(name string,filter func(string) string){
// * filteredName := filter(name)
// * mt.Println("Hello " , filteredName)
// * }

func spamFilter(name string)string{
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}

// * Anonymous function
// * sebelumnya setiap membuat function kita akan selalu memberikan sebuah nama pada function
// * namun kadang ada kalanya lebih mudah membuat function secara langsung di variable / parameter tanpa harus membuat function terlebih dahulu
// * hal tersebut dinamakan anonymous function / function tanpa nama

type Blacklist func(string) bool

func registerUser(name string,blacklist Blacklist){
	if blacklist(name) {
		fmt.Println("you are blocked", name )
	} else {
		fmt.Println("welcome", name)
	}
}

// * rekursif function
// * function yang memanggil dirinya sendiri
// * kadang dalam pekerjaan kita sering menemui kasus dimana menggunakan rekursif lebih mudah dibandingkan dengan tidak menggunakan
// * contoh kasus yg lebih mudah diselesaikan menggunakan rekursif adalah faktorial

func factorialLoop(value int) int {
	result := 1
	for i := value; i > 0; i--{
		result *= i
	}
	return result
}

// * factorial rekursif

func factorialRecursive (value int) int {
	if value == 1{
		return 1
	} else {
		return value * factorialRecursive(value - 1)
	}
}

// ! lanjut di closure.go dan seterus nya


func main(){
	sayHello()
	sayHelloTo("perrel","brown")
	result := getHello("Amba")
	fmt.Println(result)
	//! di getfullname tadi ada lastname karena tidak butuh salah 1 dari value nya maka diganti dengan _
	firstName, _ := getFullName()
	fmt.Println(firstName)
	a, _, c := getCompleteName()
	fmt.Println(a,c)

	total := sumAll(10,10,10,10,10,10)
	fmt.Println(total)
	numbers := []int{10,10,10,10}
	//! fmt.Println(sumAll(numbers...)) bisa seperti ini / cara yg dibawah
	total = sumAll(numbers...) //kenapa di tambahkan ... agar bisa menerima banyak data
	fmt.Println(total)

	goodBye := getGoodBye //! sengaja tidak diberikan () karena ingin menjadikan function sebagai variable value jika ditambahkan () seperti default nmaka dia akan memanggil fucntion nya
	fmt.Println(goodBye("perrel"))

	// 
	sayHelloWithFilter("perrel",spamFilter)

	filter := spamFilter
	sayHelloWithFilter("Anjing",filter)

	// ! anonymous function

	blacklist := func(name string) bool{
		return name == "anjing"
	}
	registerUser("ambasuke", blacklist)

	registerUser("anjing",func(name string)bool{
		return name == "anjing"
	})

	// ! rekursif function
	fmt.Println(factorialRecursive(10))
	
	fmt.Println(factorialLoop(10))
}
