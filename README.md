---

# Go Movies CRUD

## Introduction

**Go Movies CRUD** is a simple and efficient web application built with Go (Golang) that demonstrates the basic CRUD (Create, Read, Update, Delete) operations for managing a list of movies. This project serves as a practical example for learning Go and building RESTful APIs.

## Features

- **Create**: Add new movies to the collection with relevant details like title, director, release year, and genre.
- **Read**: View a list of all movies in the collection, or search for specific movies by title or director.
- **Update**: Modify the details of existing movies.
- **Delete**: Remove movies from the collection.
    

## Technologies

- **Backend**: Go (Golang)
- **Router**: Gorilla Mux
- **Version Control**: Git

## Project Structure

```
MovieCRUDAPI/
├── main.go
├── go.mod
├── go.sum
└── README.md
```

## Setup Instructions

### Prerequisites

- Go (https://golang.org/dl/)

### Installing

1. **Clone the repository:**

   ```sh
   git clone https://github.com/ArjinAlbay/go-movies-crud.git
   cd go-movies-crud
   ```

2. **Install the necessary dependencies:**

   ```sh
   go get -u github.com/gorilla/mux
   ```

### Running the Server

To start the server, run:

```sh
go run main.go
```

The server will start at `http://localhost:3000`.

## API Endpoints

### Get All Movies

- **URL:** `/movies`
- **Method:** `GET`
- **Description:** Retrieves all movies.

### Get a Movie

- **URL:** `/movies/{id}`
- **Method:** `GET`
- **Description:** Retrieves a movie by ID.

### Create a Movie

- **URL:** `/movies`
- **Method:** `POST`
- **Description:** Creates a new movie.
- **Request Body:**
 

### Update a Movie

- **URL:** `/movies/{id}`
- **Method:** `PUT`
- **Description:** Updates an existing movie by ID.
- **Request Body:**


### Delete a Movie

- **URL:** `/movies/{id}`
- **Method:** `DELETE`
- **Description:** Deletes a movie by ID.

`

### Example Movie

```json
{
  "id": "1",
  "isbn": "4598",
  "title": "The Lord of the Rings",
  "director": {
    "firstname": "Jack",
    "lastname": "Jhones"
  }
}
```

## Usage

- Open your browser and navigate to `http://localhost:3000/movies` to see all movies.
- Navigate to `http://localhost:3000/movies/{id}` to see a specific movie by ID.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request for any improvements, bug fixes, or new features.

1. Fork the repository
2. Create a new branch (`git checkout -b feature/your-feature`)
3. Commit your changes (`git commit -m 'Add some feature'`)
4. Push to the branch (`git push origin feature/your-feature`)
5. Open a pull request


## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## My Social Media Accounts


<a href="https://twitter.com/arjinalbay" target="blank"><img align="center" src="https://raw.githubusercontent.com/rahuldkjain/github-profile-readme-generator/master/src/images/icons/Social/twitter.svg" alt="arjinalbay" height="30" width="40" /></a>
<a href="https://linkedin.com/in/arjinalbay" target="blank"><img align="center" src="https://raw.githubusercontent.com/rahuldkjain/github-profile-readme-generator/master/src/images/icons/Social/linked-in-alt.svg" alt="arjinalbay" height="30" width="40" /></a>


---
