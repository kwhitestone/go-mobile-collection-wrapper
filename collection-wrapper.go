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

package main

import (
	"github.com/kwhitestone/go-mobile-collection-wrapper/mapWrapper"
	"github.com/kwhitestone/go-mobile-collection-wrapper/sliceWrapper"
	"github.com/kwhitestone/go-mobile-collection-wrapper/syncMapWrapper"
	"log"
	"os"
)

func processFile(inputPath string) {
	log.Printf("Processing file %s", inputPath)
	typesToGenerateSlices := mapWrapper.ProcessFile(inputPath)
	if len(typesToGenerateSlices)==0 {
		typesToGenerateSlices = syncMapWrapper.ProcessFile(inputPath)
	}
	sliceWrapper.ProcessFile(inputPath, typesToGenerateSlices)
}

func main() {
	log.SetFlags(0)
	log.SetPrefix("collection-wrapper: ")

	for _, path := range os.Args[1:] {
		processFile(path)
	}
}
