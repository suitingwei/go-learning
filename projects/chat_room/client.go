package main

import (
	"bufio"
	"flag"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {

	var nums int

	flag.IntVar(&nums, "n", 1, "the nums of the coconnency clients to connect the server")

	flag.Parse()

	wg := &sync.WaitGroup{}

	osSignal := make(chan os.Signal, 1)
	signal.Notify(osSignal, syscall.SIGINT)

	for i := 1; i <= nums; i++ {
		wg.Add(1)
		go connectServer(wg, osSignal)
	}
	wg.Wait()
}

func connectServer(wg *sync.WaitGroup, osChannel <-chan os.Signal) {
	defer func() {
		wg.Done()
	}()

	conn, err := net.Dial("tcp", ":8080")

	if err != nil {
		log.Fatalf("Failed to connect tcp:8080, %s\n", err.Error())
	}

	reader := bufio.NewReader(conn)

	for {
		log.Println("Client goroutine running...")
		select {
		case userSignal := <-osChannel:
			{
				log.Println("Client goroutine running in signal branch...")
				switch userSignal {
				case syscall.SIGINT:
					{
						err = conn.Close()
						if err != nil {
							log.Fatalf("Failed to close the connection")
						} else {
							log.Printf("Successfully close the connection")
						}
					}
					return
				}
			}
		default:
			{
				log.Printf("Client gorountine running in default branch...")
				msg, err := reader.ReadSring('\n')

				if err != nil {
					log.Fatalf("Failed to read msg from server: %s\n", err.Error())
				}

				log.Printf("Read msg from server: %s\n", msg)
			}
		}
	}
}
