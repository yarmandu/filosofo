package main

import (
"fmt"
"time"
)
var l int = 0
func G1(array [2] int) {//l
  time.Sleep(1000*time.Millisecond)
  var temp int
    for i := 0; i < 2 ; i++ {
                 for j := i + 1; j < 2 ; j++ {
                      if array[i] > array[j] {
                                  temp = array[i]
                                  array[i] = array[j]
                                  array[j] = temp
                  	 }
                  }
         }

  fmt.Println( array)

}
func G2(array [3]int) {
  time.Sleep(1000*time.Millisecond)
  var temp int
    for i := 0; i < 3 ; i++ {
                 for j := i + 1; j < 3 ; j++ {
                      if array[i] > array[j] {
                                  temp = array[i]
                                  array[i] = array[j]
                                  array[j] = temp
                  	 }
                  }
         }

  fmt.Println( array)
}
func G3(array [9]int) {
time.Sleep(1000*time.Millisecond)
  var temp int
    for i := 0; i < 9 ; i++ {
                 for j := i + 1; j < 9 ; j++ {
                      if array[i] > array[j] {
                                  temp = array[i]
                                  array[i] = array[j]
                                  array[j] = temp
                  	 }
                  }
         }

  fmt.Println(array)

}
func main() {
	var n int =3
	matrizA := [9] int{3, 1, 2, 5, 4, 6, 7, 9, 8}
	var m int = n*n/4
	var k int
	if 4*m != n*n {
		k = (n*n) - (3*m)
	}
	var m1 [2] int //m
	var m2 [2] int //m
	var m3 [2] int //m
	var m4 [3] int //k

	var i int = 0
	   for j:=0;i<m;j++{
	        m1[j] = matrizA[i]
	        i++
	   }
	   for j:=0;i<2*m;j++{
	        m2[j] = matrizA[i]
	        i++
	   }
	  for j:=0;i<3*m;j++{
	        m3[j] = matrizA[i]
	        i++
	   }
	  for j:=0;i<9;j++{
	        m4[j] = matrizA[i]
	        i++
	   }
     fmt.Println(matrizA)
	l=m
	go G1(m1)
	go G1(m2)
	go G1(m3)
	l=k
	go G2(m4)
	l=n*n
	go G3(matrizA)

	time.Sleep(5000*time.Millisecond)

}
