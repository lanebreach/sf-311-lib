package main

import (
	"encoding/json"
	"github.com/omegabytes/sf-311-lib/cmd/commons"
	"github.com/omegabytes/sf-311-lib/cmd/mobile311"
	"github.com/omegabytes/sf-311-lib/cmd/opensf"
	"io/ioutil"
	"log"
	"time"
)

// todo (omegabytes): we shouldnt be calling the lib functions from a main included in the library, but rather import the package to another project

func main() {
	combineReports()

	// MTA bike counter
	//mta.GetCSV("./output/bike_counters.csv", true)
	//mta.GetHourlyCounts()
	//mta.GetCountsByLocation()
}

// combineReports makes calls to both mobile311 and sf311 and combines their response body params
// we do this because some fields from mobile311 are missing from sf311 and vice versa
func combineReports() {
	reports, err := opensf.GetDataSFReports("2019-03-05T00:00:00", "2019-03-13T00:00:00")
	commons.Check(err)

	requests := make(chan opensf.ReportOpenSf, len(reports))
	for _, report := range reports {
		requests <- report
	}
	close(requests)
	log.Println(len(reports))

	combinedReports := make([]commons.Report, len(reports))

	// limiter delays the request to avoid throttling
	//limiter := time.Tick(200 * time.Millisecond)
	limiter := time.Tick(1 * time.Second)
	for req := range requests {
		<-limiter
		report, err := mobile311.Get311Report(req.ServiceRequestId)
		commons.Check(err)

		combinedReport, err := commons.CombineReports(report, req)

		log.Println(combinedReport.Description)

		combinedReports = append(combinedReports, combinedReport)
	}

	geojson, err := commons.ConvertToGeoJson(combinedReports)
	commons.Check(err)

	out, err := json.MarshalIndent(geojson, "", "  ")
	commons.Check(err)
	err = ioutil.WriteFile("output.geojson", out, 0644)
}
