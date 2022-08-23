// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20201101

import (
	"encoding/json"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/kr/pretty"
	"github.com/kylelemons/godebug/diff"
	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
	"github.com/leanovate/gopter/prop"
	"os"
	"reflect"
	"testing"
)

func Test_NetworkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbeddedARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NetworkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbeddedARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNetworkSecurityGroupSTATUSNetworkSecurityGroupSubResourceEmbeddedARM, NetworkSecurityGroupSTATUSNetworkSecurityGroupSubResourceEmbeddedARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNetworkSecurityGroupSTATUSNetworkSecurityGroupSubResourceEmbeddedARM runs a test to see if a specific instance of NetworkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbeddedARM round trips to JSON and back losslessly
func RunJSONSerializationTestForNetworkSecurityGroupSTATUSNetworkSecurityGroupSubResourceEmbeddedARM(subject NetworkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbeddedARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NetworkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbeddedARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of NetworkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbeddedARM instances for property testing -
// lazily instantiated by NetworkSecurityGroupSTATUSNetworkSecurityGroupSubResourceEmbeddedARMGenerator()
var networkSecurityGroupSTATUSNetworkSecurityGroupSubResourceEmbeddedARMGenerator gopter.Gen

// NetworkSecurityGroupSTATUSNetworkSecurityGroupSubResourceEmbeddedARMGenerator returns a generator of NetworkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbeddedARM instances for property testing.
// We first initialize networkSecurityGroupSTATUSNetworkSecurityGroupSubResourceEmbeddedARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func NetworkSecurityGroupSTATUSNetworkSecurityGroupSubResourceEmbeddedARMGenerator() gopter.Gen {
	if networkSecurityGroupSTATUSNetworkSecurityGroupSubResourceEmbeddedARMGenerator != nil {
		return networkSecurityGroupSTATUSNetworkSecurityGroupSubResourceEmbeddedARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNetworkSecurityGroupSTATUSNetworkSecurityGroupSubResourceEmbeddedARM(generators)
	networkSecurityGroupSTATUSNetworkSecurityGroupSubResourceEmbeddedARMGenerator = gen.Struct(reflect.TypeOf(NetworkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbeddedARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNetworkSecurityGroupSTATUSNetworkSecurityGroupSubResourceEmbeddedARM(generators)
	AddRelatedPropertyGeneratorsForNetworkSecurityGroupSTATUSNetworkSecurityGroupSubResourceEmbeddedARM(generators)
	networkSecurityGroupSTATUSNetworkSecurityGroupSubResourceEmbeddedARMGenerator = gen.Struct(reflect.TypeOf(NetworkSecurityGroup_STATUS_NetworkSecurityGroup_SubResourceEmbeddedARM{}), generators)

	return networkSecurityGroupSTATUSNetworkSecurityGroupSubResourceEmbeddedARMGenerator
}

// AddIndependentPropertyGeneratorsForNetworkSecurityGroupSTATUSNetworkSecurityGroupSubResourceEmbeddedARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNetworkSecurityGroupSTATUSNetworkSecurityGroupSubResourceEmbeddedARM(gens map[string]gopter.Gen) {
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForNetworkSecurityGroupSTATUSNetworkSecurityGroupSubResourceEmbeddedARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNetworkSecurityGroupSTATUSNetworkSecurityGroupSubResourceEmbeddedARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(NetworkSecurityGroupPropertiesFormatSTATUSARMGenerator())
}

func Test_NetworkSecurityGroupPropertiesFormat_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NetworkSecurityGroupPropertiesFormat_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNetworkSecurityGroupPropertiesFormatSTATUSARM, NetworkSecurityGroupPropertiesFormatSTATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNetworkSecurityGroupPropertiesFormatSTATUSARM runs a test to see if a specific instance of NetworkSecurityGroupPropertiesFormat_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForNetworkSecurityGroupPropertiesFormatSTATUSARM(subject NetworkSecurityGroupPropertiesFormat_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NetworkSecurityGroupPropertiesFormat_STATUSARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of NetworkSecurityGroupPropertiesFormat_STATUSARM instances for property testing - lazily instantiated by
// NetworkSecurityGroupPropertiesFormatSTATUSARMGenerator()
var networkSecurityGroupPropertiesFormatSTATUSARMGenerator gopter.Gen

// NetworkSecurityGroupPropertiesFormatSTATUSARMGenerator returns a generator of NetworkSecurityGroupPropertiesFormat_STATUSARM instances for property testing.
// We first initialize networkSecurityGroupPropertiesFormatSTATUSARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func NetworkSecurityGroupPropertiesFormatSTATUSARMGenerator() gopter.Gen {
	if networkSecurityGroupPropertiesFormatSTATUSARMGenerator != nil {
		return networkSecurityGroupPropertiesFormatSTATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNetworkSecurityGroupPropertiesFormatSTATUSARM(generators)
	networkSecurityGroupPropertiesFormatSTATUSARMGenerator = gen.Struct(reflect.TypeOf(NetworkSecurityGroupPropertiesFormat_STATUSARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNetworkSecurityGroupPropertiesFormatSTATUSARM(generators)
	AddRelatedPropertyGeneratorsForNetworkSecurityGroupPropertiesFormatSTATUSARM(generators)
	networkSecurityGroupPropertiesFormatSTATUSARMGenerator = gen.Struct(reflect.TypeOf(NetworkSecurityGroupPropertiesFormat_STATUSARM{}), generators)

	return networkSecurityGroupPropertiesFormatSTATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForNetworkSecurityGroupPropertiesFormatSTATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNetworkSecurityGroupPropertiesFormatSTATUSARM(gens map[string]gopter.Gen) {
	gens["ProvisioningState"] = gen.PtrOf(gen.OneConstOf(
		ProvisioningState_STATUS_Deleting,
		ProvisioningState_STATUS_Failed,
		ProvisioningState_STATUS_Succeeded,
		ProvisioningState_STATUS_Updating))
	gens["ResourceGuid"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForNetworkSecurityGroupPropertiesFormatSTATUSARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNetworkSecurityGroupPropertiesFormatSTATUSARM(gens map[string]gopter.Gen) {
	gens["DefaultSecurityRules"] = gen.SliceOf(SecurityRuleSTATUSNetworkSecurityGroupSubResourceEmbeddedARMGenerator())
	gens["FlowLogs"] = gen.SliceOf(FlowLogSTATUSSubResourceEmbeddedARMGenerator())
	gens["NetworkInterfaces"] = gen.SliceOf(NetworkInterfaceSTATUSNetworkSecurityGroupSubResourceEmbeddedARMGenerator())
	gens["SecurityRules"] = gen.SliceOf(SecurityRuleSTATUSNetworkSecurityGroupSubResourceEmbeddedARMGenerator())
	gens["Subnets"] = gen.SliceOf(SubnetSTATUSNetworkSecurityGroupSubResourceEmbeddedARMGenerator())
}

func Test_FlowLog_STATUS_SubResourceEmbeddedARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of FlowLog_STATUS_SubResourceEmbeddedARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForFlowLogSTATUSSubResourceEmbeddedARM, FlowLogSTATUSSubResourceEmbeddedARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForFlowLogSTATUSSubResourceEmbeddedARM runs a test to see if a specific instance of FlowLog_STATUS_SubResourceEmbeddedARM round trips to JSON and back losslessly
func RunJSONSerializationTestForFlowLogSTATUSSubResourceEmbeddedARM(subject FlowLog_STATUS_SubResourceEmbeddedARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual FlowLog_STATUS_SubResourceEmbeddedARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of FlowLog_STATUS_SubResourceEmbeddedARM instances for property testing - lazily instantiated by
// FlowLogSTATUSSubResourceEmbeddedARMGenerator()
var flowLogSTATUSSubResourceEmbeddedARMGenerator gopter.Gen

// FlowLogSTATUSSubResourceEmbeddedARMGenerator returns a generator of FlowLog_STATUS_SubResourceEmbeddedARM instances for property testing.
func FlowLogSTATUSSubResourceEmbeddedARMGenerator() gopter.Gen {
	if flowLogSTATUSSubResourceEmbeddedARMGenerator != nil {
		return flowLogSTATUSSubResourceEmbeddedARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFlowLogSTATUSSubResourceEmbeddedARM(generators)
	flowLogSTATUSSubResourceEmbeddedARMGenerator = gen.Struct(reflect.TypeOf(FlowLog_STATUS_SubResourceEmbeddedARM{}), generators)

	return flowLogSTATUSSubResourceEmbeddedARMGenerator
}

// AddIndependentPropertyGeneratorsForFlowLogSTATUSSubResourceEmbeddedARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForFlowLogSTATUSSubResourceEmbeddedARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}

func Test_NetworkInterface_STATUS_NetworkSecurityGroup_SubResourceEmbeddedARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NetworkInterface_STATUS_NetworkSecurityGroup_SubResourceEmbeddedARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNetworkInterfaceSTATUSNetworkSecurityGroupSubResourceEmbeddedARM, NetworkInterfaceSTATUSNetworkSecurityGroupSubResourceEmbeddedARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNetworkInterfaceSTATUSNetworkSecurityGroupSubResourceEmbeddedARM runs a test to see if a specific instance of NetworkInterface_STATUS_NetworkSecurityGroup_SubResourceEmbeddedARM round trips to JSON and back losslessly
func RunJSONSerializationTestForNetworkInterfaceSTATUSNetworkSecurityGroupSubResourceEmbeddedARM(subject NetworkInterface_STATUS_NetworkSecurityGroup_SubResourceEmbeddedARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NetworkInterface_STATUS_NetworkSecurityGroup_SubResourceEmbeddedARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of NetworkInterface_STATUS_NetworkSecurityGroup_SubResourceEmbeddedARM instances for property testing -
// lazily instantiated by NetworkInterfaceSTATUSNetworkSecurityGroupSubResourceEmbeddedARMGenerator()
var networkInterfaceSTATUSNetworkSecurityGroupSubResourceEmbeddedARMGenerator gopter.Gen

// NetworkInterfaceSTATUSNetworkSecurityGroupSubResourceEmbeddedARMGenerator returns a generator of NetworkInterface_STATUS_NetworkSecurityGroup_SubResourceEmbeddedARM instances for property testing.
// We first initialize networkInterfaceSTATUSNetworkSecurityGroupSubResourceEmbeddedARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func NetworkInterfaceSTATUSNetworkSecurityGroupSubResourceEmbeddedARMGenerator() gopter.Gen {
	if networkInterfaceSTATUSNetworkSecurityGroupSubResourceEmbeddedARMGenerator != nil {
		return networkInterfaceSTATUSNetworkSecurityGroupSubResourceEmbeddedARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNetworkInterfaceSTATUSNetworkSecurityGroupSubResourceEmbeddedARM(generators)
	networkInterfaceSTATUSNetworkSecurityGroupSubResourceEmbeddedARMGenerator = gen.Struct(reflect.TypeOf(NetworkInterface_STATUS_NetworkSecurityGroup_SubResourceEmbeddedARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNetworkInterfaceSTATUSNetworkSecurityGroupSubResourceEmbeddedARM(generators)
	AddRelatedPropertyGeneratorsForNetworkInterfaceSTATUSNetworkSecurityGroupSubResourceEmbeddedARM(generators)
	networkInterfaceSTATUSNetworkSecurityGroupSubResourceEmbeddedARMGenerator = gen.Struct(reflect.TypeOf(NetworkInterface_STATUS_NetworkSecurityGroup_SubResourceEmbeddedARM{}), generators)

	return networkInterfaceSTATUSNetworkSecurityGroupSubResourceEmbeddedARMGenerator
}

// AddIndependentPropertyGeneratorsForNetworkInterfaceSTATUSNetworkSecurityGroupSubResourceEmbeddedARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNetworkInterfaceSTATUSNetworkSecurityGroupSubResourceEmbeddedARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForNetworkInterfaceSTATUSNetworkSecurityGroupSubResourceEmbeddedARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNetworkInterfaceSTATUSNetworkSecurityGroupSubResourceEmbeddedARM(gens map[string]gopter.Gen) {
	gens["ExtendedLocation"] = gen.PtrOf(ExtendedLocationSTATUSARMGenerator())
}

func Test_SecurityRule_STATUS_NetworkSecurityGroup_SubResourceEmbeddedARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SecurityRule_STATUS_NetworkSecurityGroup_SubResourceEmbeddedARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSecurityRuleSTATUSNetworkSecurityGroupSubResourceEmbeddedARM, SecurityRuleSTATUSNetworkSecurityGroupSubResourceEmbeddedARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSecurityRuleSTATUSNetworkSecurityGroupSubResourceEmbeddedARM runs a test to see if a specific instance of SecurityRule_STATUS_NetworkSecurityGroup_SubResourceEmbeddedARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSecurityRuleSTATUSNetworkSecurityGroupSubResourceEmbeddedARM(subject SecurityRule_STATUS_NetworkSecurityGroup_SubResourceEmbeddedARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SecurityRule_STATUS_NetworkSecurityGroup_SubResourceEmbeddedARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of SecurityRule_STATUS_NetworkSecurityGroup_SubResourceEmbeddedARM instances for property testing - lazily
// instantiated by SecurityRuleSTATUSNetworkSecurityGroupSubResourceEmbeddedARMGenerator()
var securityRuleSTATUSNetworkSecurityGroupSubResourceEmbeddedARMGenerator gopter.Gen

// SecurityRuleSTATUSNetworkSecurityGroupSubResourceEmbeddedARMGenerator returns a generator of SecurityRule_STATUS_NetworkSecurityGroup_SubResourceEmbeddedARM instances for property testing.
func SecurityRuleSTATUSNetworkSecurityGroupSubResourceEmbeddedARMGenerator() gopter.Gen {
	if securityRuleSTATUSNetworkSecurityGroupSubResourceEmbeddedARMGenerator != nil {
		return securityRuleSTATUSNetworkSecurityGroupSubResourceEmbeddedARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSecurityRuleSTATUSNetworkSecurityGroupSubResourceEmbeddedARM(generators)
	securityRuleSTATUSNetworkSecurityGroupSubResourceEmbeddedARMGenerator = gen.Struct(reflect.TypeOf(SecurityRule_STATUS_NetworkSecurityGroup_SubResourceEmbeddedARM{}), generators)

	return securityRuleSTATUSNetworkSecurityGroupSubResourceEmbeddedARMGenerator
}

// AddIndependentPropertyGeneratorsForSecurityRuleSTATUSNetworkSecurityGroupSubResourceEmbeddedARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSecurityRuleSTATUSNetworkSecurityGroupSubResourceEmbeddedARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}

func Test_Subnet_STATUS_NetworkSecurityGroup_SubResourceEmbeddedARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Subnet_STATUS_NetworkSecurityGroup_SubResourceEmbeddedARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSubnetSTATUSNetworkSecurityGroupSubResourceEmbeddedARM, SubnetSTATUSNetworkSecurityGroupSubResourceEmbeddedARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSubnetSTATUSNetworkSecurityGroupSubResourceEmbeddedARM runs a test to see if a specific instance of Subnet_STATUS_NetworkSecurityGroup_SubResourceEmbeddedARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSubnetSTATUSNetworkSecurityGroupSubResourceEmbeddedARM(subject Subnet_STATUS_NetworkSecurityGroup_SubResourceEmbeddedARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Subnet_STATUS_NetworkSecurityGroup_SubResourceEmbeddedARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of Subnet_STATUS_NetworkSecurityGroup_SubResourceEmbeddedARM instances for property testing - lazily
// instantiated by SubnetSTATUSNetworkSecurityGroupSubResourceEmbeddedARMGenerator()
var subnetSTATUSNetworkSecurityGroupSubResourceEmbeddedARMGenerator gopter.Gen

// SubnetSTATUSNetworkSecurityGroupSubResourceEmbeddedARMGenerator returns a generator of Subnet_STATUS_NetworkSecurityGroup_SubResourceEmbeddedARM instances for property testing.
func SubnetSTATUSNetworkSecurityGroupSubResourceEmbeddedARMGenerator() gopter.Gen {
	if subnetSTATUSNetworkSecurityGroupSubResourceEmbeddedARMGenerator != nil {
		return subnetSTATUSNetworkSecurityGroupSubResourceEmbeddedARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSubnetSTATUSNetworkSecurityGroupSubResourceEmbeddedARM(generators)
	subnetSTATUSNetworkSecurityGroupSubResourceEmbeddedARMGenerator = gen.Struct(reflect.TypeOf(Subnet_STATUS_NetworkSecurityGroup_SubResourceEmbeddedARM{}), generators)

	return subnetSTATUSNetworkSecurityGroupSubResourceEmbeddedARMGenerator
}

// AddIndependentPropertyGeneratorsForSubnetSTATUSNetworkSecurityGroupSubResourceEmbeddedARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSubnetSTATUSNetworkSecurityGroupSubResourceEmbeddedARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}