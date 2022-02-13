/*
Copyright 2019 Marcin Ciura

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Example Go program using package przetak.
package main

import (
	"bufio"
	"fmt"
	"github.com/MarcinCiura/przetak/go"
	"os"
)

func main() {
	fmt.Printf("Przetak %s> ", przetak.VERSION)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()
	result := przetak.Evaluate(text)
	fmt.Println(result)
}
