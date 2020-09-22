package main

import "fmt"

func main() {
	// buffered channel
	ca := make(chan int, 1)

	ca <- 42
	// sending more data than the buffer can support
	ca <- 84

	fmt.Println(<-ca)
	/* this is going to show an error. To solve it simply make a bigger buffer
	with the ammount of data the chan it's going to receive
	i.e:

	ca := make(chan int, 2)

	fmt.Printl(<-ca) would print 84 */
}
