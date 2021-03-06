// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package eks

type AMITypes string

// Enum values for AMITypes
const (
	AMITypesAl2X8664    AMITypes = "AL2_x86_64"
	AMITypesAl2X8664Gpu AMITypes = "AL2_x86_64_GPU"
)

func (enum AMITypes) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum AMITypes) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ClusterStatus string

// Enum values for ClusterStatus
const (
	ClusterStatusCreating ClusterStatus = "CREATING"
	ClusterStatusActive   ClusterStatus = "ACTIVE"
	ClusterStatusDeleting ClusterStatus = "DELETING"
	ClusterStatusFailed   ClusterStatus = "FAILED"
	ClusterStatusUpdating ClusterStatus = "UPDATING"
)

func (enum ClusterStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ClusterStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ErrorCode string

// Enum values for ErrorCode
const (
	ErrorCodeSubnetNotFound            ErrorCode = "SubnetNotFound"
	ErrorCodeSecurityGroupNotFound     ErrorCode = "SecurityGroupNotFound"
	ErrorCodeEniLimitReached           ErrorCode = "EniLimitReached"
	ErrorCodeIpNotAvailable            ErrorCode = "IpNotAvailable"
	ErrorCodeAccessDenied              ErrorCode = "AccessDenied"
	ErrorCodeOperationNotPermitted     ErrorCode = "OperationNotPermitted"
	ErrorCodeVpcIdNotFound             ErrorCode = "VpcIdNotFound"
	ErrorCodeUnknown                   ErrorCode = "Unknown"
	ErrorCodeNodeCreationFailure       ErrorCode = "NodeCreationFailure"
	ErrorCodePodEvictionFailure        ErrorCode = "PodEvictionFailure"
	ErrorCodeInsufficientFreeAddresses ErrorCode = "InsufficientFreeAddresses"
)

func (enum ErrorCode) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ErrorCode) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type FargateProfileStatus string

// Enum values for FargateProfileStatus
const (
	FargateProfileStatusCreating     FargateProfileStatus = "CREATING"
	FargateProfileStatusActive       FargateProfileStatus = "ACTIVE"
	FargateProfileStatusDeleting     FargateProfileStatus = "DELETING"
	FargateProfileStatusCreateFailed FargateProfileStatus = "CREATE_FAILED"
	FargateProfileStatusDeleteFailed FargateProfileStatus = "DELETE_FAILED"
)

func (enum FargateProfileStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum FargateProfileStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type LogType string

// Enum values for LogType
const (
	LogTypeApi               LogType = "api"
	LogTypeAudit             LogType = "audit"
	LogTypeAuthenticator     LogType = "authenticator"
	LogTypeControllerManager LogType = "controllerManager"
	LogTypeScheduler         LogType = "scheduler"
)

func (enum LogType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum LogType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type NodegroupIssueCode string

// Enum values for NodegroupIssueCode
const (
	NodegroupIssueCodeAutoScalingGroupNotFound         NodegroupIssueCode = "AutoScalingGroupNotFound"
	NodegroupIssueCodeEc2securityGroupNotFound         NodegroupIssueCode = "Ec2SecurityGroupNotFound"
	NodegroupIssueCodeEc2securityGroupDeletionFailure  NodegroupIssueCode = "Ec2SecurityGroupDeletionFailure"
	NodegroupIssueCodeEc2launchTemplateNotFound        NodegroupIssueCode = "Ec2LaunchTemplateNotFound"
	NodegroupIssueCodeEc2launchTemplateVersionMismatch NodegroupIssueCode = "Ec2LaunchTemplateVersionMismatch"
	NodegroupIssueCodeIamInstanceProfileNotFound       NodegroupIssueCode = "IamInstanceProfileNotFound"
	NodegroupIssueCodeIamNodeRoleNotFound              NodegroupIssueCode = "IamNodeRoleNotFound"
	NodegroupIssueCodeAsgInstanceLaunchFailures        NodegroupIssueCode = "AsgInstanceLaunchFailures"
	NodegroupIssueCodeInstanceLimitExceeded            NodegroupIssueCode = "InstanceLimitExceeded"
	NodegroupIssueCodeInsufficientFreeAddresses        NodegroupIssueCode = "InsufficientFreeAddresses"
	NodegroupIssueCodeAccessDenied                     NodegroupIssueCode = "AccessDenied"
	NodegroupIssueCodeInternalFailure                  NodegroupIssueCode = "InternalFailure"
)

func (enum NodegroupIssueCode) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum NodegroupIssueCode) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type NodegroupStatus string

// Enum values for NodegroupStatus
const (
	NodegroupStatusCreating     NodegroupStatus = "CREATING"
	NodegroupStatusActive       NodegroupStatus = "ACTIVE"
	NodegroupStatusUpdating     NodegroupStatus = "UPDATING"
	NodegroupStatusDeleting     NodegroupStatus = "DELETING"
	NodegroupStatusCreateFailed NodegroupStatus = "CREATE_FAILED"
	NodegroupStatusDeleteFailed NodegroupStatus = "DELETE_FAILED"
	NodegroupStatusDegraded     NodegroupStatus = "DEGRADED"
)

func (enum NodegroupStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum NodegroupStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type UpdateParamType string

// Enum values for UpdateParamType
const (
	UpdateParamTypeVersion               UpdateParamType = "Version"
	UpdateParamTypePlatformVersion       UpdateParamType = "PlatformVersion"
	UpdateParamTypeEndpointPrivateAccess UpdateParamType = "EndpointPrivateAccess"
	UpdateParamTypeEndpointPublicAccess  UpdateParamType = "EndpointPublicAccess"
	UpdateParamTypeClusterLogging        UpdateParamType = "ClusterLogging"
	UpdateParamTypeDesiredSize           UpdateParamType = "DesiredSize"
	UpdateParamTypeLabelsToAdd           UpdateParamType = "LabelsToAdd"
	UpdateParamTypeLabelsToRemove        UpdateParamType = "LabelsToRemove"
	UpdateParamTypeMaxSize               UpdateParamType = "MaxSize"
	UpdateParamTypeMinSize               UpdateParamType = "MinSize"
	UpdateParamTypeReleaseVersion        UpdateParamType = "ReleaseVersion"
)

func (enum UpdateParamType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum UpdateParamType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type UpdateStatus string

// Enum values for UpdateStatus
const (
	UpdateStatusInProgress UpdateStatus = "InProgress"
	UpdateStatusFailed     UpdateStatus = "Failed"
	UpdateStatusCancelled  UpdateStatus = "Cancelled"
	UpdateStatusSuccessful UpdateStatus = "Successful"
)

func (enum UpdateStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum UpdateStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type UpdateType string

// Enum values for UpdateType
const (
	UpdateTypeVersionUpdate        UpdateType = "VersionUpdate"
	UpdateTypeEndpointAccessUpdate UpdateType = "EndpointAccessUpdate"
	UpdateTypeLoggingUpdate        UpdateType = "LoggingUpdate"
	UpdateTypeConfigUpdate         UpdateType = "ConfigUpdate"
)

func (enum UpdateType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum UpdateType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}
