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

/* Example C program using the przetak shared library. */

#include <stdio.h>
#include <string.h>

#ifdef _WIN32
# include "../../przetak.h"
#else
# include "../../libprzetak.h"
#endif

const char *przetak_version() {
  static char v[10 + 1];

  if (v[0] == '\0') {
    GoString pv = version();
    memcpy(v, pv.p, pv.n < 10 ? pv.n : 10);
  }
  return v;
}

int przetak_evaluate(const char *text) {
  GoString pt;

  pt.p = text;
  pt.n = strlen(text);
  return evaluate(pt);
}

int main() {
  char text[257];

  printf("Przetak %s> ", przetak_version());
  if (fgets(text, sizeof text, stdin) != NULL) {
    printf("%d\n", przetak_evaluate(text));
  }
  return 0;
}
