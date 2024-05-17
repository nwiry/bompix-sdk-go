package webhook

// webhookPayload representa a resposta da chave payload de uma consulta de webhook.
type webhookPayload struct {
	ID     int    `json:"id"`      // ID é o id do webhook de eventos.
	LinkID int    `json:"link_id"` // LinKID é o id do link associado ao webhook.
	URL    string `json:"url"`     // URL é o endereço que o webhook envia notificações.
}

// WebhookResponse representa a resposta de uma consulta de webhook cadastrado.
type WebhookResponse struct {
	Message string         `json:"message"`           // Message é a mensagem da resposta.
	Payload webhookPayload `json:"payload,omitempty"` // Payload contém os detalhes do webhook.
}

// WebhooksResponse representa a resposta de uma consulta de webhooks cadastrados.
type WebhooksResponse struct {
	Message string           `json:"message"`           // Message é a mensagem da resposta.
	Payload []webhookPayload `json:"payload,omitempty"` // Payload contém os detalhes dos webhooks.
}
