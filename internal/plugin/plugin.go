package plugin

import (
	"fmt"

	sdk "github.com/GoCodeAlone/workflow/plugin/external/sdk"
)

var Version = "0.0.0"

type edgeComputePlugin struct{}

func NewPlugin() sdk.PluginProvider {
	return edgeComputePlugin{}
}

func (edgeComputePlugin) Manifest() sdk.PluginManifest {
	return sdk.PluginManifest{
		Name:        "workflow-plugin-edge-compute",
		Version:     Version,
		Author:      "GoCodeAlone",
		Description: "Edge WASM provider contracts for workflow-compute",
	}
}

func (edgeComputePlugin) StepTypes() []string {
	return nil
}

func (edgeComputePlugin) CreateStep(typeName, name string, config map[string]any) (sdk.StepInstance, error) {
	return nil, fmt.Errorf("edge-compute plugin: unknown step type %q", typeName)
}
