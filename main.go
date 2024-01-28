
package main

import (
    "fmt"
    "log"
    "io"
    "os"
)

func main() {
	var FileToOpen string
	fmt.Println("what file would you like to read")
	fmt.Scan(&FileToOpen)
	fmt.Println("")
    f, err := os.Open(FileToOpen)
    if err != nil {
        log.Fatalf("unable to read file: %v", err)
    }
    defer f.Close()
    buf := make([]byte, 1024)
    for {
        n, err := f.Read(buf)
	if err == io.EOF {
		break
	}
	if err != nil {
		fmt.Println(err)
		continue
	}
	if n > 0 {
		fmt.Println(string(buf[:n]))
	}
    }
}
