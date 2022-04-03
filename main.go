package main

import (
	"fmt"
	"net/http"
	"time"
)

// Program: Status checker of different websites
// Concept: Go Rountines & channels: for concurrency.

func main() {

	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string) // channel

	for _, link := range links {

		// checkLink(link) // How can we make it concurrent?

		// $ go run main.go
		// http://google.com is up!
		// http://facebook.com is up!
		// http://stackoverflow.com is up!
		// http://golang.org is up!
		// http://amazon.com is up!

		// go checkLink(link) // Nothing happened with just adding "go" keyword. Need more.
		// Main rountine didnt wait for child routines to finish.
		// Main routines says there's nothing else for me to be down, so just exit.

		go checkLink(link, c)
	}

	// fmt.Println(<-c) // we received just one communication from channel. WHY?
	// the above call is a blocking call.
	// Therefore, when one communication is recevied, it exited after that.

	// for i := 0; i < len(links); i++ {
	// 	fmt.Println(<-c)
	// }

	// for { // Infinite loop
	// 	go checkLink(<-c, c)
	// }

	for l := range c { // This is equivalent to above for loop.
		// time.Sleep(2 * time.Second) // This is a blocking call for main routine. This means we can execute one routine every 2 sec.

		// go checkLink(l, c)

		// ------------------------------------------

		// go func() {
		// 	time.Sleep(2 * time.Second)
		// 	checkLink(l, c) // l is defined in the outer scope which changes constantly as we receive messages in channel.
		// 	// Therefore, terminal is constantly printing fb.com
		// 	// As by the name checklink starts executing, l's value is changed.
		// }()

		// Never access the variable in child routine from man routine.
		// Never ever share the variables b/w child & main.
		// ------------------------------------------

		// function literal; similar to lambda in Python; anonymous function
		go func(link string) {
			time.Sleep(2 * time.Second)
			checkLink(link, c)
		}(l) // we are passing the "l" as an arguemnt to this function literal.

	}
}

func checkLink(link string, c chan string) {

	// time.Sleep(2 * time.Second) // this would work, but won't give the answer right now.
	// B/w the main routine and this checklink, we need to put the sleep statement somewhere.

	_, err := http.Get(link) // This is a blocking call in case of serial execution.
	if err != nil {
		fmt.Println(link, "might be down!")
		// c <- "Might be down, I think!"
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	// c <- "Yep it's up"
	c <- link

}

// When a program in run, it runs inside the Go rountine.
// Go Routine: engine that executes the code. It executes the program line by line.

// There is a Go Scheduler: that runs on one CPU core. Even if you have dual core, by default,
// Go is going to attempt to run on one CPU.
// Even though we are launcing multiple Go routines, only one will run at a particular time.
// Scheduler runs one routine until it finished or makes a blocking call (like an HTTP request).
// Scheduler monitors each Go routine (when will it finish, when to start next etc.)

//					One CPU core
//					Go Scheduler
//		Go Rou1		Go Rout2	Go Rout2

// What if you havemultiple CPU cores? * By Default Go only uses one CPU core.
// We can change this behaviour. By overriding this value, each CPU core can run
// one go routine at a time.

// ------------------------------------------
// Concurrency is not Parallelism.
// concurrency: We can have multiple threads executing code. If one thread blocks, another
// one is picked up and worked on.

// Parallelism - Multiple threads executed at the exact same time. Requries multiple CPUs.

// ------------------------------------------
// CHANNELS
// are the only way to communicate b/w Go routines.
// It's like instant messaging.
// Channels are typed. Not itself, but the information which we passes through channel.
// e.g. Channel of type string, int etc.

// channel <- 5: Send value '5' into this channel.
// myNumber <- channel: Wait for a value to be sent into the channel. When we get one, assign the value to 'myNumber'.
// fmt.Println(<- channel): when we get the value, log it out immediately.
