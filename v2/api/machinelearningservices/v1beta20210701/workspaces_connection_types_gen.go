// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20210701

import (
	"fmt"
	v20210701s "github.com/Azure/azure-service-operator/v2/api/machinelearningservices/v1beta20210701storage"
	"github.com/Azure/azure-service-operator/v2/internal/reflecthelpers"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	kerrors "k8s.io/apimachinery/pkg/util/errors"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Generated from: https://schema.management.azure.com/schemas/2021-07-01/Microsoft.MachineLearningServices.json#/resourceDefinitions/workspaces_connections
type WorkspacesConnection struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              WorkspacesConnections_Spec `json:"spec,omitempty"`
	Status            WorkspaceConnection_Status `json:"status,omitempty"`
}

var _ conditions.Conditioner = &WorkspacesConnection{}

// GetConditions returns the conditions of the resource
func (connection *WorkspacesConnection) GetConditions() conditions.Conditions {
	return connection.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (connection *WorkspacesConnection) SetConditions(conditions conditions.Conditions) {
	connection.Status.Conditions = conditions
}

var _ conversion.Convertible = &WorkspacesConnection{}

// ConvertFrom populates our WorkspacesConnection from the provided hub WorkspacesConnection
func (connection *WorkspacesConnection) ConvertFrom(hub conversion.Hub) error {
	source, ok := hub.(*v20210701s.WorkspacesConnection)
	if !ok {
		return fmt.Errorf("expected machinelearningservices/v1beta20210701storage/WorkspacesConnection but received %T instead", hub)
	}

	return connection.AssignPropertiesFromWorkspacesConnection(source)
}

// ConvertTo populates the provided hub WorkspacesConnection from our WorkspacesConnection
func (connection *WorkspacesConnection) ConvertTo(hub conversion.Hub) error {
	destination, ok := hub.(*v20210701s.WorkspacesConnection)
	if !ok {
		return fmt.Errorf("expected machinelearningservices/v1beta20210701storage/WorkspacesConnection but received %T instead", hub)
	}

	return connection.AssignPropertiesToWorkspacesConnection(destination)
}

// +kubebuilder:webhook:path=/mutate-machinelearningservices-azure-com-v1beta20210701-workspacesconnection,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=machinelearningservices.azure.com,resources=workspacesconnections,verbs=create;update,versions=v1beta20210701,name=default.v1beta20210701.workspacesconnections.machinelearningservices.azure.com,admissionReviewVersions=v1

var _ admission.Defaulter = &WorkspacesConnection{}

// Default applies defaults to the WorkspacesConnection resource
func (connection *WorkspacesConnection) Default() {
	connection.defaultImpl()
	var temp interface{} = connection
	if runtimeDefaulter, ok := temp.(genruntime.Defaulter); ok {
		runtimeDefaulter.CustomDefault()
	}
}

// defaultAzureName defaults the Azure name of the resource to the Kubernetes name
func (connection *WorkspacesConnection) defaultAzureName() {
	if connection.Spec.AzureName == "" {
		connection.Spec.AzureName = connection.Name
	}
}

// defaultImpl applies the code generated defaults to the WorkspacesConnection resource
func (connection *WorkspacesConnection) defaultImpl() { connection.defaultAzureName() }

var _ genruntime.KubernetesResource = &WorkspacesConnection{}

// AzureName returns the Azure name of the resource
func (connection *WorkspacesConnection) AzureName() string {
	return connection.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-07-01"
func (connection WorkspacesConnection) GetAPIVersion() string {
	return string(APIVersionValue)
}

// GetResourceKind returns the kind of the resource
func (connection *WorkspacesConnection) GetResourceKind() genruntime.ResourceKind {
	return genruntime.ResourceKindNormal
}

// GetSpec returns the specification of this resource
func (connection *WorkspacesConnection) GetSpec() genruntime.ConvertibleSpec {
	return &connection.Spec
}

// GetStatus returns the status of this resource
func (connection *WorkspacesConnection) GetStatus() genruntime.ConvertibleStatus {
	return &connection.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.MachineLearningServices/workspaces/connections"
func (connection *WorkspacesConnection) GetType() string {
	return "Microsoft.MachineLearningServices/workspaces/connections"
}

// NewEmptyStatus returns a new empty (blank) status
func (connection *WorkspacesConnection) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &WorkspaceConnection_Status{}
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (connection *WorkspacesConnection) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(connection.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  connection.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (connection *WorkspacesConnection) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*WorkspaceConnection_Status); ok {
		connection.Status = *st
		return nil
	}

	// Convert status to required version
	var st WorkspaceConnection_Status
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	connection.Status = st
	return nil
}

// +kubebuilder:webhook:path=/validate-machinelearningservices-azure-com-v1beta20210701-workspacesconnection,mutating=false,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=machinelearningservices.azure.com,resources=workspacesconnections,verbs=create;update,versions=v1beta20210701,name=validate.v1beta20210701.workspacesconnections.machinelearningservices.azure.com,admissionReviewVersions=v1

var _ admission.Validator = &WorkspacesConnection{}

// ValidateCreate validates the creation of the resource
func (connection *WorkspacesConnection) ValidateCreate() error {
	validations := connection.createValidations()
	var temp interface{} = connection
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.CreateValidations()...)
	}
	var errs []error
	for _, validation := range validations {
		err := validation()
		if err != nil {
			errs = append(errs, err)
		}
	}
	return kerrors.NewAggregate(errs)
}

// ValidateDelete validates the deletion of the resource
func (connection *WorkspacesConnection) ValidateDelete() error {
	validations := connection.deleteValidations()
	var temp interface{} = connection
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.DeleteValidations()...)
	}
	var errs []error
	for _, validation := range validations {
		err := validation()
		if err != nil {
			errs = append(errs, err)
		}
	}
	return kerrors.NewAggregate(errs)
}

// ValidateUpdate validates an update of the resource
func (connection *WorkspacesConnection) ValidateUpdate(old runtime.Object) error {
	validations := connection.updateValidations()
	var temp interface{} = connection
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.UpdateValidations()...)
	}
	var errs []error
	for _, validation := range validations {
		err := validation(old)
		if err != nil {
			errs = append(errs, err)
		}
	}
	return kerrors.NewAggregate(errs)
}

// createValidations validates the creation of the resource
func (connection *WorkspacesConnection) createValidations() []func() error {
	return []func() error{connection.validateResourceReferences}
}

// deleteValidations validates the deletion of the resource
func (connection *WorkspacesConnection) deleteValidations() []func() error {
	return nil
}

// updateValidations validates the update of the resource
func (connection *WorkspacesConnection) updateValidations() []func(old runtime.Object) error {
	return []func(old runtime.Object) error{
		func(old runtime.Object) error {
			return connection.validateResourceReferences()
		},
		connection.validateWriteOnceProperties}
}

// validateResourceReferences validates all resource references
func (connection *WorkspacesConnection) validateResourceReferences() error {
	refs, err := reflecthelpers.FindResourceReferences(&connection.Spec)
	if err != nil {
		return err
	}
	return genruntime.ValidateResourceReferences(refs)
}

// validateWriteOnceProperties validates all WriteOnce properties
func (connection *WorkspacesConnection) validateWriteOnceProperties(old runtime.Object) error {
	oldObj, ok := old.(*WorkspacesConnection)
	if !ok {
		return nil
	}

	return genruntime.ValidateWriteOnceProperties(oldObj, connection)
}

// AssignPropertiesFromWorkspacesConnection populates our WorkspacesConnection from the provided source WorkspacesConnection
func (connection *WorkspacesConnection) AssignPropertiesFromWorkspacesConnection(source *v20210701s.WorkspacesConnection) error {

	// ObjectMeta
	connection.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec WorkspacesConnections_Spec
	err := spec.AssignPropertiesFromWorkspacesConnectionsSpec(&source.Spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesFromWorkspacesConnectionsSpec() to populate field Spec")
	}
	connection.Spec = spec

	// Status
	var status WorkspaceConnection_Status
	err = status.AssignPropertiesFromWorkspaceConnectionStatus(&source.Status)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesFromWorkspaceConnectionStatus() to populate field Status")
	}
	connection.Status = status

	// No error
	return nil
}

// AssignPropertiesToWorkspacesConnection populates the provided destination WorkspacesConnection from our WorkspacesConnection
func (connection *WorkspacesConnection) AssignPropertiesToWorkspacesConnection(destination *v20210701s.WorkspacesConnection) error {

	// ObjectMeta
	destination.ObjectMeta = *connection.ObjectMeta.DeepCopy()

	// Spec
	var spec v20210701s.WorkspacesConnections_Spec
	err := connection.Spec.AssignPropertiesToWorkspacesConnectionsSpec(&spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesToWorkspacesConnectionsSpec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status v20210701s.WorkspaceConnection_Status
	err = connection.Status.AssignPropertiesToWorkspaceConnectionStatus(&status)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesToWorkspaceConnectionStatus() to populate field Status")
	}
	destination.Status = status

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (connection *WorkspacesConnection) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: connection.Spec.OriginalVersion(),
		Kind:    "WorkspacesConnection",
	}
}

// +kubebuilder:object:root=true
// Generated from: https://schema.management.azure.com/schemas/2021-07-01/Microsoft.MachineLearningServices.json#/resourceDefinitions/workspaces_connections
type WorkspacesConnectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []WorkspacesConnection `json:"items"`
}

type WorkspaceConnection_Status struct {
	// AuthType: Authorization type of the workspace connection.
	AuthType *string `json:"authType,omitempty"`

	// Category: Category of the workspace connection.
	Category *string `json:"category,omitempty"`

	// Conditions: The observed state of the resource
	Conditions []conditions.Condition `json:"conditions,omitempty"`

	// Id: ResourceId of the workspace connection.
	Id *string `json:"id,omitempty"`

	// Name: Friendly name of the workspace connection.
	Name *string `json:"name,omitempty"`

	// Target: Target of the workspace connection.
	Target *string `json:"target,omitempty"`

	// Type: Resource type of workspace connection.
	Type *string `json:"type,omitempty"`

	// Value: Value details of the workspace connection.
	Value *string `json:"value,omitempty"`

	// ValueFormat: format for the workspace connection value
	ValueFormat *WorkspaceConnectionPropsStatusValueFormat `json:"valueFormat,omitempty"`
}

var _ genruntime.ConvertibleStatus = &WorkspaceConnection_Status{}

// ConvertStatusFrom populates our WorkspaceConnection_Status from the provided source
func (connection *WorkspaceConnection_Status) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*v20210701s.WorkspaceConnection_Status)
	if ok {
		// Populate our instance from source
		return connection.AssignPropertiesFromWorkspaceConnectionStatus(src)
	}

	// Convert to an intermediate form
	src = &v20210701s.WorkspaceConnection_Status{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = connection.AssignPropertiesFromWorkspaceConnectionStatus(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our WorkspaceConnection_Status
func (connection *WorkspaceConnection_Status) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*v20210701s.WorkspaceConnection_Status)
	if ok {
		// Populate destination from our instance
		return connection.AssignPropertiesToWorkspaceConnectionStatus(dst)
	}

	// Convert to an intermediate form
	dst = &v20210701s.WorkspaceConnection_Status{}
	err := connection.AssignPropertiesToWorkspaceConnectionStatus(dst)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusTo()")
	}

	// Update dst from our instance
	err = dst.ConvertStatusTo(destination)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusTo()")
	}

	return nil
}

var _ genruntime.FromARMConverter = &WorkspaceConnection_Status{}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (connection *WorkspaceConnection_Status) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &WorkspaceConnection_StatusARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (connection *WorkspaceConnection_Status) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(WorkspaceConnection_StatusARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected WorkspaceConnection_StatusARM, got %T", armInput)
	}

	// Set property ‘AuthType’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.AuthType != nil {
			authType := *typedInput.Properties.AuthType
			connection.AuthType = &authType
		}
	}

	// Set property ‘Category’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.Category != nil {
			category := *typedInput.Properties.Category
			connection.Category = &category
		}
	}

	// no assignment for property ‘Conditions’

	// Set property ‘Id’:
	if typedInput.Id != nil {
		id := *typedInput.Id
		connection.Id = &id
	}

	// Set property ‘Name’:
	if typedInput.Name != nil {
		name := *typedInput.Name
		connection.Name = &name
	}

	// Set property ‘Target’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.Target != nil {
			target := *typedInput.Properties.Target
			connection.Target = &target
		}
	}

	// Set property ‘Type’:
	if typedInput.Type != nil {
		typeVar := *typedInput.Type
		connection.Type = &typeVar
	}

	// Set property ‘Value’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.Value != nil {
			value := *typedInput.Properties.Value
			connection.Value = &value
		}
	}

	// Set property ‘ValueFormat’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.ValueFormat != nil {
			valueFormat := *typedInput.Properties.ValueFormat
			connection.ValueFormat = &valueFormat
		}
	}

	// No error
	return nil
}

// AssignPropertiesFromWorkspaceConnectionStatus populates our WorkspaceConnection_Status from the provided source WorkspaceConnection_Status
func (connection *WorkspaceConnection_Status) AssignPropertiesFromWorkspaceConnectionStatus(source *v20210701s.WorkspaceConnection_Status) error {

	// AuthType
	connection.AuthType = genruntime.ClonePointerToString(source.AuthType)

	// Category
	connection.Category = genruntime.ClonePointerToString(source.Category)

	// Conditions
	connection.Conditions = genruntime.CloneSliceOfCondition(source.Conditions)

	// Id
	connection.Id = genruntime.ClonePointerToString(source.Id)

	// Name
	connection.Name = genruntime.ClonePointerToString(source.Name)

	// Target
	connection.Target = genruntime.ClonePointerToString(source.Target)

	// Type
	connection.Type = genruntime.ClonePointerToString(source.Type)

	// Value
	connection.Value = genruntime.ClonePointerToString(source.Value)

	// ValueFormat
	if source.ValueFormat != nil {
		valueFormat := WorkspaceConnectionPropsStatusValueFormat(*source.ValueFormat)
		connection.ValueFormat = &valueFormat
	} else {
		connection.ValueFormat = nil
	}

	// No error
	return nil
}

// AssignPropertiesToWorkspaceConnectionStatus populates the provided destination WorkspaceConnection_Status from our WorkspaceConnection_Status
func (connection *WorkspaceConnection_Status) AssignPropertiesToWorkspaceConnectionStatus(destination *v20210701s.WorkspaceConnection_Status) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// AuthType
	destination.AuthType = genruntime.ClonePointerToString(connection.AuthType)

	// Category
	destination.Category = genruntime.ClonePointerToString(connection.Category)

	// Conditions
	destination.Conditions = genruntime.CloneSliceOfCondition(connection.Conditions)

	// Id
	destination.Id = genruntime.ClonePointerToString(connection.Id)

	// Name
	destination.Name = genruntime.ClonePointerToString(connection.Name)

	// Target
	destination.Target = genruntime.ClonePointerToString(connection.Target)

	// Type
	destination.Type = genruntime.ClonePointerToString(connection.Type)

	// Value
	destination.Value = genruntime.ClonePointerToString(connection.Value)

	// ValueFormat
	if connection.ValueFormat != nil {
		valueFormat := string(*connection.ValueFormat)
		destination.ValueFormat = &valueFormat
	} else {
		destination.ValueFormat = nil
	}

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

type WorkspacesConnections_Spec struct {
	// AuthType: Authorization type of the workspace connection.
	AuthType *string `json:"authType,omitempty"`

	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName string `json:"azureName,omitempty"`

	// Category: Category of the workspace connection.
	Category *string `json:"category,omitempty"`

	// Location: Location to deploy resource to
	Location *string `json:"location,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a machinelearningservices.azure.com/Workspace resource
	Owner *genruntime.KnownResourceReference `group:"machinelearningservices.azure.com" json:"owner,omitempty" kind:"Workspace"`

	// Tags: Name-value pairs to add to the resource
	Tags map[string]string `json:"tags,omitempty"`

	// Target: Target of the workspace connection.
	Target *string `json:"target,omitempty"`

	// Value: Value details of the workspace connection.
	Value *string `json:"value,omitempty"`

	// ValueFormat: format for the workspace connection value.
	ValueFormat *WorkspaceConnectionPropsValueFormat `json:"valueFormat,omitempty"`
}

var _ genruntime.ARMTransformer = &WorkspacesConnections_Spec{}

// ConvertToARM converts from a Kubernetes CRD object to an ARM object
func (connections *WorkspacesConnections_Spec) ConvertToARM(resolved genruntime.ConvertToARMResolvedDetails) (interface{}, error) {
	if connections == nil {
		return nil, nil
	}
	result := &WorkspacesConnections_SpecARM{}

	// Set property ‘Location’:
	if connections.Location != nil {
		location := *connections.Location
		result.Location = &location
	}

	// Set property ‘Name’:
	result.Name = resolved.Name

	// Set property ‘Properties’:
	if connections.AuthType != nil ||
		connections.Category != nil ||
		connections.Target != nil ||
		connections.Value != nil ||
		connections.ValueFormat != nil {
		result.Properties = &WorkspaceConnectionPropsARM{}
	}
	if connections.AuthType != nil {
		authType := *connections.AuthType
		result.Properties.AuthType = &authType
	}
	if connections.Category != nil {
		category := *connections.Category
		result.Properties.Category = &category
	}
	if connections.Target != nil {
		target := *connections.Target
		result.Properties.Target = &target
	}
	if connections.Value != nil {
		value := *connections.Value
		result.Properties.Value = &value
	}
	if connections.ValueFormat != nil {
		valueFormat := *connections.ValueFormat
		result.Properties.ValueFormat = &valueFormat
	}

	// Set property ‘Tags’:
	if connections.Tags != nil {
		result.Tags = make(map[string]string, len(connections.Tags))
		for key, tagsValue := range connections.Tags {
			result.Tags[key] = tagsValue
		}
	}
	return result, nil
}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (connections *WorkspacesConnections_Spec) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &WorkspacesConnections_SpecARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (connections *WorkspacesConnections_Spec) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(WorkspacesConnections_SpecARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected WorkspacesConnections_SpecARM, got %T", armInput)
	}

	// Set property ‘AuthType’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.AuthType != nil {
			authType := *typedInput.Properties.AuthType
			connections.AuthType = &authType
		}
	}

	// Set property ‘AzureName’:
	connections.SetAzureName(genruntime.ExtractKubernetesResourceNameFromARMName(typedInput.Name))

	// Set property ‘Category’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.Category != nil {
			category := *typedInput.Properties.Category
			connections.Category = &category
		}
	}

	// Set property ‘Location’:
	if typedInput.Location != nil {
		location := *typedInput.Location
		connections.Location = &location
	}

	// Set property ‘Owner’:
	connections.Owner = &genruntime.KnownResourceReference{
		Name: owner.Name,
	}

	// Set property ‘Tags’:
	if typedInput.Tags != nil {
		connections.Tags = make(map[string]string, len(typedInput.Tags))
		for key, value := range typedInput.Tags {
			connections.Tags[key] = value
		}
	}

	// Set property ‘Target’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.Target != nil {
			target := *typedInput.Properties.Target
			connections.Target = &target
		}
	}

	// Set property ‘Value’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.Value != nil {
			value := *typedInput.Properties.Value
			connections.Value = &value
		}
	}

	// Set property ‘ValueFormat’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.ValueFormat != nil {
			valueFormat := *typedInput.Properties.ValueFormat
			connections.ValueFormat = &valueFormat
		}
	}

	// No error
	return nil
}

var _ genruntime.ConvertibleSpec = &WorkspacesConnections_Spec{}

// ConvertSpecFrom populates our WorkspacesConnections_Spec from the provided source
func (connections *WorkspacesConnections_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*v20210701s.WorkspacesConnections_Spec)
	if ok {
		// Populate our instance from source
		return connections.AssignPropertiesFromWorkspacesConnectionsSpec(src)
	}

	// Convert to an intermediate form
	src = &v20210701s.WorkspacesConnections_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = connections.AssignPropertiesFromWorkspacesConnectionsSpec(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our WorkspacesConnections_Spec
func (connections *WorkspacesConnections_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*v20210701s.WorkspacesConnections_Spec)
	if ok {
		// Populate destination from our instance
		return connections.AssignPropertiesToWorkspacesConnectionsSpec(dst)
	}

	// Convert to an intermediate form
	dst = &v20210701s.WorkspacesConnections_Spec{}
	err := connections.AssignPropertiesToWorkspacesConnectionsSpec(dst)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecTo()")
	}

	// Update dst from our instance
	err = dst.ConvertSpecTo(destination)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecTo()")
	}

	return nil
}

// AssignPropertiesFromWorkspacesConnectionsSpec populates our WorkspacesConnections_Spec from the provided source WorkspacesConnections_Spec
func (connections *WorkspacesConnections_Spec) AssignPropertiesFromWorkspacesConnectionsSpec(source *v20210701s.WorkspacesConnections_Spec) error {

	// AuthType
	connections.AuthType = genruntime.ClonePointerToString(source.AuthType)

	// AzureName
	connections.AzureName = source.AzureName

	// Category
	connections.Category = genruntime.ClonePointerToString(source.Category)

	// Location
	connections.Location = genruntime.ClonePointerToString(source.Location)

	// Owner
	if source.Owner != nil {
		owner := source.Owner.Copy()
		connections.Owner = &owner
	} else {
		connections.Owner = nil
	}

	// Tags
	connections.Tags = genruntime.CloneMapOfStringToString(source.Tags)

	// Target
	connections.Target = genruntime.ClonePointerToString(source.Target)

	// Value
	connections.Value = genruntime.ClonePointerToString(source.Value)

	// ValueFormat
	if source.ValueFormat != nil {
		valueFormat := WorkspaceConnectionPropsValueFormat(*source.ValueFormat)
		connections.ValueFormat = &valueFormat
	} else {
		connections.ValueFormat = nil
	}

	// No error
	return nil
}

// AssignPropertiesToWorkspacesConnectionsSpec populates the provided destination WorkspacesConnections_Spec from our WorkspacesConnections_Spec
func (connections *WorkspacesConnections_Spec) AssignPropertiesToWorkspacesConnectionsSpec(destination *v20210701s.WorkspacesConnections_Spec) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// AuthType
	destination.AuthType = genruntime.ClonePointerToString(connections.AuthType)

	// AzureName
	destination.AzureName = connections.AzureName

	// Category
	destination.Category = genruntime.ClonePointerToString(connections.Category)

	// Location
	destination.Location = genruntime.ClonePointerToString(connections.Location)

	// OriginalVersion
	destination.OriginalVersion = connections.OriginalVersion()

	// Owner
	if connections.Owner != nil {
		owner := connections.Owner.Copy()
		destination.Owner = &owner
	} else {
		destination.Owner = nil
	}

	// Tags
	destination.Tags = genruntime.CloneMapOfStringToString(connections.Tags)

	// Target
	destination.Target = genruntime.ClonePointerToString(connections.Target)

	// Value
	destination.Value = genruntime.ClonePointerToString(connections.Value)

	// ValueFormat
	if connections.ValueFormat != nil {
		valueFormat := string(*connections.ValueFormat)
		destination.ValueFormat = &valueFormat
	} else {
		destination.ValueFormat = nil
	}

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

// OriginalVersion returns the original API version used to create the resource.
func (connections *WorkspacesConnections_Spec) OriginalVersion() string {
	return GroupVersion.Version
}

// SetAzureName sets the Azure name of the resource
func (connections *WorkspacesConnections_Spec) SetAzureName(azureName string) {
	connections.AzureName = azureName
}

// +kubebuilder:validation:Enum={"JSON"}
type WorkspaceConnectionPropsValueFormat string

const WorkspaceConnectionPropsValueFormatJSON = WorkspaceConnectionPropsValueFormat("JSON")

func init() {
	SchemeBuilder.Register(&WorkspacesConnection{}, &WorkspacesConnectionList{})
}
