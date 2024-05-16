package auth

// AuthResponse representa a resposta da solicitação de autenticação da API.
type AuthResponse struct {
	Message string `json:"message"` // Message é uma mensagem de status da resposta.
	Payload struct {
		Token     string `json:"token"`      // Token é o token de acesso retornado pela API.
		TokenType string `json:"token_type"` // TokenType é o tipo de token retornado pela API.
		ExpiresIn int64  `json:"expires_in"` // ExpiresIn é o tempo de expiração do token de acesso em segundos.
	} `json:"payload"` // Payload contém os detalhes do token de acesso.
}
