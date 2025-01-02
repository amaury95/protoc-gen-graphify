// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The protoc-gen-go binary is a protoc plugin to generate Go code for
// both proto2 and proto3 versions of the protocol buffer language.
//
// For more information about the usage of this plugin, see:
// https://protobuf.dev/reference/go/go-generated.
package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/amaury95/protoc-gen-graphify/internal/version"
	gengo "google.golang.org/protobuf/cmd/protoc-gen-go/internal_gengo"
	"google.golang.org/protobuf/compiler/protogen"
)

const genGoDocURL = "https://protobuf.dev/reference/go/go-generated"
const grpcDocURL = "https://grpc.io/docs/languages/go/quickstart/#regenerate-grpc-code"

func main() {
	if len(os.Args) == 2 && os.Args[1] == "--version" {
		fmt.Fprintf(os.Stdout, "%v %v\n", filepath.Base(os.Args[0]), version.String())
		os.Exit(0)
	}
	if len(os.Args) == 2 && os.Args[1] == "--help" {
		fmt.Fprintf(os.Stdout, "See "+genGoDocURL+" for usage information.\n")
		os.Exit(0)
	}

	var (
		flags   flag.FlagSet
		plugins = flags.String("plugins", "", "deprecated option")
	)
	protogen.Options{
		ParamFunc: flags.Set,
	}.Run(func(gen *protogen.Plugin) error {
		if *plugins != "" {
			return errors.New("protoc-gen-graphify: plugins are not supported; use 'protoc --go-grpc_out=...' to generate gRPC\n\n" +
				"See " + grpcDocURL + " for more information")
		}
		for _, f := range gen.Files {
			if f.Generate {
				messages := walkMessages(f.Messages)
				updateMessageNames(messages...)

				g := gengo.GenerateFile(gen, f)

				g.Write([]byte("// " + strings.Join(os.Args, " ") + "\n"))

				// if flags.Lookup("generate_graphql_schema").Value.String() != "false" {
				generateGraphql(g, f, messages...)
				// }
				// if flags.Lookup("generate_object_schema").Value.String() != "false" {
				generateObjectSchema(g, f, messages...)
				// }
				// if flags.Lookup("generate_unmarshaler").Value.String() != "false" {
				generateUnmarshaler(g, f, messages...)
				// }
			}
		}
		gen.SupportedFeatures = gengo.SupportedFeatures
		return nil
	})
}

func walkMessages(messages []*protogen.Message) []*protogen.Message {
	var result []*protogen.Message
	for _, message := range messages {
		result = append(result, message)
		result = append(result, walkMessages(message.Messages)...)
	}
	return result
}
