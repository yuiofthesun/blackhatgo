package main

import (
  "fmt"
  "net"
  "sort"
)

func worker(ports, results chan int) {
  for port:= range ports {
    address := fmt.Sprintf("scanme.nmap.org:%d", port)
    conn, err := net.Dial("tcp", address)
    if err != nil {
     results <- 0
     continue
    }
    conn.Close()
    results <- port
  }
}

func main(){
  ports := make(chan int, 100)
  results := make(chan int)
  var openports []int
  for port := 0; port < cap(ports); port++ {
    go worker(ports, results)
  }
  go func() {
    for port := 0; port <= 65535; port ++ {
      ports <- port
    }
  }()

  for port := 0; port < 65535; port++ {
    port := <-results
    if port !=0 {
      openports = append(openports, port)
    }
  }
  close(ports)
  close(results)
  sort.Ints(openports)
  for _, port := range openports {
    fmt.Printf("Connection successful at %d\n", port)
  }
}
