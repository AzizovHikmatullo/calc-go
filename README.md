# Calculator Service

## Описание
Этот проект представляет собой веб-сервис для вычисления арифметических выражений. Сервис принимает HTTP POST-запрос с выражением и возвращает результат вычислений. В случае ошибок сервис сообщает о причине.

### Основной функционал:
1. Принимает арифметическое выражение через POST-запрос на эндпоинт `/api/v1/calculate`.
2. Возвращает результат вычисления или соответствующую ошибку.

### Пример использования:
**Запрос:**
```bash
curl --location 'http://localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
  "expression": "2+2*2"
}'
```
**Успешный ответ:**
```json
{
  "result": 6
}
```

**Пример ошибки 422 (некорректное выражение):**
```bash
curl --location 'http://localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
  "expression": "2+abc"
}'
```
**Ответ:**
```json
{
  "error": "invalid character"
}
```

```bash
curl --location 'http://localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
  "expression": "2+"
}'
```
**Ответ:**
```json
{
  "error": "invalid expression"
}
```

---

## Запуск проекта
### Требования
- Go 1.20+

### Команда запуска
```bash
go run ./cmd/...
```

После выполнения команды сервис будет запущен на `http://localhost:8080`.

---

## Тестирование

В проекте реализованы тесты для модуля вычислений. Чтобы запустить тесты, выполните следующую команду:
```bash
go test ./pkg/calculation/...
```

---

## Contributing

А зачем?