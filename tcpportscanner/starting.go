package main

import (
  "fmt"
  "net"
)

func main() {
  var foo string = "hello, world";
  fmt.Println(foo);

  _, err := net.Dial("tcp", "scanme.nmap.org:80")
  if err == nil{
    fmt.Println("Connection established")
  }
}

