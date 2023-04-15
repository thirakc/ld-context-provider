package ldcontext

import "ld-context-provider/pkg/internal/common"

const (
	InvalidRequestErrorCode = iota + common.LDContext
	SaveLDContextErrorCode
	SaveResultErrorCode
	Base64DecodeStringErrorCode
	ContentIsNilErrorCode
)
