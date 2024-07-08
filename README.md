# API in Go using the GIN framework

## Functionalities

* GET: books (all or by id)
* POST: insert new books on the shelf
* PATCH: checkout or return a book

## Adding the new book

* Use the file *body.json*

```
$ curl localhost:8080/books --include --header "Content-Type: application/json" -d @body.json --request "POST"
```

## Accessing books by id

* If client tries to access something that does not exist, it returns "Book not found." in JSON.

```
curl localhost:8080/books/<ID>
```

## Checking out a book

```
curl localhost:8080/checkout?id=<id> --request "PATCH"
```

