// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "time"

// ScheduleEntity undocumented
type ScheduleEntity struct {
	// StartDateTime undocumented
	StartDateTime *time.Time `json:"startDateTime,omitempty"`
	// EndDateTime undocumented
	EndDateTime *time.Time `json:"endDateTime,omitempty"`
	// Theme undocumented
	Theme *ScheduleEntityTheme `json:"theme,omitempty"`
}