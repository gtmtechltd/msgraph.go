// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "time"

// DeviceHealthAttestationState undocumented
type DeviceHealthAttestationState struct {
	// LastUpdateDateTime The Timestamp of the last update.
	LastUpdateDateTime *string `json:"lastUpdateDateTime,omitempty"`
	// ContentNamespaceURL The DHA report version. (Namespace version)
	ContentNamespaceURL *string `json:"contentNamespaceUrl,omitempty"`
	// DeviceHealthAttestationStatus The DHA report version. (Namespace version)
	DeviceHealthAttestationStatus *string `json:"deviceHealthAttestationStatus,omitempty"`
	// ContentVersion The HealthAttestation state schema version
	ContentVersion *string `json:"contentVersion,omitempty"`
	// IssuedDateTime The DateTime when device was evaluated or issued to MDM
	IssuedDateTime *time.Time `json:"issuedDateTime,omitempty"`
	// AttestationIdentityKey TWhen an Attestation Identity Key (AIK) is present on a device, it indicates that the device has an endorsement key (EK) certificate.
	AttestationIdentityKey *string `json:"attestationIdentityKey,omitempty"`
	// ResetCount The number of times a PC device has hibernated or resumed
	ResetCount *int `json:"resetCount,omitempty"`
	// RestartCount The number of times a PC device has rebooted
	RestartCount *int `json:"restartCount,omitempty"`
	// DataExcutionPolicy DEP Policy defines a set of hardware and software technologies that perform additional checks on memory
	DataExcutionPolicy *string `json:"dataExcutionPolicy,omitempty"`
	// BitLockerStatus On or Off of BitLocker Drive Encryption
	BitLockerStatus *string `json:"bitLockerStatus,omitempty"`
	// BootManagerVersion The version of the Boot Manager
	BootManagerVersion *string `json:"bootManagerVersion,omitempty"`
	// CodeIntegrityCheckVersion The version of the Boot Manager
	CodeIntegrityCheckVersion *string `json:"codeIntegrityCheckVersion,omitempty"`
	// SecureBoot When Secure Boot is enabled, the core components must have the correct cryptographic signatures
	SecureBoot *string `json:"secureBoot,omitempty"`
	// BootDebugging When bootDebugging is enabled, the device is used in development and testing
	BootDebugging *string `json:"bootDebugging,omitempty"`
	// OperatingSystemKernelDebugging When operatingSystemKernelDebugging is enabled, the device is used in development and testing
	OperatingSystemKernelDebugging *string `json:"operatingSystemKernelDebugging,omitempty"`
	// CodeIntegrity  When code integrity is enabled, code execution is restricted to integrity verified code
	CodeIntegrity *string `json:"codeIntegrity,omitempty"`
	// TestSigning When test signing is allowed, the device does not enforce signature validation during boot
	TestSigning *string `json:"testSigning,omitempty"`
	// SafeMode Safe mode is a troubleshooting option for Windows that starts your computer in a limited state
	SafeMode *string `json:"safeMode,omitempty"`
	// WindowsPE Operating system running with limited services that is used to prepare a computer for Windows
	WindowsPE *string `json:"windowsPE,omitempty"`
	// EarlyLaunchAntiMalwareDriverProtection ELAM provides protection for the computers in your network when they start up
	EarlyLaunchAntiMalwareDriverProtection *string `json:"earlyLaunchAntiMalwareDriverProtection,omitempty"`
	// VirtualSecureMode VSM is a container that protects high value assets from a compromised kernel
	VirtualSecureMode *string `json:"virtualSecureMode,omitempty"`
	// PcrHashAlgorithm Informational attribute that identifies the HASH algorithm that was used by TPM
	PcrHashAlgorithm *string `json:"pcrHashAlgorithm,omitempty"`
	// BootAppSecurityVersion The security version number of the Boot Application
	BootAppSecurityVersion *string `json:"bootAppSecurityVersion,omitempty"`
	// BootManagerSecurityVersion The security version number of the Boot Application
	BootManagerSecurityVersion *string `json:"bootManagerSecurityVersion,omitempty"`
	// TpmVersion The security version number of the Boot Application
	TpmVersion *string `json:"tpmVersion,omitempty"`
	// Pcr0 The measurement that is captured in PCR[0]
	Pcr0 *string `json:"pcr0,omitempty"`
	// SecureBootConfigurationPolicyFingerPrint Fingerprint of the Custom Secure Boot Configuration Policy
	SecureBootConfigurationPolicyFingerPrint *string `json:"secureBootConfigurationPolicyFingerPrint,omitempty"`
	// CodeIntegrityPolicy The Code Integrity policy that is controlling the security of the boot environment
	CodeIntegrityPolicy *string `json:"codeIntegrityPolicy,omitempty"`
	// BootRevisionListInfo The Boot Revision List that was loaded during initial boot on the attested device
	BootRevisionListInfo *string `json:"bootRevisionListInfo,omitempty"`
	// OperatingSystemRevListInfo The Operating System Revision List that was loaded during initial boot on the attested device
	OperatingSystemRevListInfo *string `json:"operatingSystemRevListInfo,omitempty"`
	// HealthStatusMismatchInfo This attribute appears if DHA-Service detects an integrity issue
	HealthStatusMismatchInfo *string `json:"healthStatusMismatchInfo,omitempty"`
	// HealthAttestationSupportedStatus This attribute indicates if DHA is supported for the device
	HealthAttestationSupportedStatus *string `json:"healthAttestationSupportedStatus,omitempty"`
}