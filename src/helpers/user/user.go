package user

import (
	"fmt"
	"encoding/json"
	"bufio"
	"os"
)


type Person struct {
	Name string
	IsKnown bool
	Transactions []Transaction
}

type Transaction struct {
	Date string
	Price float32
	Exchange string
	Amount float32
}

func (p *Person) BuildProfile () {

	defaultBasePath := "data/users/"
	userFileLocation := p.Name
	defaultFileExtension := ".txt"

	filePath := defaultBasePath + userFileLocation + defaultFileExtension


	userData, err := os.Open(filePath)
	if err != nil {
	    //log.Fatal(err)
	    fmt.Println("Welcome! I will get you set up.")
	    p.CreateInitialUserStorage()
	    p.Init()
		p.LoadTransactions()
	} else {
		p.IsKnown = true

		fmt.Println("Welcome back " + p.Name)

		loadedUser := &Person{ Name : p.Name }
		var amountExtraTransactions float32

		fmt.Println("How many extra transactions would you like to save?")
		fmt.Scanln(&amountExtraTransactions)

		if (amountExtraTransactions == 0) {
			// Scan all transactions and load them into the struct
			p.LoadTransactions()
		} else if (amountExtraTransactions >= 1 ) {

			loadedUser.SaveTransactions(amountExtraTransactions)
			p.LoadTransactions()
		}
	}
	// Done. Go to calculate extra profits in main.go.
	defer userData.Close()
}

func (p *Person) LoadTransactions () {
	//Load all transactions

	defaultBasePath := "data/users/"
	userFileLocation := p.Name
	defaultFileExtension := ".txt"

	filePath := defaultBasePath + userFileLocation + defaultFileExtension

	userData, readErr := os.Open(filePath)
	scanner := bufio.NewScanner(userData)

	if (readErr != nil) {
		fmt.Println(readErr)
	}

	for scanner.Scan() {
	    readLine := []byte(scanner.Text())

	    var decoded Transaction = decodeFromJSON(readLine)

	    p.Transactions = append(p.Transactions, decoded)
	}

	if err := scanner.Err(); err != nil {
	    panic(err)
	}

	defer userData.Close()
}

func encodeToJSON (t Transaction) ([]byte) {
	b, err := json.Marshal(t)
	if err != nil {
		panic(err)
	}

	return b
}

func decodeFromJSON (b []byte) (Transaction) {
	var decoded Transaction

	err := json.Unmarshal(b, &decoded)
	if err != nil {
		panic(err)
	}

	return decoded
}

func (p *Person) SaveTransactions (AmountOfTransactions float32) {

	for i := 0.0; i < float64(AmountOfTransactions); i++ {
		var Date string
		var Price float32
		var Exchange string
		var Amount float32

		fmt.Println("At what exchange did you buy?")
		fmt.Scanln(&Exchange)
		fmt.Println("At what date did you buy in?")
		fmt.Scanln(&Date)
		fmt.Println("At what USD price did you buy the whole amount?")
		fmt.Scanln(&Price)
		fmt.Println("What was the total amount you got for your investment?")
		fmt.Scanln(&Amount)

		p.Transactions = append(p.Transactions, Transaction{Date : Date, Price : Price, Exchange: Exchange, Amount: Amount})

		if ((i +1) < float64(AmountOfTransactions)) {
			fmt.Println("Going to the next transaction...\n")
		}
	}

	defaultBasePath := "data/users/"
	userFileLocation := p.Name
	defaultFileExtension := ".txt"

	filePath := defaultBasePath + userFileLocation + defaultFileExtension

	for _, transaction := range p.Transactions {
		var encodedTransaction []byte = encodeToJSON(transaction)

		//Add lines
		f, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, 0600)
		if err != nil {
		    panic(err)
		}

		defer f.Close()


		if _, err = f.Write(encodedTransaction); err != nil {
		    panic(err)
		}

		if _, err = f.WriteString("\r\n"); err != nil {
		    panic(err)
		}
	}
}

func (p *Person) CreateInitialUserStorage () {
	//Save userdetails to a file

	defaultBasePath := "data/users/"
	userFileLocation := p.Name
	defaultFileExtension := ".txt"

	filePath := defaultBasePath + userFileLocation + defaultFileExtension

	// Create the file
	fo, err := os.Create(filePath)
	if err != nil {
	    panic(err)
	}
	// close fo on exit and check for its returned error
	defer func() {
	    if err := fo.Close(); err != nil {
	        panic(err)
	    }
	}()
}

func (p *Person) Init () {
	// Run functions to set up the user profile
	var  AmountOfTransactions int
	fmt.Println("How many transactions would you like to save?")
	fmt.Scanln(&AmountOfTransactions)

	p.SaveTransactions(float32(AmountOfTransactions))
}