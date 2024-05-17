package request

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/nwiry/bompix-sdk-go/environment"
	"github.com/nwiry/bompix-sdk-go/source/bperror"
)

// Request representa uma solicitação à API Bompix.
type Request struct {
	Environment    *environment.Environment // Environment é a configuração do ambiente para a solicitação.
	Body           interface{}              // Body é o corpo da solicitação.
	UseAccessToken bool                     // UseAccessToken indica se o token de acesso deve ser usado na solicitação.
	Route          string                   // Route é o caminho da URL da solicitação.
}

// DoRequest executa uma solicitação HTTP e preenche os dados de resposta na estrutura response.
func (r *Request) DoRequest(method string, response interface{}) (*bperror.ResponseError, error) {
	var b []byte
	if r.Body != nil {
		var err error
		b, err = json.Marshal(r.Body)
		if err != nil {
			return nil, fmt.Errorf("erro ao codificar corpo da requisição para JSON: %v", err)
		}
	}

	req, err := http.NewRequest(method, environment.ENDPOINT_API_URL+r.Route, bytes.NewBuffer(b))
	if err != nil {
		return nil, fmt.Errorf("erro ao criar requisição HTTP: %v", err)
	}

	// Verifica se deve usar o token de acesso na solicitação.
	if !r.UseAccessToken {
		// Caso não, utiliza a autenticação básica com a chave da API e o segredo.
		authHeader := "Basic " + base64.StdEncoding.EncodeToString([]byte(r.Environment.ApiKey+":"+r.Environment.ApiSecret))
		req.Header.Set("Authorization", authHeader)
	} else {
		// Caso sim, utiliza o token de acesso no cabeçalho de autorização.
		req.Header.Set("Authorization", "Bearer "+*r.Environment.AccessToken)
	}

	req.Header.Set("Content-Type", "application/json")

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("erro ao executar a requisição HTTP: %v", err)
	}
	defer resp.Body.Close()

	// Verifica se a resposta foi bem-sucedida (código de status entre 200 e 299)
	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
		if response != nil {
			// Decodifica o JSON da resposta e preenche a estrutura response
			err = json.NewDecoder(resp.Body).Decode(response)
			if err != nil {
				return nil, fmt.Errorf("erro ao decodificar resposta JSON: %v", err)
			}
		} else {
			return nil, nil // Nenhum retorno de resposta foi passado para receber dados
		}
	} else {
		// Se a resposta não for bem-sucedida, tenta decodificar um possível JSON de erro
		var errorResponse bperror.ResponseError
		errorResponse.Headers = resp.Header
		errorResponse.StatusCode = resp.StatusCode
		err = json.NewDecoder(resp.Body).Decode(&errorResponse)
		if err == nil && errorResponse.Message != "" {
			return &errorResponse, nil
		}
		// Se não foi possível decodificar um JSON de erro, retorna um erro padrão com o status code
		return nil, fmt.Errorf("erro na requisição HTTP: %s", resp.Status)
	}

	return nil, nil
}
