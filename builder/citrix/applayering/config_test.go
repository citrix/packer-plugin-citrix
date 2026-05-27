package applayering

import (
	"testing"
)

func newTestBuilder() *Builder {
	return &Builder{StrategyFactory: NewBuilderStrategy}
}

func TestBuilder_Prepare_Valid(t *testing.T) {
	raw := map[string]any{
		"elm_server":          "http://elm.example.com",
		"elm_username":        "admin",
		"elm_password":        "secret",
		"insecure_connection": true,
		"wait_for_ip_timeout": 30,
		"create_platform_layer": map[string]any{
			"hypervisor_platform":            "vmware",
			"provisioning_platform":          "MCS",
			"broker_platform":                "None",
			"os_layer_name":                  "Win10",
			"os_layer_version_name":          "v1",
			"layer_name":                     "PlatformLayer",
			"packaging_disk_file_name":       "disk.vhd",
			"platform_connector_config_name": "Connector",
			"icon_id":                        int64(456),
			"version_name":                   "RevInfo",
			"version_description":            "Revision Description",
			"version_size_gb":                int32(10),
			"comment":                        "Reason",
			"skip_cleanup_on_failure":        false,
		},
		"communicator":   "winrm",
		"winrm_username": "Administrator",
		"winrm_password": "vm_password",
	}

	b := newTestBuilder()
	_, _, err := b.Prepare(raw)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	if b.config.ELMServer != "http://elm.example.com" {
		t.Errorf("expected ELMServer to be set, got: %v", b.config.ELMServer)
	}
	if b.config.CreatePlatform == nil {
		t.Fatalf("expected CreatePlatform to be set")
	}
	if b.config.CreatePlatform.HypervisorPlatform != "vmware" {
		t.Errorf("expected HypervisorPlatform=vmware, got: %v", b.config.CreatePlatform.HypervisorPlatform)
	}
	if b.config.CreatePlatform.LayerName != "PlatformLayer" {
		t.Errorf("expected LayerName=PlatformLayer, got: %v", b.config.CreatePlatform.LayerName)
	}
}

func TestBuilder_Prepare_Invalid(t *testing.T) {
	raw := map[string]any{
		// Missing required elm_username/elm_password and no operation block
		"elm_server": int32(789),
	}

	b := newTestBuilder()
	_, _, err := b.Prepare(raw)
	if err == nil {
		t.Fatalf("expected error for invalid config, got nil")
	}
}

func TestBuilder_Prepare_NoOperationBlock(t *testing.T) {
	raw := map[string]any{
		"elm_server":          "http://elm.example.com",
		"elm_username":        "admin",
		"elm_password":        "secret",
		"insecure_connection": true,
		"communicator":        "winrm",
		"winrm_username":      "Administrator",
		"winrm_password":      "vm_password",
		// no operation block
	}

	b := newTestBuilder()
	_, _, err := b.Prepare(raw)
	if err == nil {
		t.Fatalf("expected error for missing operation block, got nil")
	}
}
