package payment

import (
	"net/http"

	"github.com/nwiry/bompix-sdk-go/environment"
	"github.com/nwiry/bompix-sdk-go/request"
	"github.com/nwiry/bompix-sdk-go/source/bperror"
)

// route é o caminho da rota de autenticação na API.
const route string = "/payment"
// TypePix representa o tipo de pagamento como "pix"
const TypePix = "pix"

// Create envia uma solicitação para criar um novo pagamento usando o ambiente fornecido.
func Create(env *environment.Environment, payment *PaymentRequest) (paymentResponse *PaymentResponse, bperr *bperror.ResponseError, err error) {
	// Verifica se um token de acesso foi atribuído ao ambiente.
	err = env.RequireAccessToken()
	if err != nil {
		return nil, nil, err
	}

	// Cria uma nova solicitação para a rota de criação de pagamento.
	r := request.Request{
		Environment:    env,
		Body:           payment,
		Route:          route,
		UseAccessToken: true,
	}

	// Inicializa a resposta do pagamento.
	paymentResponse = &PaymentResponse{}
	// Executa a solicitação de criação de pagamento.
	bperr, err = r.DoRequest(http.MethodPost, paymentResponse)

	// Se não houver erros na solicitação e na resposta, retorna o resultado
	return
}

// GetUuid envia uma solicitação para obter um pagamento pelo UUID fornecido, usando o ambiente fornecido.
func GetUuid(env *environment.Environment, uuid string) (paymentResponse *PaymentResponseUuid, bperr *bperror.ResponseError, err error) {
	// Verifica se um token de acesso foi atribuído ao ambiente.
	err = env.RequireAccessToken()
	if err != nil {
		return nil, nil, err
	}

	// Cria uma nova solicitação para a rota de obtenção de pagamento.
	r := request.Request{
		Environment:    env,
		Route:          route + "/" + uuid,
		UseAccessToken: true,
	}

	// Inicializa a resposta da busca de pagamento.
	paymentResponse = &PaymentResponseUuid{}
	// Executa a solicitação de obtenção de pagamento.
	bperr, err = r.DoRequest(http.MethodGet, paymentResponse)

	// Se não houver erros na solicitação e na resposta, retorna o resultado
	return
}

// GetPayments realiza uma solicitação para obter uma lista de pagamentos.
func GetPayments(env *environment.Environment) (paymentResponse *PaymentsResponse, bperr *bperror.ResponseError, err error) {
	// Verifica se um token de acesso foi atribuído ao ambiente.
	err = env.RequireAccessToken()
	if err != nil {
		return nil, nil, err
	}

	// Cria uma nova solicitação para a rota de obtenção de pagamento.
	r := request.Request{
		Environment:    env,
		Route:          route,
		UseAccessToken: true,
	}

	// Inicializa a resposta da busca de pagamento.
	paymentResponse = &PaymentsResponse{}
	// Executa a solicitação de obtenção de pagamento.
	bperr, err = r.DoRequest(http.MethodGet, paymentResponse)

	// Se não houver erros na solicitação e na resposta, retorna o resultado
	return
}