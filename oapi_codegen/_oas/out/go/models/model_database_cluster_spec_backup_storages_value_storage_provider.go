package models

// DatabaseClusterSpecBackupStoragesValueStorageProvider - BackupStorageProviderSpec represents set of settings to configure cloud provider.
type DatabaseClusterSpecBackupStoragesValueStorageProvider struct {

	Bucket string `json:"bucket,omitempty"`

	// A container name is a valid DNS name that conforms to the Azure naming rules.
	ContainerName string `json:"containerName,omitempty"`

	CredentialsSecret string `json:"credentialsSecret"`

	EndpointUrl string `json:"endpointUrl,omitempty"`

	Prefix string `json:"prefix,omitempty"`

	Region string `json:"region,omitempty"`

	// STANDARD, NEARLINE, COLDLINE, ARCHIVE for GCP Hot (Frequently accessed or modified data), Cool (Infrequently accessed or modified data), Archive (Rarely accessed or modified data) for Azure.
	StorageClass string `json:"storageClass,omitempty"`
}
