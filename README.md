# Weather API - CEP to Temperature

API RESTful em Go que recebe um CEP brasileiro, identifica a cidade e retorna a temperatura atual em Celsius, Fahrenheit e Kelvin.

[![Go Version](https://img.shields.io/badge/Go-1.25-blue.svg)](https://golang.org/)
[![Cloud Run](https://img.shields.io/badge/Google%20Cloud-Run-4285F4.svg)](https://cloud.google.com/run)

## 🚀 Deploy

**URL do serviço**: `https://weather-api-nccto6oxnq-uc.a.run.app`

## 📋 Endpoints

### GET /weather/{cep}

Retorna a temperatura para o CEP informado.

**Parâmetros:**
- `cep` (path): CEP brasileiro com 8 dígitos

**Respostas:**

| Status | Descrição | Body |
|--------|-----------|------|
| 200 | Sucesso | `{"temp_C": 28.5, "temp_F": 83.3, "temp_K": 301.5}` |
| 422 | CEP inválido | `{"message": "invalid zipcode"}` |
| 404 | CEP não encontrado | `{"message": "can not find zipcode"}` |

**Exemplos:**

```bash
# CEP válido
curl https://weather-api-nccto6oxnq-uc.a.run.app/weather/01310100

# CEP inválido (formato)
curl https://weather-api-nccto6oxnq-uc.a.run.app/weather/123

# CEP não encontrado
curl https://weather-api-nccto6oxnq-uc.a.run.app/weather/99999999
```

## 🛠️ Desenvolvimento Local

### Pré-requisitos

- Go 1.25+
- Docker (opcional)
- API Key do [WeatherAPI](https://www.weatherapi.com/)

### Executar

```bash
# Clonar repositório
git clone https://github.com/markuscandido/go-expert-desafio-cloud-run.git
cd go-expert-desafio-cloud-run

# Configurar variáveis de ambiente
export WEATHER_API_KEY=sua-api-key

# Executar
make run
# ou
go run cmd/api/main.go
```

### Docker

```bash
# Build e run
docker-compose up --build

# Ou manualmente
docker build -t weather-api .
docker run -p 8080:8080 -e WEATHER_API_KEY=xxx weather-api
```

### Testes

```bash
# Executar todos os testes
make test

# Com cobertura
make test-cover
```

### REST Client

Você também pode testar a API usando o arquivo `api/api.http` (requer extensão REST Client no VS Code).

## 📁 Estrutura do Projeto

```
├── api/
│   └── api.http             # Chamadas de teste (REST Client)
├── cmd/api/main.go          # Entry point
├── configs/                  # Configurações
├── internal/
│   ├── domain/              # Entidades e interfaces
│   ├── usecase/             # Regras de negócio
│   └── infra/               # Implementações (HTTP, APIs externas)
├── docs/                    # Documentação
├── Dockerfile               # Container multi-stage
└── docker-compose.yml       # Ambiente local
```

## 🌡️ Conversões de Temperatura

- **Celsius → Fahrenheit**: `F = C × 1.8 + 32`
- **Celsius → Kelvin**: `K = C + 273`

## 🔗 APIs Utilizadas

- [ViaCEP](https://viacep.com.br/) - Consulta de CEP
- [WeatherAPI](https://www.weatherapi.com/) - Dados meteorológicos

## 📖 Documentação

- [Requisitos](docs/1.Requisitos.md)
- [Arquitetura](docs/2.Arquitetura.md)
- [APIs Externas](docs/3.APIs-Externas.md)
- [Estrutura do Projeto](docs/4.Estrutura-Projeto.md)
- [Conversões de Temperatura](docs/5.Conversoes-Temperatura.md)
- [Deploy no Cloud Run](docs/6.Deploy-CloudRun.md)
- [Plano de Implementação](docs/7.Plano-Implementacao.md)

## ☁️ Deploy no Google Cloud Run

```bash
gcloud run deploy weather-api \
  --source . \
  --region=southamerica-east1 \
  --allow-unauthenticated \
  --set-env-vars="WEATHER_API_KEY=sua-api-key"
```

## 📝 Licença

MIT
