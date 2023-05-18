package models

// KubernetesCluster - kubernetes object
type KubernetesCluster struct {

	Name string `json:"name,omitempty"`

	Kubeconfig string `json:"kubeconfig,omitempty"`
}
