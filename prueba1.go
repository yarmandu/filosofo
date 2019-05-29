package main
import "time"
import "fmt"
func main(){
  go fmt.Printf("New routine")
  time.Sleep(100*time.Millisecond)
  fmt.Printf("Main routine")
}
