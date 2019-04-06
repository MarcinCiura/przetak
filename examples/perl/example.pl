#!/usr/bin/perl

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

# Example Perl program using the przetak shared library.

package My::Przetak;

use FFI::Platypus;

my $libprzetak = '../../libprzetak.so';
if ($^O eq 'darwin') {
  $libprzetak = '../../libprzetak.dylib';
} elsif ($^O eq 'MSWin32') {
  $libprzetak = '../../przetak.dll';
}
my $ffi = FFI::Platypus->new( lib => $libprzetak );
$ffi->attach( ['evaluatePP' => 'evaluate'] => ['string*', 'sint32*', 'sint32*'] => 'void', sub {
  my($fn, $text) = @_;
  my $result = 0;
  $fn->( \$text, \(length $text), \$result );
  return $result;
});

package main;

print( "Przetak> " );
$text = readline STDIN;
print( My::Przetak::evaluate( $text ), "\n" );
