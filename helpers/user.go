//Methods:
	//Check user
	//Save user
	//Update user

package main

import (
    "io/ioutil"
)

type user struct {
  Name string
  Transactions slice?
}

type transaction struct {
	Date string
	Price float32
	Exchange string
	Amount float32
}

func (u *user) Exists () {
	defaultBasePath := "users/"
	userFileLocation := user.Name
	defaultFileExtension := ".txt"

	// read the whole file at once
	b, err := ioutil.ReadFile(defaultBasePath + userFileLocation + defaultFileExtension)
	if err != nil {
		fmt.printLn("Welcome! I will get you set up.")
	    //panic(err)
	} else {
		fmt.printLn("Welcome back " + user.Name)
		bool extraTransactions
		float32 amountExtraTransactions

		fmt.Println("Do you want to save extra transactions?")
		fmt.Scanln(&extraTransactions)

		if (extraTransactions) {
			fmt.Println("How many transactions would you like to save?")
			fmt.Scanln(&amountExtraTransactions)

			user.SaveTransactions(AmountOfTransactions)
		} else {

		}
	}

/*	// write the whole body at once
	err = ioutil.WriteFile("output.txt", b, 0644)
	if err != nil {
	    panic(err)
	}*/


	// check if exists
	if userExists {
		fmt.printLn("Welcome back " + user.Name)
		return true
	} else {
		
		return false
	}
}

func (u *user) LoadTransactions () {
	//Load all transactions

	defaultBasePath := "users/"
	userFileLocation := user.Name
	defaultFileExtension := ".txt"

	// read the whole file at once
	b, err := ioutil.ReadFile(defaultBasePath + userFileLocation + defaultFileExtension)
	if err != nil {
	    panic(err)
	} else {

	}
}


func (u *user) Save () {
	//Save userdetails to a file
}

func (u *user) Update () {
	//Update user details (transactions, etc.)
	// Keep it easy, just read the whole file, update the transaction details (overwrite all), overwrite the file
}

func (u *user, transactionDetails) SaveTransactions () {

	for := i; i == AmountOfTransactions; i++ {
		string Exchange
		string Date
		float32 Price
		float32 Amount

		fmt.Println("How many transactions would you like to save?")
		fmt.Scanln(&Exchange)
		fmt.Println("At what date did you buy in?")
		fmt.Scanln(&Date)
		fmt.Println("At what EUR price did you buy the whole amount?")
		fmt.Scanln(&Price)
		fmt.Println("What was the total amount you got for your investment?")
		fmt.Scanln(&Amount)

		make interface{Date, Exchange, Price, Amount}
		user.transactions.push()??? := transaction{transactionDetails[]}
	}

	
}

func (user *user) Init () {
	// Run functions to set up the user profile
	int AmountOfTransactions
	fmt.Println("How many transactions would you like to save?")
	fmt.Scanln(&AmountOfTransactions)

	user.SaveTransactions(AmountOfTransactions)
}