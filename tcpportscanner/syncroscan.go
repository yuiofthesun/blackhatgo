package main

import (
  "fmt"
  "net"
  "sync"
)

func main(){
  var waitgroup sync.WaitGroup
  for i := 0; i <= 65535; i++ {
   waitgroup.Add(1)
   go func(port int) {
     defer waitgroup.Done()
     //address := fmt.Sprintf("scanme.nmap.org:%d", port)
     address := fmt.Sprintf("127.0.0.1:%d", port)
     conn, err:= net.Dial("tcp", address)
     if err != nil{
       return
     }
     conn.Close()
     fmt.Println("Connection successful at ", port)
   }(i)
  }
  waitgroup.Wait()
}
