package main

import (
	"flag"
	"fmt"
	"github.com/thorduri/pushover"
	"log"
	"os"
	"runtime"
	"strconv"
	"sync"
	"syscall"
	"time"
)

func checkPid(pid string, wg *sync.WaitGroup) {
	defer wg.Done()

	pidn, err := strconv.Atoi(pid)
	if err != nil {
		return
	}

	for {
		time.Sleep(1 * time.Second)
		err := syscall.Kill(pidn, syscall.Signal(0))
		if err != nil {
			return
		}
	}
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	config := NewConfig()

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage:\n  ppush [pids]\n\n")
		flag.PrintDefaults()
	}
	flag.StringVar(&config.App, "a", config.App, "app token")
	flag.StringVar(&config.User, "u", config.User, "user token")
	flag.StringVar(&config.Message, "m", config.Message, "push message")
	flag.Parse()

	po, err := pushover.NewPushover(config.App, config.User)
	if err != nil {
		log.Fatal(err)
	}

	wg := new(sync.WaitGroup)

	pids := flag.Args()

	for _, pid := range pids {
		wg.Add(1)
		go checkPid(pid, wg)
	}

	wg.Wait()

	err = po.Message(config.Message)
	if err != nil {
		log.Fatal(err)
	}
}
