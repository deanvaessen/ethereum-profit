//Methods
	// Get from all markets the price, report back

package communicate

var httpClient = &http.Client{Timeout: 10 * time.Second}

func getJson(url string, target interface{}) error {
		r, err := httpClient.Get(url)
		if err != nil {
				return err
		}
		defer r.Body.Close()

		return json.NewDecoder(r.Body).Decode(target)
}

type PricesRequest struct {
		answer string
}

func communicate() {
	ProcessedRequest := new(PricesRequest)
	getJson("https://api.kraken.com/0/public/Ticker?pair=ETHEUR", ProcessedRequest)
	println(ProcessedRequest.answer)

	//return ProcessedRequest.answer

/*    foo1 := new(Foo) // or &Foo{}
		getJson("http://example.com", foo1)
		println(foo1.Bar)

		// alternately:

		foo2 := Foo{}
		getJson("http://example.com", &foo2)
		println(foo2.Bar)*/
}

