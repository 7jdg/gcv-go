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
		Content string `json:"content,omitempty"`

		Source *struct {
			GcsImageUri string `json:"gcsImageUri,omitempty"`
		} `json:"source,omitempty"`
	} `json:"image,omitempty"`

	Features Features `json:"features,omitempty"`

	ImageContext struct {
		LatLongRect struct {
			MinLatLng LatLng `json:"minLatLng,omitempty"`

			MaxLatLng LatLng `json:"maxLatLng,omitempty"`
		} `json:"latLongRect,omitempty"`

		LanguageHints []string `json:"languageHints,omitempty"` // for a complete list: https://cloud.google.com/translate/v2/using_rest#language-params
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

type Response struct {
	FaceAnnotations []FaceAnnotation `json:"faceAnnotations,omitempty"`

	LandmarkAnnotations []EntityAnnotation `json:"landmarkAnnotations,omitempty"`

	LogoAnnotations []EntityAnnotation `json:"logoAnnotations,omitempty"`

	LabelAnnotations []EntityAnnotation `json:"labelAnnotations,omitempty"`

	TextAnnotations []EntityAnnotation `json:"textAnnotations,omitempty"`

	SafeSearchAnnotation SafeSearchAnnotation `json:"safeSearchAnnotation,omitempty"`

	ImagePropertiesAnnotation ImageProperties `json:"imagePropertiesAnnotation,omitempty"`

	Error Status `json:"error,omitempty"`
}

type Vertex struct {
	X int64 `json:"x"`
	Y int64 `json:"y"`
}

type Landmark struct {
	Type     FaceLandmarkType `json:"type"`
	Position Position         `json:"position"`
}

type FaceLandmarkType string

const (
	UNKNOWN_LANDMARK             FaceLandmarkType = "UNKNOWN_LANDMARK"             // Unknown face landmark detected. Should not be filled.
	LEFT_EYE                     FaceLandmarkType = "LEFT_EYE"                     // Left eye.
	RIGHT_EYE                    FaceLandmarkType = "RIGHT_EYE"                    // Right eye.
	LEFT_OF_LEFT_EYEBROW         FaceLandmarkType = "LEFT_OF_LEFT_EYEBROW"         // Left of left eyebrow.
	RIGHT_OF_LEFT_EYEBROW        FaceLandmarkType = "RIGHT_OF_LEFT_EYEBROW"        // Right of left eyebrow.
	LEFT_OF_RIGHT_EYEBROW        FaceLandmarkType = "LEFT_OF_RIGHT_EYEBROW"        // Left of right eyebrow.
	RIGHT_OF_RIGHT_EYEBROW       FaceLandmarkType = "RIGHT_OF_RIGHT_EYEBROW"       // Right of right eyebrow.
	MIDPOINT_BETWEEN_EYES        FaceLandmarkType = "MIDPOINT_BETWEEN_EYES"        // Midpoint between eyes.
	NOSE_TIP                     FaceLandmarkType = "NOSE_TIP"                     // Nose tip.
	UPPER_LIP                    FaceLandmarkType = "UPPER_LIP"                    // Upper lip.
	LOWER_LIP                    FaceLandmarkType = "LOWER_LIP"                    // Lower lip.
	MOUTH_LEFT                   FaceLandmarkType = "MOUTH_LEFT"                   // Mouth left.
	MOUTH_RIGHT                  FaceLandmarkType = "MOUTH_RIGHT"                  // Mouth right.
	MOUTH_CENTER                 FaceLandmarkType = "MOUTH_CENTER"                 // Mouth center.
	NOSE_BOTTOM_RIGHT            FaceLandmarkType = "NOSE_BOTTOM_RIGHT"            // Nose, bottom right.
	NOSE_BOTTOM_LEFT             FaceLandmarkType = "NOSE_BOTTOM_LEFT"             // Nose, bottom left.
	NOSE_BOTTOM_CENTER           FaceLandmarkType = "NOSE_BOTTOM_CENTER"           // Nose, bottom center.
	LEFT_EYE_TOP_BOUNDARY        FaceLandmarkType = "LEFT_EYE_TOP_BOUNDARY"        // Left eye, top boundary.
	LEFT_EYE_RIGHT_CORNER        FaceLandmarkType = "LEFT_EYE_RIGHT_CORNER"        // Left eye, right corner.
	LEFT_EYE_BOTTOM_BOUNDARY     FaceLandmarkType = "LEFT_EYE_BOTTOM_BOUNDARY"     // Left eye, bottom boundary.
	LEFT_EYE_LEFT_CORNER         FaceLandmarkType = "LEFT_EYE_LEFT_CORNER"         // Left eye, left corner.
	RIGHT_EYE_TOP_BOUNDARY       FaceLandmarkType = "RIGHT_EYE_TOP_BOUNDARY"       // Right eye, top boundary.
	RIGHT_EYE_RIGHT_CORNER       FaceLandmarkType = "RIGHT_EYE_RIGHT_CORNER"       // Right eye, right corner.
	RIGHT_EYE_BOTTOM_BOUNDARY    FaceLandmarkType = "RIGHT_EYE_BOTTOM_BOUNDARY"    // Right eye, bottom boundary.
	RIGHT_EYE_LEFT_CORNER        FaceLandmarkType = "RIGHT_EYE_LEFT_CORNER"        // Right eye, left corner.
	LEFT_EYEBROW_UPPER_MIDPOINT  FaceLandmarkType = "LEFT_EYEBROW_UPPER_MIDPOINT"  // Left eyebrow, upper midpoint.
	RIGHT_EYEBROW_UPPER_MIDPOINT FaceLandmarkType = "RIGHT_EYEBROW_UPPER_MIDPOINT" // Right eyebrow, upper midpoint.
	LEFT_EAR_TRAGION             FaceLandmarkType = "LEFT_EAR_TRAGION"             // Left ear tragion.
	RIGHT_EAR_TRAGION            FaceLandmarkType = "RIGHT_EAR_TRAGION"            // Right ear tragion.
	LEFT_EYE_PUPIL               FaceLandmarkType = "LEFT_EYE_PUPIL"               // Left eye pupil.
	RIGHT_EYE_PUPIL              FaceLandmarkType = "RIGHT_EYE_PUPIL"              // Right eye pupil.
	FOREHEAD_GLABELLA            FaceLandmarkType = "FOREHEAD_GLABELLA"            // Forehead glabella.
	CHIN_GNATHION                FaceLandmarkType = "CHIN_GNATHION"                // Chin gnathion.
	CHIN_LEFT_GONION             FaceLandmarkType = "CHIN_LEFT_GONION"             // Chin left gonion.
	CHIN_RIGHT_GONION            FaceLandmarkType = "CHIN_RIGHT_GONION"            //
)

//

type Position struct {
	X int64 `json:"x"`
	Y int64 `json:"y"`
	Z int64 `json:"z"`
}

//

type Likelihood string

const (
	UNKNOWN       Likelihood = "UNKNOWN"       // Unknown likelihood.
	VERY_UNLIKELY Likelihood = "VERY_UNLIKELY" // The image very unlikely belongs to the vertical specified.
	UNLIKELY      Likelihood = "UNLIKELY"      // The image unlikely belongs to the vertical specified.
	POSSIBLE      Likelihood = "POSSIBLE"      // The image possibly belongs to the vertical specified.
	LIKELY        Likelihood = "LIKELY"        // The image likely belongs to the vertical specified.
	VERY_LIKELY   Likelihood = "VERY_LIKELY"   // The image very likely belongs to the vertical specified.
)

//

type LocationInfo struct {
	LatLng LatLng `json:"latLng"`
}

type LatLng struct {
	Latitude  float64 `json:"latitude,omitempty"`  // The latitude in degrees. It must be in the range [-90.0, +90.0].
	Longitude float64 `json:"longitude,omitempty"` // The longitude in degrees. It must be in the range [-180.0, +180.0].
}

type Property struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

//

type FaceAnnotation struct {
	BoundingPoly struct {
		Vertices []Vertex `json:"vertices"`
	} `json:"boundingPoly"`
	FdBoundingPoly struct {
		Vertices []Vertex `json:"vertices"`
	} `json:"fdBoundingPoly"`
	Landmarks             []Landmark `json:"landmarks"`
	RollAngle             float64    `json:"rollAngle"`
	PanAngle              float64    `json:"panAngle"`
	TiltAngle             float64    `json:"tiltAngle"`
	DetectionConfidence   float64    `json:"detectionConfidence"`
	LandmarkingConfidence float64    `json:"landmarkingConfidence"`

	JoyLikelihood          Likelihood `json:"joyLikelihood"`
	SorrowLikelihood       Likelihood `json:"sorrowLikelihood"`
	AngerLikelihood        Likelihood `json:"angerLikelihood"`
	SurpriseLikelihood     Likelihood `json:"surpriseLikelihood"`
	UnderExposedLikelihood Likelihood `json:"underExposedLikelihood"`
	BlurredLikelihood      Likelihood `json:"blurredLikelihood"`
	HeadwearLikelihood     Likelihood `json:"headwearLikelihood"`
}

type EntityAnnotation struct {
	Mid          string  `json:"mid"`
	Locale       string  `json:"locale"`
	Description  string  `json:"description"`
	Score        float64 `json:"score"`
	Confidence   float64 `json:"confidence"`
	Topicality   float64 `json:"topicality"`
	BoundingPoly struct {
		Vertices []Vertex `json:"vertices"`
	} `json:"boundingPoly"`
	Locations  []LocationInfo `json:"locations"`
	Properties []Property     `json:"properties"`
}

type SafeSearchAnnotation struct {
	Adult    Likelihood `json:"adult"`
	Spoof    Likelihood `json:"spoof"`
	Medical  Likelihood `json:"medical"`
	Violence Likelihood `json:"violence"`
}

type ImageProperties struct {
	DominantColors DominantColorsAnnotation `json:"dominantColors"`
}

type DominantColorsAnnotation struct {
	Colors []ColorInfo `json:"colors"`
}

type ColorInfo struct {
	Color         Color   `json:"color"`
	Score         float64 `json:"score"`
	PixelFraction float64 `json:"pixelFraction"`
}

type Color struct {
	Red   float64 `json:"red"`
	Green float64 `json:"green"`
	Blue  float64 `json:"blue"`
	Alpha float64 `json:"alpha"`
}

//

type Status struct {
	Code    int64    `json:"code"`
	Message string   `json:"message"`
	Details []Detail `json:"details"`
}

type Detail struct {
	Id   int64  `json:"id"`
	Type string `json:"@type"`
}
