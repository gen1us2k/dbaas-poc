package models

// DatabaseClusterSpecMonitoring - Monitoring contains a monitoring settings.
type DatabaseClusterSpecMonitoring struct {

	ContainerSecurityContext DatabaseClusterSpecBackupContainerSecurityContext `json:"containerSecurityContext,omitempty"`

	// PullPolicy describes a policy for if/when to pull a container image
	ImagePullPolicy string `json:"imagePullPolicy,omitempty"`

	Pmm DatabaseClusterSpecMonitoringPmm `json:"pmm,omitempty"`

	Resources DatabaseClusterSpecBackupResources `json:"resources,omitempty"`

	RuntimeClassName string `json:"runtimeClassName,omitempty"`
}
