package models

// DatabaseClusterSpecBackupStoragesValuePodSecurityContextSeLinuxOptions - The SELinux context to be applied to all containers. If unspecified, the container runtime will allocate a random SELinux context for each container.  May also be set in SecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence for that container. Note that this field cannot be set when spec.os.name is windows.
type DatabaseClusterSpecBackupStoragesValuePodSecurityContextSeLinuxOptions struct {

	// Level is SELinux level label that applies to the container.
	Level string `json:"level,omitempty"`

	// Role is a SELinux role label that applies to the container.
	Role string `json:"role,omitempty"`

	// Type is a SELinux type label that applies to the container.
	Type string `json:"type,omitempty"`

	// User is a SELinux user label that applies to the container.
	User string `json:"user,omitempty"`
}
