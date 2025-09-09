package generator

import (
	"bytes"
	"go/types"
	"log"

	"github.com/formalco/rpcglue/provider"

	"golang.org/x/tools/imports"
)

type GenerateInput struct {
	Provider    provider.Provider
	PackageName string
	Service     string

	Funcs []*types.Func
}

func Generate(in GenerateInput) ([]byte, error) {
	data := TemplateData{
		Package: in.PackageName,
		Service: in.Service,
	}

	emptyStruct := types.NewStruct(nil, nil)

	resolver := newResolver()
	for _, f := range in.Funcs {
		argT := in.Provider.GetArgType(f)
		replyT := in.Provider.GetReplyType(f)

		var argType string
		if !types.Identical(argT, emptyStruct) {
			argType = resolver.GetTypeString(argT)
		}

		actualReplyType := resolver.GetTypeString(replyT)
		var replyType string
		if !types.Identical(replyT, emptyStruct) {
			replyType = actualReplyType
		}

		data.Methods = append(data.Methods, MethodTemplate{
			Name:            f.Name(),
			ArgType:         argType,
			ReplyType:       replyType,
			ActualReplyType: actualReplyType,
		})
	}

	data.Imports = resolver.GetImports()

	var src bytes.Buffer
	err := tmpl.Execute(&src, data)
	if err != nil {
		log.Printf("failed to render template: %s", err.Error())
		return nil, err
	}

	formatted, err := imports.Process("client.go", src.Bytes(), nil)
	if err != nil {
		log.Printf("failed to format code: \nCODE:\n%s \nERR:\n%s", src.String(), err.Error())
		return nil, err
	}

	return formatted, err
}
