package pointer

// StringPtr return the pointer to v
func StringPtr(v string) *string {
	return &v
}

// IntPtr return the pointer to v
func IntPtr(v int) *int {
	return &v
}

// Float32Ptr return the pointer to v
func Float32Ptr(v float32) *float32 {
	return &v
}
