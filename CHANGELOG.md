# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

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
