/*
Copyright 2024.

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

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ExporterConfigSpec defines the desired state of ExporterConfig
type ExporterConfigSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of ExporterConfig. Edit exporterconfig_types.go to remove/update
	IngressClassName *string `json:"ingressClassName,omitempty" protobuf:"bytes,1,opt,name=ingressClassName"`
	Cluster          `json:"cluster,omitempty" protobuf:"bytes,2,opt,name=cluster"`
	Traefik          `json:"traefik,omitempty" protobuf:"bytes,3,opt,name=traefik"`
	PrintDebug       `json:"printDebug,omitempty" protobuf:"bytes,4,opt,name=printDebug"`
}
type PrintDebug struct {
	Enable bool `json:"enable,omitempty" protobuf:"varint,1,opt,name=enable"`
	HTTPOk bool `json:"httpok,omitempty" protobuf:"varint,2,opt,name=httpok"`
}
type Cluster struct {
	Ingress        `json:"ingress,omitempty" protobuf:"bytes,1,opt,name=ingress"`
	ServiceName    string `json:"serviceName,omitempty" protobuf:"bytes,2,opt,name=serviceName"`
	RootCAFilename string `json:"rootCAFilename,omitempty" protobuf:"bytes,3,opt,name=rootCAFilename"`
}
type Ingress struct {
	Address string `json:"address,omitempty" protobuf:"bytes,1,opt,name=address"`
	HTTP    Port   `json:"http,omitempty" protobuf:"bytes,2,opt,name=http"`
	HTTPS   Port   `json:"https,omitempty" protobuf:"bytes,3,opt,name=https"`
}
type Traefik struct {
	EntrypointName `json:"entrypoint,omitempty" protobuf:"bytes,1,opt,name=entrypoint"`
}
type EntrypointName struct {
	HTTP  string `json:"http,omitempty" protobuf:"bytes,1,opt,name=http"`
	HTTPS string `json:"https,omitempty" protobuf:"bytes,2,opt,name=https"`
}
type Port struct { // Should be fetchable from service
	Port     string `json:"port" protobuf:"bytes,1,opt,name=port"`
	Protocol string `json:"protocol" protobuf:"bytes,2,opt,name=protocol"`
}

// ExporterConfigStatus defines the observed state of ExporterConfig
type ExporterConfigStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	LoadBalancer Ingress `json:"loadbalancer,omitempty" protobuf:"bytes,1,opt,name=loadbalancer"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
//+kubebuilder:resource:path=exporterconfigs,scope=Cluster

// ExporterConfig is the Schema for the exporterconfigs API
type ExporterConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ExporterConfigSpec   `json:"spec,omitempty"`
	Status ExporterConfigStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ExporterConfigList contains a list of ExporterConfig
type ExporterConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ExporterConfig `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ExporterConfig{}, &ExporterConfigList{})
}
