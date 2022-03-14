# Spelling Corrector
 
> NOTE: This is a project for learning purposes!

The solution for this project was based on the proposal made on the following website: http://norvig.com/spell-correct.html and some ideas from https://cxwangyi.wordpress.com/2012/02/15/peter-norvigs-spelling-corrector-in-go/ as well.

> The difference is that the algorithm has been trained using the `Spanish` language.

## Training

Most of the training was made using free versions of books in `Spanish`

## Development

Build the project with

    go build -o spelling .

## Test the project

    http://localhost:8080/spelling?word={YOUR_WORD}
