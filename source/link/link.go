package link

import (
	"fmt"
	"net/http"

	"github.com/nwiry/bompix-sdk-go/environment"
	"github.com/nwiry/bompix-sdk-go/request"
	"github.com/nwiry/bompix-sdk-go/source/bperror"
)

// route é o caminho da rota de autenticação na API.
const route string = "/link"

// Create envia uma solicitação para criar um novo link de pagamentos usando o ambiente fornecido.
func Create(env *environment.Environment, link *LinkRequest) (response *LinkResponse, bperr *bperror.ResponseError, err error) {
	// Verifica se um token de acesso foi atribuído ao ambiente.
	err = env.RequireAccessToken()
	if err != nil {
		return nil, nil, err
	}

	// Cria uma nova solicitação para a rota de criação de link.
	r := request.Request{
		Environment:    env,
		Body:           link,
		Route:          route,
		UseAccessToken: true,
	}

	// Inicializa a resposta do link.
	response = &LinkResponse{}

	// Executa a solicitação de criação do link.
	bperr, err = r.DoRequest(http.MethodPost, response)

	return
}

// Delete realiza uma solicitação para deletar um link de pagamentos
func Delete(env *environment.Environment, id int) (*bperror.ResponseError, error) {
	// Verifica se um token de acesso foi atribuído ao ambiente.
	err := env.RequireAccessToken()
	if err != nil {
		return nil, err
	}

	// Cria uma nova solicitação para a rota de deletar link.
	r := request.Request{
		Environment:    env,
		Route:          fmt.Sprintf("%s/%d", route, id),
		UseAccessToken: true,
	}

	// Executa a solicitação para deletar o link.
	return r.DoRequest(http.MethodDelete, nil)
}

// GetLink realiza uma solicitação para obter um link pelo id fornecido.
func GetLink(env *environment.Environment, id int) (response *LinkResponse, bperr *bperror.ResponseError, err error) {
	// Verifica se um token de acesso foi atribuído ao ambiente.
	err = env.RequireAccessToken()
	if err != nil {
		return nil, nil, err
	}

	// Cria uma nova solicitação para a rota de obter link por id.
	r := request.Request{
		Environment:    env,
		Route:          fmt.Sprintf("%s/%d", route, id),
		UseAccessToken: true,
	}

	// Inicializa a resposta do link.
	response = &LinkResponse{}

	// Executa a solicitação para obter o link.
	bperr, err = r.DoRequest(http.MethodGet, response)

	return
}

// GetLinkSlug realiza uma solicitação para obter um link pelo slug fornecido.
func GetLinkSlug(env *environment.Environment, slug string) (response *LinkResponse, bperr *bperror.ResponseError, err error) {
	// Verifica se um token de acesso foi atribuído ao ambiente.
	err = env.RequireAccessToken()
	if err != nil {
		return nil, nil, err
	}

	// Cria uma nova solicitação para a rota de obter link por id.
	r := request.Request{
		Environment:    env,
		Route:          route + "/" + slug,
		UseAccessToken: true,
	}

	// Inicializa a resposta do link.
	response = &LinkResponse{}

	// Executa a solicitação para obter o link.
	bperr, err = r.DoRequest(http.MethodGet, response)

	return
}

// GetLinks realiza uma solicitação para obter uma lista de links.
func GetLinks(env *environment.Environment) (response *LinksResponse, bperr *bperror.ResponseError, err error) {
	// Verifica se um token de acesso foi atribuído ao ambiente.
	err = env.RequireAccessToken()
	if err != nil {
		return nil, nil, err
	}

	// Cria uma nova solicitação para a rota de obter links.
	r := request.Request{
		Environment:    env,
		Route:          route,
		UseAccessToken: true,
	}

	// Inicializa a resposta do link.
	response = &LinksResponse{}

	// Executa a solicitação para obter os links.
	bperr, err = r.DoRequest(http.MethodGet, response)

	return
}
