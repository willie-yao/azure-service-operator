// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20201101

// Deprecated version of PublicIPAddress_STATUS_PublicIPAddress_SubResourceEmbedded. Use v1beta20201101.PublicIPAddress_STATUS_PublicIPAddress_SubResourceEmbedded instead
type PublicIPAddress_STATUS_PublicIPAddress_SubResourceEmbedded_ARM struct {
	Etag             *string                                     `json:"etag,omitempty"`
	ExtendedLocation *ExtendedLocation_STATUS_ARM                `json:"extendedLocation,omitempty"`
	Id               *string                                     `json:"id,omitempty"`
	Location         *string                                     `json:"location,omitempty"`
	Name             *string                                     `json:"name,omitempty"`
	Properties       *PublicIPAddressPropertiesFormat_STATUS_ARM `json:"properties,omitempty"`
	Sku              *PublicIPAddressSku_STATUS_ARM              `json:"sku,omitempty"`
	Tags             map[string]string                           `json:"tags,omitempty"`
	Type             *string                                     `json:"type,omitempty"`
	Zones            []string                                    `json:"zones,omitempty"`
}

// Deprecated version of PublicIPAddressPropertiesFormat_STATUS. Use v1beta20201101.PublicIPAddressPropertiesFormat_STATUS instead
type PublicIPAddressPropertiesFormat_STATUS_ARM struct {
	DdosSettings             *DdosSettings_STATUS_ARM                                        `json:"ddosSettings,omitempty"`
	DnsSettings              *PublicIPAddressDnsSettings_STATUS_ARM                          `json:"dnsSettings,omitempty"`
	IdleTimeoutInMinutes     *int                                                            `json:"idleTimeoutInMinutes,omitempty"`
	IpAddress                *string                                                         `json:"ipAddress,omitempty"`
	IpConfiguration          *IPConfiguration_STATUS_PublicIPAddress_SubResourceEmbedded_ARM `json:"ipConfiguration,omitempty"`
	IpTags                   []IpTag_STATUS_ARM                                              `json:"ipTags,omitempty"`
	MigrationPhase           *PublicIPAddressPropertiesFormat_MigrationPhase_STATUS          `json:"migrationPhase,omitempty"`
	NatGateway               *NatGateway_STATUS_PublicIPAddress_SubResourceEmbedded_ARM      `json:"natGateway,omitempty"`
	ProvisioningState        *ProvisioningState_STATUS                                       `json:"provisioningState,omitempty"`
	PublicIPAddressVersion   *IPVersion_STATUS                                               `json:"publicIPAddressVersion,omitempty"`
	PublicIPAllocationMethod *IPAllocationMethod_STATUS                                      `json:"publicIPAllocationMethod,omitempty"`
	PublicIPPrefix           *SubResource_STATUS_ARM                                         `json:"publicIPPrefix,omitempty"`
	ResourceGuid             *string                                                         `json:"resourceGuid,omitempty"`
}

// Deprecated version of PublicIPAddressSku_STATUS. Use v1beta20201101.PublicIPAddressSku_STATUS instead
type PublicIPAddressSku_STATUS_ARM struct {
	Name *PublicIPAddressSku_Name_STATUS `json:"name,omitempty"`
	Tier *PublicIPAddressSku_Tier_STATUS `json:"tier,omitempty"`
}

// Deprecated version of DdosSettings_STATUS. Use v1beta20201101.DdosSettings_STATUS instead
type DdosSettings_STATUS_ARM struct {
	DdosCustomPolicy   *SubResource_STATUS_ARM                 `json:"ddosCustomPolicy,omitempty"`
	ProtectedIP        *bool                                   `json:"protectedIP,omitempty"`
	ProtectionCoverage *DdosSettings_ProtectionCoverage_STATUS `json:"protectionCoverage,omitempty"`
}

// Deprecated version of IPConfiguration_STATUS_PublicIPAddress_SubResourceEmbedded. Use v1beta20201101.IPConfiguration_STATUS_PublicIPAddress_SubResourceEmbedded instead
type IPConfiguration_STATUS_PublicIPAddress_SubResourceEmbedded_ARM struct {
	Id *string `json:"id,omitempty"`
}

// Deprecated version of IpTag_STATUS. Use v1beta20201101.IpTag_STATUS instead
type IpTag_STATUS_ARM struct {
	IpTagType *string `json:"ipTagType,omitempty"`
	Tag       *string `json:"tag,omitempty"`
}

// Deprecated version of NatGateway_STATUS_PublicIPAddress_SubResourceEmbedded. Use v1beta20201101.NatGateway_STATUS_PublicIPAddress_SubResourceEmbedded instead
type NatGateway_STATUS_PublicIPAddress_SubResourceEmbedded_ARM struct {
	Id *string `json:"id,omitempty"`
}

// Deprecated version of PublicIPAddressDnsSettings_STATUS. Use v1beta20201101.PublicIPAddressDnsSettings_STATUS instead
type PublicIPAddressDnsSettings_STATUS_ARM struct {
	DomainNameLabel *string `json:"domainNameLabel,omitempty"`
	Fqdn            *string `json:"fqdn,omitempty"`
	ReverseFqdn     *string `json:"reverseFqdn,omitempty"`
}

// Deprecated version of PublicIPAddressSku_Name_STATUS. Use v1beta20201101.PublicIPAddressSku_Name_STATUS instead
type PublicIPAddressSku_Name_STATUS string

const (
	PublicIPAddressSku_Name_STATUS_Basic    = PublicIPAddressSku_Name_STATUS("Basic")
	PublicIPAddressSku_Name_STATUS_Standard = PublicIPAddressSku_Name_STATUS("Standard")
)

// Deprecated version of PublicIPAddressSku_Tier_STATUS. Use v1beta20201101.PublicIPAddressSku_Tier_STATUS instead
type PublicIPAddressSku_Tier_STATUS string

const (
	PublicIPAddressSku_Tier_STATUS_Global   = PublicIPAddressSku_Tier_STATUS("Global")
	PublicIPAddressSku_Tier_STATUS_Regional = PublicIPAddressSku_Tier_STATUS("Regional")
)