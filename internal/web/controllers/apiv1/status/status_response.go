package status

import "github.com/Dubrovsky18/backend-trainee-assignment-2023/internal/app/build"

type Response struct {
	ID     string      `jsonapi:"primary,status"`
	Status string      `jsonapi:"attr,status"`
	Build  *build.Info `jsonapi:"attr,build"`
}
