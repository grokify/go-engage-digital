/*
 * Engage Digital API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package engagedigital

type GetAllRolesResponse struct {
	Count   int32  `json:"count,omitempty"`
	Limit   int32  `json:"limit,omitempty"`
	Offset  int32  `json:"offset,omitempty"`
	Records []Role `json:"records,omitempty"`
}
