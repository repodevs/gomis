# GOMIS - Golang (mini) Service

This repo contains a simple service with the following specifications.

Specifications:   
[x] Add Note
[x] Get Note by ID
[x] Get All Notes
[x] Delete Note by ID
[ ] Update Note (TODO)

---

## Getting Started
> NOTE: Make sure you have Golang and PostgreSQL installed.

1. Clone this repository in to `src/github.com/repodevs` on your `$GOPATH`.
2. Export environment variables on your terminal.
```sh
export DB_HOST="<your_db_host>"
export DB_PORT="<your_db_port>"
export DB_USER="<your_db_user>"
export DB_PASS="<your_db_password>"
export DB_NAME="<your_db_name>"
``` 
3. Run the App with `go run main.go` or you can build the app with `go build` then `./gomis`.


---

## Test the app with `httpie`

### Add Notes
`http POST 127.0.0.1:9090/notes title='My Note' content='my note content' author='YourName'`

```
HTTP/1.1 201 Created
Content-Length: 51
Content-Type: application/json
Date: Mon, 13 Jan 2020 05:24:27 GMT

{
    "data": {
        "id": 20
    },
    "message": "Success",
    "status": 201
}
```

### Get Note by ID
`http 127.0.0.1:9090/notes/20`

```
HTTP/1.1 200 OK
Content-Length: 119
Content-Type: application/json
Date: Mon, 13 Jan 2020 05:26:10 GMT

{
    "data": {
        "author": "repodevs",
        "content": "My Notes Content",
        "id": 20,
        "title": "My Notes"
    },
    "message": "Success",
    "status": 200
}

```

### Get All Notes
`http 127.0.0.1:9090/notes`

```
HTTP/1.1 200 OK
Content-Length: 193
Content-Type: application/json
Date: Mon, 13 Jan 2020 05:28:09 GMT

{
    "data": [
        {
            "author": "repodevs",
            "content": "content note 2",
            "id": 2,
            "title": "note2"
        },
        {
            "author": "repodevs",
            "content": "My Notes Content",
            "id": 20,
            "title": "My Notes"
        }
    ],
    "message": "Success",
    "status": 200
}
```

### Delete Note by ID
`http DELETE 127.0.0.1:9090/notes/2`

```
HTTP/1.1 204 No Content
Content-Type: application/json
Date: Mon, 13 Jan 2020 05:29:17 GMT

```

### Update Note
TODO

---

