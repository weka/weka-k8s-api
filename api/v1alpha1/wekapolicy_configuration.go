package v1alpha1

// ConfigurationPayload defines the parameters for the configuration policy: operator-wide settings
type ConfigurationPayload struct {
	// Csi configures the embedded CSI deployment.
	Csi *CsiSpec `json:"csi,omitempty"`
}

// CsiSpec groups settings of the embedded CSI deployment
type CsiSpec struct {
	// MetricsEnabled controls the Prometheus metrics endpoints of the CSI controller, the CSI node
	// plugin and the controller sidecars. Enabled by default.
	MetricsEnabled *bool `json:"metricsEnabled,omitempty"`
	// SkipGarbageCollection stops the CSI controller from reclaiming deleted volume directories in
	// the background. False by default.
	SkipGarbageCollection *bool `json:"skipGarbageCollection,omitempty"`
	// FsGroupPolicy sets fsGroupPolicy on the CSIDriver object, which decides whether Kubernetes
	// reapplies a pod's fsGroup to the volume. Default value is File.
	// +kubebuilder:validation:Enum=File;None;ReadWriteOnceWithFSType
	FsGroupPolicy *string `json:"fsGroupPolicy,omitempty"`
}
