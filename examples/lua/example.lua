#!/usr/bin/luajit

--[[
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
--]]

local libprzetak = "../../libprzetak.so"
if ffi.os == "OSX" then
  libprzetak = "../../libprzetak.dylib"
elseif ffi.os == "Windows" then
  libprzetak = "../../przetak.dll"
end
local ffi = require("ffi")
local przetak = ffi.load(libprzetak)

ffi.cdef([[
  typedef int GoInt32;
  typedef long long GoInt64;
  typedef GoInt64 GoInt;
  typedef struct { const char *p; GoInt n; } GoString;

  extern GoString version();
  extern GoInt32 evaluate(GoString text);
]])

function version()
  local pv = przetak.version()
  return ffi.string(pv.p, pv.n)
end

function evaluate(text)
  local pt = ffi.new("GoString", {text, string.len(text)})
  return przetak.evaluate(pt)
end

io.write(string.format("Przetak %s> ", version()))
io.write(string.format("%d\n", evaluate(io.read())))
