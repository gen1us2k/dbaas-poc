package models

// DatabaseClusterSpecLoadBalancer - LoadBalancer contains a load balancer settings. For PXC it's haproxy or proxysql. For PSMDB it's mongos.
type DatabaseClusterSpecLoadBalancer struct {

	Annotations map[string]string `json:"annotations,omitempty"`

	Configuration string `json:"configuration,omitempty"`

	// Service Type string describes ingress methods for a service
	ExposeType string `json:"exposeType,omitempty"`

	Image string `json:"image,omitempty"`

	LoadBalancerSourceRanges []string `json:"loadBalancerSourceRanges,omitempty"`

	Resources DatabaseClusterSpecBackupResources `json:"resources,omitempty"`

	Size int32 `json:"size,omitempty"`

	// ServiceExternalTrafficPolicyType describes how nodes distribute service traffic they receive on one of the Service's \"externally-facing\" addresses (NodePorts, ExternalIPs, and LoadBalancer IPs).
	TrafficPolicy string `json:"trafficPolicy,omitempty"`

	// LoadBalancerType contains supported loadbalancers. It can be proxysql or haproxy for PXC clusters, mongos for PSMDB clusters or pgbouncer for Postgresql clusters.
	Type string `json:"type,omitempty"`
}
