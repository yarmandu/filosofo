package main

import "fmt"
func mostrar1(){
  for f:=1;f<=1000;f++{
    fmt.Print("1-")
  }
}
func mostrar2(){
  for f:=1;f<=1000;f++{
    fmt.Print("2-")
  }
}
func main(){
  go mostrar1()
  go mostrar2()
  var continua string
  fmt.Scan(&continua)
}
