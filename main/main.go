package main

import (
	"fmt"
	"github.com/giskook/dtu/conf"
	"github.com/giskook/dtu/reactor"
	"log"
	"os"
	"os/signal"
	"runtime"
	"syscall"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	conf, err := conf.ReadConfig("./conf.json")
	checkError(err)
	rc := reactor.NewReactor(conf)
	err = rc.Start()
	checkError(err)
	// catchs system signal
	chSig := make(chan os.Signal)
	signal.Notify(chSig, syscall.SIGINT, syscall.SIGTERM)
	fmt.Println("Signal: ", <-chSig)
	rc.Stop()
}
func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
