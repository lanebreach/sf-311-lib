package lanebreach

import "github.com/omegabytes/sf-311-lib/cmd/commons"

// todo (omegabytes): this will eventually become a robust categorization system, but for now is just a scratch pad

type BlockageType struct {
	PrivateVehicle  int // 18
	Unknown         int // 45
	Lyft            int // 4
	Uber            int // 6
	UberOrLyft      int // 18
	CommuterShuttle int // 3
	USPS            int // 3
	Taxi            int // 2
	FedEx           int // 1
	UPS             int // 4
	MovingTruck     int // 1
	SFPD            int //
	SchoolBus       int // 1
	BusOther        int // 1
	CommercialTruck int // 20
	CommercialVan   int // 7
	CommercialCar   int // 4
	SFMTA           int // 1
	RentalTruck     int // 1
}

type BlockageClass struct {
	Unknown    int // 45
	Commercial int // 37
	TNC        int // 28
	Private    int // 18
	Government int // 5
	Shuttle    int // 4
	Taxi       int // 2
}

type TNC struct {
	Company      string // lyft, uber, etc
	LicensePlate string // plate
	ImageURL     string // if available, link to image sent in report
}

type RideshareCompany struct {
	TotalReports     int              // total number of reports
	Reports          []commons.Report // associated reports
	NumFreqOffenders int              // number of drivers reported in a bike lane more than 5 times
}

/* meta stuff */
// HasImage        285
// UserDescription 195

/* some reports include descriptors */
// Unattended
// FoodDelivery
// FoodPickup
// School
// Church
// AutoShop
// PsngrLoadUnload
// Contractor
// GoodsDelivery
// PrivateTrash

/* some people mention numbers */
// two
// three

/* some other frequency indicators */
// persistent problem area
// again

/* some reports mention businesses */
// PropositionChicken
// SlackHQ

/* some reports mention business names on vehicles */
// wedriveu       1
// amazon         1
// pacifico       1
// sf paratransit 1
// innopak        1
// SFMTA          1
// Penske         1
// YellowCab      1
// CASAMIGOS      1

/* we should keep track of transportation justice indicators */
// sf paratransit 1
// disabled tag   1

// our reports: 24

/* NOTES/OBSERVATIONS
- I could tell one vehicle was a lyft because I could see the phone open to Lyft's driver UI
- ticket trikes are a fucking scourge
- people offer opinions on the lane effectiveness: "too wide", "bollards don't do anything"



*/
