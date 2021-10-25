package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

func main() {
	file := os.Args[1]
	buf := make([]byte, 1024)
	offset := 0

	f, err := os.Open(file)
	defer f.Close()
	if err != nil {
		fmt.Printf("open file %s failed: %v\n", file, err)
		return
	}

	for {
		n, err := f.Read(buf)
		if err != nil && err != io.EOF {
			fmt.Printf("read file %s failed: %v\n", file, err)
			return
		}
		if n > 1 {
			if n != len(buf) {
				n--
			}
			fmt.Printf("%s", string(buf[:n]))
			offset += n
		}

		if n != len(buf) {
			time.Sleep(100 * time.Millisecond)
			f.Close()
			f, err = os.Open(file)
			if err != nil {
				fmt.Printf("open file %s failed: %v\n", file, err)
				return
			}
			f.Seek(int64(offset), 0)
		}
	}

}
