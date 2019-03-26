package mobile311

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Report311 struct {
	Address           string `json:"address,omitempty"`            // "address" : "Intersection of Folsom St & Zeno Pl",
	Description       string `json:"description,omitempty"`        // "description" : "USPS again",
	Lat               string `json:"lat,omitempty"`                // "lat" : 37.788153999999999,
	Long              string `json:"long,omitempty"`               // "long" : -122.3934,
	MediaUrl          string `json:"media_url,omitempty"`          // "media_url" : "http:\/\/mobile311.sfgov.org\/media\/san_francisco\/report\/photos\/5c84534552469dc41e00d717\/photo_20190309_155829.jpg",
	RequestedDatetime string `json:"requested_datetime,omitempty"` // "requested_datetime" : "2019-03-09T15:59:12-08:00",
	ServiceCode       string `json:"service_code,omitempty"`       // "service_code" : "5a6b5ac2d0521c1134854b01"
	ServiceName       string `json:"service_name,omitempty"`       // "service_name" : "Blocked Driveway & Illegal Parking",
	ServiceRequestId  string `json:"service_request_id,omitempty"` // "service_request_id" : "10582616",
	Status            string `json:"status,omitempty"`             // "status" : "closed",
	StatusNotes       string `json:"status_notes,omitempty"`       // "status_notes" : "Comment Noted. The report has been logged and will help the City collect data on double parking and bike lane violations to determine target areas and future enforcement efforts. Thank you.",
	UpdatedDatetime   string `json:"updated_datetime,omitempty"`   // "updated_datetime" : "2019-03-09T16:00:36-08:00",
}

func Get311Report(caseID string) (Report311, error) {
	url := "http://mobile311.sfgov.org/open311/v2/requests/" + caseID + ".json"
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("Accept", "application/json")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	report, err := report311FromHttpResponse(res)
	if err != nil {
		return Report311{}, err
	}
	return report, err
}

func report311FromHttpResponse(resp *http.Response) (Report311, error) {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Report311{}, err
	}

	var data []Report311
	json.Unmarshal(body, &data)

	return data[0], nil
}
