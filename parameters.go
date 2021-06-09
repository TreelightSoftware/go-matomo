package matomo

// Parameters are the content that gets sent to the API. If the field is nil, it is skipped. If it isn't nil, it will be
// automatically encoded and added to the body of the request. sendImage will be set to false
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
	ActionName string `json:"action_name" matomo:"action_name"`
	URL        string `json:"url" matomo:"url"`
	UserID     string `json:"_id" matomo:"_id"`
	Rand       int64  `json:"rand" matomo:"rand"` // generated at call time if not provided
	APIV       int64  `json:"apiv" matomo:"apiv"` // always set to 1
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

func (params *Parameters) encode() map[string]string {
	ret := map[string]string{}

	return ret
}
