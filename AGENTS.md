# workflow-plugin-edge-compute

This repo owns Workflow provider contracts for light/edge WASM compute such as
edge lambda and edge CDN request filtering.

- Keep workflow-compute scheduling, leases, capabilities, proof, and artifact
  semantics in `GoCodeAlone/workflow-compute`.
- Keep generic Workflow dispatch/wait/catalog plumbing in
  `workflow-plugin-compute`.
- This plugin owns only edge provider identities, schemas, and
  `ProviderContract` declarations.
- Do not add product capture, BMW, or application-specific workflow logic.
- Run `GOWORK=off go test ./...` before pushing.
