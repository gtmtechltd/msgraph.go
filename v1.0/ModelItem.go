// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "time"

// ItemActionStat undocumented
type ItemActionStat struct {
	// Object is the base model of ItemActionStat
	Object
	// ActionCount undocumented
	ActionCount *int `json:"actionCount,omitempty"`
	// ActorCount undocumented
	ActorCount *int `json:"actorCount,omitempty"`
}

// ItemActivity undocumented
type ItemActivity struct {
	// Entity is the base model of ItemActivity
	Entity
	// Access undocumented
	Access *AccessAction `json:"access,omitempty"`
	// ActivityDateTime undocumented
	ActivityDateTime *time.Time `json:"activityDateTime,omitempty"`
	// Actor undocumented
	Actor *IdentitySet `json:"actor,omitempty"`
	// DriveItem undocumented
	DriveItem *DriveItem `json:"driveItem,omitempty"`
}

// ItemActivityStat undocumented
type ItemActivityStat struct {
	// Entity is the base model of ItemActivityStat
	Entity
	// StartDateTime undocumented
	StartDateTime *time.Time `json:"startDateTime,omitempty"`
	// EndDateTime undocumented
	EndDateTime *time.Time `json:"endDateTime,omitempty"`
	// Access undocumented
	Access *ItemActionStat `json:"access,omitempty"`
	// Create undocumented
	Create *ItemActionStat `json:"create,omitempty"`
	// Delete undocumented
	Delete *ItemActionStat `json:"delete,omitempty"`
	// Edit undocumented
	Edit *ItemActionStat `json:"edit,omitempty"`
	// Move undocumented
	Move *ItemActionStat `json:"move,omitempty"`
	// IsTrending undocumented
	IsTrending *bool `json:"isTrending,omitempty"`
	// IncompleteData undocumented
	IncompleteData *IncompleteData `json:"incompleteData,omitempty"`
	// Activities undocumented
	Activities []ItemActivity `json:"activities,omitempty"`
}

// ItemAnalytics undocumented
type ItemAnalytics struct {
	// Entity is the base model of ItemAnalytics
	Entity
	// ItemActivityStats undocumented
	ItemActivityStats []ItemActivityStat `json:"itemActivityStats,omitempty"`
	// AllTime undocumented
	AllTime *ItemActivityStat `json:"allTime,omitempty"`
	// LastSevenDays undocumented
	LastSevenDays *ItemActivityStat `json:"lastSevenDays,omitempty"`
}

// ItemAttachment undocumented
type ItemAttachment struct {
	// Attachment is the base model of ItemAttachment
	Attachment
	// Item undocumented
	Item *OutlookItem `json:"item,omitempty"`
}

// ItemBody undocumented
type ItemBody struct {
	// Object is the base model of ItemBody
	Object
	// ContentType undocumented
	ContentType *BodyType `json:"contentType,omitempty"`
	// Content undocumented
	Content *string `json:"content,omitempty"`
}

// ItemPreviewInfo undocumented
type ItemPreviewInfo struct {
	// Object is the base model of ItemPreviewInfo
	Object
	// GetURL undocumented
	GetURL *string `json:"getUrl,omitempty"`
	// PostParameters undocumented
	PostParameters *string `json:"postParameters,omitempty"`
	// PostURL undocumented
	PostURL *string `json:"postUrl,omitempty"`
}

// ItemReference undocumented
type ItemReference struct {
	// Object is the base model of ItemReference
	Object
	// DriveID undocumented
	DriveID *string `json:"driveId,omitempty"`
	// DriveType undocumented
	DriveType *string `json:"driveType,omitempty"`
	// ID undocumented
	ID *string `json:"id,omitempty"`
	// Name undocumented
	Name *string `json:"name,omitempty"`
	// Path undocumented
	Path *string `json:"path,omitempty"`
	// ShareID undocumented
	ShareID *string `json:"shareId,omitempty"`
	// SharepointIDs undocumented
	SharepointIDs *SharepointIDs `json:"sharepointIds,omitempty"`
	// SiteID undocumented
	SiteID *string `json:"siteId,omitempty"`
}
