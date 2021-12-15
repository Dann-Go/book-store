# Books Store

###Book Store is a simple golang web application. It uses gin framework and postgres DB.

## Installation

###Clone this repository onto your system.

## Usage

Before statring using it you have to set few envs.\

HOST (host of db)\
DBPORT (port of db)\
USERNAME (name of db user)\
PASSWORD (password foe db)\
DBNAME (db name)\
SSLMODE \
SERVPORT (server port)\
MODE (starting mode (debug ir release))

You don't have to create tebles in your db on your own. App will do it for you by using migrations.\
If your mode set to debug it will automatically create test data in your DB using seeds.

After setting all envs yoc can run 
```bash
make build
```
It will build a binar for you.\
To build a docker container with docker compose run\
```bash
docker-compose up --build [name]
```

##Tests
To run tests use
```bash
make test
```

## Endpoints
- Add /books
- GetAll /books
- GetById /books/:id
- Update /books/:id
- Delete /books/:id

## Update Request Example

```text
host:port/books/:?id=
```
JSON Body of UPDATE request
```json lines
{
  "id": 4,
  "title": "Мы",
  "authors":["Евгений Замятин"],
  "year":"1924"
}
```

## JSON Body example

```json lines
{
  "title": "Мы",
  "authors":["Евгений Замятин"],
  "year": "2000-01-01"
}
```
