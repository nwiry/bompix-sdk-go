package environment

import (
	"errors"
)

// ENDPOINT_API_URL é a URL base da API Bompix.
const ENDPOINT_API_URL string = "https://app.bompix.com.br/api/v1"

// Environment representa as configurações do ambiente para interagir com a API Bompix.
type Environment struct {
	ApiKey      string   // ApiKey é a chave da API.
	ApiSecret   string   // ApiSecret é o segredo da API.
	AccessToken *string  // AccessToken é o token de acesso à API.
}

// New cria uma nova instância do ambiente com as credenciais fornecidas.
func New(credentials ...string) (e *Environment) {
	e = &Environment{} // Inicializa uma nova instância do ambiente.
	l := len(credentials)
	if l < 2 {
		panic("informe um tipo válido para apiKey e apiSecret")
	}
	e.ApiKey = credentials[0] // Define a chave da API.
	e.ApiSecret = credentials[1] // Define o segredo da API.
	if l > 2 {
		e.AccessToken = &credentials[2] // Se um token de acesso foi fornecido, define-o.
	}
	return
}

// RequireAccessToken verifica se um token de acesso foi atribuído ao ambiente.
// Retorna um erro se nenhum token de acesso foi definido.
func (e *Environment) RequireAccessToken() error {
	if e.AccessToken == nil {
		return errors.New("nenhum token foi atribuído à estrutura do ambiente, por favor, realize a autenticação na API")
	}
	return nil
}
