package envprops

import (
	"context"
	"encoding/json"
	"fmt"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	secretmanagerpb "google.golang.org/genproto/googleapis/cloud/secretmanager/v1"
)

// MustReadSecret is similar to ReadSecret, but panics if it errors
func MustReadSecret(secretName, secretVersion, projectID string) map[string]interface{} {
	secret, err := ReadSecret(secretName, secretVersion, projectID)
	if err != nil {
		panic(err)
	}
	return secret
}

// ReadSecret from the KMS. The secret must be a JSON object
func ReadSecret(secretName, secretVersion, projectID string) (map[string]interface{}, error) {
	ctx := context.Background()
	client, err := secretmanager.NewClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to setup client: %w", err)
	}

	secretQualifiedName := fmt.Sprintf("projects/%s/secrets/%s/versions/%s", projectID, secretName, secretVersion)
	accessRequest := &secretmanagerpb.AccessSecretVersionRequest{
		Name: secretQualifiedName,
	}

	secretResult, err := client.AccessSecretVersion(ctx, accessRequest)
	if err != nil {
		return nil, fmt.Errorf("failed to access secret version: %w", err)
	}

	result := make(map[string]interface{})
	err = json.Unmarshal(secretResult.Payload.Data, &result)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling secret result. is the secret a JSON object?: %w", err)
	}
	return result, nil
}
