## Example Templates

This folder contains working HCL2 examples for all supported operation types.

Each example requires the following environment variables to be set before running:

```shell
export PKR_VAR_elm_password="<your-elm-password>"
export PKR_VAR_vm_password="<your-ce-vm-password>"
```

Update `elm_server` and the layer/version names in each template to match your environment before running.

### Available Examples

| File | Operation |
|------|-----------|
| `create_app_finalize.pkr.hcl` | Create a new app layer |
| `create_app_no_platform_finalize.pkr.hcl` | Create a new app layer without a platform layer |
| `create_platform_finalize.pkr.hcl` | Create a new platform layer |
| `revision_app_finalize.pkr.hcl` | Revision an existing app layer |
| `revision_os_finalize.pkr.hcl` | Revision an existing OS layer |
| `revision_platform_finalize.pkr.hcl` | Revision an existing platform layer |
| `all_final.pkr.hcl` | All operation types in a single file |
| `all_runonce_final.pkr.hcl` | All operation types with `handle_run_once = true` |

### Running an Example

```shell
packer init .
packer build create_app_finalize.pkr.hcl
```
