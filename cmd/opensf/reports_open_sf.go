package opensf

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type ReportOpenSf struct {
	Address            string `json:"address,omitempty"`
	AgencyResponsible  string `json:"agency_responsible,omitempty"`
	ClosedDate         string `json:"closed_date,omitempty"`
	Lat                string `json:"lat,omitempty"`
	Long               string `json:"long,omitempty"`
	MediaUrl           string `json:"media_url,omitempty"`
	Neighborhood       string `json:"neighborhoods_sffind_boundaries,omitempty"`
	Point              Point  `json:"point,omitempty"`
	PoliceDistrict     string `json:"police_district,omitempty"`
	RequestedDatetime  string `json:"requested_datetime,omitempty"`
	ServiceDetails     string `json:"service_details,omitempty"`
	ServiceName        string `json:"service_name,omitempty"`
	ServiceRequestId   string `json:"service_request_id,omitempty"`
	ServiceSubtype     string `json:"service_subtype,omitempty"`
	Source             string `json:"source,omitempty"`
	StatusDescription  string `json:"status_description,omitempty"`
	StatusNotes        string `json:"status_notes,omitempty"`
	Street             string `json:"street,omitempty"`
	SupervisorDistrict string `json:"supervisor_district,omitempty"`
	UpdatedDatetime    string `json:"updated_datetime,omitempty"`
}

type Point struct {
	Coordinates []float64 `json:"coordinates,omitempty"`
	Type        string    `json:"type,omitempty"`
}

func GetDataSFReports(start_date string, end_date string) ([]ReportOpenSf, error) {
	u := "https://data.sfgov.org/resource/ktji-gk7t.json"
	req, err := http.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("X-App_Token", "6a8b90bfef1ef751b2d161679f936b6e")

	q := req.URL.Query()
	q.Add("$where", "requested_datetime > '"+start_date+"'")
	q.Add("service_subtype", "Blocking_Bicycle_Lane")
	q.Add("$limit", "600000")
	req.URL.RawQuery = q.Encode()

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()

	reports, err := reportsFromHttpResponse(res)
	if err != nil {
		return nil, err
	}

	return reports, nil
}

func reportsFromHttpResponse(resp *http.Response) ([]ReportOpenSf, error) {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var data []ReportOpenSf
	json.Unmarshal(body, &data)

	return data, nil
}
