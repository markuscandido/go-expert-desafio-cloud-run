# Weather API (Go Expert - Cloud Run Challenge)

## Project Overview

This project is a RESTful API developed in Go that retrieves current weather information based on a Brazilian Zipcode (CEP). It identifies the location associated with the CEP and returns the temperature in Celsius, Fahrenheit, and Kelvin.

The application is designed to be deployed on **Google Cloud Run** and follows strict architectural patterns.

## Tech Stack

*   **Language:** Go 1.25+
*   **Architecture:** Clean Architecture / Hexagonal Architecture
*   **External APIs:**
    *   [ViaCEP](https://viacep.com.br/) (Location lookup)
    *   [WeatherAPI](https://www.weatherapi.com/) (Weather data)
*   **Infrastructure:** Docker, Docker Compose, Google Cloud Run
*   **Dependencies:** Standard Library (`net/http`) - No external frameworks used for core logic.

## Architecture

The project follows **Clean Architecture** principles to ensure separation of concerns and testability:

*   **`cmd/api/`**: Application entry point.
*   **`internal/domain/`**: Core business entities (`Location`, `Weather`) and repository interfaces.
*   **`internal/usecase/`**: Application business rules (`GetWeatherByZipcode`). Orquestrates data flow between entities and repositories.
*   **`internal/infra/`**: Implementation of interfaces (Adapters).
    *   `external/`: Clients for ViaCEP and WeatherAPI.
    *   `web/`: HTTP Handlers.

## Key Features

1.  **CEP Validation**: Validates if the input is a valid 8-digit string.
2.  **Location Lookup**: Resolves CEP to City/State.
3.  **Weather Retrieval**: Fetches current temperature for the resolved city.
4.  **Temperature Conversion**: auto-converts Celsius to Fahrenheit and Kelvin.
    *   F = C * 1.8 + 32
    *   K = C + 273

## Building and Running

The project includes a `Makefile` to simplify common tasks.

### Prerequisites

*   Go 1.25+
*   Docker (optional)
*   WeatherAPI Key (set as `WEATHER_API_KEY` env var)

### Common Commands

*   **Run Locally:**
    ```bash
    export WEATHER_API_KEY=your_key
    make run
    ```
*   **Build:**
    ```bash
    make build
    ```
*   **Run with Docker:**
    ```bash
    make docker-build
    make docker-run
    ```
*   **Run with Docker Compose:**
    ```bash
    make up
    ```

## Testing

*   **Run Unit Tests:**
    ```bash
    make test
    ```
*   **Run Tests with Coverage:**
    ```bash
    make test-cover
    ```

## Development Conventions

*   **Dependency Injection:** Dependencies are injected manually (e.g., via constructors like `NewClient`, `NewUseCase`).
*   **Error Handling:** Custom errors are defined in `internal/domain/errors.go`.
*   **Configuration:** 12-factor app principles; configuration via environment variables.
*   **Formatting:** Standard `go fmt`.

## API Endpoints

### `GET /weather/{cep}`

*   **Success (200):** `{"temp_C": 28.5, "temp_F": 83.3, "temp_K": 301.5}`
*   **Invalid CEP (422):** `{"message": "invalid zipcode"}`
*   **Not Found (404):** `{"message": "can not find zipcode"}`
