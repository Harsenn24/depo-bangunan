# Porto Harsenn API Documentation

&nbsp;

## Endpoints :

List of available endpoints:

- `POST /customer`
- `POST /login`
- `GET /customer`
- `GET /customer:id`
- `PUT /customer:id`
- `DELETE /customer:id`

Routes below need authentication:

- `POST /order`
- `GET /order`
- `GET /order:id`
- `PUT /order:id`
- `DELETE /customer:id`

&nbsp;

## 1. POST /customer

Request:

- body:

```json
{
  "name": "string",
  "email": "string",
  "password": "string"
}
```

_Response (201 - Created)_

```json
{
  "status": "integer",
  "message": "string",
  "data": "Success create customer"
}
```

_Response (400 - Bad Request)_

```json
{
  "status": 400,
  "message": "error",
  "data": "custom"
}
```

&nbsp;

## 2. PUT /login

Request:

- body:

```json
{
  "email": "string",
  "password": "string"
}
```

_Response (200 - OK)_

```json
{
  "status": "integer",
  "message": "string",
  "data": "token"
}
```

&nbsp;

## 3. get /customer

_Response (200 - OK)_

```json
{
  "status": "integer",
  "message": "string",
  "page": "integer",
  "total_page": "integer",
  "total_data": "integer",
  "limit": "integer",
  "data": [
    {
      "name": "string",
      "email": "string",
      "id": "integer"
    },
    .......
  ]
}
```

_Response (400 - Bad Request)_

&nbsp;

## 4. GET /customer/:id

Request:

- params:

```json
{
  "id": "int"
}
```

_Response (200 - Ok)_

```json
{
  "status": "integer",
  "message": "string",
  "data": {
    "name": "string",
    "email": "string",
    "id": "integer"
  }
}
```

&nbsp;

## 5. PUT /customer/:id

Request:

- params:

```json
{
  "id": "int"
}
```

- body:

```json
{
  "name": "string",
  "email": "string"
}
```

_Response (200 - OK)_

```json
{
  "status": "integer",
  "message": "string",
  "data": "string"
}
```

&nbsp;

## 6. DELETE /customer/:id

Request:

- params:

```json
{
  "id": "integer"
}
```

_Response (200 - OK)_

```json
{
  "status": "integer",
  "message": "string",
  "data": "string"
}
```

&nbsp;

## 7. POST /order

Request:

- Headers

```json
{
  "authorization": "string"
}
```

- Body:

```json
{
  "product": "string",
  "quantity": "integer",
  "price": "integer"
}
```

_Response (200 - OK)_

```json
{
  "status": "integer",
  "message": "string",
  "data": "string"
}
```

&nbsp;

## 8. GET /order

Request:

- Headers

```json
{
  "authorization": "string"
}
```

_Response (200 - OK)_

```json
{
  "status": "integer",
  "message": "string",
  "page": "integer",
  "total_page": "integer",
  "total_data": "integer",
  "limit": "integer",
  "data": [
    {
      "id": "integer",
      "product": "string",
      "quantity": "integer",
      "price": "integer",
      "Customer": {
        "name": "string"
      }
    },
    .......
  ]
}
```

&nbsp;

## 9. GET /order/:id

Request:

- Headers

```json
{
  "authorization": "string"
}
```

- Params

```json
{
  "id": "integer"
}
```

_Response (200 - OK)_

```json
{
  "status": "integer",
  "message": "string",
  "data": {
    "id": "integer",
    "product": "string",
    "quantity": "integer",
    "price": "integer",
    "Customer": {
      "name": "string"
    }
  }
}
```

&nbsp;

## 10. PUT /order/:id

Request:

- Headers

```json
{
  "authorization": "string"
}
```

- Params

```json
{
  "id": "integer"
}
```

_Response (200 - OK)_

```json
{
  "status": "integer",
  "message": "string",
  "data": "string"
}
```

&nbsp;

## 11. DELETE /order/:id

Request:

- Headers

```json
{
  "authorization": "string"
}
```

- Params

```json
{
  "id": "integer"
}
```

_Response (200 - OK)_

```json
{
  "status": "integer",
  "message": "string",
  "data": "string"
}
```

&nbsp;

## Global Error

_Response (500 - Internal Server Error)_

```json
{
  "status": "integer",
  "message": "string",
  "data": "custom error"
}
```
