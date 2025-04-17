# Specifications

## Requests

> POST /v1/todos

request:

```json
  {
    "title": "Atividade",
    "dueDate": "2025-11-17 10:29"
  }
```

response:

```json
  {
    "id": "620bbf09-2030-48d3-b407-7a2a0f278774",
    "title": "Atividade",
    "dueDate": "2025-11-17 10:29",
    "updatedAt": "2025-11-17 10:29",
    "createdAt": "2025-11-17 10:29",
    "status": "CREATED"
  }
```

> POST /v1/todos/:id/comments

request:

```json
  {
    "comment": "Algum comentario"
  }
```

> POST /v1/todos/:id/progress

```json
  {
    "comment": "Algum comentario",
    "progress": "0.23"
  }
```

> PATCH /v1/todos/:id

```json
  {
    "title": "Titulo atualizado"
  }
```

> PUT /v1/todos/:id/status

```json
  {
    "status": "STARTED"
  }
```

> GET /v1/todos/:id

```json
  {
    "id": "190ab47e-b4f5-4e95-96ca-4089b26e2a38",
    "title": "Atividade",
    "dueDate": "2025-04-17 10:35",
    "createdAt": "2025-04-17 10:35",
    "updatedAt": "2025-04-17 10:35",
    "progress": "50.00",
    "timeDelayed": "0",
    "timeRunning": "0",
    "comments": [
      {
        "comment": "Algum comentario legal",
        "commentedAt": "2025-04-17 10:36"
      }
    ]
  }
```

> GET /v1/todos

```json
  [
    {
      "status": "CREATED",
      "quantity": 1,
      "todos": [
        {
          "id": "a072d0d8-2bd8-4be0-bb15-3b455375d78c",
          "title": "Atividade",
          "createdAt": "2025-04-17 11:12",
          "timeToDueDate": "15000"
        }
      ]
    },
    {
      "status": "ACTIVE",
      "quantity": 1,
      "todos": [
        {
          "id": "a072d0d8-2bd8-4be0-bb15-3b455375d78c",
          "title": "Atividade",
          "createdAt": "2025-04-17 11:12",
          "startedAt": "2025-04-17 11:12",
          "timeRunning": "100",
          "timeDelayed": "0",
          "timeToDueDate": "15000"
        }
      ]
    },
    {
      "status": "COMPLETED",
      "quantity": 1,
      "todos": [
        {
          "id": "a072d0d8-2bd8-4be0-bb15-3b455375d78c",
          "title": "Atividade finalizada",
          "createdAt": "2025-04-17 11:12",
          "createdAt": "2025-04-17 11:12",
          "finishedAt": "2025-04-17 11:14",
          "timeToDueDate": "0"
        }
      ]
    }
  ]
```
