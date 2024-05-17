package webhook

import (
	"fmt"
	"net/http"

	"github.com/nwiry/bompix-sdk-go/environment"
	"github.com/nwiry/bompix-sdk-go/request"
	"github.com/nwiry/bompix-sdk-go/source/bperror"
)

// route é o caminho da rota de autenticação na API.
const route string = "/webhook"

// Create envia uma solicitação para criar um novo webhook de eventos usando o ambiente fornecido.
func Create(env *environment.Environment, webhook WebhookRequest) (response *WebhookResponse, bperr *bperror.ResponseError, err error) {
	// Verifica se um token de acesso foi atribuído ao ambiente.
	err = env.RequireAccessToken()
	if err != nil {
		return nil, nil, err
	}

	// Cria uma nova solicitação para a rota de criar webhook.
	r := request.Request{
		Environment:    env,
		Body:           webhook,
		Route:          route,
		UseAccessToken: true,
	}

	// Inicializa a resposta do webhook.
	response = &WebhookResponse{}

	// Executa a solicitação de criação do webhook.
	bperr, err = r.DoRequest(http.MethodPost, response)

	return
}

// GetWebhook realiza uma solicitação para obter um webhook pelo id fornecido.
func GetWebhook(env *environment.Environment, id int) (response *WebhooksResponse, bperr *bperror.ResponseError, err error) {
	// Verifica se um token de acesso foi atribuído ao ambiente.
	err = env.RequireAccessToken()
	if err != nil {
		return nil, nil, err
	}

	// Cria uma nova solicitação para a rota de obter o webhook por id
	r := request.Request{
		Environment:    env,
		Route:          fmt.Sprintf("%s/%d", route, id),
		UseAccessToken: true,
	}

	// Inicializa a resposta do webhook.
	response = &WebhooksResponse{}

	// Executa a solicitação de criação do webhook.
	bperr, err = r.DoRequest(http.MethodGet, response)

	return
}

// GetWebhooks realiza uma solicitação para obter uma lista de webhooks.
func GetWebhooks(env *environment.Environment, webhook WebhookRequest) (response *WebhooksResponse, bperr *bperror.ResponseError, err error) {
	// Verifica se um token de acesso foi atribuído ao ambiente.
	err = env.RequireAccessToken()
	if err != nil {
		return nil, nil, err
	}

	// Cria uma nova solicitação para a rota de obter webhooks
	r := request.Request{
		Environment:    env,
		Body:           webhook,
		Route:          route,
		UseAccessToken: true,
	}

	// Inicializa a resposta do webhook.
	response = &WebhooksResponse{}

	// Executa a solicitação de criação do webhook.
	bperr, err = r.DoRequest(http.MethodGet, response)

	return
}

// Delete realiza uma solicitação para deletar um webhook de eventos associado à um link.
func Delete(env *environment.Environment, id int) (*bperror.ResponseError, error) {
	// Verifica se um token de acesso foi atribuído ao ambiente.
	err := env.RequireAccessToken()
	if err != nil {
		return nil, err
	}

	// Cria uma nova solicitação para a rota de deletar webhook.
	r := request.Request{
		Environment:    env,
		Route:          fmt.Sprintf("%s/%d", route, id),
		UseAccessToken: true,
	}

	// Executa a solicitação para deletar o webhook.
	return r.DoRequest(http.MethodDelete, nil)
}
