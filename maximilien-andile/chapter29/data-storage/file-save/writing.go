// data-storage/file-save/writing/main.go
package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	s := [][]string{
		{"age", "genre", "name"},
		{"23", "M", "Hendrick"},
		{"65", "F", "Stephany"},
	}
	// Open the file or create it
	f, err := os.Create("/home/serge/go_work_with_files/myFile.csv")
	defer f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	var buffer bytes.Buffer
	// iterate over the slice
	for _, data := range s {
		buffer.WriteString(fmt.Sprintf("%s,%s,%s\n", data[0], data[1], data[2]))
	}
	n, err := f.Write(buffer.Bytes())
	fmt.Printf("%d bytes written\n", n)
	if err != nil {
		fmt.Println(err)
		return
	}
}
