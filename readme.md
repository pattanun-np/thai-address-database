# Thai Address Database

## Introduction

The Thai Address Database is a Go library designed to facilitate the management and querying of geographical data for Thailand, including provinces, districts, and subdistricts. This tool provides functionalities to search through these entities based on various attributes in both Thai and English languages.

## Features

- Load and parse address data from JSON files.
- Search addresses by name in Thai and English.
- Retrieve detailed information on provinces, districts, and subdistricts based on various criteria like ID or name.

## Requirements

- Go 1.15 or higher

## Installation

To install the Thai Address Database

your machine. Follow these steps to download and install the package:

1. **Ensure Go is Installed:**
   - You can check if Go is installed by running `go version` in your terminal. If Go is not installed, you can download it from [the official Go website](https://golang.org/dl/).

2. **Get the Package:**
   - Use the `go get` command to download the Thai Address Database package to your Go workspace or module:
     ```bash
     go get github.com/yourusername/thaiaddressdatabase
     ```

3. **Import the Package:**
   - Include the package in your Go projects by importing it:
     ```go
     import "github.com/yourusername/thaiaddressdatabase"
     ```

## Usage

### Initializing the Database

Before performing any searches, you must initialize the database to load the data from JSON files:

```go
addressDB, err := thaiaddressdatabase.Init()
if err != nil {
    log.Fatalf("Failed to initialize address database: %v", err)
}
```

### Performing Searches

To search for address data, use the `Search` method:

```go
results := addressDB.Search("Bangkok")
for _, result := range results {
    fmt.Println(result)
}
```

## Testing

To run the tests included in the Thai Address Database:

1. Navigate to the project directory.
2. Run the following command:
   ```bash
   go test
   ```

## Contributing

Contributions to the Thai Address Database are welcome! Please refer to the CONTRIBUTING.md file for more details on how to submit pull requests, report issues, or request new features.

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Contact

For any questions or concerns regarding the Thai Address Database, please contact:

- Email: pattanan@standupcode.tech
- GitHub: [https://github.com/pattanunNP](https://github.com/pattanunNP)

---

This README provides a clear and concise guide for anyone interested in using or contributing to the Thai Address Database project. Adjust the content as necessary to fit the specifics of your project or organizational guidelines.