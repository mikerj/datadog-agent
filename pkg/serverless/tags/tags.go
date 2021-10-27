// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016-present Datadog, Inc.

package tags

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/DataDog/datadog-agent/pkg/util/log"
)

const (
	qualifierEnvVar = "AWS_LAMBDA_FUNCTION_VERSION"
	envEnvVar       = "DD_ENV"
	versionEnvVar   = "DD_VERSION"
	serviceEnvVar   = "DD_SERVICE"
	runtimeVar      = "AWS_EXECUTION_ENV"
	memorySizeVar   = "AWS_LAMBDA_FUNCTION_MEMORY_SIZE"

	bootstrapPath = "/var/runtime/bootstrap"

	traceOriginMetadataKey   = "_dd.origin"
	traceOriginMetadataValue = "lambda"
	computeStatsKey          = "_dd.compute_stats"
	computeStatsValue        = "1"
	functionARNKey           = "function_arn"
	functionNameKey          = "functionname"
	regionKey                = "region"
	accountIDKey             = "account_id"
	awsAccountKey            = "aws_account"
	resourceKey              = "resource"
	executedVersionKey       = "executedversion"
	extensionVersionKey      = "dd_extension_version"
	envKey                   = "env"
	versionKey               = "version"
	serviceKey               = "service"
	runtimeKey               = "runtime"
	memorySizeKey            = "memorysize"
	architectureKey          = "architecture"
)

// currentExtensionVersion represents the current version of the Datadog Lambda Extension.
// It is applied to all telemetry as a tag.
// It is replaced at build time with an actual version number.
var currentExtensionVersion = "xxx"

// BuildTagMap builds a map of tag based on the arn and user defined tags
func BuildTagMap(arn string, configTags []string) map[string]string {
	tags := make(map[string]string)

	architecture := ResolveRuntimeArch()
	tags = setIfNotEmpty(tags, architectureKey, architecture)

	runtime, err := detectRuntime(bootstrapPath, runtimeVar)
	if err != nil {
		log.Debug(err)
		runtime = "custom"
	}

	tags = setIfNotEmpty(tags, runtimeKey, runtime)
	tags = setIfNotEmpty(tags, memorySizeKey, os.Getenv(memorySizeVar))

	tags = setIfNotEmpty(tags, envKey, os.Getenv(envEnvVar))
	tags = setIfNotEmpty(tags, versionKey, os.Getenv(versionEnvVar))
	tags = setIfNotEmpty(tags, serviceKey, os.Getenv(serviceEnvVar))

	for _, tag := range configTags {
		splitTags := strings.Split(tag, ",")
		for _, singleTag := range splitTags {
			tags = addTag(tags, singleTag)
		}
	}

	tags = setIfNotEmpty(tags, traceOriginMetadataKey, traceOriginMetadataValue)
	tags = setIfNotEmpty(tags, computeStatsKey, computeStatsValue)
	tags = setIfNotEmpty(tags, functionARNKey, arn)
	tags = setIfNotEmpty(tags, extensionVersionKey, currentExtensionVersion)

	parts := strings.Split(arn, ":")
	if len(parts) < 6 {
		return tags
	}

	tags = setIfNotEmpty(tags, regionKey, parts[3])
	tags = setIfNotEmpty(tags, awsAccountKey, parts[4])
	tags = setIfNotEmpty(tags, accountIDKey, parts[4])
	tags = setIfNotEmpty(tags, functionNameKey, parts[6])
	tags = setIfNotEmpty(tags, resourceKey, parts[6])

	qualifier := os.Getenv(qualifierEnvVar)
	if len(qualifier) > 0 {
		if qualifier != "$LATEST" {
			tags = setIfNotEmpty(tags, resourceKey, fmt.Sprintf("%s:%s", parts[6], qualifier))
			tags = setIfNotEmpty(tags, executedVersionKey, qualifier)
		}
	}

	return tags
}

// BuildTagsFromMap builds an array of tag based on map of tags
func BuildTagsFromMap(tags map[string]string) []string {
	tagsMap := make(map[string]string)
	tagBlackList := []string{traceOriginMetadataKey, computeStatsKey}
	for k, v := range tags {
		tagsMap[k] = v
	}
	for _, blackListKey := range tagBlackList {
		delete(tagsMap, blackListKey)
	}
	tagsArray := make([]string, 0, len(tagsMap))
	for key, value := range tagsMap {
		tagsArray = append(tagsArray, fmt.Sprintf("%s:%s", key, value))
	}
	return tagsArray
}

// BuildTracerTags builds a map of tag from an existing map of tag removing useless tags for traces
func BuildTracerTags(tags map[string]string) map[string]string {
	tagsMap := make(map[string]string)
	tagBlackList := []string{resourceKey}
	for k, v := range tags {
		tagsMap[k] = v
	}
	for _, blackListKey := range tagBlackList {
		delete(tagsMap, blackListKey)
	}
	return tagsMap
}

// AddColdStartTag appends the cold_start tag to existing tags
func AddColdStartTag(tags []string, coldStart bool) []string {
	tags = append(tags, fmt.Sprintf("cold_start:%v", coldStart))
	return tags
}

func setIfNotEmpty(tagMap map[string]string, key string, value string) map[string]string {
	if key != "" && value != "" {
		tagMap[key] = strings.ToLower(value)
	}
	return tagMap
}

func addTag(tagMap map[string]string, tag string) map[string]string {
	extract := strings.Split(tag, ":")
	if len(extract) == 2 {
		tagMap[strings.ToLower(extract[0])] = strings.ToLower(extract[1])
	}
	return tagMap
}

func detectRuntime(path string, runtimeVarName string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", fmt.Errorf("could not detect the runtime: bootstrap file cannot be read")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		currentLine := scanner.Text()
		if strings.Contains(currentLine, runtimeVarName) {
			return extractRuntimeFromLine(currentLine)
		}
	}
	return "", fmt.Errorf("could not detect the runtime: cannot find the runtime value in the bootstrap file")
}

func extractRuntimeFromLine(line string) (string, error) {
	parts := strings.Split(line, "=")
	if len(parts) == 2 {
		prefixedRuntime := strings.TrimSpace(parts[1])
		runtime := strings.Replace(prefixedRuntime, "AWS_Lambda_", "", 1)
		if runtime != prefixedRuntime {
			return runtime, nil
		}
	}
	return "", fmt.Errorf("could not detect the runtime: invalid format found")
}
