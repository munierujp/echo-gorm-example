package response

// NOTE: see https://tools.ietf.org/html/rfc7807

type Problem struct {
	Type   string `json:"type"`
	Title  string `json:"title"`
	Detail string `json:"detail,omitempty"`
}
