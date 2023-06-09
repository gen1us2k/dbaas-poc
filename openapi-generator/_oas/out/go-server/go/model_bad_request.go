/*
 * Percona Everest schema
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi


import (
	"encoding/json"
)



// BadRequest - Bad request response
type BadRequest struct {

	Message string `json:"message,omitempty"`
}

// UnmarshalJSON sets *m to a copy of data while respecting defaults if specified.
func (m *BadRequest) UnmarshalJSON(data []byte) error {

	type Alias BadRequest // To avoid infinite recursion
    return json.Unmarshal(data, (*Alias)(m))
}

// AssertBadRequestRequired checks if the required fields are not zero-ed
func AssertBadRequestRequired(obj BadRequest) error {
	return nil
}

// AssertBadRequestConstraints checks if the values respects the defined constraints
func AssertBadRequestConstraints(obj BadRequest) error {
	return nil
}
