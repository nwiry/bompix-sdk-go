package link

// LinkRequest representa uma solicitação de criação de link de pagamentos.
type LinkRequest struct {
	Slug string `json:"slug"` // Slug é o nome do link.
}
