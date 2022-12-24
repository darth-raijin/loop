package dtos

import "github.com/darth-raijin/borealis/api/models/enums"

type Question struct {
	question string
	rating   enums.Rating
}
