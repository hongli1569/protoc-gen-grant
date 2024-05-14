package generate

import (
	"fmt"

	"github.com/hongli1569/protoc-gen-grant/annotations"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
)

const (
	annotationsPackage = protogen.GoImportPath("github.com/hongli1569/protoc-gen-grant/annotations")
)

type ServiceGrant struct {
	ServiceName string
	Grant       []*annotations.Grant
}

func GenerateFile(gen *protogen.Plugin, file *protogen.File) *protogen.GeneratedFile {
	filename := file.GeneratedFilenamePrefix + "_grant.pb.go"
	g := gen.NewGeneratedFile(filename, file.GoImportPath)

	g.P("// Code generated by protoc-gen-grant. DO NOT EDIT.")

	if file.Proto.GetOptions().GetDeprecated() {
		g.P("// ", file.Desc.Path(), " is a deprecated file.")
	} else {
		g.P("// source: ", file.Desc.Path())
	}
	g.P()

	g.P("package ", file.GoPackageName)
	g.P()

	generateFileContent(gen, file, g)

	return g
}

func generateFileContent(gen *protogen.Plugin, file *protogen.File, g *protogen.GeneratedFile) {
	if len(file.Services) == 0 {
		return
	}

	g.P("// This is a compile-time assertion to ensure that this generated file")
	g.P("var _ = new(", annotationsPackage.Ident("Grant"), ")")

	for _, service := range file.Services {
		genService(gen, file, g, service)
	}
}

func genService(_ *protogen.Plugin, file *protogen.File, g *protogen.GeneratedFile, service *protogen.Service) {
	serviceGrant := &ServiceGrant{
		ServiceName: service.GoName,
		Grant:       make([]*annotations.Grant, 0),
	}

	for _, method := range service.Methods {
		grant, ok := proto.GetExtension(method.Desc.Options(), annotations.E_Grant).(*annotations.Grant)
		if grant == nil || !ok {
			continue
		}
		grant.FullName = fmt.Sprintf("%s_%s_FullMethodName", service.GoName, method.GoName)
		serviceGrant.Grant = append(serviceGrant.Grant, grant)
	}

	g.P(execTpl(serviceGrant))
}