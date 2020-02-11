package policy

// ErrorHandlingPolicy controls error hadnling behavior when taking snapshots.
type ErrorHandlingPolicy struct {
	// IgnoreFileErrors controls whether or not snapshot operation should terminate when a file throws an error on being read
	IgnoreFileErrors *bool `json:"ignoreFileErrors,omitempty"`

	// IgnoreDirectoryErrors controls whether or not snapshot operation should terminate when a directory throws an error on being read or opened
	IgnoreDirectoryErrors *bool `json:"ignoreDirectoryErrors,omitempty"`
}

// Merge applies default values from the provided policy.
func (p *ErrorHandlingPolicy) Merge(src ErrorHandlingPolicy) {
	if p.IgnoreFileErrors == nil && src.IgnoreFileErrors != nil {
		p.IgnoreFileErrors = newBool(*src.IgnoreFileErrors)
	}

	if p.IgnoreDirectoryErrors == nil && src.IgnoreDirectoryErrors != nil {
		p.IgnoreDirectoryErrors = newBool(*src.IgnoreDirectoryErrors)
	}
}

// defaultErrorHandlingPolicy is the default error handling policy.
var defaultErrorHandlingPolicy = ErrorHandlingPolicy{
	IgnoreFileErrors:      newBool(false),
	IgnoreDirectoryErrors: newBool(false),
}

func newBool(b bool) *bool {
	return &b
}
