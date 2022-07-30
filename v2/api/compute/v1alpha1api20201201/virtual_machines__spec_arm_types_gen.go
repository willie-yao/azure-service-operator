// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20201201

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

// Deprecated version of VirtualMachines_Spec. Use v1beta20201201.VirtualMachines_Spec instead
type VirtualMachines_SpecARM struct {
	ExtendedLocation *ExtendedLocationARM                `json:"extendedLocation,omitempty"`
	Identity         *VirtualMachineIdentityARM          `json:"identity,omitempty"`
	Location         *string                             `json:"location,omitempty"`
	Name             string                              `json:"name,omitempty"`
	Plan             *PlanARM                            `json:"plan,omitempty"`
	Properties       *VirtualMachines_Spec_PropertiesARM `json:"properties,omitempty"`
	Tags             map[string]string                   `json:"tags,omitempty"`
	Zones            []string                            `json:"zones,omitempty"`
}

var _ genruntime.ARMResourceSpec = &VirtualMachines_SpecARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2020-12-01"
func (machines VirtualMachines_SpecARM) GetAPIVersion() string {
	return string(APIVersionValue)
}

// GetName returns the Name of the resource
func (machines *VirtualMachines_SpecARM) GetName() string {
	return machines.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Compute/virtualMachines"
func (machines *VirtualMachines_SpecARM) GetType() string {
	return "Microsoft.Compute/virtualMachines"
}

// Deprecated version of VirtualMachineIdentity. Use v1beta20201201.VirtualMachineIdentity instead
type VirtualMachineIdentityARM struct {
	Type *VirtualMachineIdentityType `json:"type,omitempty"`
}

// Deprecated version of VirtualMachines_Spec_Properties. Use v1beta20201201.VirtualMachines_Spec_Properties instead
type VirtualMachines_Spec_PropertiesARM struct {
	AdditionalCapabilities  *AdditionalCapabilitiesARM                         `json:"additionalCapabilities,omitempty"`
	AvailabilitySet         *SubResourceARM                                    `json:"availabilitySet,omitempty"`
	BillingProfile          *BillingProfileARM                                 `json:"billingProfile,omitempty"`
	DiagnosticsProfile      *DiagnosticsProfileARM                             `json:"diagnosticsProfile,omitempty"`
	EvictionPolicy          *VirtualMachinesSpecPropertiesEvictionPolicy       `json:"evictionPolicy,omitempty"`
	ExtensionsTimeBudget    *string                                            `json:"extensionsTimeBudget,omitempty"`
	HardwareProfile         *HardwareProfileARM                                `json:"hardwareProfile,omitempty"`
	Host                    *SubResourceARM                                    `json:"host,omitempty"`
	HostGroup               *SubResourceARM                                    `json:"hostGroup,omitempty"`
	LicenseType             *string                                            `json:"licenseType,omitempty"`
	NetworkProfile          *VirtualMachines_Spec_Properties_NetworkProfileARM `json:"networkProfile,omitempty"`
	OsProfile               *VirtualMachines_Spec_Properties_OsProfileARM      `json:"osProfile,omitempty"`
	PlatformFaultDomain     *int                                               `json:"platformFaultDomain,omitempty"`
	Priority                *VirtualMachinesSpecPropertiesPriority             `json:"priority,omitempty"`
	ProximityPlacementGroup *SubResourceARM                                    `json:"proximityPlacementGroup,omitempty"`
	SecurityProfile         *SecurityProfileARM                                `json:"securityProfile,omitempty"`
	StorageProfile          *StorageProfileARM                                 `json:"storageProfile,omitempty"`
	VirtualMachineScaleSet  *SubResourceARM                                    `json:"virtualMachineScaleSet,omitempty"`
}

// Deprecated version of BillingProfile. Use v1beta20201201.BillingProfile instead
type BillingProfileARM struct {
	MaxPrice *float64 `json:"maxPrice,omitempty"`
}

// Deprecated version of DiagnosticsProfile. Use v1beta20201201.DiagnosticsProfile instead
type DiagnosticsProfileARM struct {
	BootDiagnostics *BootDiagnosticsARM `json:"bootDiagnostics,omitempty"`
}

// Deprecated version of HardwareProfile. Use v1beta20201201.HardwareProfile instead
type HardwareProfileARM struct {
	VmSize *HardwareProfileVmSize `json:"vmSize,omitempty"`
}

// Deprecated version of SecurityProfile. Use v1beta20201201.SecurityProfile instead
type SecurityProfileARM struct {
	EncryptionAtHost *bool                        `json:"encryptionAtHost,omitempty"`
	SecurityType     *SecurityProfileSecurityType `json:"securityType,omitempty"`
	UefiSettings     *UefiSettingsARM             `json:"uefiSettings,omitempty"`
}

// Deprecated version of StorageProfile. Use v1beta20201201.StorageProfile instead
type StorageProfileARM struct {
	DataDisks      []DataDiskARM      `json:"dataDisks,omitempty"`
	ImageReference *ImageReferenceARM `json:"imageReference,omitempty"`
	OsDisk         *OSDiskARM         `json:"osDisk,omitempty"`
}

// Deprecated version of VirtualMachineIdentityType. Use v1beta20201201.VirtualMachineIdentityType instead
// +kubebuilder:validation:Enum={"None","SystemAssigned","SystemAssigned, UserAssigned","UserAssigned"}
type VirtualMachineIdentityType string

const (
	VirtualMachineIdentityTypeNone                       = VirtualMachineIdentityType("None")
	VirtualMachineIdentityTypeSystemAssigned             = VirtualMachineIdentityType("SystemAssigned")
	VirtualMachineIdentityTypeSystemAssignedUserAssigned = VirtualMachineIdentityType("SystemAssigned, UserAssigned")
	VirtualMachineIdentityTypeUserAssigned               = VirtualMachineIdentityType("UserAssigned")
)

// Deprecated version of VirtualMachines_Spec_Properties_NetworkProfile. Use v1beta20201201.VirtualMachines_Spec_Properties_NetworkProfile instead
type VirtualMachines_Spec_Properties_NetworkProfileARM struct {
	NetworkInterfaces []VirtualMachines_Spec_Properties_NetworkProfile_NetworkInterfacesARM `json:"networkInterfaces,omitempty"`
}

// Deprecated version of VirtualMachines_Spec_Properties_OsProfile. Use v1beta20201201.VirtualMachines_Spec_Properties_OsProfile instead
type VirtualMachines_Spec_Properties_OsProfileARM struct {
	AdminPassword               *string                  `json:"adminPassword,omitempty"`
	AdminUsername               *string                  `json:"adminUsername,omitempty"`
	AllowExtensionOperations    *bool                    `json:"allowExtensionOperations,omitempty"`
	ComputerName                *string                  `json:"computerName,omitempty"`
	CustomData                  *string                  `json:"customData,omitempty"`
	LinuxConfiguration          *LinuxConfigurationARM   `json:"linuxConfiguration,omitempty"`
	RequireGuestProvisionSignal *bool                    `json:"requireGuestProvisionSignal,omitempty"`
	Secrets                     []VaultSecretGroupARM    `json:"secrets,omitempty"`
	WindowsConfiguration        *WindowsConfigurationARM `json:"windowsConfiguration,omitempty"`
}

// Deprecated version of BootDiagnostics. Use v1beta20201201.BootDiagnostics instead
type BootDiagnosticsARM struct {
	Enabled    *bool   `json:"enabled,omitempty"`
	StorageUri *string `json:"storageUri,omitempty"`
}

// Deprecated version of DataDisk. Use v1beta20201201.DataDisk instead
type DataDiskARM struct {
	Caching                 *DataDiskCaching          `json:"caching,omitempty"`
	CreateOption            *DataDiskCreateOption     `json:"createOption,omitempty"`
	DetachOption            *DataDiskDetachOption     `json:"detachOption,omitempty"`
	DiskSizeGB              *int                      `json:"diskSizeGB,omitempty"`
	Image                   *VirtualHardDiskARM       `json:"image,omitempty"`
	Lun                     *int                      `json:"lun,omitempty"`
	ManagedDisk             *ManagedDiskParametersARM `json:"managedDisk,omitempty"`
	Name                    *string                   `json:"name,omitempty"`
	ToBeDetached            *bool                     `json:"toBeDetached,omitempty"`
	Vhd                     *VirtualHardDiskARM       `json:"vhd,omitempty"`
	WriteAcceleratorEnabled *bool                     `json:"writeAcceleratorEnabled,omitempty"`
}

// Deprecated version of ImageReference. Use v1beta20201201.ImageReference instead
type ImageReferenceARM struct {
	Id        *string `json:"id,omitempty"`
	Offer     *string `json:"offer,omitempty"`
	Publisher *string `json:"publisher,omitempty"`
	Sku       *string `json:"sku,omitempty"`
	Version   *string `json:"version,omitempty"`
}

// Deprecated version of LinuxConfiguration. Use v1beta20201201.LinuxConfiguration instead
type LinuxConfigurationARM struct {
	DisablePasswordAuthentication *bool                  `json:"disablePasswordAuthentication,omitempty"`
	PatchSettings                 *LinuxPatchSettingsARM `json:"patchSettings,omitempty"`
	ProvisionVMAgent              *bool                  `json:"provisionVMAgent,omitempty"`
	Ssh                           *SshConfigurationARM   `json:"ssh,omitempty"`
}

// Deprecated version of OSDisk. Use v1beta20201201.OSDisk instead
type OSDiskARM struct {
	Caching                 *OSDiskCaching             `json:"caching,omitempty"`
	CreateOption            *OSDiskCreateOption        `json:"createOption,omitempty"`
	DiffDiskSettings        *DiffDiskSettingsARM       `json:"diffDiskSettings,omitempty"`
	DiskSizeGB              *int                       `json:"diskSizeGB,omitempty"`
	EncryptionSettings      *DiskEncryptionSettingsARM `json:"encryptionSettings,omitempty"`
	Image                   *VirtualHardDiskARM        `json:"image,omitempty"`
	ManagedDisk             *ManagedDiskParametersARM  `json:"managedDisk,omitempty"`
	Name                    *string                    `json:"name,omitempty"`
	OsType                  *OSDiskOsType              `json:"osType,omitempty"`
	Vhd                     *VirtualHardDiskARM        `json:"vhd,omitempty"`
	WriteAcceleratorEnabled *bool                      `json:"writeAcceleratorEnabled,omitempty"`
}

// Deprecated version of UefiSettings. Use v1beta20201201.UefiSettings instead
type UefiSettingsARM struct {
	SecureBootEnabled *bool `json:"secureBootEnabled,omitempty"`
	VTpmEnabled       *bool `json:"vTpmEnabled,omitempty"`
}

// Deprecated version of VaultSecretGroup. Use v1beta20201201.VaultSecretGroup instead
type VaultSecretGroupARM struct {
	SourceVault       *SubResourceARM       `json:"sourceVault,omitempty"`
	VaultCertificates []VaultCertificateARM `json:"vaultCertificates,omitempty"`
}

// Deprecated version of VirtualMachines_Spec_Properties_NetworkProfile_NetworkInterfaces. Use v1beta20201201.VirtualMachines_Spec_Properties_NetworkProfile_NetworkInterfaces instead
type VirtualMachines_Spec_Properties_NetworkProfile_NetworkInterfacesARM struct {
	Id         *string                                 `json:"id,omitempty"`
	Properties *NetworkInterfaceReferencePropertiesARM `json:"properties,omitempty"`
}

// Deprecated version of WindowsConfiguration. Use v1beta20201201.WindowsConfiguration instead
type WindowsConfigurationARM struct {
	AdditionalUnattendContent []AdditionalUnattendContentARM `json:"additionalUnattendContent,omitempty"`
	EnableAutomaticUpdates    *bool                          `json:"enableAutomaticUpdates,omitempty"`
	PatchSettings             *PatchSettingsARM              `json:"patchSettings,omitempty"`
	ProvisionVMAgent          *bool                          `json:"provisionVMAgent,omitempty"`
	TimeZone                  *string                        `json:"timeZone,omitempty"`
	WinRM                     *WinRMConfigurationARM         `json:"winRM,omitempty"`
}

// Deprecated version of AdditionalUnattendContent. Use v1beta20201201.AdditionalUnattendContent instead
type AdditionalUnattendContentARM struct {
	ComponentName *AdditionalUnattendContentComponentName `json:"componentName,omitempty"`
	Content       *string                                 `json:"content,omitempty"`
	PassName      *AdditionalUnattendContentPassName      `json:"passName,omitempty"`
	SettingName   *AdditionalUnattendContentSettingName   `json:"settingName,omitempty"`
}

// Deprecated version of DiffDiskSettings. Use v1beta20201201.DiffDiskSettings instead
type DiffDiskSettingsARM struct {
	Option    *DiffDiskSettingsOption    `json:"option,omitempty"`
	Placement *DiffDiskSettingsPlacement `json:"placement,omitempty"`
}

// Deprecated version of DiskEncryptionSettings. Use v1beta20201201.DiskEncryptionSettings instead
type DiskEncryptionSettingsARM struct {
	DiskEncryptionKey *KeyVaultSecretReferenceARM `json:"diskEncryptionKey,omitempty"`
	Enabled           *bool                       `json:"enabled,omitempty"`
	KeyEncryptionKey  *KeyVaultKeyReferenceARM    `json:"keyEncryptionKey,omitempty"`
}

// Deprecated version of LinuxPatchSettings. Use v1beta20201201.LinuxPatchSettings instead
type LinuxPatchSettingsARM struct {
	PatchMode *LinuxPatchSettingsPatchMode `json:"patchMode,omitempty"`
}

// Deprecated version of ManagedDiskParameters. Use v1beta20201201.ManagedDiskParameters instead
type ManagedDiskParametersARM struct {
	DiskEncryptionSet  *DiskEncryptionSetParametersARM          `json:"diskEncryptionSet,omitempty"`
	Id                 *string                                  `json:"id,omitempty"`
	StorageAccountType *ManagedDiskParametersStorageAccountType `json:"storageAccountType,omitempty"`
}

// Deprecated version of NetworkInterfaceReferenceProperties. Use v1beta20201201.NetworkInterfaceReferenceProperties instead
type NetworkInterfaceReferencePropertiesARM struct {
	Primary *bool `json:"primary,omitempty"`
}

// Deprecated version of PatchSettings. Use v1beta20201201.PatchSettings instead
type PatchSettingsARM struct {
	EnableHotpatching *bool                   `json:"enableHotpatching,omitempty"`
	PatchMode         *PatchSettingsPatchMode `json:"patchMode,omitempty"`
}

// Deprecated version of SshConfiguration. Use v1beta20201201.SshConfiguration instead
type SshConfigurationARM struct {
	PublicKeys []SshPublicKeyARM `json:"publicKeys,omitempty"`
}

// Deprecated version of VaultCertificate. Use v1beta20201201.VaultCertificate instead
type VaultCertificateARM struct {
	CertificateStore *string `json:"certificateStore,omitempty"`
	CertificateUrl   *string `json:"certificateUrl,omitempty"`
}

// Deprecated version of VirtualHardDisk. Use v1beta20201201.VirtualHardDisk instead
type VirtualHardDiskARM struct {
	Uri *string `json:"uri,omitempty"`
}

// Deprecated version of WinRMConfiguration. Use v1beta20201201.WinRMConfiguration instead
type WinRMConfigurationARM struct {
	Listeners []WinRMListenerARM `json:"listeners,omitempty"`
}

// Deprecated version of DiskEncryptionSetParameters. Use v1beta20201201.DiskEncryptionSetParameters instead
type DiskEncryptionSetParametersARM struct {
	Id *string `json:"id,omitempty"`
}

// Deprecated version of KeyVaultKeyReference. Use v1beta20201201.KeyVaultKeyReference instead
type KeyVaultKeyReferenceARM struct {
	KeyUrl      *string         `json:"keyUrl,omitempty"`
	SourceVault *SubResourceARM `json:"sourceVault,omitempty"`
}

// Deprecated version of KeyVaultSecretReference. Use v1beta20201201.KeyVaultSecretReference instead
type KeyVaultSecretReferenceARM struct {
	SecretUrl   *string         `json:"secretUrl,omitempty"`
	SourceVault *SubResourceARM `json:"sourceVault,omitempty"`
}

// Deprecated version of SshPublicKey. Use v1beta20201201.SshPublicKey instead
type SshPublicKeyARM struct {
	KeyData *string `json:"keyData,omitempty"`
	Path    *string `json:"path,omitempty"`
}

// Deprecated version of WinRMListener. Use v1beta20201201.WinRMListener instead
type WinRMListenerARM struct {
	CertificateUrl *string                `json:"certificateUrl,omitempty"`
	Protocol       *WinRMListenerProtocol `json:"protocol,omitempty"`
}
