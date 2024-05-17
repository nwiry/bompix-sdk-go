package bperror

import "net/http"

// ResponseError representa um erro retornado pela API Bompix.
type ResponseError struct {
	Headers    http.Header // Headers contém os cabeçalhos retornados pela API na requisição.
	Message    string      `json:"message"` // Message é a mensagem de erro retornada pela API.
	StatusCode int         // StatusCode é o código de erro retornado pela API.
}
