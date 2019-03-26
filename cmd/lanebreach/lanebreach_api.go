package lanebreach

import "github.com/omegabytes/sf-311-lib/cmd/MTA"

type LaneBreachReport struct {
	Address            string   `json:"address,omitempty"`
	AgencyResponsible  string   `json:"agency_responsible,omitempty"`
	ClosedDate         string   `json:"closed_date,omitempty"`
	Description        string   `json:"description,omitempty"`
	Lat                string   `json:"lat,omitempty"`
	Long               string   `json:"long,omitempty"`
	MediaUrl           string   `json:"media_url,omitempty"`
	Neighborhood       string   `json:"neighborhoods_sffind_boundaries,omitempty"`
	Point              Point    `json:"point,omitempty"`
	PoliceDistrict     string   `json:"police_district,omitempty"`
	RequestedDatetime  string   `json:"requested_datetime,omitempty"`
	ServiceCode        string   `json:"service_code,omitempty"`
	ServiceDetails     string   `json:"service_details,omitempty"`
	ServiceName        string   `json:"service_name,omitempty"`
	ServiceRequestId   string   `json:"service_request_id,omitempty"`
	ServiceSubtype     string   `json:"service_subtype,omitempty"`
	Source             string   `json:"source,omitempty"`
	Status             string   `json:"status,omitempty"`
	StatusDescription  string   `json:"status_description,omitempty"`
	Street             string   `json:"street,omitempty"`
	StatusNotes        string   `json:"status_notes,omitempty"`
	SupervisorDistrict string   `json:"supervisor_district,omitempty"`
	UpdatedDatetime    string   `json:"updated_datetime,omitempty"`
	MetaData           MetaData `json:"meta_data,omitempty"`
}

type Point struct {
	Coordinates []float64 `json:"coordinates,omitempty"`
	Type        string    `json:"type,omitempty"`
}

type MetaData struct {
	BikewayNetworkId string `json:"bikeway_network_id,omitempty"`
}

type BikeLane struct {
	Barrier                string          `json:"barrier,omitempty"`    // identifies the barrier type for separated bikeways (CONCRETE, PARKING, SAFE HIT POSTS)
	BackInAngledParking    string          `json:"baip,omitempty"`       // indicates whether the bike lane is adjacent to back in angled parking
	Counter                mta.BikeCounter `json:"counter,omitempty"`    // associated bike counter, if any. Derived.
	Contraflow             string          `json:"contraflow,omitempty"` // facility travels in opposite direction of vehicular flow;
	CreatedUs              string          `json:"created_us,omitempty"` // userId of the user who created the record
	DateCreated            string          `json:"date_creat,omitempty"` // date the record was created
	DateLast               string          `json:"date_last_,omitempty"` // todo: look up date_last_
	LaneTravelDirection    string          `json:"dir,omitempty"`        // cardinal direction of the flow of bike traffic
	BlockDirectionality    string          `json:"direct,omitempty"`     // number of directions on that block (1 way or 2 way)
	Double                 string          `json:"double,omitempty"`     // 0 = one side of the street; 1 = other side of the street. Possibly deprecated by LaneTravelDirection
	FacilityType           string          `json:"facility_t,omitempty"` // identifies the bikeway facility based on national bikeway functional classification
	FromStreet             string          `json:"from_st,omitempty"`    // start of the block
	FiscalYearInstalled    string          `json:"fy,omitempty"`         // fiscal year installed
	FiscalQuarterInstalled string          `json:"qtr,omitempty"`        // fiscal quarter installed
	Geometry               string          `json:"geom,omitempty"`       // todo: look up geom
	GlobalId               string          `json:"global_id,omitempty"`  // todo: look up global_id
	Greenwave              string          `json:"greenwave,omitempty"`  // segment is part of greenwave
	Id                     int             `json:"id,omitempty"`         // todo: look up id
	InstallMonth           string          `json:"install_mo,omitempty"` // calendar month the segment was installed
	InstallYear            string          `json:"install_yr,omitempty"` // calendar year the segment was installed
	LastEdited             string          `json:"last_edite,omitempty"` // userID of the user who last edited the record
	Length                 string          `json:"length,omitempty"`     // length of feature in miles
	BikeRouteNumber        string          `json:"number,omitempty"`     // historical bike route number
	Notes                  string          `json:"notes,omitempty"`      // additional notes
	ObjectId               string          `json:"objectid,omitempty"`   // todo: look up object_id
	Raised                 string          `json:"raised,omitempty"`     // bikeway is raised up from street level
	ShapeLength            string          `json:"shape_len,omitempty"`  // length of feature in feet
	Sharrow                int             `json:"sharrow,omitempty"`    // number of sharrows present
	SmallSweeper           string          `json:"sm_sweeper,omitempty"` // indicates if the lane requires a small street sweeper
	Street                 string          `json:"street,omitempty"`     // street name and type in proper case
	StreetCenterlineId     string          `json:"cnn,omitempty"`        // unique ID correlating to SF's street centerline dataset
	StreetName             string          `json:"streetname,omitempty"` // street name in all caps
	SurfaceTreatments      string          `json:"surface_tr,omitempty"` // additional surface treatments
	Symbology              string          `json:"symbology,omitempty"`  // used to symbolize different facility types (SEPARATED BIKEWAY, BIKE LANE, BIKE ROUTE, NEIGHBORWAY, BIKE PATH)
	TimeCreated            string          `json:"time_creat,omitempty"` // time the record was created
	TimeLast               string          `json:"time_last_,omitempty"` // todo look up time_last_
	ToStreet               string          `json:"to_st,omitempty"`      // end of the block
	UpdateYear             string          `json:"update_yr,omitempty"`  // year the feature was last upgraded
	UpdateMonth            string          `json:"update_mo,omitempty"`  // month the feature was last upgraded
}
