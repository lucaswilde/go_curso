package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.com",
		"http://amazon.com",
		"http://google.com",
	}

	// channel
	c := make(chan string)
	for _, link := range links {
		go checkLink(link, c)
	}

	// espera até receber uma mensagem do channel, blocking call
	// for i := 0; i < len(links); i++ {
	// 	fmt.Println(<-c)
	// }

	// infinite loop
	// for {
	// 	// assim que um link for verificado uma nova rotina é dispara para verifica-lo novamente
	// 	go checkLink(<-c, c)
	// }
	// outra forma mais clara de implementar a mesma coisa
	for l := range c {
		// assim que um link for verificado uma nova rotina é dispara para verifica-lo novamente
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}

}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down")
		// envia mensagem pelo channel
		c <- link
		return
	}
	fmt.Println(link, "is up!")
	// envia mensagem pelo channel
	c <- link
}
