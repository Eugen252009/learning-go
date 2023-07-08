# Project #1

## Bookshelf RESTFUL API

## How to Start

```bash
go run main.go
```

### Get All Movies

```curl
curl localhost:8000/movies
```

```json
[
    {
        "id":"1",
        "isbn":"438227",
        "title":"Movie One",
        "director":
        {
            "firstname":"John",
            "lastname":"Doe"
        }
        },
        {
        "id":"2",
        "isbn":"45455",
        "title":"Movie two",
        "director":
        {
            "firstname":"Steven",
            "lastname":"Smith"
            }
    }
]
```

### GET MOVIE By ID

```curl
curl localhost:8000/movies/{id}
```

```json
    {
        "id":"1",
        "isbn":"438227",
        "title":"Movie One",
        "director":
        {
            "firstname":"John",
            "lastname":"Doe"
        }
    }    
```

### Create Movie

```curl
curl -L -X PUT 'http://localhost:8000/movies/1'
```

```json
{
    "id": "1",
    "isbn": "",
    "title": "",
    "director": null
}
```

### Update Movie

```curl

curl -L -X PUT 'http://localhost:8000/movies/1' \
-H 'Content-Type: application/json' \
-d '  {
        "id": "1",
        "isbn": "438227",
        "title": "Movie three",
        "director": {
            "firstname": "DDD",
            "lastname": "CCC"
        }
    }'
```

```json
{
    "id": "1",
    "isbn": "438227",
    "title": "Movie three",
    "director": {
        "firstname": "DDD",
        "lastname": "CCC"
    }
}
```

### Delete Movie

```curl
curl -L -X PUT 'http://localhost:8000/movies/1'
```

```json
{
    "id": "1",
    "isbn": "",
    "title": "",
    "director": null
}
```
