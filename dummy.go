package tools

import "github.com/bblfsh/sdk/uast"

type Dummy struct{}

func (d Dummy) Exec(*uast.Node) error {
	println("It works! You can now proceed with another tool :)")
	return nil
}
