package link

// linkPayload representa a resposta da chave payload de uma consulta de link.
type linkPayload struct {
	ID   int    `json:"id"`   // ID é o id do link.
	Slug string `json:"slug"` // Slug é o nome do link.
	URL  string `json:"url"`  // URL é o endereço final para utilizar o link de pagamentos.
}

// WebhookResponse representa a resposta de uma consulta de link cadastrado.
type LinkResponse struct {
	Message string      `json:"message"`           // Message é a mensagem da resposta.
	Payload linkPayload `json:"payload,omitempty"` // Payload contém os detalhes dos links.
}

// WebhooksResponse representa a resposta de uma consulta de links cadastrados.
type LinksResponse struct {
	Message string        `json:"message"`           // Message é a mensagem da resposta.
	Payload []linkPayload `json:"payload,omitempty"` // Payload contém um slice de links.
}
