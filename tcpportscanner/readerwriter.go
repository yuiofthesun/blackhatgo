package main

import (
  "fmt"
  "log"
  "os"
  "io"
)

//FooReader defines an io.Reader to tead from stdin

type FooReader struct{}

//Read reads data from stdin

func (fooReader *FooReader) Read(b []byte) (int, error) {
  fmt.Print("in >  ")
  return os.Stdin.Read(b)
}

//FooWriter defines an io.Writer to write Stdout

type FooWriter struct{}

//Write writes data to Stdout

func (fooWriter *FooWriter) Write(b []byte) (int, error) {
  fmt.Print("out> ")
  return os. Stdout.Write(b)
}

//func Copy(dst io.Writer, src io.Reader) (written int64, error)

func main() {
  //Instantiate reader and writer
  var (
    reader FooReader
    writer FooWriter
  )

  if _, err := io.Copy(&writer, &reader); err != nil {
    log.Fatalln("Unable to read/write data")
  }
}
