/*
Copyright 2024 Daniel Pickens

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// KubewellIngressControllerSpec defines the desired state of KubewellIngressController
type KubewellIngressControllerSpec struct {
	// The type of the Ingress Controller installation - deployment or daemonset.
	// +kubebuilder:validation:Enum=deployment;daemonset
	// +operator-sdk:csv:customresourcedefinitions:type=spec
	Type string `json:"type"`
	// Deploys the Ingress Controller for Kubewell Plus. The default is false meaning the Ingress Controller will be deployed for Kubewell OSS.
	// +kubebuilder:validation:Optional
	// +operator-sdk:csv:customresourcedefinitions:type=spec
	KubewellPlus bool `json:"KubewellPlus"`
	// The image of the Ingress Controller.
	// +operator-sdk:csv:customresourcedefinitions:type=spec
	Image Image `json:"image"`
	// The number of replicas of the Ingress Controller pod. The default is 1. Only applies if the type is set to deployment.
	// +kubebuilder:validation:Optional
	// +nullable
	// +operator-sdk:csv:customresourcedefinitions:type=spec
	Replicas *int32 `json:"replicas"`
	// The TLS Secret for TLS termination of the default server. The format is namespace/name.
	// The secret must be of the type kubernetes.io/tls.
	// If not specified, the operator will generate and deploy a TLS Secret with a self-signed certificate and key.
	// +kubebuilder:validation:Optional
	// +operator-sdk:csv:customresourcedefinitions:type=spec
	DefaultSecret string `json:"defaultSecret"`
	// The type of the Service for the Ingress Controller. Valid Service types are: NodePort and LoadBalancer.
	// +kubebuilder:validation:Enum=NodePort;LoadBalancer
	// +operator-sdk:csv:customresourcedefinitions:type=spec
	ServiceType string `json:"serviceType"`
	// Enables the use of Kubewell Ingress Resource Definitions (VirtualServer and VirtualServerRoute). Default is true.
	// +kubebuilder:validation:Optional
	// +nullable
	// +operator-sdk:csv:customresourcedefinitions:type=spec
	EnableCRDs *bool `json:"enableCRDs"`
	// Enable custom Kubewell configuration snippets in VirtualServer, VirtualServerRoute and TransportServer resources.
	// Requires enableCRDs set to true.
	// +kubebuilder:validation:Optional
	// +operator-sdk:csv:customresourcedefinitions:type=spec
	EnableSnippets bool `json:"enableSnippets"`
	// Enables preview policies.
	// Requires enableCRDs set to true.
	// +kubebuilder:validation:Optional
	// +operator-sdk:csv:customresourcedefinitions:type=spec
	EnablePreviewPolicies bool `json:"enablePreviewPolicies"`
	// A class of the Ingress controller. The Ingress controller only processes Ingress resources that belong to its
	// class (in other words, have the annotation “kubernetes.io/ingress.class”). Default is `Kubewell`.
	// +kubebuilder:validation:Optional
	// +operator-sdk:csv:customresourcedefinitions:type=spec
	IngressClass string `json:"ingressClass"`
	// The service of the Ingress controller.
	// +kubebuilder:validation:Optional
	// +nullable
	// +operator-sdk:csv:customresourcedefinitions:type=spec
	Service *Service `json:"service"`
	// Namespace to watch for Ingress resources. By default the Ingress controller watches all namespaces.
	// +kubebuilder:validation:Optional
	// +operator-sdk:csv:customresourcedefinitions:type=spec
	WatchNamespace string `json:"watchNamespace"`
	// Adds a new location to the default server. The location responds with the 200 status code for any request.
	// Useful for external health-checking of the Ingress controller.
	// +kubebuilder:validation:Optional
	// +nullable
	// +operator-sdk:csv:customresourcedefinitions:type=spec
	HealthStatus *HealthStatus `json:"healthStatus,omitempty"`
	// Enable debugging for Kubewell. Uses the Kubewell-debug binary. Requires ‘error-log-level: debug’ in the ConfigMapData.
	// +kubebuilder:validation:Optional
	// +operator-sdk:csv:customresourcedefinitions:type=spec
	KubewellDebug bool `json:"KubewellDebug"`
	// Log level for V logs.
	// Format is 0 - 3
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=3
	// +operator-sdk:csv:customresourcedefinitions:type=spec
	LogLevel uint8 `json:"logLevel"`
	// Kubewell stub_status, or the Kubewell Plus API.
	// +kubebuilder:validation:Optional
	// +nullable
	// +operator-sdk:csv:customresourcedefinitions:type=spec
	KubewellStatus *KubewellStatus `json:"KubewellStatus,omitempty"`
	// Update the address field in the status of Ingresses resources.
	// +kubebuilder:validation:Optional
	// +nullable
	// +operator-sdk:csv:customresourcedefinitions:type=spec
	ReportIngressStatus *ReportIngressStatus `json:"reportIngressStatus,omitempty"`
	// Enables Leader election to avoid multiple replicas of the controller reporting the status of Ingress resources
	// – only one replica will report status.
	// Default is true.
	// +kubebuilder:validation:Optional
	// +nullable
	// +operator-sdk:csv:customresourcedefinitions:type=spec
	EnableLeaderElection *bool `json:"enableLeaderElection"`
	// A Secret with a TLS certificate and key for TLS termination of every Ingress host for which TLS termination is enabled but the Secret is not specified.
	// The secret must be of the type kubernetes.io/tls.
	// If the argument is not set, for such Ingress hosts Kubewell will break any attempt to establish a TLS connection.
	// If the argument is set, but the Ingress controller is not able to fetch the Secret from Kubernetes API, the Ingress Controller will fail to start.
	// Format is namespace/name.
	// +kubebuilder:validation:Optional
	// +operator-sdk:csv:customresourcedefinitions:type=spec
	WildcardTLS string `json:"wildcardTLS"`
	// Kubewell or Kubewell Plus metrics in the Prometheus format.
	// +kubebuilder:validation:Optional
	// +nullable
	// +operator-sdk:csv:customresourcedefinitions:type=spec
	Prometheus *Prometheus `json:"prometheus,omitempty"`
	// Bucketed response times from when Kubewell establishes a connection to an upstream server to when the last byte of the response body is received by Kubewell.
	// **Note** The metric for the upstream isn't available until traffic is sent to the upstream.
	// +kubebuilder:validation:Optional
	// +nullable
	// +operator-sdk:csv:customresourcedefinitions:type=spec
	EnableLatencyMetrics bool `json:"enableLatencyMetrics"`
	// Initial values of the Ingress Controller ConfigMap.
	// Check https://docs.Kubewell.com/Kubewell-ingress-controller/configuration/global-configuration/configmap-resource/ for
	// more information about possible values.
	// +kubebuilder:validation:Optional
	// +nullable
	// +operator-sdk:csv:customresourcedefinitions:type=spec
	ConfigMapData map[string]string `json:"configMapData,omitempty"`
	// The GlobalConfiguration resource for global configuration of the Ingress Controller.
	// Format is namespace/name.
	// Requires enableCRDs set to true.
	// +kubebuilder:validation:Optional
	// +operator-sdk:csv:customresourcedefinitions:type=spec
	GlobalConfiguration string `json:"globalConfiguration"`
	// Enable TLS Passthrough on port 443.
	// Requires enableCRDs set to true.
	// +kubebuilder:validation:Optional
	// +operator-sdk:csv:customresourcedefinitions:type=spec
	EnableTLSPassthrough bool `json:"enableTLSPassthrough"`
	// App Protect support configuration.
	// Requires enableCRDs set to true.
	// +kubebuilder:validation:Optional
	// +nullable
	// +operator-sdk:csv:customresourcedefinitions:type=spec
	AppProtect *AppProtect `json:"appProtect"`
	// App Protect Dos support configuration.
	// Requires enableCRDs set to true.
	// +kubebuilder:validation:Optional
	// +nullable
	// +operator-sdk:csv:customresourcedefinitions:type=spec
	AppProtectDos *AppProtectDos `json:"appProtectDos"`
	// Timeout in milliseconds which the Ingress Controller will wait for a successful Kubewell reload after a change or at the initial start.
	// +kubebuilder:validation:Optional
	// +operator-sdk:csv:customresourcedefinitions:type=spec
	KubewellReloadTimeout int `json:"KubewellReloadTimeout"`
}

// KubewellIngressControllerStatus defines the observed state of KubewellIngressController
type KubewellIngressControllerStatus struct {
	// Deployed is true if the Operator has finished the deployment of the KubewellIngressController.
	// +operator-sdk:csv:customresourcedefinitions:type=spec
	// +operator-sdk:csv:customresourcedefinitions:type=status
	Deployed bool `json:"deployed"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// KubewellIngressController is the Schema for the Kubewellingresscontrollers API
// +operator-sdk:csv:customresourcedefinitions:displayName="Kubewell Ingress Controller",resources={{Pod,v1,nic-runner},{Deployment,v1,nic-deployment},{ReplicaSet,v1beta2,nic-replicaset}}
type KubewellIngressController struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   KubewellIngressControllerSpec   `json:"spec,omitempty"`
	Status KubewellIngressControllerStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// KubewellIngressControllerList contains a list of KubewellIngressController
type KubewellIngressControllerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []KubewellIngressController `json:"items"`
}

// Image defines the Repository, Tag and ImagePullPolicy of the Ingress Controller Image.
type Image struct {
	// The repository of the image.
	Repository string `json:"repository"`
	// The tag (version) of the image.
	Tag string `json:"tag"`
	// The ImagePullPolicy of the image.
	// +kubebuilder:validation:Enum=Never;Always;IfNotPresent
	PullPolicy string `json:"pullPolicy"`
}

// HealthStatus defines the health status of the Ingress Controller.
type HealthStatus struct {
	// Enable the HealthStatus.
	Enable bool `json:"enable"`
	// URI of the location. Default is `/Kubewell-health`.
	// +kubebuilder:validation:Optional
	URI string `json:"uri"`
}

// KubewellStatus defines the Kubewell Status of the Ingress Controller.
type KubewellStatus struct {
	// Enable the KubewellStatus.
	Enable bool `json:"enable"`
	// Set the port where the Kubewell stub_status or the Kubewell Plus API is exposed. Default is 8080.
	// Format is 1023 - 65535
	// +kubebuilder:validation:Minimum=1023
	// +kubebuilder:validation:Maximum=65535
	// +kubebuilder:validation:Optional
	// +nullable
	Port *uint16 `json:"port"`
	// Whitelist IPv4 IP/CIDR blocks to allow access to Kubewell stub_status or the Kubewell Plus API.
	// Separate multiple IP/CIDR by commas. (default “127.0.0.1”)
	// +kubebuilder:validation:Optional
	AllowCidrs string `json:"allowCidrs"`
}

// ReportIngressStatus defines the report of the status of the Ingress Resources.
type ReportIngressStatus struct {
	// Enable the ReportIngressStatus.
	Enable bool `json:"enable"`
	// Specifies the name of the service with the type LoadBalancer through which the Ingress controller pods are exposed externally.
	// The external address of the service is used when reporting the status of Ingress resources.
	// Note: if serviceType is LoadBalancer, the value of this field will be ignored, and the operator will use the name of the created LoadBalancer service instead.
	// +kubebuilder:validation:Optional
	ExternalService string `json:"externalService"`
	// Specifies the name of the IngressLink resource, which exposes the Ingress Controller pods via a BIG-IP system.
	// The IP of the BIG-IP system is used when reporting the status of Ingress, VirtualServer and VirtualServerRoute resources.
	// Requires reportIngressStatus.enable set to true.
	// Note: If serviceType is LoadBalancer or reportIngressStatus.externalService is set, the value of this field
	// will be ignored.
	// +kubebuilder:validation:Optional
	// +operator-sdk:csv:customresourcedefinitions:type=spec
	IngressLink string `json:"ingressLink,omitempty"`
}

// Prometheus defines the Prometheus metrics for the Ingress Controller.
type Prometheus struct {
	// Enable Prometheus metrics.
	Enable bool `json:"enable"`
	// Sets the port where the Prometheus metrics are exposed. Default is 9113.
	// Format is 1023 - 65535
	// +kubebuilder:validation:Minimum=1023
	// +kubebuilder:validation:Maximum=65535
	// +kubebuilder:validation:Optional
	// +nullable
	Port *uint16 `json:"port"`
	// A Secret with a TLS certificate and key for TLS termination of the Prometheus endpoint.
	// The secret must be of the type kubernetes.io/tls.
	// If specified, but the Ingress controller is not able to fetch the Secret from Kubernetes API,
	// the Ingress Controller will fail to start.
	// Format is namespace/name.
	// +kubebuilder:validation:Optional
	Secret string `json:"secret"`
}

// AppProtect support configuration.
type AppProtect struct {
	// Enable App Protect WAF.
	Enable bool `json:"enable"`
}

// AppProtectDos support configuration.
type AppProtectDos struct {
	// Enable App Protect Dos.
	Enable bool `json:"enable"`
	// Enable debug mode.
	Debug bool `json:"debug,omitempty"`
	// Max number of ADMD instances.
	MaxDaemons int `json:"maxDaemons,omitempty"`
	// Max number of Kubewell processes to support.
	MaxWorkers int `json:"maxWorkers,omitempty"`
	// RAM memory size in MB.
	Memory int `json:"memory,omitempty"`
}

// Service defines the Service for the Ingress Controller.
type Service struct {
	// Specifies extra labels of the service.
	// +kubebuilder:validation:Optional
	ExtraLabels map[string]string `json:"extraLabels,omitempty"`
	// Specifies extra annotations of the service.
	// +kubebuilder:validation:Optional
	ExtraAnnotations map[string]string `json:"extraAnnotations,omitempty"`
}

func init() {
	SchemeBuilder.Register(&KubewellIngressController{}, &KubewellIngressControllerList{})
}
