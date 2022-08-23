// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20180601

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

func Test_Configuration_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Configuration_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForConfigurationSTATUSARM, ConfigurationSTATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForConfigurationSTATUSARM runs a test to see if a specific instance of Configuration_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForConfigurationSTATUSARM(subject Configuration_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Configuration_STATUSARM
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

// Generator of Configuration_STATUSARM instances for property testing - lazily instantiated by
// ConfigurationSTATUSARMGenerator()
var configurationSTATUSARMGenerator gopter.Gen

// ConfigurationSTATUSARMGenerator returns a generator of Configuration_STATUSARM instances for property testing.
// We first initialize configurationSTATUSARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ConfigurationSTATUSARMGenerator() gopter.Gen {
	if configurationSTATUSARMGenerator != nil {
		return configurationSTATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForConfigurationSTATUSARM(generators)
	configurationSTATUSARMGenerator = gen.Struct(reflect.TypeOf(Configuration_STATUSARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForConfigurationSTATUSARM(generators)
	AddRelatedPropertyGeneratorsForConfigurationSTATUSARM(generators)
	configurationSTATUSARMGenerator = gen.Struct(reflect.TypeOf(Configuration_STATUSARM{}), generators)

	return configurationSTATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForConfigurationSTATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForConfigurationSTATUSARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForConfigurationSTATUSARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForConfigurationSTATUSARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(ConfigurationPropertiesSTATUSARMGenerator())
}

func Test_ConfigurationProperties_STATUSARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ConfigurationProperties_STATUSARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForConfigurationPropertiesSTATUSARM, ConfigurationPropertiesSTATUSARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForConfigurationPropertiesSTATUSARM runs a test to see if a specific instance of ConfigurationProperties_STATUSARM round trips to JSON and back losslessly
func RunJSONSerializationTestForConfigurationPropertiesSTATUSARM(subject ConfigurationProperties_STATUSARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ConfigurationProperties_STATUSARM
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

// Generator of ConfigurationProperties_STATUSARM instances for property testing - lazily instantiated by
// ConfigurationPropertiesSTATUSARMGenerator()
var configurationPropertiesSTATUSARMGenerator gopter.Gen

// ConfigurationPropertiesSTATUSARMGenerator returns a generator of ConfigurationProperties_STATUSARM instances for property testing.
func ConfigurationPropertiesSTATUSARMGenerator() gopter.Gen {
	if configurationPropertiesSTATUSARMGenerator != nil {
		return configurationPropertiesSTATUSARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForConfigurationPropertiesSTATUSARM(generators)
	configurationPropertiesSTATUSARMGenerator = gen.Struct(reflect.TypeOf(ConfigurationProperties_STATUSARM{}), generators)

	return configurationPropertiesSTATUSARMGenerator
}

// AddIndependentPropertyGeneratorsForConfigurationPropertiesSTATUSARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForConfigurationPropertiesSTATUSARM(gens map[string]gopter.Gen) {
	gens["AllowedValues"] = gen.PtrOf(gen.AlphaString())
	gens["DataType"] = gen.PtrOf(gen.AlphaString())
	gens["DefaultValue"] = gen.PtrOf(gen.AlphaString())
	gens["Description"] = gen.PtrOf(gen.AlphaString())
	gens["Source"] = gen.PtrOf(gen.AlphaString())
	gens["Value"] = gen.PtrOf(gen.AlphaString())
}