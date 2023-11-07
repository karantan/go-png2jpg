# go-png2jpg

`go-png2jpg` is a simple and efficient web API written in Go that converts PNG images to JPG format.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. 

### Prerequisites

What things you need to install the software and how to install them:

- Go (1.20.x)
- Any additional dependencies will be managed by `go mod`.

### Installing

A step by step series of examples that tell you how to get a development environment running:

1. Clone the repository to your local machine:
```bash
git clone https://github.com/karantan/go-png2jpg.git
```

2. Navigate to the cloned repository:
```bash
cd go-png2jpg
```

3. Install the dependencies:
```bash
go mod tidy
```

4. Run the service:
```bash
go run main.go
```

The API should now be running on `http://localhost:8080`.

## Usage

To convert an image from PNG to JPG see [openapi.yml](openapi.yml) for the API specs.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
