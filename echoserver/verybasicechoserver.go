package main

import (
  "io"
  "log"
  "net"
)

func echo(conn net.Conn) {
  defer conn.Close()

  //Create a buffer to store received data.
  buffer := make([]byte, 512)
  for {
    //receive data via conn.Read into a buffer

    size, err := conn.Read(buffer[0:])
    if err == io.EOF {
      log.Println("Client disconnected")
      break
    }
    if err != nil {
      log.Println("Unexpected error")
      break
    }
    log.Printf("Received %d bytes: %s\n", size, string(buffer))

    //send data via conn.Write
    log.Println("Writing data")
    if _,err := conn.Write(buffer[0:size]); err != nil {
      log.Fatalln("Unable to write data")
    }
  }
}

func main() {
  //Bind to TCP port 20080 on all interfaces
  listener, err := net.Listen("tcp", ":20080")
  if err != nil {
    log.Fatalln("Unable to bind to port")
  }
  log.Println("Listening on 0.0.0.0:20080")
  for {
    //Wait for connection Create net.Conn on connection established
    conn, err := listener.Accept()
    log.Println("Received connection")
    if err != nil {
      log.Fatalln("Unable to accept connection")
    }
    //Handle the connection. Using goroutine for concurrency
    go echo(conn)
  }
}
