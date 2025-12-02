# City Weather Cloud Run ☁️🌡️

Sistema em Go que recebe um CEP brasileiro, identifica a localização e retorna as temperaturas atuais formatadas em Celsius, Fahrenheit e Kelvin. Projeto preparado para deploy no Google Cloud Run.

## 📋 Pré-requisitos

Para executar este projeto, você precisará de uma chave de API da [WeatherAPI](https://www.weatherapi.com/).

Crie um arquivo `.env` na raiz do projeto:

```env
WEATHER_API_KEY=sua_chave_aqui
```
> **⚠️ Observação Importante:** O arquivo `.env` foi incluído no repositório propositalmente para facilitar a correção e execução do projeto sem necessidade de configuração de ambiente. Em um cenário real de produção, este arquivo deve ser ignorado (`.gitignore`) por conter dados sensíveis.

## 🚀 Como Executar

O servidor roda por padrão na porta **8081**.

### Opção 1: Via Docker (Recomendado)

```bash
# 1. Construir a imagem
docker build -t weather-service .

# 2. Rodar o container
docker run -p 8081:8081 -e WEATHER_API_KEY=sua_chave weather-service
```

### Opção 2: Localmente (Via Makefile)

```bash
# Instalar dependências
make deps

# Compilar e executar a aplicação
make run
```

## 📝 Testes Manuais (api.http)

Para facilitar a validação e correção, o projeto inclui o arquivo `api.http` na raiz. Ele contém requisições prontas para testar os cenários de sucesso e erro.

**Como usar:**
1. Instale a extensão **REST Client** no VS Code.
2. Abra o arquivo `api.http`.
3. Clique no botão "Send Request" que aparecerá acima de cada URL (ex: CEP de São Paulo, Rio de Janeiro, CEP Inválido, etc).
4. A resposta aparecerá em uma aba lateral.

> **⚠️ Nota:** Certifique-se de que a aplicação está rodando (via Docker ou `make run`) antes de disparar as requisições.

## 📡 Endpoints da API

**GET** `/weather/{cep}`

* **cep**: Código postal de 8 dígitos.

#### Exemplos de Resposta

**Sucesso (200 OK):**
```json
{
  "temp_C": 30.3,
  "temp_F": 86.54,
  "temp_K": 303.3
}
```

**Erro (422 Unprocessable Entity):**
```text
{
  "message": "invalid zipcode"
}
```
**Erro (404 Not Found):**
```text
{
  "message": "can not find zipcode"
}
```

## ☁️ Google Cloud Run

URL de Produção para correção acadêmica: `https://city-weather-cloud-run-620205445610.us-central1.run.app`

**Atenção:** A aplicação não responde na raiz (`/`). É necessário complementar o endereço com a rota `/weather/` seguida do CEP.

**Formato:**
`https://<URL-DO-CLOUD-RUN>/weather/<CEP>`

**Exemplo real:**
```bash
https://city-weather-cloud-run-620205445610.us-central1.run.app/weather/01310100
```

> **Nota:** O CEP deve conter **apenas números**, sem espaços, traços ou caracteres especiais.
