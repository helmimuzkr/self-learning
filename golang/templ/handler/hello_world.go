package handler

import (
	"github.com/helmimuzkr/templ/types"
	"fmt"
) 

func HelloWorldHandler(ctx *types.Context) error {
	val := ctx.Request.PathValue("params")
	if val != "hello_world" {
		return fmt.Errorf("invalid praams")
	}

	fmt.Fprintln(ctx.ResponseWriter, val)

	return nil
}
