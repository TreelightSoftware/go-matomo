package matomo

import (
	"fmt"
	"math/rand"
	"net/url"
	"time"
)

// Parameters are the content that gets sent to the API. If the field is nil, it is skipped. If it isn't nil, it will be
// automatically encoded and added to the body of the request. sendImage will be set to false. We have a matomo tag, but
// aren't currently using it and just using hard coded values.
type Parameters struct {
	RecommendedParameters     *RecommendedParameters
	UserParameters            *UserParameters
	ActionParameters          *ActionParameters
	PagePerformanceParameters *PagePerformanceParameters
	EventTrackingParameters   *EventTrackingParameters
	ContentTrackingParameters *ContentTrackingParameters
	EcommerceParameters       *EcommerceParameters
}

// RecommendedParameters are the recommended parameters that really should be provided on each call if available
type RecommendedParameters struct {
	ActionName *string `json:"action_name" matomo:"action_name"`
	URL        *string `json:"url" matomo:"url"`
	UserID     *string `json:"_id" matomo:"_id"`
	Rand       *int64  `json:"rand" matomo:"rand"` // generated at call time if not provided
	APIV       *int64  `json:"apiv" matomo:"apiv"` // always set to 1
}

// UserParameters are user specific parameters for the event
type UserParameters struct {
	URLRef           *string      `json:"urlref" matomo:"fla"`
	CVar             *string      `json:"_cvar" matomo:"_cvar"`
	IDVC             *int64       `json:"_idvc" matomo:"_idvc"`
	ViewTS           *int64       `json:"_viewts" matomo:"_viewts"`
	IDTS             *int64       `json:"_idts" matomo:"_idts"`
	CampaignName     *string      `json:"_rcn" matomo:"_rcn"`
	CampaignKeyword  *string      `json:"_rck" matomo:"_rck"`
	Resolution       *string      `json:"res" matomo:"res"`
	CurrentHour      *string      `json:"h" matomo:"h"`
	CurrentMinute    *string      `json:"m" matomo:"m"`
	CurrentSecond    *string      `json:"s" matomo:"s"`
	UserPlugins      *UserPlugins `json:"plugins" matomo:"-"`
	CookiesSupported *bool        `json:"cookie" matomo:"cookie"`
	UserAgent        *string      `json:"ua" matomo:"ua"`
	Lang             *string      `json:"lang" matomo:"lang"`
	UID              *string      `json:"uid" matomo:"uid"`
	NewVisit         *bool        `json:"new_visit" matomo:"new_visit"`
}

// UserPlugins is a sub-struct of capabilities for a user
type UserPlugins struct {
	Flash       *bool `json:"fla" matomo:"fla"`
	Java        *bool `json:"java" matomo:"java"`
	Director    *bool `json:"dir" matomo:"dir"`
	Quicktime   *bool `json:"qt" matomo:"qt"`
	RealPlayer  *bool `json:"realp" matomo:"realp"`
	PDF         *bool `json:"pdf" matomo:"pdf"`
	WMA         *bool `json:"wma" matomo:"wma"`
	Gears       *bool `json:"gears" matomo:"gears"`
	Silverlight *bool `json:"ag" matomo:"ag"`
}
type ActionParameters struct {
}

type PagePerformanceParameters struct {
}

type EventTrackingParameters struct {
}

type ContentTrackingParameters struct {
}

type EcommerceParameters struct {
}

// stringPtr converts a static string to a pointer for use in the api
func stringPtr(input string) *string {
	return &input
}

// int64Ptr converts a static int64 to a pointer for use in the api
func int64Ptr(input int64) *int64 {
	return &input
}

// boolPtr converts a static bool to a pointer for use in the api
func boolPtr(input bool) *bool {
	return &input
}

//
// below, we set up all of the encoders for the structs to convert them into
// map[string]string for embedding in the URL

func (params *Parameters) encode() map[string]string {
	ret := map[string]string{}
	if params == nil {
		return ret
	}
	// loop through the fields
	if params.RecommendedParameters != nil {
		subRet := params.RecommendedParameters.encode()
		for k, v := range subRet {
			ret[k] = v
		}
	}

	return ret
}

func (params *RecommendedParameters) encode() map[string]string {
	ret := map[string]string{}
	if params == nil {
		return ret
	}
	// set the required constants
	params.APIV = int64Ptr(1)
	if params.Rand == nil {
		params.Rand = int64Ptr(rand.Int63n(99999999999999999))
	}
	// loop through the fields
	ret["apiv"] = url.QueryEscape(fmt.Sprintf("%v", *params.APIV))
	ret["rand"] = url.QueryEscape(fmt.Sprintf("%v", *params.Rand))
	if params.ActionName != nil {
		ret["action_name"] = url.QueryEscape(*params.ActionName)
	}
	if params.UserID != nil {
		ret["_id"] = url.QueryEscape(*params.UserID)
	}
	if params.URL != nil {
		ret["url"] = url.QueryEscape(*params.URL)
	}

	return ret
}

func (params *UserParameters) encode() map[string]string {
	ret := map[string]string{}
	if params == nil {
		return ret
	}
	now := time.Now()
	// conver tthe fields
	if params.URLRef != nil {
		ret["urlref"] = url.QueryEscape(*params.URLRef)
	}
	if params.CVar != nil {
		ret["_cvar"] = url.QueryEscape(*params.CVar)
	}
	if params.IDVC != nil {
		ret["_idvc"] = url.QueryEscape(fmt.Sprintf("%v", *params.IDVC))
	}
	if params.ViewTS != nil {
		ret["_viewts"] = url.QueryEscape(fmt.Sprintf("%v", *params.ViewTS))
	}
	if params.IDTS != nil {
		ret["_idts"] = url.QueryEscape(fmt.Sprintf("%v", *params.IDTS))
	}
	if params.CampaignName != nil {
		ret["_rcn"] = url.QueryEscape(*params.CampaignName)
	}
	if params.CampaignKeyword != nil {
		ret["_rck"] = url.QueryEscape(*params.CampaignKeyword)
	}
	if params.Resolution != nil {
		ret["res"] = url.QueryEscape(*params.Resolution)
	}
	if params.CurrentHour != nil {
		ret["h"] = url.QueryEscape(*params.CurrentHour)
	} else {
		ret["h"] = url.QueryEscape(now.Format("15"))
	}
	if params.CurrentMinute != nil {
		ret["m"] = url.QueryEscape(*params.CurrentMinute)
	} else {
		ret["m"] = url.QueryEscape(now.Format("04"))
	}
	if params.CurrentSecond != nil {
		ret["s"] = url.QueryEscape(*params.CurrentSecond)
	} else {
		ret["s"] = url.QueryEscape(now.Format("05"))
	}
	if params.CookiesSupported != nil {
		if *params.CookiesSupported {
			ret["cookie"] = url.QueryEscape("1")
		} else {
			ret["cookie"] = url.QueryEscape("0")
		}
	}
	if params.Lang != nil {
		ret["lang"] = url.QueryEscape(*params.Lang)
	}
	if params.UID != nil {
		ret["uid"] = url.QueryEscape(*params.UID)
	}
	if params.NewVisit != nil {
		ret["new_visit"] = url.QueryEscape("1")
	}

	// now plugins
	plugins := params.UserPlugins.encode()
	for k, v := range plugins {
		ret[k] = v
	}

	return ret
}

func (params *UserPlugins) encode() map[string]string {
	ret := map[string]string{}
	if params == nil {
		return ret
	}
	if params.Flash != nil {
		if *params.Flash {
			ret["fla"] = url.QueryEscape("1")
		} else {
			ret["fla"] = url.QueryEscape("0")
		}
	}
	if params.Java != nil {
		if *params.Java {
			ret["java"] = url.QueryEscape("1")
		} else {
			ret["java"] = url.QueryEscape("0")
		}
	}
	if params.Director != nil {
		if *params.Director {
			ret["dir"] = url.QueryEscape("1")
		} else {
			ret["dir"] = url.QueryEscape("0")
		}
	}
	if params.Quicktime != nil {
		if *params.Quicktime {
			ret["qt"] = url.QueryEscape("1")
		} else {
			ret["qt"] = url.QueryEscape("0")
		}
	}
	if params.RealPlayer != nil {
		if *params.RealPlayer {
			ret["realp"] = url.QueryEscape("1")
		} else {
			ret["realp"] = url.QueryEscape("0")
		}
	}
	if params.PDF != nil {
		if *params.PDF {
			ret["pdf"] = url.QueryEscape("1")
		} else {
			ret["pdf"] = url.QueryEscape("0")
		}
	}
	if params.WMA != nil {
		if *params.WMA {
			ret["wma"] = url.QueryEscape("1")
		} else {
			ret["wma"] = url.QueryEscape("0")
		}
	}
	if params.Gears != nil {
		if *params.Gears {
			ret["gears"] = url.QueryEscape("1")
		} else {
			ret["gears"] = url.QueryEscape("0")
		}
	}
	if params.Silverlight != nil {
		if *params.Silverlight {
			ret["ag"] = url.QueryEscape("1")
		} else {
			ret["ag"] = url.QueryEscape("0")
		}
	}

	return ret
}
