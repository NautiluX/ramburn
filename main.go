package main

import (
	"flag"
	"fmt"
	"runtime"
	"time"
)

func main() {

	var mib = flag.Int("s", 8192, "size to allocate in MiB")
	var seconds = flag.Int("t", 10, "seconds until all memory is allocated")
	var pause = flag.Bool("p", false, "pause after memory allocation")
	flag.Parse()
	blockSize := 1024 * 1024
	sleepTime := time.Duration(float64(1000**seconds/(*mib))) * time.Millisecond

	fmt.Printf("Allocating %d MiB in %d seconds. Sleeping %v seconds between blocks.\n", *mib, *seconds, sleepTime)
	memEater := []string{}
	for i := 0; i < *mib; i++ {
		memEater = append(memEater, string(make([]byte, blockSize)))
		time.Sleep(sleepTime)
	}
	runtime.GC()
	fmt.Printf("Allocated size: %v MiB\n", len(memEater))

	for *pause {
		fmt.Println("Going to sleep. Press CTRL+C to quit")
		time.Sleep(60 * time.Second)
	}
}
