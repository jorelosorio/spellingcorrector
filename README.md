# Spelling Corrector

[![Tests](https://github.com/jorelosorio/spellingcorrector/actions/workflows/tests.yml/badge.svg)](https://github.com/jorelosorio/spellingcorrector/actions/workflows/tests.yml)
[![Coverage Status](https://coveralls.io/repos/github/jorelosorio/spellingcorrector/badge.svg?branch=main)](https://coveralls.io/github/jorelosorio/spellingcorrector?branch=main)
[![Go Reference](https://pkg.go.dev/badge/github.com/jorelosorio/spellingcorrector.svg)](https://pkg.go.dev/github.com/jorelosorio/spellingcorrector)

A spelling corrector for the Spanish language or create your own.

The solution for this project was based on the proposal made on the following website: http://norvig.com/spell-correct.html and some ideas from https://cxwangyi.wordpress.com/2012/02/15/peter-norvigs-spelling-corrector-in-go/ as well.

> The built-in data was trained using the `Spanish` language.

## Try it

Use it now with a Docker instance. It will open the `8080` port to access the service.

    docker pull jorelosorio/spellingcorrector:latest

    docker run --name spellingcorrector -d -p 8080:80 -t jorelosorio/spellingcorrector:latest

Try it using the following example:

    http://localhost:8080/spelling?word=espanol

## Tools

- GoLang `1.17.x`
- Docker
- Visual Studio Code `Optional!`
    > It requires a `Remote - Containers` extension. for more information please refers to: https://code.visualstudio.com/docs/remote/containers#_getting-started

## Development

This project contains a `Dockerfile` file with all required dependencies to run it using `Visual Studio Code` + `Remote - Containers` extension.
However, if you want to make it run locally in your development machine, please follow the instructions below.

### Install Go

Install it from https://go.dev/dl/

### Build the `Example/Service`

> Make sure the port `80` is currently free. **Optionally could be changed in the code!**

    go build -o ./bin/ ./examples/service.go

Then run the service

    ./bin/service ./dictionaries/es.dic

### Training

Most of the training was made using free versions of books in `Spanish`. However, if you like to train for a new language you can use the following functions

```go
package main

import (
    sc "github.com/jorelosorio/spellingcorrector"
)

func main() {
    dic, _ := sc.NewDictionary("{YOUR_PATH_TO_DICTIONARY}", sc.ESAlphabet) // Or ENAlphabet
    dic.TrainFromTextFile("{YOUR_INPUT_TEXT}")
}
```

> Call `TrainFromTextFile` function as many times you wish with different inputs.

### Build Docker

To build the docker image use `.dockers/Dockerfile.deploy` and the command

    docker build -f Dockerfile.deploy -t jorelosorio/spellingcorrector:latest .

To run the docker image

    docker run --name spellingcorrector -d -p 8080:80 -t jorelosorio/spellingcorrector:latest

Test the `spelling corrector` from the docker image

    http://localhost:8080/spelling?word=espanol
