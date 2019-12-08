/*
 * RingCentral Engage Voice API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package engagevoice

type LeadsLoader struct {
	Description string `json:"description,omitempty"`
	// allows you to insert a lead to the top of the dialer cache for immediate dialing if you want a normal insert then do not add this parameter.
	DialPriority      string `json:"dialPriority,omitempty"`
	DuplicateHandling string `json:"duplicateHandling,omitempty"`
	ListState         string `json:"listState,omitempty"`
	TimeZoneOption    string `json:"timeZoneOption,omitempty"`
	UploadLeads       []Lead `json:"uploadLeads,omitempty"`
}
