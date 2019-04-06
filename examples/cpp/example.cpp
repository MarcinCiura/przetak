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

// Example C++ program using the przetak shared library.

#include <iostream>
#include <string>

#ifdef _WIN32
# include "../../przetak.h"
#else
# include "../../libprzetak.h"
#endif

namespace przetak {

const std::string version() {
  GoString pv = ::version();
  return std::string(pv.p, pv.n);
}

int evaluate(const std::string& text) {
  GoString pt;
  pt.p = text.data();
  pt.n = text.size();
  return ::evaluate(pt);
}

}  // namespace przetak

int main() {
  std::cout << "Przetak " << przetak::version() << "> ";
  std::string text;
  std::getline(std::cin, text);
  std::cout << przetak::evaluate(text) << std::endl;
}
