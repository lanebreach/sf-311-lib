package mta

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"github.com/omegabytes/sf-311-lib/cmd/commons"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

/*
- https://stats.sfmta.com/t/public/views/AutomatedBicycleCounters/HOURLYTABLE.csv
- https://stats.sfmta.com/t/public/views/AutomatedBicycleCounters/MONTHLYTABLE.csv
- https://stats.sfmta.com/t/public/views/AutomatedBicycleCounters/AVGWEEKDAYTABLE.csv
- https://stats.sfmta.com/t/public/views/AutomatedBicycleCounters/NEWTABLE.csv ("counters most recently installed")
*/

type BikeCounter struct {
	CounterLocation   string `json:"counter_location,omitempty"`         // 2nd St North of Townsend SB
	Days              string `json:"days,omitempty"`                     // SUNDAY
	HourOfCollection  string `json:"hour_of_collection_time,omitempty"`  // 0
	MonthOfCollection string `json:"month_of_collection_time,omitempty"` // January
	TotalBikeCount    string `json:"total_bike_count,omitempty"`         // 1,000
	YearOfCollection  string `json:"year_of_collection_time,omitempty"`  // 2016
}

func GetHourlyCounts() {
	GetCSV("https://stats.sfmta.com/t/public/views/AutomatedBicycleCounters/HOURLYTABLE.csv", true)
}

func GetMonthlyCounts() {}

func GetAverageWeekdayCounts() {}

func GetCountsByLocation() {
	counters := readCountersFromFile("./bike_counts.json")
	locations, _ := getListOfLocations(counters)
	log.Println(locations)
}

func GetCSV(path string, headers bool) ([]BikeCounter, error) {
	resp, err := http.Get(path)
	//if err != nil {
	//	return nil, err
	//}
	defer resp.Body.Close()

	csvFile, _ := os.Open(path)
	reader := csv.NewReader(bufio.NewReader(csvFile))
	var counters []BikeCounter

	for {
		line, err := reader.Read()

		if headers {
			headers = false
			continue
		}

		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		counters = append(counters, BikeCounter{
			CounterLocation:   line[0],
			Days:              line[1],
			HourOfCollection:  line[2],
			MonthOfCollection: line[3],
			YearOfCollection:  line[4],
			TotalBikeCount:    line[5],
		})
	}
	counterJson, _ := json.Marshal(counters)
	//fmt.Println(string(counterJson))
	err = ioutil.WriteFile("bike_counts.json", counterJson, 0644)
	commons.Check(err)
	return counters, nil
}

func readCSVFromUrl(url string) ([][]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	reader := csv.NewReader(resp.Body)
	reader.Comma = ';'
	data, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	return data, nil
}

func readCountersFromFile(path string) []BikeCounter {
	jsonFile, err := os.Open(path)
	commons.Check(err)
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var counters []BikeCounter
	json.Unmarshal(byteValue, &counters)

	return counters
}

func getListOfLocations(counters []BikeCounter) (map[string]int, error) {
	m := make(map[string]int)

	replacer := strings.NewReplacer(",", "")

	for _, counter := range counters {
		if strings.Contains(counter.TotalBikeCount, ",") {
			counter.TotalBikeCount = replacer.Replace(counter.TotalBikeCount)
		}
		count, err := strconv.Atoi(counter.TotalBikeCount)
		if err != nil {
			log.Println(err)
		}
		if m[counter.CounterLocation] == 0 {
			m[counter.CounterLocation] = count
		} else {
			m[counter.CounterLocation] += count
		}
	}
	return m, nil
}
