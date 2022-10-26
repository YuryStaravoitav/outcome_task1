package mystructs

import (
	"fmt"
	"os"
	"sort"

	"github.com/gocarina/gocsv"
)

type AdExposureData struct {
	UserId     string   `csv:"user_id"`
	TimeStamp  DateTime `csv:"timestamp"`
	CreativeId string   `csv:"creative_id"`
}

func ReadAdExposureDataFile() []*AdExposureData {
	adExposureFile, err := os.OpenFile("ad_exposures.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer adExposureFile.Close()

	ads := []*AdExposureData{}

	if err := gocsv.UnmarshalFile(adExposureFile, &ads); err != nil { // Load clients from file
		panic(err)
	}

	return ads

}

func SortAds(ads []*AdExposureData) {
	sort.Slice(ads, func(i, j int) bool {
		if ads[i].UserId == ads[j].UserId {
			return ads[i].TimeStamp.Before(ads[j].TimeStamp.Time)
		}
		return ads[i].UserId < ads[j].UserId
	})
}

func MapOfUserIds(ads []*AdExposureData) map[string][]*AdExposureData {
	mapOfAds := make(map[string][]*AdExposureData)
	for _, ad := range ads {
		mapOfAds[ad.UserId] = append(mapOfAds[ad.UserId], ad)
	}
	return mapOfAds
}

func PrintMapAds(ads map[string][]*AdExposureData) {
	for key, ad := range ads {
		for _, item := range ad {
			fmt.Printf("user_id: %s, timestamp: %s, creative_id: %s", key, item.TimeStamp.String(), item.CreativeId)
			fmt.Println()
		}

	}
}
