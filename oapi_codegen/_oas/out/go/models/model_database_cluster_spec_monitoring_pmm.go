package models

// DatabaseClusterSpecMonitoringPmm - PMMSpec contains PMM settings.
type DatabaseClusterSpecMonitoringPmm struct {

	Image string `json:"image,omitempty"`

	Login string `json:"login,omitempty"`

	Password string `json:"password,omitempty"`

	PublicAddress string `json:"publicAddress,omitempty"`

	ServerHost string `json:"serverHost,omitempty"`

	ServerUser string `json:"serverUser,omitempty"`
}
