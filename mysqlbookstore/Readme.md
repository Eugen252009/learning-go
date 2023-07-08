# Project #2

## Docker Required

```bash
docker-compose up -d
go run cmd/main.go
```

Import this Postman config to test the API

```json
{
 "info": {
  "_postman_id": "34055c2a-f4ad-4b7e-b214-8dbe8f483a80",
  "name": "GO-Learning #2",
  "schema": "https://schema.getpostman.com/json/collection/v2.1.0/collection.json"
 },
 "item": [
  {
   "name": "GET",
   "request": {
    "method": "GET",
    "header": [],
    "url": {
     "raw": "http://localhost:8000/book/",
     "protocol": "http",
     "host": [
      "localhost"
     ],
     "port": "8000",
     "path": [
      "book",
      ""
     ]
    }
   },
   "response": []
  },
  {
   "name": "GET by ID",
   "protocolProfileBehavior": {
    "disableBodyPruning": true
   },
   "request": {
    "method": "GET",
    "header": [],
    "body": {
     "mode": "raw",
     "raw": "{\n    \"ID\": 0,\n    \"CreatedAt\": \"0001-01-01T00:00:00Z\",\n    \"UpdatedAt\": \"0001-01-01T00:00:00Z\",\n    \"DeletedAt\": null,\n    \"name\": \"\",\n    \"author\": \"\",\n    \"publication\": \"\"\n}",
     "options": {
      "raw": {
       "language": "json"
      }
     }
    },
    "url": {
     "raw": "http://localhost:8000/book/1",
     "protocol": "http",
     "host": [
      "localhost"
     ],
     "port": "8000",
     "path": [
      "book",
      "1"
     ]
    }
   },
   "response": []
  },
  {
   "name": "POST",
   "request": {
    "method": "POST",
    "header": [],
    "body": {
     "mode": "raw",
     "raw": "{\n    \"ID\": 1,\n    \"name\": \"Harry Potter\",\n    \"author\": \"Joanne K. Rowling\",\n    \"publication\": \"Carlsen Verlag\"\n}",
     "options": {
      "raw": {
       "language": "json"
      }
     }
    },
    "url": {
     "raw": "http://localhost:8000/book/",
     "protocol": "http",
     "host": [
      "localhost"
     ],
     "port": "8000",
     "path": [
      "book",
      ""
     ]
    }
   },
   "response": []
  },
  {
   "name": "UPDATE",
   "request": {
    "method": "PUT",
    "header": [],
    "body": {
     "mode": "raw",
     "raw": " {\n        \"name\": \"Harry Potter\",\n        \"author\": \"Joanne K. Rowling\",\n        \"publication\": \"Carlsen Verlag\"\n    }",
     "options": {
      "raw": {
       "language": "json"
      }
     }
    },
    "url": {
     "raw": "http://localhost:8000/book/2",
     "protocol": "http",
     "host": [
      "localhost"
     ],
     "port": "8000",
     "path": [
      "book",
      "2"
     ]
    }
   },
   "response": []
  },
  {
   "name": "DELETE",
   "request": {
    "method": "DELETE",
    "header": [],
    "url": {
     "raw": "http://localhost:8000/book/1",
     "protocol": "http",
     "host": [
      "localhost"
     ],
     "port": "8000",
     "path": [
      "book",
      "1"
     ]
    }
   },
   "response": []
  }
 ]
}
```
