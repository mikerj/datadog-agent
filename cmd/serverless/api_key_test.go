// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016-present Datadog, Inc.

package main

import (
	"bytes"
	"errors"
	"os"
	"testing"

	"github.com/aws/aws-sdk-go/service/kms"
	"github.com/aws/aws-sdk-go/service/kms/kmsiface"
	"github.com/stretchr/testify/assert"
)

// mockEncryptedAPIKeyBase64 represents an API key encrypted with KMS and encoded as a base64 string
const mockEncryptedAPIKeyBase64 = "MjIyMjIyMjIyMjIyMjIyMg=="

// mockEncryptedAPIKey represents the encrypted API key after it has been decoded from base64
const mockDecodedEncryptedAPIKey = "2222222222222222"

// expectedDecryptedAPIKey represents the true value of the API key after decryption by KMS
const expectedDecryptedAPIKey = "1111111111111111"

// mockSecretsManagerAPIKeyArn represents a SecretsManager Arn passed in via DD_API_KEY_SECRET_ARN environment variable
const mockSecretsManagerAPIKeyArn string = "arn:aws:secretsmanager:us-west-2:123456789012:secret:DatadogAppKeySecret"

// mockMalformedSecretsManagerAPIKeyArn represents a SecretsManager Arn formatted incorrectly that doesn't match regex.
const mockMalformedSecretsManagerAPIKeyArn string = "arn:aws:secretsmanager:uswest2:123456789012:secret:DatadogAppKeySecret"

// mockFunctionName represents the name of the current function
var mockFunctionName = "my-Function"

type mockKMSClientWithEncryptionContext struct {
	kmsiface.KMSAPI
}

func (mockKMSClientWithEncryptionContext) Decrypt(params *kms.DecryptInput) (*kms.DecryptOutput, error) {
	encryptionContextPointer, exists := params.EncryptionContext[encryptionContextKey]
	if !exists {
		return nil, errors.New("InvalidCiphertextException")
	}
	if *encryptionContextPointer != mockFunctionName {
		return nil, errors.New("InvalidCiphertextException")
	}
	if bytes.Equal(params.CiphertextBlob, []byte(mockDecodedEncryptedAPIKey)) {
		return &kms.DecryptOutput{
			Plaintext: []byte(expectedDecryptedAPIKey),
		}, nil
	}
	return nil, errors.New("KMS error")
}

type mockKMSClientNoEncryptionContext struct {
	kmsiface.KMSAPI
}

func (mockKMSClientNoEncryptionContext) Decrypt(params *kms.DecryptInput) (*kms.DecryptOutput, error) {
	if params.EncryptionContext[encryptionContextKey] != nil {
		return nil, errors.New("InvalidCiphertextException")
	}
	if bytes.Equal(params.CiphertextBlob, []byte(mockDecodedEncryptedAPIKey)) {
		return &kms.DecryptOutput{
			Plaintext: []byte(expectedDecryptedAPIKey),
		}, nil
	}
	return nil, errors.New("KMS error")
}

func TestDecryptKMSWithEncryptionContext(t *testing.T) {
	os.Setenv(functionNameEnvVar, mockFunctionName)
	defer os.Setenv(functionNameEnvVar, "")

	client := mockKMSClientWithEncryptionContext{}
	result, _ := decryptKMS(client, mockEncryptedAPIKeyBase64)
	assert.Equal(t, expectedDecryptedAPIKey, result)
}

func TestDecryptKMSNoEncryptionContext(t *testing.T) {
	client := mockKMSClientNoEncryptionContext{}
	result, _ := decryptKMS(client, mockEncryptedAPIKeyBase64)
	assert.Equal(t, expectedDecryptedAPIKey, result)
}

func TestExtractRegionFromSecretsManagerArn(t *testing.T) {
	result, _ := extractRegionFromSecretsManagerArn(mockSecretsManagerAPIKeyArn)
	assert.Equal(t, result, "us-west-2")
}

func TestExtractRegionFromMalformedSecretsManagerArn(t *testing.T) {
	result, err := extractRegionFromSecretsManagerArn(mockMalformedSecretsManagerAPIKeyArn)
	assert.Equal(t, result, "")
	assert.Error(t, err, "Couldn't extract region from arn: arn:aws:secretsmanager:uswest2:123456789012:secret:DatadogAppKeySecret")
}
