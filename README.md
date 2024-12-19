# Calculator Service

## Описание
Этот проект представляет собой веб-сервис для вычисления арифметических выражений. Сервис принимает HTTP POST-запрос с выражением и возвращает результат вычислений. В случае ошибок сервис сообщает о причине.

### Основной функционал:
1. Принимает арифметическое выражение через POST-запрос на эндпоинт `/api/v1/calculate`.
2. Возвращает результат вычисления или соответствующую ошибку.

### Примеры использования:

<sub> Рекомендуется использовать Git Bash/WSL для успешного тестирования проекта</sub>

#### Успешный результат
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

#### Неправильный запрос
**Запрос:**
```bash
curl --location 'http://localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
  "expression": "2+2+"
}'
```
**Успешный ответ:**
```json
{
  "error":"invalid expression"
}
```

Всего программа может возвращать 6 видов ошибок:

<table>
	<thead>
		<tr>
			<th>№</th>
			<th>Пример тела запроса/Метод запроса</th>
			<th>Ответ сервера</th>
			<th>Описание ошибки</th>
		</tr>
	</thead>
	<tbody>
		<tr>
			<td>1</td>
			<td>{"expression": "2+abc"}/POST</td>
			<td>invalid character</td>
			<td>В выражении присутствуют недопустимые символы.</td>
		</tr>
		<tr>
			<td>2</td>
			<td>{"expression": "2+3)"}/POST</td>
			<td>unmatched parentheses</td>
			<td>В выражении присутствует открывающая или закрывающая скобка без пары.</td>
		</tr>
		<tr>
			<td>3</td>
			<td>{"expression": "2/0"}/POST</td>
			<td>division by zero</td>
			<td>Попытка деления на ноль.</td>
		</tr>
		<tr>
			<td>4</td>
			<td>{"expression": "2+2"}/GET</td>
			<td>method not allowed</td>
			<td>Неверный HTTP-метод (должен быть POST).</td>
		</tr>
		<tr>
			<td>5</td>
			<td>{"express"/POST</td>
			<td>expression is not valid</td>
			<td>Неверный формат тела запроса.</td>
		</tr>
		<tr>
			<td>6</td>
			<td>{"expression": "2+"}/POST</td>
			<td>invalid expression</td>
			<td>Некорректное выражение.</td>
		</tr>
	</tbody>
</table>

---

## Запуск проекта
### Требования
- Go 1.20+

### Команда запуска
```bash
go run ./cmd/...
```

После выполнения команды сервис будет запущен на `http://localhost:8080`.

Также вы можете указать любой другой порт в переменной среды вашей системы для запуска на этом порте. Пример для Git Bash

```bash
export PORT=7777 && go run ./cmd/...
```

---

## Тестирование

В проекте реализованы тесты для модуля вычислений, а также для основного модуля. Чтобы запустить тесты, можете выполнить следующие команды:

```bash
go test ./pkg/calculation/...
```

```bash
go test ./internal/application/...
```

Также вы можете проверить покрытие тестами с помощью следующих команд:

```bash
cd pkg/calculation && go test -cover
```

```bash
cd internal/application && go test -cover
```

---

## Связь с автором

Если у вас возникнут вопросы по поводу проекта, вы можете написать мне в телеграм: [@azizov_hikmatullo](https://azizov_hikmatullo.t.me)

---

## Contributing

А зачем? :)