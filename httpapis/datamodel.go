package httpapis

// PushReqest push接口的请求结构
type PushReqest struct {
	PushSourceName     string         `json:"push_source_name"`
	DF                 map[string]any `json:"df"`
	AllowRegistryCache bool           `json:"allow_registry_cache"`
	To                 string         `json:"to"`
}

type GetOnlineFeaturesHttpRespMetaData struct {
	FeatureNames []string `json:"feature_names"`
}
type GetOnlineFeaturesHttpRespResult struct {
	EventTimestamps []string `json:"event_timestamps"`
	Statuses        []string `json:"statuses"`
	Values          []any    `json:"values"`
}

type GetOnlineFeaturesHttpResp struct {
	MetaData *GetOnlineFeaturesHttpRespMetaData `json:"metadata"`
	Results  []*GetOnlineFeaturesHttpRespResult `json:"results"`
}

type GetOnlineFeaturesHttpReq struct {
	FullFeatureNames bool           `json:"full_feature_names"`
	Entities         map[string]any `json:"entities"`
	FeatureService   string         `json:"feature_service,omitempty"`
	Features         []string       `json:"features,omitempty"`
}
