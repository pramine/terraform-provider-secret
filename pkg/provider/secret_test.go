package provider

import (
	"testing"

	"github.com/hashicorp/terraform/helper/schema"

	"github.com/hashicorp/terraform/config"
	"github.com/hashicorp/terraform/terraform"
)

var testProvider *schema.Provider
var testProviders map[string]terraform.ResourceProvider

func init() {
	testProvider = Secret()
	testProviders = map[string]terraform.ResourceProvider{
		"secret": testProvider,
	}
}

func TestProvider(t *testing.T) {
	if err := Secret().InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestBackendKmsConfig(t *testing.T) {
	cfg, err := config.LoadFile("test-fixtures/backend_kms.tf")
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	if len(cfg.ProviderConfigs) != 1 {
		t.Fatalf("error loading provider config")
	}
}

func TestProvider_impl(t *testing.T) {
	var _ terraform.ResourceProvider = Secret()
}

func TestValidateBackendType(t *testing.T) {
	validBackends := []string{"kms", "gpg"}
	for _, v := range validBackends {
		_, err := validateBackendType(v, "backend")
		if err != nil {
			t.Errorf("unexpected err: %s", err)
		}
	}
	_, err := validateBackendType("invalid", "backend")
	if err == nil {
		t.Errorf("expected error for backend type 'invalid', got none")
	}

}
