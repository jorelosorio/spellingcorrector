# Spelling Corrector
 
> NOTE: This is a project for learning purposes!

The solution for this project was based on the proposal made on the following website: http://norvig.com/spell-correct.html and some ideas from https://cxwangyi.wordpress.com/2012/02/15/peter-norvigs-spelling-corrector-in-go/ as well.

> The difference is that the algorithm has been trained using the `Spanish` language.

## Try it

To use this project using a `Docker` image, please run the following commands

    docker pull jorelosorio/spelling-corrector:latest

    docker run --name spelling-corrector -d -p 8080:80 -t jorelosorio/spelling-corrector:latest

Try it using the following example:

    http://localhost:8080/spelling?word=espanol

## Training

Most of the training was made using free versions of books in `Spanish`.

## Development

Build the project, if you want to run it manually

> Make sure the port `80` is currently free.

    go build -o spelling .

Then run the server

    ./spelling

> NOTE: This execution assumes a `Linux` OS.

### Build Docker

To build the docker image use `Dockerfile.deploy` and the command

    docker build -f Dockerfile.deploy -t jorelosorio/spelling-corrector:latest .

To run the docker image

    docker run --name spelling-corrector -d -p 8080:80 -t jorelosorio/spelling-corrector:latest

Test the `spelling corrector` from the docker image

    http://localhost:8080/spelling?word={YOUR_WORD}
