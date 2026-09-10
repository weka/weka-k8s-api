package v1alpha1

// OperatorDefaultsPayload defines the parameters for the operator-defaults policy: operator-wide defaults
type OperatorDefaultsPayload struct {
	AdhocOpResources *PodResourcesSpec `json:"adhocOpResources,omitempty"`
}
