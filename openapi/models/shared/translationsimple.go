// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// TranslationSimple - translation simple record
type TranslationSimple struct {
	Language *string `json:"language,omitempty"`
}

func (o *TranslationSimple) GetLanguage() *string {
	if o == nil {
		return nil
	}
	return o.Language
}
