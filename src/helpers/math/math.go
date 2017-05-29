package math
import (
	"helpers/user"
)


func Average (prices []float32) float32 {
	var sum float32 = 0.0
	for _, num := range prices {
		sum += num
	}
	return float32(sum) / float32(len(prices))

}

func Profit (allTransactions []user.Transaction, currentPriceAverage float32) float32 {
	var totalEth float32 = 0.0
	var totalCost float32 = 0.0

	var currentMarketValue float32 = 0.0
	var profitLoss float32 = 0.0

	for _, transaction := range allTransactions {
		totalEth += transaction.Amount
		totalCost += transaction.Price
	}

	currentMarketValue = totalEth * currentPriceAverage
	profitLoss = currentMarketValue - totalCost

	return profitLoss
}