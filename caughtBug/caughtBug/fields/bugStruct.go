package fields

type BugInformation struct {
	UniqueId        string `json:"uniqueid"`
	BugDescription  string `json:"bugdescription"`
	BugTopic        string `json:"bugTopic"`
	ApplicationName string `json:"applicationName"`
	StillPresent    string `json:"stillPresent"`
	PostedBy        string `json:"postedBy"`
}

type BugFilter struct {
	UniqueId        *string `json:"uniqueid"`
	BugDescription  *string `json:"bugdescription"`
	BugTopic        *string `json:"bugTopic"`
	ApplicationName *string `json:"applicationName"`
	StillPresent    *string `json:"stillPresent"`
	PostedBy        *string `json:"postedBy"`
}
