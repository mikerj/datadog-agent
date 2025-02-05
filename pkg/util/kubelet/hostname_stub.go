// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016-present Datadog, Inc.

//go:build !kubelet
// +build !kubelet

package kubelet

import (
	"context"
	"fmt"
)

// GetHostname builds a hostname from the kubernetes nodename and an optional cluster-name
func GetHostname(ctx context.Context) (string, error) {
	return "", fmt.Errorf("kubelet hostname provider is not enabled")
}

// IsAgentKubeHostNetwork returns true if the agent is running on a POD with hostNetwork
func IsAgentKubeHostNetwork() (bool, error) {
	return true, nil
}
