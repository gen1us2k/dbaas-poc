package models

// DatabaseClusterSpecBackupResourcesClaimsInner - ResourceClaim references one entry in PodSpec.ResourceClaims.
type DatabaseClusterSpecBackupResourcesClaimsInner struct {

	// Name must match the name of one entry in pod.spec.resourceClaims of the Pod where this field is used. It makes that resource available inside a container.
	Name string `json:"name"`
}
