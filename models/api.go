package models

type ShortenUrlRequest struct {
	OriginalURL string `json:"original_url"`
}

type ShortenUrlResponse struct {
	ShortUrl string `json:"short_url"`
}

type ShortenUrlRedirect struct {
	Location string
}

// type HttpResponse struct {
// 	Code        int     `json:"code"`
// 	Message     string  `json:"message"`
// 	ErrorDetail *string `json:"error_detail,omitempty"`
// 	TraceID     *string `json:"trace_id,omitempty"`
// }
