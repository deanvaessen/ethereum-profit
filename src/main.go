package main

import (
	"fmt"
	"helpers/user"
	"helpers/math"
	"helpers/communicate"
	//"time"
)

/*func f (n int) {
	for i := 0; i <10; i++ {
		fmt.Println(n, ":", i)
	}
}

func pinger(c chan string) {
	for i := 0; ; i++ {
		c <- "ping"
	}
}

func ponger(c chan string) {
	for i := 0; ; i++ {
		c <- "pong"
	}
}

func printer(c chan string) {
	for {
		msg := <- c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}*/


func main() {

	// Define a username so we can store your data
	var user string
	fmt.Println("Please give me your name so I can see if I need to set you up or can load your details.")
	fmt.Scanlnt(&user)

		// Check if we already have this user
		userKnown := user.Exists()

		if userKnown {
			allTransactions := user.LoadAllTransactions()
		} else {
			user.Init()
			allTransactions := user.LoadAllTransactions()
		}

		// Every X seconds (user can define):
		var updateFrequency float64
		fmt.Println("How often would you like me to call to check the current price avarage?")
		fmt.Scanlnt(&updateFrequency)

		ticker := time.NewTicker(updateFrequency * time.Second)
		quit := make(chan struct{})

		go func() {
				for {
					 select {
						case <- ticker.C:
							// Calc current market avarage value
							var prices slice = commMod()

							// Calc profit
							var avarage float32 = mathMod.average(prices)
							var profit float32 = mathMod.profit(allTransactions, avarage)

							// Print both
							fmt.printLn("Current avarage price for 1 ETH per 1 EUR is: ", avarage)
							fmt.printLn("Your profits: ", profit)
						case <- quit:
								ticker.Stop()
								return
						}
				}
		 }()

		 var quitInput string
		 fmt.Scanln(&quitInput)







	//n := 10
	//f(n)

//1
/*	go f(0)
	var input string
	fmt.Scanln(&input)*/

//2
/*	for i := 0; i < 10; i++ {
		go f(i)
	}
	var input string
	fmt.Scanln(&input)*/



// Extra, to understand scanln
/*	var input string

	fmt.Println("Who will inherit this planet after we are gone ? ")

	fmt.Scanln(&input) //<--- here

	fmt.Println("Your answer : ", input)*/


// Channels
	//1
/*	var c chan string = make(chan string)

	go pinger(c)
	go printer(c)

	var input string
	fmt.Scanln(&input)*/

//2
/*	var c chan string = make(chan string)

	go pinger(c)
	go ponger(c)
	go printer(c)

	var input string
	fmt.Scanln(&input)*/

}
