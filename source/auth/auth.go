package auth

import (
	"net/http"

	"github.com/nwiry/bompix-sdk-go/environment"
	"github.com/nwiry/bompix-sdk-go/request"
	"github.com/nwiry/bompix-sdk-go/source/bperror"
)

// route é o caminho da rota de autenticação na API.
const route string = "/auth"

// Login realiza uma solicitação de login para autenticar o usuário na API.
func Login(env *environment.Environment) (auth *AuthResponse, bperr *bperror.ResponseError, err error) {
	// Cria uma nova solicitação para a rota de autenticação.
	r := request.Request{
		Environment: env,
		Route:       route,
	}

	// Inicializa a resposta de autenticação.
	auth = &AuthResponse{}

	// Executa a solicitação de login.
	bperr, err = r.DoRequest(http.MethodPost, auth)

	// Se não houver erros na solicitação e na resposta,
	// atualiza o token de acesso no ambiente.
	if bperr == nil && err == nil {
		env.AccessToken = &auth.Payload.Token
	}

	return
}

// Logout realiza uma solicitação de logout para desautenticar o usuário na API.
func Logout(env *environment.Environment) (*bperror.ResponseError, error) {
	// Verifica se um token de acesso foi atribuído ao ambiente.
	err := env.RequireAccessToken()
	if err != nil {
		return nil, err
	}

	// Cria uma nova solicitação para a rota de autenticação.
	r := request.Request{
		Environment: env,
		Route:       route,
	}

	// Executa a solicitação de logout.
	return r.DoRequest(http.MethodPost, nil)
}
