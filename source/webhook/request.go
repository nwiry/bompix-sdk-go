package webhook

// WebhookRequest representa uma solicitação de criação de webhook de eventos para um link correspondente.
type WebhookRequest struct {
	LinkID int    `json:"link_id"` // LinkID é o Id do link associado
	URL    string `json:"url"`     // URL é
}
