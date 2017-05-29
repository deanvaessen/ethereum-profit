package main

import (
	"fmt"
	"helpers/user"
	"helpers/math"
	"helpers/communicate"
	"time"
)

func main() {

	// Define a username so we can store your data
	var newUser string
	fmt.Println("Please give me your name so I can see if I need to set you up or can load your details.")
	fmt.Scanln(&newUser)

		// Check if we already have this user
		connectedUser := &user.Person{Name: newUser}
		connectedUser.BuildProfile()

		// Every X seconds (user can define):
		var updateFrequency float32
		fmt.Println("Every how many seconds would you like me to call to check the current price avarage?")
		fmt.Scanln(&updateFrequency)

		ticker := time.NewTicker(time.Duration(updateFrequency) * time.Second)
		quit := make(chan struct{})

		go func() {
				for {
					 select {
						case <- ticker.C:
							// Calc current market avarage value
							var prices[]float32 = communicate.GetPrices()

							// Calc profit
							var avarage float32 = math.Average(prices)
							var profit float32 = math.Profit(connectedUser.Transactions, avarage)

							// Print both
							fmt.Println("Current avarage price for 1 ETH per 1 USD is: ", avarage)
							fmt.Println("Your profits: ", profit)
						case <- quit:
								ticker.Stop()
								return
						}
				}
		 }()

		 var quitInput string
		 fmt.Scanln(&quitInput)

}
