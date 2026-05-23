package main

import (
	"github.com/GoCodeAlone/workflow-plugin-edge-compute/internal/plugin"
	sdk "github.com/GoCodeAlone/workflow/plugin/external/sdk"
)

func main() {
	sdk.Serve(plugin.NewPlugin())
}
