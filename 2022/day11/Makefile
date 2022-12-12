all: input

example: gen_example main

input: gen_input main

main:
	go run .

gen_example:
	go run gen/gen.go < example_input

gen_input:
	go run gen/gen.go < input
