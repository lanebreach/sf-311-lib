package commons

import (
	"encoding/json"
	"github.com/omegabytes/sf-311-lib/cmd/mobile311"
	"github.com/omegabytes/sf-311-lib/cmd/opensf"
	"io/ioutil"
	"log"
	"os"
)

type Report struct {
	Address            string `json:"address,omitempty"` // "address" : "Intersection of Folsom St & Zeno Pl",
	AgencyResponsible  string `json:"agency_responsible,omitempty"`
	ClosedDate         string `json:"closed_date,omitempty"`
	Description        string `json:"description,omitempty"` // "description" : "USPS again",
	Lat                string `json:"lat,omitempty"`         // "lat" : 37.788153999999999,
	Long               string `json:"long,omitempty"`        // "long" : -122.3934,
	MediaUrl           string `json:"media_url,omitempty"`   // "media_url" : "http:\/\/mobile311.sfgov.org\/media\/san_francisco\/report\/photos\/5c84534552469dc41e00d717\/photo_20190309_155829.jpg",
	Neighborhood       string `json:"neighborhoods_sffind_boundaries,omitempty"`
	Point              Point  `json:"point,omitempty"`
	PoliceDistrict     string `json:"police_district,omitempty"`
	RequestedDatetime  string `json:"requested_datetime,omitempty"` // "requested_datetime" : "2019-03-09T15:59:12-08:00",
	ServiceCode        string `json:"service_code,omitempty"`       // "service_code" : "5a6b5ac2d0521c1134854b01"
	ServiceDetails     string `json:"service_details,omitempty"`
	ServiceName        string `json:"service_name,omitempty"`       // "service_name" : "Blocked Driveway & Illegal Parking",
	ServiceRequestId   string `json:"service_request_id,omitempty"` // "service_request_id" : "10582616",
	ServiceSubtype     string `json:"service_subtype,omitempty"`
	Source             string `json:"source,omitempty"`
	Status             string `json:"status,omitempty"` // "status" : "closed",
	StatusDescription  string `json:"status_description,omitempty"`
	Street             string `json:"street,omitempty"`
	StatusNotes        string `json:"status_notes,omitempty"` // "status_notes" : "Comment Noted. The report has been logged and will help the City collect data on double parking and bike lane violations to determine target areas and future enforcement efforts. Thank you.",
	SupervisorDistrict string `json:"supervisor_district,omitempty"`
	UpdatedDatetime    string `json:"updated_datetime,omitempty"` // "updated_datetime" : "2019-03-09T16:00:36-08:00",
}

type Point struct {
	Coordinates []float64 `json:"coordinates,omitempty"`
	Type        string    `json:"type,omitempty"`
}

func CombineReports(mobile311 mobile311.Report311, openSF opensf.ReportOpenSf) (Report, error) {
	combinedReport := Report{}

	combinedReport = Report{
		Address:            openSF.Address,
		AgencyResponsible:  openSF.AgencyResponsible,
		ClosedDate:         openSF.ClosedDate,
		Description:        mobile311.Description,
		Lat:                openSF.Lat,
		Long:               openSF.Long,
		MediaUrl:           openSF.MediaUrl,
		Neighborhood:       openSF.Neighborhood,
		Point:              Point(openSF.Point),
		PoliceDistrict:     openSF.PoliceDistrict,
		RequestedDatetime:  openSF.RequestedDatetime,
		ServiceCode:        mobile311.ServiceCode,
		ServiceDetails:     openSF.ServiceDetails,
		ServiceName:        openSF.ServiceName,
		ServiceRequestId:   openSF.ServiceRequestId,
		ServiceSubtype:     openSF.ServiceSubtype,
		Source:             openSF.Source,
		Status:             mobile311.Status,
		StatusDescription:  openSF.StatusDescription,
		StatusNotes:        openSF.StatusNotes,
		Street:             openSF.Street,
		SupervisorDistrict: openSF.SupervisorDistrict,
		UpdatedDatetime:    openSF.UpdatedDatetime,
	}

	return combinedReport, nil
}

func ReadCombinedReportsFromFile(path string) {
	jsonFile, err := os.Open(path) // "./output/output.json"
	Check(err)
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var reports []Report
	json.Unmarshal(byteValue, &reports)

	converted, err := ConvertToGeoJson(reports)
	out, err := json.MarshalIndent(converted, "", "")
	Check(err)

	err = ioutil.WriteFile("out2.geojson", out, 0644)
}

func Check(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
