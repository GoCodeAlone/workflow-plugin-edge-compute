package plugin

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/GoCodeAlone/workflow-compute/pkg/protocol"
	sdk "github.com/GoCodeAlone/workflow/plugin/external/sdk"
)

func TestManifestAndNoWorkflowSteps(t *testing.T) {
	manifest := NewPlugin().Manifest()
	if manifest.Name != "workflow-plugin-edge-compute" {
		t.Fatalf("manifest name: %q", manifest.Name)
	}
	if steps := NewPlugin().(interface{ StepTypes() []string }).StepTypes(); len(steps) != 0 {
		t.Fatalf("edge provider plugin should not own workflow steps yet: %v", steps)
	}
}

func TestProviderContractsAreStrictWASMEdgeContracts(t *testing.T) {
	contracts := ProviderContracts()
	if len(contracts) != 2 {
		t.Fatalf("contract count: got %d", len(contracts))
	}
	seen := map[string]bool{}
	for _, contract := range contracts {
		if err := contract.Validate(); err != nil {
			t.Fatalf("contract %s invalid: %v", contract.ID, err)
		}
		seen[contract.ProviderID] = true
		if contract.PluginID != "workflow-plugin-edge-compute" {
			t.Fatalf("wrong plugin boundary: %+v", contract)
		}
		text := strings.ToLower(contract.ID + contract.PluginID + contract.ProviderID + contract.ContractID)
		if strings.Contains(text, "product") || strings.Contains(text, "capture") || strings.Contains(text, "bmw") {
			t.Fatalf("edge contract leaked product/BMW domain: %+v", contract)
		}
		if len(contract.RuntimeContract.Profiles) != 1 {
			t.Fatalf("runtime profiles: %+v", contract.RuntimeContract.Profiles)
		}
		runtime := contract.RuntimeContract.Profiles[0]
		if runtime.RuntimeProfile != protocol.RuntimeProfileWASMComponent ||
			runtime.ExecutionSecurityTier != protocol.ExecutionWASMCapability ||
			runtime.ExecutorProvider != "wasm-component" ||
			runtime.WASM.ABI != "wasm-export-i32-v1" ||
			runtime.WASM.ComponentDigest == "" ||
			runtime.WASM.Filesystem != "forbidden" ||
			runtime.WASM.NativeHostUpdates != "forbidden" {
			t.Fatalf("runtime contract: %+v", runtime)
		}
	}
	if !seen["edge-lambda"] || !seen["edge-cdn-filter"] {
		t.Fatalf("providers missing: %+v", seen)
	}
}

func TestPluginManifestExposesContractFiles(t *testing.T) {
	data, err := os.ReadFile("../../plugin.json")
	if err != nil {
		t.Fatalf("read plugin.json: %v", err)
	}
	var manifest struct {
		Capabilities struct {
			StepTypes []string `json:"stepTypes"`
		} `json:"capabilities"`
		Contracts []struct {
			ID     string `json:"id"`
			Path   string `json:"path"`
			Schema string `json:"schema"`
		} `json:"contracts"`
	}
	if err := json.Unmarshal(data, &manifest); err != nil {
		t.Fatalf("decode plugin.json: %v", err)
	}
	if len(manifest.Capabilities.StepTypes) != 0 {
		t.Fatalf("step types belong outside this provider contract plugin: %v", manifest.Capabilities.StepTypes)
	}
	if len(manifest.Contracts) != 2 {
		t.Fatalf("manifest contracts: %+v", manifest.Contracts)
	}
	for _, contract := range manifest.Contracts {
		if _, err := os.Stat(filepath.Join("../..", contract.Path)); err != nil {
			t.Fatalf("contract path %q: %v", contract.Path, err)
		}
		if _, err := os.Stat(filepath.Join("../..", contract.Schema)); err != nil {
			t.Fatalf("schema path %q: %v", contract.Schema, err)
		}
	}
}

func TestProviderContractsImplementExternalSDKPlugin(t *testing.T) {
	if _, ok := NewPlugin().(sdk.PluginProvider); !ok {
		t.Fatal("plugin must implement sdk.PluginProvider")
	}
}
