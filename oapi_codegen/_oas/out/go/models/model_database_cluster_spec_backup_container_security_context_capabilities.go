package models

// DatabaseClusterSpecBackupContainerSecurityContextCapabilities - The capabilities to add/drop when running containers. Defaults to the default set of capabilities granted by the container runtime. Note that this field cannot be set when spec.os.name is windows.
type DatabaseClusterSpecBackupContainerSecurityContextCapabilities struct {

	// Added capabilities
	Add []string `json:"add,omitempty"`

	// Removed capabilities
	Drop []string `json:"drop,omitempty"`
}
