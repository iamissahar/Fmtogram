package types

type PassportElementError struct {
	Source      string   `json:"source,omitempty"`
	Type        string   `json:"type,omitempty"`
	FieldName   string   `json:"field_name,omitempty"`
	DataHash    *string  `json:"data_hash,omitempty"`
	ElementHash *string  `json:"element_hash,omitempty"`
	FileHashes  []string `json:"file_hashes,omitempty"`
	Message     string   `json:"message,omitempty"`
}
