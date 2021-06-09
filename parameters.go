package matomo

import (
	"fmt"
	"math/rand"
	"net/url"
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

type UserParameters struct {
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
