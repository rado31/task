# Task

## Steps to initialize a project

- to create .env file, run a command: ```cp env.sample .env```
- open the .env file and fill with your values
- to initialize a database, run a command: ```make db```
- to run a server, run a command: ```make dev```

## Quick API docs

### 1. Get user

```http
GET http://<your-host>:<your-port>/users/:id
```

- Response

```json
{
	"id": 1,
	"name": "string",
	"surname": "string",
	"patronymic": "string",
	"age": 20,
	"gender": "male",
	"nationality": "string"
}
```

### 2. Get all users

```http
GET http://<your-host>:<your-port>/users/all
```

- Query params

key | value | description
--- | --- | ---
page | integer > 0 | optional, default value is 1
count | integer > 0 | optional, default value is 10
order_by | string | required, possible values: id, age, name, surname
desc | string | required, possible values: asc, desc

- Response

```json
[
	{
		"id": 1,
		"name": "string",
		"surname": "string",
		"patronymic": "string",
		"age": 10,
		"gender": "male",
		"nationality": "string"
	},
	{
		"id": 2,
		"name": "string",
		"surname": "string",
		"patronymic": "string",
		"age": 15,
		"gender": "female",
		"nationality": "string"
	}
]
```

### 3. Create user

```http
POST http://<your-host>:<your-port>/users
```

- Body

```json
{
    "name": "string, required",
    "surname": "string, required",
    "patronymic": "string, optional"
}
```

- Response ```201 Created```

### 4. Update user

```http
PUT http://<your-host>:<your-port>/users
```

- Body

```json
{
    "id": "integer, required",
    "name": "string, required",
    "surname": "string, required",
    "patronymic": "string, optional",
    "age": "integer, required",
    "gender": "string, required",
    "nationality": "string, required"
}
```

- Response ```200 Success```

### 5. Delete user

```http
DELETE http://<your-host>:<your-port>/users/:id
```

- Response ```200 Success```

## There is API collection (exported from Insomnia v4) in "docs" folder, so you can import and check it
