package fositestore

import "encoding/json"

// marshalScopes serializes scopes to JSON text as pixelfed stores them: ["read","write"].
func marshalScopes(scopes []string) string {
	if len(scopes) == 0 {
		return "[]"
	}
	b, _ := json.Marshal(scopes)
	return string(b)
}

// unmarshalScopes parses pixelfed's JSON text scopes field.
func unmarshalScopes(text string) []string {
	if text == "" {
		return nil
	}
	var scopes []string
	_ = json.Unmarshal([]byte(text), &scopes)
	return scopes
}
