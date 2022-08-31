// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20200601

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

// Deprecated version of Domains_Topics_Spec. Use v1beta20200601.Domains_Topics_Spec instead
type Domains_Topics_SpecARM struct {
	Location *string           `json:"location,omitempty"`
	Name     string            `json:"name,omitempty"`
	Tags     map[string]string `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &Domains_Topics_SpecARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2020-06-01"
func (topics Domains_Topics_SpecARM) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
func (topics *Domains_Topics_SpecARM) GetName() string {
	return topics.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.EventGrid/domains/topics"
func (topics *Domains_Topics_SpecARM) GetType() string {
	return "Microsoft.EventGrid/domains/topics"
}