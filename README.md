## Архитектура микросервисов
[![msg.png](https://i.postimg.cc/wvCHx9GS/msg.png)](https://postimg.cc/N9kZpvY4)

## Документация по API
### Отправить сообщение для записи его в БД
#### Запрос

```http
POST http://5.35.12.248:5001/messages
```

**Пример тела запроса (JSON):**
```json
{
    "text":"Hello!"
}
```

#### Ответ

```json
{
    "id": 3
}
```
### Посмотреть статистику по обработанным сообщениям
#### Запрос

```http
GET http://5.35.12.248:5001/messages/statistic
```

#### Ответ

```json
{
    "processed-messages": 2,
    "total-messages": 3
}
```
