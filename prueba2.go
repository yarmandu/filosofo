package main
import "fmt"
import "sync"
func foo(wg*sync.WaitGroup){
  fmt.Printf("New routine")
  wg.Done()
}
func main(){
  var wg sync.WaitGroup
  wg.Add(1)
  go foo(&wg)
  wg.Wait()
  fmt.Printf("Main routine")
}
