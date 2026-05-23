# workflow-plugin-edge-compute

Workflow provider contract plugin for light/edge WASM workloads on
`workflow-compute`.

This plugin proves the provider boundary for edge compute:

- `workflow-compute` owns execution, routing, browser-worker/edge capability
  profiles, artifact proof, and scheduling.
- `workflow-plugin-compute` owns generic Workflow dispatch/wait/catalog plumbing.
- `workflow-plugin-edge-compute` owns edge provider identities and
  `ProviderContract` records for edge lambda and edge CDN filter workloads.

Both providers require digest-bound WASM components, the `wasm-component`
executor, the `wasm-capability` security tier, artifact-hash proof, forbidden
filesystem/native host updates, and relay/offline network compatibility.

## Development

```sh
GOWORK=off go test ./...
```
