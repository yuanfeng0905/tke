/*
 * Harbor API
 *
 * These APIs provide services for manipulating Harbor project.
 *
 * API version: 2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package client

type HarborExecution struct {
	// The ID of execution
	Id int32 `json:"id,omitempty"`
	// The vendor type of execution
	VendorType string `json:"vendor_type,omitempty"`
	// The vendor id of execution
	VendorId int32 `json:"vendor_id,omitempty"`
	// The status of execution
	Status string `json:"status,omitempty"`
	// The status message of execution
	StatusMessage string `json:"status_message,omitempty"`
	Metrics *HarborMetrics `json:"metrics,omitempty"`
	// The trigger of execution
	Trigger string `json:"trigger,omitempty"`
	ExtraAttrs *map[string]interface{} `json:"extra_attrs,omitempty"`
	// The start time of execution
	StartTime string `json:"start_time,omitempty"`
	// The end time of execution
	EndTime string `json:"end_time,omitempty"`
}