package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.com",
		"http://amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	for l := range c {
		//go checkLink(l, c)

		go func(l string) {
			time.Sleep(time.Millisecond * 5000)
			checkLink(l, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link + " might be down!")

		c <- link

		return
	}

	fmt.Println(link + " is up!")

	//time.Sleep(time.Millisecond * 5000)

	c <- link
}
