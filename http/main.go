package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

//As we are implementing the below method, which is for Writer, our custom struct can be passed as a Writer.
func (logWriter) Write(b []byte) (int, error) {
	fmt.Println(string(b))
	fmt.Println("Completed writing byes: ", len(b))
	return len(b), nil
}

func main() {
	resp, err := http.Get("http://google.com/")

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	log := logWriter{}

	// bs := make([]byte, 12000)
	// resp.Body.Read(bs)
	// fmt.Println(string(bs))

	// io.Copy(os.Stdout, resp.Body)
	//Here, instead of 'os.Stdout', we are using our custom struct 'log', which has custom implementation for 'Write'
	io.Copy(log, resp.Body)
}
