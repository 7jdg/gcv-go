package gcvgo

import (
	"net/http"
)

type Credentials struct {
	APIkey string
}

type Client struct {
	httpClient  *http.Client
	Credentials Credentials
}

type Request struct {
	Image struct {
		Content string `json:"content"`

		Source struct {
			GcsImageUri string `json:"gcsImageUri"`
		} `json:"source"`
	} `json:"image,omitempty"`

	Features Features `json:"features,omitempty"`

	ImageContext struct {
		LatLongRect struct {
			MinLatLng struct {
				Latitude  float64 "latitude"  // The latitude in degrees. It must be in the range [-90.0, +90.0].
				Longitude float64 "longitude" // The longitude in degrees. It must be in the range [-180.0, +180.0].
			} "minLatLng"

			MaxLatLng struct {
				Latitude  float64 "latitude"  // The latitude in degrees. It must be in the range [-90.0, +90.0].
				Longitude float64 "longitude" // The longitude in degrees. It must be in the range [-180.0, +180.0].
			} "maxLatLng"
		} `json:"latLongRect"`

		LanguageHints []string `json:"languageHints"` // for a complete list: https://cloud.google.com/translate/v2/using_rest#language-params
	} `json:"imageContext,omitempty"`
}

type Requests []Request

//

type Features []Feature

type Feature struct {
	Type       FeatureType `json:"type"`
	MaxResults int64       `json:"maxResults"`
}

//

type FeatureType string

const (
	TYPE_UNSPECIFIED      FeatureType = "TYPE_UNSPECIFIED"      // Unspecified feature type.
	FACE_DETECTION        FeatureType = "FACE_DETECTION"        // Run face detection.
	LANDMARK_DETECTION    FeatureType = "LANDMARK_DETECTION"    // Run landmark detection.
	LOGO_DETECTION        FeatureType = "LOGO_DETECTION"        // Run logo detection.
	LABEL_DETECTION       FeatureType = "LABEL_DETECTION"       // Run label detection.
	TEXT_DETECTION        FeatureType = "TEXT_DETECTION"        // Run OCR.
	SAFE_SEARCH_DETECTION FeatureType = "SAFE_SEARCH_DETECTION" // Run various computer vision models to compute image safe-search properties.
	IMAGE_PROPERTIES      FeatureType = "IMAGE_PROPERTIES"      // Compute a set of properties about the image (such as the image's dominant colors).
)

//

type Response struct{}
