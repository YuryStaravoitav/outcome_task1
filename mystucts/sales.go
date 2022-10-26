package mystructs

import (
	"os"
	"sort"

	"github.com/gocarina/gocsv"
)

type SalesData struct {
	UserId    string   `csv:"user_id"`
	TimeStamp DateTime `csv:"timestamp"`
	Amount    float64  `csv:"amount"`
}

func ReadSalesDataFile() []*SalesData {
	salesFile, err := os.OpenFile("sales_data.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer salesFile.Close()

	sales := []*SalesData{}

	if err := gocsv.UnmarshalFile(salesFile, &sales); err != nil { // Load clients from file
		panic(err)
	}

	return sales
}

func SortSales(sales []*SalesData) {
	sort.Slice(sales, func(i, j int) bool {
		if sales[i].UserId == sales[j].UserId {
			return sales[i].TimeStamp.Before(sales[j].TimeStamp.Time)
		}
		return sales[i].UserId < sales[j].UserId
	})
}
