// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// DeliveryOptimizationBandwidthPercentage undocumented
type DeliveryOptimizationBandwidthPercentage struct {
	DeliveryOptimizationBandwidth
	// MaximumBackgroundBandwidthPercentage Specifies the maximum background download bandwidth that Delivery Optimization uses across all concurrent download activities as a percentage of available download bandwidth (0-100). Valid values 0 to 100
	MaximumBackgroundBandwidthPercentage *int `json:"maximumBackgroundBandwidthPercentage,omitempty"`
	// MaximumForegroundBandwidthPercentage Specifies the maximum foreground download bandwidth that Delivery Optimization uses across all concurrent download activities as a percentage of available download bandwidth (0-100). Valid values 0 to 100
	MaximumForegroundBandwidthPercentage *int `json:"maximumForegroundBandwidthPercentage,omitempty"`
}