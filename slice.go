// sering menggunakan tipe data slice slice adalah potongan dari data array
// slice mirip seperti array tapi ukuran slice bisa berubah
// slice dan array selalu terkoneksi dimana slice adalah data yg mengakses sebagian / seluruh data di array

package main

import "fmt"

func main(){
	// tipe data slice memiliki 3 data yaitu pointer length dan capacity
	// pointer adalah penunjuk data pertama di array pada slice
	// length adalah panjang dari slice 
	// capacity adalah kapasitas dari slice dimana length tidak boleh lebih dari capacity

	// array[low:high] membuat slice dari arrray dimulai index low sampai index high
	// array[low:] membuat slide dari array dimulai dari index low sampai index akhir di array 
	// array[:high] membuat slice dari array dimulai dari index 0 sampai index sebelum high
	// array[:] membuat slice dari array dimulai index 0 sampai index akhir di array 

	jomoks  := [...]string{"amba","rusdi","fuad","mr ironi","perrel","lilnas"}
	slice1 := jomoks[4:6]

	fmt.Println(slice1)

	slice2 := jomoks[:3]
	fmt.Println(slice2)

	slice3 := jomoks[3:]
	fmt.Println(slice3)
	
	slice4 := jomoks[:]
	fmt.Println(slice4)

	// fucntion slice 
	// len(slice) untuk mendapatkan panjang
	// cap(slice) untuk mendapatkan kapasitas
	// append(slice,data) membuat slice baru dengan menambah data ke posisi terakhir slice jika kapasitas sudah penuh maka akan membuat array baru
	// make([]TypeData,length,capacity) membuat slice baru
	// copy(destination , source) menyalin slice dari source ke destination
	
	
	// append slice 

	days := [...]string {"senin","selasa","rabu","kamis","jumat","sabtu","minggu"}
	daySlice1 := days[5:]
	daySlice1[0] = "sabtu baru"
	daySlice1[1] = "minggu baru"

	fmt.Println(days)

	daySlice2 := append(daySlice1,"libur baru")
	daySlice2[0] = "sabtu lama"
	fmt.Println(daySlice2)
	fmt.Println(days)

	// make slice (membuat slice dari awal)
	// 2 panjang 5 kapasitas kenapa harus ditentukan panjang dan kapasitas nya ketika append dia akan membuat slice baru tapi karena cap 5 masih mempunyai array yg sama agar ketika apppend tidak membuat array baru

	newSlice := make([]string,2,5)
	newSlice[0] = "rizz"
	newSlice[1] = "gyatt"
	// cara menambah data di dalam array tidak bisa langsung [2] = "..." tapi harus menggunakan append karena sudah ditentukan panjang nya 

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	newSlice2 := append(newSlice,"skibidi")
	fmt.Println(newSlice2)
	fmt.Println(len(newSlice2))
	fmt.Println(cap(newSlice2))

	newSlice2[0] = "budi"
	// rizz nya tereplace menjadi budi ini karena masih menggunakan array yg sama kode diatas hanya menimpa value yg diawal
	fmt.Println(newSlice2)
	fmt.Println(newSlice)

	// copy slice

	fromSlice := days[:]
	toSlice := make([]string,len(fromSlice),cap(fromSlice))
	
	copy(toSlice,fromSlice)

	fmt.Println(toSlice)
	fmt.Println(fromSlice)

	// warning ketika membuat array jika salah maka yg kita buat adalah slice / bahkkan kebalikanya
	// contoh

	iniArray :=[...]int{1,2,3,4,5}
	iniSlice := []int{1,2,3,4,5}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)

	// slice jika dideklarasi tidak ada jumlah data nya
}
