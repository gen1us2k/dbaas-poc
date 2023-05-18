package models

// DatabaseClusterSpecBackupScheduleInner - BackupSchedule represents set of settings to configure backup schedule.
type DatabaseClusterSpecBackupScheduleInner struct {

	CompressionLevel int32 `json:"compressionLevel,omitempty"`

	CompressionType string `json:"compressionType,omitempty"`

	Enabled bool `json:"enabled,omitempty"`

	Keep int32 `json:"keep,omitempty"`

	Name string `json:"name,omitempty"`

	Schedule string `json:"schedule,omitempty"`

	StorageName string `json:"storageName,omitempty"`
}
