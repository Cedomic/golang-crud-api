# Try A New Language Series: GOLANG

In the `Try A New Language Series` I take a new programming language that looks promising and implement a simple CRUD API in it.

In this part I took the programming language `Golang`.

This project uses the `gorilla/mux` router and `gorm` as an ORM.

The API can store artists and their songs in a database (Postgres) and lets you perform basic CRUD operations.

## Usage

You need Go installed on your machine.

    # Download the project
    go get github.com/cedomic/golang-crud-api

    # Build the project
    cd golang-crud-api
    go build

    # Run the project
    ./golang-crud-api

The API is available at `http://localhost:3000`

## API

- /artists

        GET => Get all artists
            Response => [
                {
                    "ID": 1,
                    "CreatedAt": "2020-03-31T11:39:36.175657Z",
                    "UpdatedAt": "2020-03-31T12:05:34.453087Z",
                    "DeletedAt": null,
                    "name": "Example Artist",
                    "description": "This is a description",
                    "songs": null
                }
            ]

        POST => Create a new artist
            Body => {
                "name": "Example Artist",
                "description": "This is a description"
            }

            Response => {
                "ID": 1,
                "CreatedAt": "2020-03-31T11:39:36.175657Z",
                "UpdatedAt": "2020-03-31T12:05:34.453087Z",
                "DeletedAt": null,
                "name": "Example Artist",
                "description": "This is a description",
                "songs": null
            }

* /artists/{artistName}

        GET => Get an artist by name
            Response => {
                    "ID": 1,
                    "CreatedAt": "2020-03-31T11:39:36.175657Z",
                    "UpdatedAt": "2020-03-31T12:05:34.453087Z",
                    "DeletedAt": null,
                    "name": "Example Artist",
                    "description": "This is a description",
                    "songs": null
            }

        PUT => Update an artist
            Body => {
                "name": "Updated Example Artist",
                "description": "This is a description"
            }

            Response => {
                "ID": 1,
                "CreatedAt": "2020-03-31T11:39:36.175657Z",
                "UpdatedAt": "2020-03-31T14:40:17.026605+02:00",
                "DeletedAt": null,
                "name": "Updated Example Artist",
                "description": "This is a description",
                "songs": null
            }

        DELETE => Delete an artist by name

* /artists/{artistName}/songs

        GET => Get all songs from an artist by name
            Response => [
                {
                    "ID": 1,
                    "CreatedAt": "2020-03-31T12:35:42.548864Z",
                    "UpdatedAt": "2020-03-31T12:35:42.548864Z",
                    "DeletedAt": null,
                    "title": "good morning",
                    "artist_id": 1,
                    "genre": ["Indie Poptism"],
                    "length": 180
                }
            ]

        POST => Create a new song for an artist by name
            Body => {
                "title": "good morning",
                "release": "2020-02-14T00:00:00.000000Z",
                "genre": ["Indie Poptism"],
                "length": 180
            }

            Response => {
                "ID": 1,
                "CreatedAt": "2020-03-31T12:35:42.548864Z",
                "UpdatedAt": "2020-03-31T12:35:42.548864Z",
                "DeletedAt": null,
                "title": "good morning",
                "artist_id": 1,
                "genre": ["Indie Poptism"],
                "length": 180
            }

* /artists/{artistName}/songs/{songTitle}

        GET => Get a song from an artist by name and song title
            Response => {
                "ID": 1,
                "CreatedAt": "2020-03-31T12:35:42.548864Z",
                "UpdatedAt": "2020-03-31T12:35:42.548864Z",
                "DeletedAt": null,
                "title": "good morning",
                "artist_id": 1,
                "genre": ["Indie Poptism"],
                "length": 180
            }

        PUT => Update a song from an artist by name and song title
            Body => {
                "title": "updated song title",
                "release": "2020-02-14T00:00:00.000000Z",
                "genre": ["Indie Poptism"],
                "length": 180
            }

            Response => {
                "ID": 1,
                "CreatedAt": "2020-03-31T12:35:42.548864Z",
                "UpdatedAt": "2020-03-31T12:35:42.548864Z",
                "DeletedAt": null,
                "title": "updated song title",
                "artist_id": 1,
                "genre": ["Indie Poptism"],
                "length": 180
            }

        DELETE => Delete a song from artist by name and song title
