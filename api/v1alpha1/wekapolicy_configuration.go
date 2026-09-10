package v1alpha1

// ConfigurationPayload defines the parameters for the configuration policy: operator-wide settings
type ConfigurationPayload struct {
	// AdhocPods configures the pods the operator creates to run one-off operations.
	AdhocPods *AdhocPodsSpec `json:"adhocPods,omitempty"`
}

// AdhocPodsSpec groups settings for ad-hoc operation pods
type AdhocPodsSpec struct {
	// Resources sets the requests and limits of the pod
	Resources *PodResourcesSpec `json:"resources,omitempty"`
}
