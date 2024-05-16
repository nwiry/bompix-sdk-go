package payment

// PaymentRequest representa uma solicitação de pagamento.
type PaymentRequest struct {
	Amount  float64         `json:"amount"`  // Amount é o valor do pagamento.
	Type    string          `json:"type"`    // Type é o tipo de pagamento.
	LinkID  int             `json:"link_id"` // LinkID é o ID do link associado ao pagamento.
	Message *PaymentMessage `json:"message"` // Message é a mensagem associada ao pagamento.
}

// PaymentMessage representa a mensagem associada a um pagamento.
type PaymentMessage struct {
	Name  string `json:"name,omitempty"`  // Name é o nome associado à mensagem do pagamento.
	Text  string `json:"text,omitempty"`  // Text é o texto da mensagem do pagamento.
	Email string `json:"email,omitempty"` // Email é o e-mail associado à mensagem do pagamento.
}
