package models

// DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaimResources - resources represents the minimum resources the volume should have. If RecoverVolumeExpansionFailure feature is enabled users are allowed to specify resource requirements that are lower than previous value but must still be higher than capacity recorded in the status field of the claim. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#resources
type DatabaseClusterSpecBackupStoragesValueVolumeSpecPersistentVolumeClaimResources struct {

	// Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container.   This is an alpha field and requires enabling the DynamicResourceAllocation feature gate.   This field is immutable. It can only be set for containers.
	Claims []DatabaseClusterSpecBackupResourcesClaimsInner `json:"claims,omitempty"`

	// Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/
	Limits map[string]DatabaseClusterSpecBackupResourcesLimitsValue `json:"limits,omitempty"`

	// Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/
	Requests map[string]DatabaseClusterSpecBackupResourcesLimitsValue `json:"requests,omitempty"`
}
