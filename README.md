## Документация по API

### Отправить сообщение для записи его в БД

#### Запрос

```http
POST http://45.89.188.162:5001/messages
```

**Пример тела запроса (JSON):**

```json
{
  "text": "Hello!"
}
```

#### Ответ

```json
{
    "id": 2,
    "text": "Hello!",
    "processed": false,
    "created_at": "2024-07-22T16:20:50.769995Z"
}
```

### Посмотреть статистику по обработанным сообщениям

#### Запрос

```http
GET http://45.89.188.162:5001/messages/statistic
```

#### Ответ

```json
{
  "processed-messages": 2,
  "total-messages": 3
}
```
