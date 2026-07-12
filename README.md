# Donate Processor Simulator (Go)

[English](#english) | [Русский](#русский)

---

## English

Simulator of a backend server for processing and converting donations in Go.

A simple and reliable backend server simulator for processing and converting incoming donations for a streaming platform.

### Project Description

The program simulates the flow of incoming donations in different currencies, performs strict data validation, and converts successful transactions into rubles (RUB). 

#### Implemented Features:
* **Currency validation:** Only `USD` and `EUR` are supported. Any other currencies are rejected by the security system.
* **Security limits:** Minimum donation is 1 unit, maximum one-time donation is 5,000 units. All transactions outside these limits are blocked.
* **Error Handling:** Validation errors are logged but do not cause the server to crash  processing of subsequent donations continues normally (`continue`).
* **Guaranteed session closure:** Uses the `defer` mechanism to log session closure under any shutdown scenario.

### Code Architecture

The logic is divided into independent layers:
1. `randomCurrency()`  a slice-based random currency generator for simulating real traffic.
2. `convertDonateToRUB()`  an isolated validation and conversion function.
3. `main()`  the entry point that simulates an infinite transaction processing loop with a 1-second interval.

### How to Run Locally

Make sure you have [Go](https://go.dev) installed.

1. Clone the repository:
   ```bash
   git clone https://github.com/pepesoul/donate-converter-go.git
   ```
2. Navigate to the project directory:
   ```bash
   cd donate-converter-go
   ```
3. Run the application:
   ```bash
   go run main.go
   ```

---

## Русский

Симулятор бэкенд-сервера для обработки и конвертации донатов на Go.

Простой и надежный симулятор бэкенд-сервера для обработки и конвертации входящих донатов для стриминговой платформы.

### Описание проекта

Программа имитирует поток входящих донатов в различных валютах, выполняет строгую валидацию данных и конвертирует успешные транзакции в рубли (RUB).

#### Реализованный функционал:
* **Валидация валюты:** Поддерживаются только `USD` и `EUR`. Любые другие валюты отклоняются системой безопасности.
* **Лимиты безопасности:** Минимальный донат 1 единица, максимальный разовый донат 5 000 единиц. Все транзакции вне лимитов блокируются.
* **Обработка ошибок:** Ошибки валидации логируются, но не вызывают падение сервера обработка следующих донатов продолжается в штатном режиме (`continue`).
* **Гарантированное закрытие сессии:** Использование механизма `defer` для логирования закрытия сессии при любых сценариях завершения работы.

### Архитектура кода

Логика разделена на независимые слои:
1. `randomCurrency()`  генератор случайной валюты на основе срезов (slice) для симуляции реального трафика.
2. `convertDonateToRUB()`  изолированная функция валидации и конвертации.
3. `main()`  точка входа, симулирующая бесконечный цикл обработки транзакций с интервалом в 1 секунду.

### Как запустить проект локально

Убедитесь, что у вас установлен [Go](https://go.dev).

1. Склонируйте репозиторий:
   ```bash
   git clone https://github.com/pepesoul/donate-converter-go.git
   ```
2. Перейдите в директорию проекта:
   ```bash
   cd donate-converter-go
   ```
3. Запустите приложение:
   ```bash
   go run main.go
   ```
