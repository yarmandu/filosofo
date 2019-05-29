package main

import "fmt"
//import "sync"
import "sort"
func main() {

//var wg sync.WaitGroup
size :=8
ar :=[]int{4,2,1,3,10,9,7,5}
//for i:=0;i<size;i++{
  //fmt.Scanln(&ar[i])
//}
fmt.Println(ar)
for j:=0; j<size;j=j+2{
  sort.Ints(ar[j:j+2])
}
fmt.Println(ar)
sort.Ints(ar)
fmt.Println(ar)
//wg.Wait()


  //numeros := [...]int {18, 8, 21, 96, 97}

	//for i := len(numeros) - 1; i >= 0; i-- {
		//fmt.Println(numeros[i])
	//}
}
