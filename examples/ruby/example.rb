#!/usr/bin/ruby

# Copyright 2019 Marcin Ciura
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# Example Ruby program using the przetak shared library.

require 'ffi'

module Przetak
  extend FFI::Library
  libprzetak = '../../libprzetak.so'
  if RUBY_PLATFORM =~ /darwin/
    libprzetak = '../../libprzetak.dylib'
  elsif RUBY_PLATFORM =~ /win32/
    libprzetak = '../../przetak.dll'
  end
  ffi_lib libprzetak
  class GoString < FFI::Struct
    layout :p, :pointer,
           :n, :long_long
  end
  attach_function :version, [], GoString.by_value
  attach_function :evaluate, [GoString.by_value], :int
end

def version()
  pv = Przetak.version()
  return pv[:p].read_string(pv[:n])
end

def evaluate(text)
  pt = Przetak::GoString.new
  pt[:p] = FFI::MemoryPointer.from_string(text)
  pt[:n] = text.bytesize
  return Przetak.evaluate(pt)
end

def main
  print "Przetak #{version()}> "
  print "#{evaluate(STDIN.gets)}\n"
end

if $PROGRAM_NAME == __FILE__
  main()
end
