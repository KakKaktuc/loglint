# loglint

**Loglint** — линтер для Go, проверяющий корректность сообщений логов, написанных через `slog` и `go.uber.org/zap`.  
Он помогает избежать ошибок в логах и выявлять потенциально чувствительные данные (пароли, токены, API-ключи).

---

## 1 Сборка и запуск

### A. Сборка бинаря

```powershell
# В корне проекта
go build -o loglint.exe ./cmd/loglint
```
### B. Запуск линтера на реальных файлах
```
.\loglint.exe .\realtest\
```

### C. Использование golangci-lint
На Windows golangci-lint напрямую не видит exe, поэтому рекомендуется запускать бинарь отдельно.
На Linux/macOS можно подключить через wrapper

Сборка wrapper:
```
go build -o loglint-wrapper ./golangci-lint-plugin/loglint_wrapper.go
```
Добавление в golangci.yml:
```
linters:
  enable:
    - loglint

linters-settings:
  loglint:
    type: path
    path: ./loglint-wrapper
    description: "Checks slog and zap log messages"
```
Запуск
```
golangci-lint run
```
---

## 2 Пример использования

Пример реального файла ./realtest/test.go

Запуск
```
.\loglint.exe .\realtest
```
Пример вывода
```
C:\Users\home\Documents\projects\loglint\realtest\test.go:27:12: log message must contain only lowercase english letters and allowed characters
C:\Users\home\Documents\projects\loglint\realtest\test.go:28:13: log message must contain only lowercase english letters and allowed characters
C:\Users\home\Documents\projects\loglint\realtest\test.go:35:12: log message must contain only english ascii characters with lowercase letters and allowed symbols
C:\Users\home\Documents\projects\loglint\realtest\test.go:38:12: log message must contain only english ascii characters with lowercase letters and allowed symbols
C:\Users\home\Documents\projects\loglint\realtest\test.go:41:14: log message may contain sensitive data
C:\Users\home\Documents\projects\loglint\realtest\test.go:42:15: log message may contain sensitive data
C:\Users\home\Documents\projects\loglint\realtest\test.go:43:14: log message may contain sensitive data
```

