# Spelling Corrector

The solution for this project was based on the proposal made on the following website: http://norvig.com/spell-correct.html and some ideas from https://cxwangyi.wordpress.com/2012/02/15/peter-norvigs-spelling-corrector-in-go/ as well.

> The difference is that the algorithm has been trained using the `Spanish` language.

## Try it

Use it now with a Docker instance. It will open the `8080` port to access the service.

    docker pull jorelosorio/spelling-corrector:latest

    docker run --name spelling-corrector -d -p 8080:80 -t jorelosorio/spelling-corrector:latest

Try it using the following example:

    http://localhost:8080/spelling?word=espanol

## Tools

- GoLang `1.17`
- Docker
- Visual Studio Code `Optional!`
    > It requires a `Remote - Containers` extension. for more information please refers to: https://code.visualstudio.com/docs/remote/containers#_getting-started

## Development

This project contains a `Dockerfile` file with all required dependencies to run it using `Visual Studio Code` + `Remote - Containers` extension.
However, if you want to make it run locally in your development machine, please follow the instructions below.

### Install Go

Install it from https://go.dev/dl/

### Build the Service

> Make sure the port `80` is currently free. **Optionally could be changed in the code!**

    go build -o spelling .

Then run the service

    ./spelling

### Training

Most of the training was made using free versions of books in `Spanish`. However, if you like to train for a new language or add extra words to the existing `dic` file, you can use the following main

    ```go
    package main

    import (
        "spelling-corrector/helpers"
    )

    func main() {
        helpers.TrainFromFile("./texts/book.txt")
    }
    ```

Call `TrainFromFile` function as many times you wish with different textbooks.

> To start from scratch please `delete` first the document file `dic` otherwise you will be mixing words if you use other languages. **Only keep `dic` if you are going to append more data in `Spanish`**

### Build Docker

To build the docker image use `Dockerfile.deploy` and the command

    docker build -f Dockerfile.deploy -t jorelosorio/spelling-corrector:latest .

To run the docker image

    docker run --name spelling-corrector -d -p 8080:80 -t jorelosorio/spelling-corrector:latest

Test the `spelling corrector` from the docker image

    http://localhost:8080/spelling?word={YOUR_WORD}
