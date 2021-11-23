// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210515storage

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// +kubebuilder:rbac:groups=documentdb.azure.com,resources=sqldatabasecontainers,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=documentdb.azure.com,resources={sqldatabasecontainers/status,sqldatabasecontainers/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
//Storage version of v1alpha1api20210515.SqlDatabaseContainer
//Generated from: https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/resourceDefinitions/databaseAccounts_sqlDatabases_containers
type SqlDatabaseContainer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DatabaseAccountsSqlDatabasesContainers_Spec `json:"spec,omitempty"`
	Status            SqlContainerGetResults_Status               `json:"status,omitempty"`
}

var _ conditions.Conditioner = &SqlDatabaseContainer{}

// GetConditions returns the conditions of the resource
func (sqlDatabaseContainer *SqlDatabaseContainer) GetConditions() conditions.Conditions {
	return sqlDatabaseContainer.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (sqlDatabaseContainer *SqlDatabaseContainer) SetConditions(conditions conditions.Conditions) {
	sqlDatabaseContainer.Status.Conditions = conditions
}

var _ genruntime.KubernetesResource = &SqlDatabaseContainer{}

// AzureName returns the Azure name of the resource
func (sqlDatabaseContainer *SqlDatabaseContainer) AzureName() string {
	return sqlDatabaseContainer.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-05-15"
func (sqlDatabaseContainer SqlDatabaseContainer) GetAPIVersion() string {
	return "2021-05-15"
}

// GetResourceKind returns the kind of the resource
func (sqlDatabaseContainer *SqlDatabaseContainer) GetResourceKind() genruntime.ResourceKind {
	return genruntime.ResourceKindNormal
}

// GetSpec returns the specification of this resource
func (sqlDatabaseContainer *SqlDatabaseContainer) GetSpec() genruntime.ConvertibleSpec {
	return &sqlDatabaseContainer.Spec
}

// GetStatus returns the status of this resource
func (sqlDatabaseContainer *SqlDatabaseContainer) GetStatus() genruntime.ConvertibleStatus {
	return &sqlDatabaseContainer.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.DocumentDB/databaseAccounts/sqlDatabases/containers"
func (sqlDatabaseContainer *SqlDatabaseContainer) GetType() string {
	return "Microsoft.DocumentDB/databaseAccounts/sqlDatabases/containers"
}

// NewEmptyStatus returns a new empty (blank) status
func (sqlDatabaseContainer *SqlDatabaseContainer) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &SqlContainerGetResults_Status{}
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (sqlDatabaseContainer *SqlDatabaseContainer) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(sqlDatabaseContainer.Spec)
	return &genruntime.ResourceReference{
		Group:     group,
		Kind:      kind,
		Namespace: sqlDatabaseContainer.Namespace,
		Name:      sqlDatabaseContainer.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (sqlDatabaseContainer *SqlDatabaseContainer) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*SqlContainerGetResults_Status); ok {
		sqlDatabaseContainer.Status = *st
		return nil
	}

	// Convert status to required version
	var st SqlContainerGetResults_Status
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	sqlDatabaseContainer.Status = st
	return nil
}

// Hub marks that this SqlDatabaseContainer is the hub type for conversion
func (sqlDatabaseContainer *SqlDatabaseContainer) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (sqlDatabaseContainer *SqlDatabaseContainer) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: sqlDatabaseContainer.Spec.OriginalVersion,
		Kind:    "SqlDatabaseContainer",
	}
}

// +kubebuilder:object:root=true
//Storage version of v1alpha1api20210515.SqlDatabaseContainer
//Generated from: https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/resourceDefinitions/databaseAccounts_sqlDatabases_containers
type SqlDatabaseContainerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SqlDatabaseContainer `json:"items"`
}

//Storage version of v1alpha1api20210515.DatabaseAccountsSqlDatabasesContainers_Spec
type DatabaseAccountsSqlDatabasesContainers_Spec struct {
	//AzureName: The name of the resource in Azure. This is often the same as the name
	//of the resource in Kubernetes but it doesn't have to be.
	AzureName       string               `json:"azureName"`
	Location        *string              `json:"location,omitempty"`
	Options         *CreateUpdateOptions `json:"options,omitempty"`
	OriginalVersion string               `json:"originalVersion"`

	// +kubebuilder:validation:Required
	Owner       genruntime.KnownResourceReference `group:"documentdb.azure.com" json:"owner" kind:"SqlDatabase"`
	PropertyBag genruntime.PropertyBag            `json:"$propertyBag,omitempty"`
	Resource    *SqlContainerResource             `json:"resource,omitempty"`
	Tags        map[string]string                 `json:"tags,omitempty"`
}

var _ genruntime.ConvertibleSpec = &DatabaseAccountsSqlDatabasesContainers_Spec{}

// ConvertSpecFrom populates our DatabaseAccountsSqlDatabasesContainers_Spec from the provided source
func (databaseAccountsSqlDatabasesContainersSpec *DatabaseAccountsSqlDatabasesContainers_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == databaseAccountsSqlDatabasesContainersSpec {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(databaseAccountsSqlDatabasesContainersSpec)
}

// ConvertSpecTo populates the provided destination from our DatabaseAccountsSqlDatabasesContainers_Spec
func (databaseAccountsSqlDatabasesContainersSpec *DatabaseAccountsSqlDatabasesContainers_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == databaseAccountsSqlDatabasesContainersSpec {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(databaseAccountsSqlDatabasesContainersSpec)
}

//Storage version of v1alpha1api20210515.SqlContainerGetResults_Status
type SqlContainerGetResults_Status struct {
	Conditions  []conditions.Condition                     `json:"conditions,omitempty"`
	Id          *string                                    `json:"id,omitempty"`
	Location    *string                                    `json:"location,omitempty"`
	Name        *string                                    `json:"name,omitempty"`
	Options     *OptionsResource_Status                    `json:"options,omitempty"`
	PropertyBag genruntime.PropertyBag                     `json:"$propertyBag,omitempty"`
	Resource    *SqlContainerGetProperties_Status_Resource `json:"resource,omitempty"`
	Tags        map[string]string                          `json:"tags,omitempty"`
	Type        *string                                    `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &SqlContainerGetResults_Status{}

// ConvertStatusFrom populates our SqlContainerGetResults_Status from the provided source
func (sqlContainerGetResultsStatus *SqlContainerGetResults_Status) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == sqlContainerGetResultsStatus {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(sqlContainerGetResultsStatus)
}

// ConvertStatusTo populates the provided destination from our SqlContainerGetResults_Status
func (sqlContainerGetResultsStatus *SqlContainerGetResults_Status) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == sqlContainerGetResultsStatus {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(sqlContainerGetResultsStatus)
}

//Storage version of v1alpha1api20210515.SqlContainerGetProperties_Status_Resource
type SqlContainerGetProperties_Status_Resource struct {
	AnalyticalStorageTtl     *int                             `json:"analyticalStorageTtl,omitempty"`
	ConflictResolutionPolicy *ConflictResolutionPolicy_Status `json:"conflictResolutionPolicy,omitempty"`
	DefaultTtl               *int                             `json:"defaultTtl,omitempty"`
	Etag                     *string                          `json:"_etag,omitempty"`
	Id                       *string                          `json:"id,omitempty"`
	IndexingPolicy           *IndexingPolicy_Status           `json:"indexingPolicy,omitempty"`
	PartitionKey             *ContainerPartitionKey_Status    `json:"partitionKey,omitempty"`
	PropertyBag              genruntime.PropertyBag           `json:"$propertyBag,omitempty"`
	Rid                      *string                          `json:"_rid,omitempty"`
	Ts                       *float64                         `json:"_ts,omitempty"`
	UniqueKeyPolicy          *UniqueKeyPolicy_Status          `json:"uniqueKeyPolicy,omitempty"`
}

//Storage version of v1alpha1api20210515.SqlContainerResource
//Generated from: https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/SqlContainerResource
type SqlContainerResource struct {
	AnalyticalStorageTtl     *int                      `json:"analyticalStorageTtl,omitempty"`
	ConflictResolutionPolicy *ConflictResolutionPolicy `json:"conflictResolutionPolicy,omitempty"`
	DefaultTtl               *int                      `json:"defaultTtl,omitempty"`
	Id                       *string                   `json:"id,omitempty"`
	IndexingPolicy           *IndexingPolicy           `json:"indexingPolicy,omitempty"`
	PartitionKey             *ContainerPartitionKey    `json:"partitionKey,omitempty"`
	PropertyBag              genruntime.PropertyBag    `json:"$propertyBag,omitempty"`
	UniqueKeyPolicy          *UniqueKeyPolicy          `json:"uniqueKeyPolicy,omitempty"`
}

//Storage version of v1alpha1api20210515.ConflictResolutionPolicy
//Generated from: https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/ConflictResolutionPolicy
type ConflictResolutionPolicy struct {
	ConflictResolutionPath      *string                `json:"conflictResolutionPath,omitempty"`
	ConflictResolutionProcedure *string                `json:"conflictResolutionProcedure,omitempty"`
	Mode                        *string                `json:"mode,omitempty"`
	PropertyBag                 genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20210515.ConflictResolutionPolicy_Status
type ConflictResolutionPolicy_Status struct {
	ConflictResolutionPath      *string                `json:"conflictResolutionPath,omitempty"`
	ConflictResolutionProcedure *string                `json:"conflictResolutionProcedure,omitempty"`
	Mode                        *string                `json:"mode,omitempty"`
	PropertyBag                 genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20210515.ContainerPartitionKey
//Generated from: https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/ContainerPartitionKey
type ContainerPartitionKey struct {
	Kind        *string                `json:"kind,omitempty"`
	Paths       []string               `json:"paths,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Version     *int                   `json:"version,omitempty"`
}

//Storage version of v1alpha1api20210515.ContainerPartitionKey_Status
type ContainerPartitionKey_Status struct {
	Kind        *string                `json:"kind,omitempty"`
	Paths       []string               `json:"paths,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	SystemKey   *bool                  `json:"systemKey,omitempty"`
	Version     *int                   `json:"version,omitempty"`
}

//Storage version of v1alpha1api20210515.IndexingPolicy
//Generated from: https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/IndexingPolicy
type IndexingPolicy struct {
	Automatic        *bool                  `json:"automatic,omitempty"`
	CompositeIndexes [][]CompositePath      `json:"compositeIndexes,omitempty"`
	ExcludedPaths    []ExcludedPath         `json:"excludedPaths,omitempty"`
	IncludedPaths    []IncludedPath         `json:"includedPaths,omitempty"`
	IndexingMode     *string                `json:"indexingMode,omitempty"`
	PropertyBag      genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	SpatialIndexes   []SpatialSpec          `json:"spatialIndexes,omitempty"`
}

//Storage version of v1alpha1api20210515.IndexingPolicy_Status
type IndexingPolicy_Status struct {
	Automatic        *bool                    `json:"automatic,omitempty"`
	CompositeIndexes [][]CompositePath_Status `json:"compositeIndexes,omitempty"`
	ExcludedPaths    []ExcludedPath_Status    `json:"excludedPaths,omitempty"`
	IncludedPaths    []IncludedPath_Status    `json:"includedPaths,omitempty"`
	IndexingMode     *string                  `json:"indexingMode,omitempty"`
	PropertyBag      genruntime.PropertyBag   `json:"$propertyBag,omitempty"`
	SpatialIndexes   []SpatialSpec_Status     `json:"spatialIndexes,omitempty"`
}

//Storage version of v1alpha1api20210515.UniqueKeyPolicy
//Generated from: https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/UniqueKeyPolicy
type UniqueKeyPolicy struct {
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	UniqueKeys  []UniqueKey            `json:"uniqueKeys,omitempty"`
}

//Storage version of v1alpha1api20210515.UniqueKeyPolicy_Status
type UniqueKeyPolicy_Status struct {
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	UniqueKeys  []UniqueKey_Status     `json:"uniqueKeys,omitempty"`
}

//Storage version of v1alpha1api20210515.CompositePath
//Generated from: https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/CompositePath
type CompositePath struct {
	Order       *string                `json:"order,omitempty"`
	Path        *string                `json:"path,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20210515.CompositePath_Status
type CompositePath_Status struct {
	Order       *string                `json:"order,omitempty"`
	Path        *string                `json:"path,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20210515.ExcludedPath
//Generated from: https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/ExcludedPath
type ExcludedPath struct {
	Path        *string                `json:"path,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20210515.ExcludedPath_Status
type ExcludedPath_Status struct {
	Path        *string                `json:"path,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20210515.IncludedPath
//Generated from: https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/IncludedPath
type IncludedPath struct {
	Indexes     []Indexes              `json:"indexes,omitempty"`
	Path        *string                `json:"path,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20210515.IncludedPath_Status
type IncludedPath_Status struct {
	Indexes     []Indexes_Status       `json:"indexes,omitempty"`
	Path        *string                `json:"path,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20210515.SpatialSpec
//Generated from: https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/SpatialSpec
type SpatialSpec struct {
	Path        *string                `json:"path,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Types       []string               `json:"types,omitempty"`
}

//Storage version of v1alpha1api20210515.SpatialSpec_Status
type SpatialSpec_Status struct {
	Path        *string                `json:"path,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Types       []string               `json:"types,omitempty"`
}

//Storage version of v1alpha1api20210515.UniqueKey
//Generated from: https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/UniqueKey
type UniqueKey struct {
	Paths       []string               `json:"paths,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20210515.UniqueKey_Status
type UniqueKey_Status struct {
	Paths       []string               `json:"paths,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20210515.Indexes
//Generated from: https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/Indexes
type Indexes struct {
	DataType    *string                `json:"dataType,omitempty"`
	Kind        *string                `json:"kind,omitempty"`
	Precision   *int                   `json:"precision,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20210515.Indexes_Status
type Indexes_Status struct {
	DataType    *string                `json:"dataType,omitempty"`
	Kind        *string                `json:"kind,omitempty"`
	Precision   *int                   `json:"precision,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

func init() {
	SchemeBuilder.Register(&SqlDatabaseContainer{}, &SqlDatabaseContainerList{})
}