/*
Percona Everest schema

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// DatabaseClusterSpecBackupStoragesValueVolumeSpecEmptyDirSizeLimit sizeLimit is the total amount of local storage required for this EmptyDir volume. The size limit is also applicable for memory medium. The maximum usage on memory medium EmptyDir would be the minimum value between the SizeLimit specified here and the sum of memory limits of all containers in a pod. The default is nil which means that the limit is undefined. More info: http://kubernetes.io/docs/user-guide/volumes#emptydir
type DatabaseClusterSpecBackupStoragesValueVolumeSpecEmptyDirSizeLimit struct {
	int32 *int32
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *DatabaseClusterSpecBackupStoragesValueVolumeSpecEmptyDirSizeLimit) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into int32
	err = json.Unmarshal(data, &dst.int32);
	if err == nil {
		jsonint32, _ := json.Marshal(dst.int32)
		if string(jsonint32) == "{}" { // empty struct
			dst.int32 = nil
		} else {
			return nil // data stored in dst.int32, return on the first match
		}
	} else {
		dst.int32 = nil
	}

	// try to unmarshal JSON data into string
	err = json.Unmarshal(data, &dst.string);
	if err == nil {
		jsonstring, _ := json.Marshal(dst.string)
		if string(jsonstring) == "{}" { // empty struct
			dst.string = nil
		} else {
			return nil // data stored in dst.string, return on the first match
		}
	} else {
		dst.string = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(DatabaseClusterSpecBackupStoragesValueVolumeSpecEmptyDirSizeLimit)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *DatabaseClusterSpecBackupStoragesValueVolumeSpecEmptyDirSizeLimit) MarshalJSON() ([]byte, error) {
	if src.int32 != nil {
		return json.Marshal(&src.int32)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableDatabaseClusterSpecBackupStoragesValueVolumeSpecEmptyDirSizeLimit struct {
	value *DatabaseClusterSpecBackupStoragesValueVolumeSpecEmptyDirSizeLimit
	isSet bool
}

func (v NullableDatabaseClusterSpecBackupStoragesValueVolumeSpecEmptyDirSizeLimit) Get() *DatabaseClusterSpecBackupStoragesValueVolumeSpecEmptyDirSizeLimit {
	return v.value
}

func (v *NullableDatabaseClusterSpecBackupStoragesValueVolumeSpecEmptyDirSizeLimit) Set(val *DatabaseClusterSpecBackupStoragesValueVolumeSpecEmptyDirSizeLimit) {
	v.value = val
	v.isSet = true
}

func (v NullableDatabaseClusterSpecBackupStoragesValueVolumeSpecEmptyDirSizeLimit) IsSet() bool {
	return v.isSet
}

func (v *NullableDatabaseClusterSpecBackupStoragesValueVolumeSpecEmptyDirSizeLimit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDatabaseClusterSpecBackupStoragesValueVolumeSpecEmptyDirSizeLimit(val *DatabaseClusterSpecBackupStoragesValueVolumeSpecEmptyDirSizeLimit) *NullableDatabaseClusterSpecBackupStoragesValueVolumeSpecEmptyDirSizeLimit {
	return &NullableDatabaseClusterSpecBackupStoragesValueVolumeSpecEmptyDirSizeLimit{value: val, isSet: true}
}

func (v NullableDatabaseClusterSpecBackupStoragesValueVolumeSpecEmptyDirSizeLimit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDatabaseClusterSpecBackupStoragesValueVolumeSpecEmptyDirSizeLimit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


