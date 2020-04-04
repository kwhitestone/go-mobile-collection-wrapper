// Copyright 2014 Brett Slatkin
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package syncMapWrapper

import (
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"strings"
)

func loadFile(inputPath string) (string, []GeneratedType) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, inputPath, nil, parser.ParseComments)
	if err != nil {
		log.Fatalf("Could not parse file: %s", err)
	}

	packageName := identifyPackage(f)
	if packageName == "" {
		log.Fatalf("Could not determine package name of %s", inputPath)
	}

	collectionTypes := make(map[string][]string)
	for _, decl := range f.Decls {
		kTypeNames, vTypeName, ok := identifyCollectionType(decl)
		if ok {
			collectionTypes[vTypeName] = kTypeNames
		}
	}

	var types []GeneratedType
	for vTypeName, kTypeNames := range collectionTypes {
		for _, kTypeName := range kTypeNames {
			collectionType := GeneratedType{
				Key:      kTypeName,
				KeyTitle: strings.Title(kTypeName),
				Value:    vTypeName,
			}
			types = append(types, collectionType)
		}
	}

	return packageName, types
}

func identifyPackage(f *ast.File) string {
	if f.Name == nil {
		return ""
	}
	return f.Name.Name
}

func identifyCollectionType(decl ast.Decl) (keyTypeNames []string, valueTypeName string, match bool) {
	genDecl, ok := decl.(*ast.GenDecl)
	if !ok {
		return
	}
	if genDecl.Doc == nil {
		return
	}

	found := false
	for _, comment := range genDecl.Doc.List {
		if strings.Contains(comment.Text, "@sync-map-wrapper") {
			i := strings.Index(comment.Text, "@sync-map-wrapper")
			rest := comment.Text[i:]
			comps := strings.SplitAfter(rest, ":")
			if len(comps) > 1 {
				keys := comps[1]
				keyTypeNames = strings.Split(keys, ",")
				found = true
				break
			}
		}
	}
	if !found {
		return
	}

	for _, spec := range genDecl.Specs {
		if typeSpec, ok := spec.(*ast.TypeSpec); ok {
			if typeSpec.Name != nil {
				valueTypeName = typeSpec.Name.Name
				break
			}
		}
	}
	if valueTypeName == "" {
		return
	}

	match = true
	return
}
