//Methods
	// Calc avarage
	// Calc profit

package math
/*
import "fmt"

type rect struct {
    width, height int
}
func (r *rect) area() int {
    return r.width * r.height
}
func (r rect) perim() int {
    return 2*r.width + 2*r.height
}


func main() {
    r := rect{width: 10, height: 5}
    fmt.Println("area: ", r.area())
    fmt.Println("perim:", r.perim())
    rp := &r
    fmt.Println("area: ", rp.area())
    fmt.Println("perim:", rp.perim())
}
*/

/*
func avg(rest ...float64) float64 { //<----- here!
	sum := 0.0
	for _, num := range rest {
		sum += num
	}
	return float64(sum) / float64(len(rest))
}

func main() {
	fmt.Printf("%0.2f\n", avg(1, 2))
	fmt.Printf("%0.2f\n", avg(1, 2, 3, 4, 5, 6, 7, 8, 9))
	fmt.Printf("%0.2f\n", avg(99, 98, 87, 110, 121, 2))
}

*/

//var avarage float32 = mathMod.avarage(prices)
//var profit float32 = mathMod.profit(allTransactions, avarage)


import "fmt"

type portfolioMath struct {}


func (pM *portfolioMath) average (prices ...float32) float32 {
	sum := 0.0
	for _, num := range prices {
		sum += num
	}
	return float32(sum) / float32(len(prices))

}

func (pM *portfolioMath) profit (allTransactions interface, currentPriceAverage float32) float32 {
	totalEth := 0.0
	totalCost := 0.0

	currentMarketValue := 0.0
	profitLoss := 0.0

	for _, transaction := range allTranactions {
		for _, i := range transaction {
			totalEth += Amount
		}
	}

	currentMarketValue = totalEth * currentPriceAverage
	profitLoss = currentPriceAverage - currentMarketValue

	return profitLoss
}

func main() {
/*    r := portfolioMath{width: 10, height: 5}
    fmt.Println("area: ", r.area())
    fmt.Println("perim:", r.perim())
    rp := &r
    fmt.Println("area: ", rp.area())
    fmt.Println("perim:", rp.perim())*/
}