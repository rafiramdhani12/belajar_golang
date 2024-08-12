/*
! pointer di method
* walaupun method akan menempel di struct , tapi sebenarnya data struct yg diakases didalam method adalah pass by value
* sangat direkomendasikan menggunakan pointer di method , sehingga tidak boros memory karena harus selalu diduplikasi ketika memanggil method
*/

package main

import "fmt"

type Man struct{
	Name string
}

func (man *Man) Married(){
	man.Name = "Mr. " + man.Name
}

func main(){
	amba := Man{"ambasuke"}
	amba.Married()

	fmt.Println(amba.Name)
}
