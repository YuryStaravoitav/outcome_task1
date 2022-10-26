package mystructs

import (
	"fmt"
	"os"
)

type Outcome struct {
	CreativeId string
	Count      int
	Amount     float64
}

func CalculateOutcome(ads map[string][]*AdExposureData, sales []*SalesData) map[string]*Outcome {
	outcomes := make(map[string]*Outcome)
	for _, sale := range sales {
		if ads[sale.UserId] == nil {
			continue
		}
		creativeId := getCreativeIdForSale(ads[sale.UserId], sale)
		if outcomes[creativeId] == nil {
			out := &Outcome{CreativeId: creativeId, Count: 1, Amount: sale.Amount}
			outcomes[creativeId] = out
		} else {
			out := outcomes[creativeId]
			out.Count++
			out.Amount += sale.Amount
		}
	}

	return outcomes
}

func PrintOutcome(outcomes map[string]*Outcome) {
	count, amount := getTotal(outcomes)
	fmt.Println("dimension,value,num_purchasers,total_sales")
	fmt.Printf("overall,overall,%d,%f", count, amount)
	fmt.Println()
	for _, out := range outcomes {
		fmt.Printf("CreativeId: %s, count: %d, amount: %f", out.CreativeId, out.Count, out.Amount)
		fmt.Println()
	}
}

func getCreativeIdForSale(ads []*AdExposureData, sale *SalesData) string {
	for i := 0; i < len(ads)-2; i++ {
		if ads[i].TimeStamp.Time.Before(sale.TimeStamp.Time) && ads[i+1].TimeStamp.After(sale.TimeStamp.Time) {
			return ads[i].CreativeId
		}

	}
	return ads[len(ads)-1].CreativeId
}

func getTotal(outcomes map[string]*Outcome) (int64, float64) {
	totalCount := 0
	totalAmount := 0.0
	for _, out := range outcomes {
		totalCount += out.Count
		totalAmount += out.Amount
	}

	return int64(totalCount), totalAmount

}

func OutcomesToCsv(outcomes map[string]*Outcome) {
	outcomeFiles, err := os.OpenFile("outcome.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer outcomeFiles.Close()

	count, amount := getTotal(outcomes)

	outcomeFiles.WriteString("dimension,value,num_purchasers,total_sales\n")
	outcomeFiles.WriteString(fmt.Sprintf("overall,overall,%d,$%.2f\n", count, amount))

	for key, out := range outcomes {
		outcomeFiles.WriteString(fmt.Sprintf("creative_id,%s,%d,$%.2f\n", key, out.Count, out.Amount))
	}

}
