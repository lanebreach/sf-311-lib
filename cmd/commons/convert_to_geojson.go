package commons

import "fmt"

type FeatureCollection struct {
	Type     string     `json:"type,omitempty"`
	Features []*Feature `json:"features,omitempty"`
}

type Feature struct {
	Type       string `json:"type,omitempty"`
	Geometry   Point  `json:"geometry,omitempty"`
	Properties Report `json:"properties,omitempty"`
}

// ConvertToGeoJson embeds the properties of a given 311 report into a geojson object conforming to RFC 7946
func ConvertToGeoJson(reports []Report) (*FeatureCollection, error) {
	fmt.Println(len(reports))
	converted := make([]*Feature, len(reports))
	collection := &FeatureCollection{
		Type: "FeatureCollection",
	}

	for i, report := range reports {
		cp := Feature{
			Type:       "Feature",
			Geometry:   report.Point,
			Properties: report,
		}

		converted[i] = &cp
		fmt.Println(converted)
	}
	collection.Features = converted

	return collection, nil
}
