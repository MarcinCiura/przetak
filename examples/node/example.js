#!/usr/bin/node

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

// Example Node.js program using the przetak shared library.

const ffi = require("ffi");
const ref = require("ref");
const StructType = require("ref-struct")

const CString = ref.types.CString;
const int32 = ref.types.int32;
const longlong = ref.types.longlong;

const GoString = StructType({
  p: CString,
  n: longlong
});

var libprzetak = "../../libprzetak.so";
if (process.platform === "darwin") {
  libprzetak = "../../libprzetak.dylib";
} else if (process.platform === "win32") {
  libprzetak = "../../przetak.dll";
}
const przetak = ffi.Library(libprzetak, {
  version: [GoString, []],
  evaluate: [int32, [GoString]]
});

function version() {
  var pv = przetak.version();
  return pv["p"].substring(0, pv["n"]);
}

function evaluate(text) {
  var b = Buffer.from(text, "utf8");
  var pt = new GoString();
  pt["p"] = b;
  pt["n"] = b.length;
  return przetak.evaluate(pt);
}

const readline = require("readline");
const rl = readline.createInterface({
  input: process.stdin,
  output: process.stdout
});

rl.question(`Przetak ${version()}> `, (text) => {
  console.log(evaluate(text));
  rl.close();
});
