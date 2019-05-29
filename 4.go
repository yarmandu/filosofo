package main

import "fmt"
import "time"
func g1(){
  for i := 1; i <= 5; i++ {
    time.Sleep(1000*time.Millisecond)
    fmt.Printf("%d",i)
  }
}
func g2(){
  for i := 'a'; i <='e'; i++ {
    time.Sleep(2000*time.Millisecond)
    fmt.Printf("%c",i)
  }
}
func main(){
  go g1()
  go g2()
  time.Sleep(12000*time.Millisecond)
}
