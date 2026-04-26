package sample

// SampleRequest is the payload for creating a sample resource.
type SampleRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

// SampleResponse is returned for sample resources.
type SampleResponse struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
