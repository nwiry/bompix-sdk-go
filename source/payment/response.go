package payment

// PaymentResponse representa a resposta de uma solicitação de pagamento.
type PaymentResponse struct {
	Message string `json:"message"` // Message é a mensagem da resposta.
	Payload struct {
		UUID        string          `json:"uuid"`         // UUID é o identificador único do pagamento.
		LinkID      interface{}     `json:"link_id"`      // LinkID é o ID do link associado ao pagamento.
		Amount      interface{}     `json:"amount"`       // Amount é o valor do pagamento.
		Message     *PaymentMessage `json:"message"`      // Message é a mensagem associada ao pagamento.
		Paid        bool            `json:"paid"`         // Paid indica se o pagamento foi concluído.
		Qrcode      string          `json:"qrcode"`       // Qrcode é o código QR do pagamento.
		PixDuration string          `json:"pix_duration"` // PixDuration é a duração do PIX.
		QrcodePNG   string          `json:"qrcode_png"`   // QrcodePNG é o código QR do pagamento em formato PNG.
		Type        string          `json:"type"`         // Type é o tipo de pagamento.
	} `json:"payload"` // Payload contém os detalhes do pagamento.
}

// PaymentResponseGet representa a resposta de uma consulta individual de pagamento.
type PaymentResponseUuid struct {
	Message string `json:"message"` // Message é a mensagem da resposta.
	Payload struct {
		UUID       string          `json:"uuid"`        // UUID é o identificador único do pagamento.
		LinksID    interface{}     `json:"links_id"`    // LinkID é o ID do link associado ao pagamento.
		Amount     interface{}     `json:"amount"`      // Amount é o valor do pagamento.
		Message    *PaymentMessage `json:"message"`     // Message é a mensagem associada ao pagamento.
		Paid       bool            `json:"paid"`        // Paid indica se o pagamento foi concluído.
		QrcodePNG  string          `json:"qrcode_png"`  // QrcodePNG é o código QR do pagamento em formato PNG.
		LastUpdate string          `json:"last_update"` // Type é o tipo de pagamento.
	} `json:"payload"` // Payload contém os detalhes do pagamento.
}

// PaymentResponse representa a resposta de uma consulta de pagamentos múltiplos.
type PaymentsResponse struct {
	Message string `json:"message"` // Message é a mensagem da resposta.
	Payload []struct {
		UUID     string          `json:"uuid"`     // UUID é o identificador único do pagamento.
		LinksID  interface{}     `json:"links_id"` // LinkID é o ID do link associado ao pagamento.
		Amount   interface{}     `json:"amount"`   // Amount é o valor do pagamento.
		Message  *PaymentMessage `json:"message"`  // Message é a mensagem associada ao pagamento.
		Paid     bool            `json:"paid"`     // Paid indica se o pagamento foi concluído.
		DateTime string          `json:"date_time"`
	} `json:"payload,omitempty"` // Payload contém os detalhes do pagamento.
}
