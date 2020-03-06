# PhaseDto

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApprovalGate** | **string** |  | [optional] [default to null]
**Conflicts** | [**[]ConflictDto**](ConflictDto.md) |  | [optional] [default to null]
**Description** | **string** |  | [optional] [default to null]
**Environments** | [**[]EnvironmentDto**](EnvironmentDto.md) |  | [optional] [default to null]
**ExecutionData** | [***PhaseExecutionDto**](PhaseExecutionDto.md) |  | [optional] [default to null]
**FreezePeriods** | [**[]FreezePeriodDto**](FreezePeriodDto.md) |  | [optional] [default to null]
**Id** | **int64** |  | [optional] [default to null]
**IsDisabled** | **bool** |  | [optional] [default to null]
**IsNeedApproval** | **bool** |  | [optional] [default to null]
**IsProduction** | **bool** |  | [optional] [default to null]
**MaintenanceWindows** | [**[]MaintenanceWindowDto**](MaintenanceWindowDto.md) |  | [optional] [default to null]
**MilestonePhaseRelations** | [**[]MilestonePhaseRelationDto**](MilestonePhaseRelationDto.md) |  | [optional] [default to null]
**Name** | **string** |  | [optional] [default to null]
**NextPhase** | [***IdentifiableDto**](IdentifiableDto.md) |  | [optional] [default to null]
**OnFailureFailedPhase** | [***IdentifiableDto**](IdentifiableDto.md) |  | [optional] [default to null]
**OwnerParties** | [**[]PartyDto**](PartyDto.md) |  | [optional] [default to null]
**PercentageCompleted** | **float32** |  | [optional] [default to null]
**PlannedEndDate** | **int64** |  | [optional] [default to null]
**PlannedStartDate** | **int64** |  | [optional] [default to null]
**PreviousPhase** | [***IdentifiableDto**](IdentifiableDto.md) |  | [optional] [default to null]
**Recurrence** | [***RecurrenceDto**](RecurrenceDto.md) |  | [optional] [default to null]
**Release** | [***ReleaseDto**](ReleaseDto.md) |  | [optional] [default to null]
**ReleaseId** | **int64** |  | [optional] [default to null]
**SkipTasksWithUnchangedApplications** | **bool** |  | [optional] [default to null]
**Tasks** | [**[]TaskDto**](TaskDto.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


