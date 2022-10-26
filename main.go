package main

import (
	mystructs "anonym.task1/mystucts"
)

func main() {
	ads := mystructs.ReadAdExposureDataFile()
	mystructs.SortAds(ads)
	mapOfAds := mystructs.MapOfUserIds(ads)
	//	mystructs.PrintMapAds(mapOfAds)

	sales := mystructs.ReadSalesDataFile()
	mystructs.SortSales(sales)

	outcome := mystructs.CalculateOutcome(mapOfAds, sales)
	mystructs.PrintOutcome(outcome)
	mystructs.OutcomesToCsv(outcome)
}
