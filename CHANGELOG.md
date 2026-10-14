# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [1.8.0]

### Breaking

#### Stack client methods

- `stacks.Client.ReadStack` takes a new `*ReadStackQueryParams` argument after `wfGrp`. Pass `nil`
  to keep the previous behaviour:

  ```go
  client.Stacks.ReadStack(ctx, org, stack, wfGrp, nil)
  ```

- `stacks.Client.DeleteStack` takes a new `*DeleteStackRequest` argument after `wfGrp`. Pass `nil`
  to keep the previous behaviour.
- `stacks.Client.UpdateStack` now returns `*GeneratedStackUpdateResponse` instead of
  `*GeneratedStackCreateResponse`.

#### Stack request types

- `Stack` (the create request) uses plain types instead of `*core.Optional[T]` for `ResourceName`,
  `Description`, `Tags`, `TemplateGroupId`, `Actions`, `ContextTags`, `UserSchedules`,
  `EnvironmentVariables` and `Operations`. A create request never needs to send an explicit `null`,
  so leaving a field `nil` is enough to omit it:

  ```go
  // Before
  Description: sgsdkgo.Optional("my stack"),
  // After
  Description: sgsdkgo.String("my stack"),
  ```

  `PatchedStack` (the update request) still uses `*core.Optional[T]`, so an update can still clear a
  field with `sgsdkgo.Null[T]()`.
- `Stack.WorkflowsConfig` is now `*StackWorkflowsConfig` and `PatchedStack.WorkflowsConfig` is now
  `*core.Optional[StackWorkflowsConfig]` (both used `WorkflowsConfig`).
- `TemplatesConfig` is removed from `Stack` and `PatchedStack`.
- Stacks use their own types for deployment platform config and mini steps, so the shared
  `DeploymentPlatformConfig`, `MiniStepsSchema` and `WebhookTypes` are unchanged for other clients:

  | Field | Now |
  | --- | --- |
  | `Stack.DeploymentPlatformConfig` | `[]*StackDeploymentPlatformConfig` |
  | `Stack.MiniSteps` | `*StackMiniStepsSchema` |
  | `PatchedStack.DeploymentPlatformConfig` | `*core.Optional[[]*StackDeploymentPlatformConfig]` |
  | `PatchedStack.MiniSteps` | `*core.Optional[StackMiniStepsSchema]` |
  | `StackActionParameters.DeploymentPlatformConfig` | `[]*StackDeploymentPlatformConfig` |
  | `stacktemplaterevisions.StackTemplateRevisionWorkflow.DeploymentPlatformConfig` | `[]*sgsdkgo.StackDeploymentPlatformConfig` |
  | `stacktemplaterevisions.StackTemplateRevisionWorkflow.MiniSteps` | `*sgsdkgo.StackMiniStepsSchema` |

  `StackDeploymentPlatformConfig` has `Kind *DeploymentPlatformConfigKindEnum` and
  `Config *DeploymentPlatformConfigConfig` (`IntegrationId`, `ProfileName`). `StackMiniStepsSchema`
  is `MiniStepsSchema` with `Webhooks *StackWebhookTypes`, whose lists hold typed `*Webhook` entries
  instead of `map[string]interface{}`.

  `Actions` is also used by `Template`, `StackTemplate` and `StackDataResponse` (the `templates`
  client), so `Actions[...].Order[...].Parameters.DeploymentPlatformConfig` on those types is
  `[]*StackDeploymentPlatformConfig` as well.

#### Stack response types

- `GeneratedStackGetResponse.Msg` (read) and `GeneratedStackCreateResponseData.Stack` (create) are
  now `*StackData`. They were `*GeneratedStackGetResponseMsg` and
  `*GeneratedStackCreateResponseDataStack`.
- `GeneratedStackCreateResponseDataWorkflows.EnvironmentVariables` is now `interface{}` (it was
  `[]interface{}`).

#### Input schemas

- `InputSchemas` no longer has `Id`, `Description` or `IsCommitted`; these fields don't exist in the
  API. This affects stack and workflow template revisions and templates.

### Added

- `ReadStackQueryParams.RefreshWorkflowsConfig`: sends `refresh_workflows_config=true`, so the
  stack's `WorkflowsConfig` is rebuilt from its current workflows instead of the stored copy.
- `DeleteStackRequest.ForceDelete`: sends `force_delete=true`, which deletes the stack together with
  the workflows in it.
- `StackData`, `StackWorkflowsConfig` and `StackWorkflowsConfigWorkflow`, including
  `StackWorkflowsConfigWorkflow.WorkflowId`.
- `StackDeploymentPlatformConfig`, `StackMiniStepsSchema`, `StackWebhookTypes`, `Webhook` and
  `DeploymentPlatformConfigConfig`.
- `PatchedStack.UpgradeMode` (`UpgradeModeEnum`: `PRESERVE_SETTINGS`, `RESET_TO_TEMPLATE`),
  `PatchedStack.UpdateWorkflowsFromConfig`, `PatchedStack.Id` and `PatchedStack.IsActive`.
- `Stack.IsActive` and `Stack.ReplaceMovedResource`.
- `ListAllStacksRequest.TemplateGroupId` and `ListAllStacksRequest.ParentTemplateGroupId` filters.
- `GeneratedStackUpdateResponse`.
- `DeploymentPlatformConfigKindEnumAzureManagedIdOidc`.
- `WfType` and `ParallelExecution` on `stacktemplaterevisions.StackTemplateRevisionWorkflow` and on
  `workflowtemplaterevisions.CreateWorkflowTemplateRevisionsRequest` and
  `UpdateWorkflowTemplateRevisionRequest`; `ParallelExecution` on `WorkflowsConfigWorkflow`.

### Changed

- Request payloads: slices and maps on the stack, stack template and stack template revision types
  are now tagged `omitzero` instead of `omitempty`. A `nil` slice or map is still left out, but an
  empty non-nil one (`[]T{}`, `map[K]V{}`) is now sent as `[]` / `{}` instead of being dropped. Set
  the field to `nil` to leave it out.

## [1.7.0]

### Breaking

- `workflowtemplaterevisions.UserSchedules.Inputs` is now `*UserSchedulesInputs` (it was an
  anonymous struct value). A schedule with no inputs now has `Inputs == nil`, so code that reads
  `schedule.Inputs.<Field>` panics unless it checks for `nil` first:

  ```go
  if schedule.Inputs != nil && schedule.Inputs.TerraformAction != nil {
      // ...
  }
  ```

  Code that built the old anonymous struct with a literal no longer compiles; use
  `&workflowtemplaterevisions.UserSchedulesInputs{...}` instead.
- `UserSchedulesInputs.Ministeps` is now `*Ministeps` (it was a struct value). Check it for `nil`
  before reading its fields, like `Inputs`.

### Changed

- The inputs of a user schedule are now the named type `workflowtemplaterevisions.UserSchedulesInputs`,
  with the same fields as before.
- Request payloads: a schedule whose `Inputs` is `nil` is sent without an `inputs` key, and inputs
  whose `Ministeps` is `nil` are sent without a `MiniSteps` key. Previously every schedule was sent
  with at least `"inputs": {"MiniSteps": {}}`. Setting `Inputs` or `Ministeps` to a pointer to an
  empty struct sends an explicit `{}`.
