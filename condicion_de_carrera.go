package main

import "fmt"
func goroutine1(){
  for f:=1;f<=100;f++{
    fmt.Print("1-")
  }
}
func goroutine2(){
  for f:=1;f<=100;f++{
    fmt.Print("2-")
  }
}
func goroutine3(){
  for f:=1;f<=100;f++{
    fmt.Print("3-")
  }
}
func main(){
  go goroutine1()
  go goroutine2()
  go goroutine3()
  var input string
  fmt.Scanln(&input)
  fmt.Println("done")

}
