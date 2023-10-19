package tvdb

// stolen from aws-sdk-go
// https://github.com/aws/aws-sdk-go/blob/main/aws/convert_types.go

// String returns a pointer to the string value passed in.
func String(v string) *string {
	return &v
}

// StringValue returns the value of the string pointer passed in or
// "" if the pointer is nil.
func StringValue(v *string) string {
	if v != nil {
		return *v
	}
	return ""
}