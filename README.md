# BomPix Golang SDK - v1.0.0

Golang SDK para utilização dos métodos da API Rest BomPix de forma prática e fácil.

## Dependências

- Go (1.21.6+)

### Instalação

Execute o comando abaixo no terminal da sua aplicação para instalar a última versão da SDK
```sh
go get github.com/nwiry/bompix-sdk-go
```
### Uso da SDK

###### Autenticação via API:
---
```go
import (

	"github.com/nwiry/bompix-sdk-go/environment"
)

	env := environment.New("{API_KEY}", "{API_SECRET}")
	
	session, _, _ := auth.Login(env)

	fmt.Printf("Mensagem: %s", session.Message)
	fmt.Printf("Token: %s", session.Payload.Token)

	fmt.Printf("Expiração em segundos: %d",session.Payload.ExpiresIn)
  // [...]
```
---
###### Autenticação com dados salvos:
---
```go
import (

	"github.com/nwiry/bompix-sdk-go/environment"
)

	env := environment.New("{API_KEY}", "{API_SECRET}")
	// Realize as validações de tempo para garantir que o token que você guardou não expirou
	*env.AccessToken = "{TOKEN_SALVO}"
  // [...]
```

###### Criando links de pagamento:
----
```go

import (

	"github.com/nwiry/bompix-sdk-go/source/link"
)

	pLink, _, _ := link.Create(env, &link.LinkRequest{
		Slug: "nometestego2",
	})

	fmt.Printf("Link ID: %d\n", pLink.Payload.ID)
	fmt.Printf("Link de Pagamento: %s\n", pLink.Payload.URL)
	fmt.Printf("Nome do Link: %s\n", pLink.Payload.Slug)
// [...]
```

###### Obtendo links:
----
```go
import (

	"github.com/nwiry/bompix-sdk-go/source/link"
)


	// Obtendo uma lista de links
	links, _, _ := link.GetLinks(env)

	// Obtendo um link via ID:
	pLink, _, _ := link.GetLink(env, 1)

	// Obtendo um link via Nome:
	pLink, _, _ := link.GetLinkSlug(env, "nomedolink")

	// Caso a consulta não seja de multiplos links, o for não é necessario. Utilize a variavel diretamente.
	for _, pLink := range links.Payload {
		fmt.Printf("ID: %d\n", pLink.ID)
		fmt.Printf("Nome: %s\n", pLink.Slug)
		fmt.Printf("URL: %s\n", pLink.URL)
	}
// [...]
```

###### Gerando Pagamentos:
----
```go
import (

	"github.com/nwiry/bompix-sdk-go/source/payment"
)


	linkID := int(ID)
	payment, _, _ := payment.Create(env, &payment.PaymentRequest{
		Amount: 4.98,
		Type:   payment.TypePix,
		LinkID: linkID,
		Message: &payment.PaymentMessage{
			Name:  "Nome Teste",
			Email: "emailteste@example.com",
			Text:  "Teste",
		},
	})

  fmt.Printf("UUID: %s\n", payment.Payload.UUID)
	fmt.Printf("Código copia e cola: %s\n", payment.Payload.Qrcode)
	fmt.Printf("QR Code Image: %s\n", payment.Payload.QrcodePNG)
	fmt.Printf("Duração do pix em segundos: %s\n", payment.Payload.PixDuration)
// [...]
```

###### Obtendo pagamentos:
----
```go
import (

	"github.com/nwiry/bompix-sdk-go/source/payment"
)
	// Obtendo um pagamento especifico via UUID
	paymentResult, _, _ := payment.GetUuid(env, "{UUID}")

	// Obtendo uma lista de pagamentos
	payments, _, _ := payment.GetPayments(env)
// [...]
```

###### Criando e manuseando Webhooks:
----
```go
import (

	"github.com/nwiry/bompix-sdk-go/source/webhook"
)

	linkID := int(ID)
	// Criando um webhook
	nWebhook, _, _ := webhook.Create(env, &webhook.WebhookRequest{
		LinkID: linkID, // ID do link de pagamentos associado
		URL: "https://yourstore.app/notification_url", // URL que ira receber as notificações de pagamento
	})

	// Obtendo um webhook via id
	nWebhook, _, _ := webhook.GetWebhook(env, ID)

	// Obtendo uma lista de webhooks
	webhooks, _, _ := webhook.GetWebhooks(env)

	// Deletando um webhook
	webhook.Delete(env, ID)
// [...]
```

#### Tratando Execeções

> Durante o uso dos métodos para interação com as rotas da API, você pode se deparar com erros durante o caminho, por isso, é importante que você verifique se o resultado das estruturas `error` e `bperror.ResponseError` são do tipo `nil`, para tomar as ações devidas corretamente em sua aplicação
----
###### Exemplo de tratamento de exceção:
---
```go
// Tentativa de gerar um pagamento, obtendo as variaveis de erro da api e de requisição
payment, bperr, err := payment.Create(env, &payment.PaymentRequest{
		Amount: 4.98,
		Type:   payment.TypePix,
		LinkID: linkID,
		Message: &payment.PaymentMessage{
			Name:  "Nome Teste",
			Email: "emailteste@example.com",
			Text:  "Teste",
		},
	})

	if bperr != nil {
		log.Fatal(bperr) // A API do BomPix retornou um erro. Verifique se algum dado não está inválido ou faltante
	}

	if err != nil {
		log.Fatal(err) // Ocorreu um erro na sua requisição, verifique o formato dos dados enviados e realize a correção
	}

// Chegamos até aqui sem erro, prossiga com o funcionamento da sua aplicação
```
