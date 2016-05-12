package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"runtime"
	"strconv"
	"strings"

	"github.com/gorilla/pat"
)

const (
	PORT         = "1080"
	VERSION_FILE = "version.dat"
)

var version string

func calculateFibonacci(n int) uint64 {
	a := uint64(0)
	b := uint64(1)
	// Iterate until desired position in sequence.
	for i := 0; i < n; i++ {
		// Use temporary variable to swap values.
		temp := a
		a = b
		b = temp + a
	}
	return a
}

func DataHandler(w http.ResponseWriter, r *http.Request) {
	input := r.URL.Query().Get(":num")
	num, err := strconv.Atoi(input)

	docker := os.Getenv("DOCKER_NAME")

	if err == nil {
		result := calculateFibonacci(num)
		fmt.Fprintf(w, "{\"Version\" : \""+version+"\", \"ResponseBy\" : \""+docker+"\", \"Number\" : \""+strconv.Itoa(num)+"\", \"Result\": \""+strconv.FormatUint(result, 10)+"\"}")
	} else {
		fmt.Fprintf(w, "{\"Version\" : \""+version+"\", \"ResponseBy\" : \""+docker+"\", \"Error\": \"Entered number is invalid\"}")
	}
}

func readVersionFile() {
	v, err := ioutil.ReadFile(VERSION_FILE)

	if err != nil {
		version = "0.0"
	} else {
		version = strings.Replace(string(v), "\n", "", -1)
	}
}

func main() {
	// Use all available processing power
	runtime.GOMAXPROCS(runtime.NumCPU() - 1)

	readVersionFile()

	// Serve and handle requests
	r := pat.New()
	r.Get("/fibo/{num}", DataHandler)
	http.Handle("/", r)
	fmt.Println("Serving on localhost:", PORT)
	http.ListenAndServe(":"+PORT, nil)
}
