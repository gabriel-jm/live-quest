package id

import gonanoid "github.com/matoous/go-nanoid/v2"

const alphabet string = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const size int = 21

func NanoId() string {
	id, err := gonanoid.Generate(alphabet, size)

	if err != nil {
		panic("Error generating nano id")
	}

	return id
}
