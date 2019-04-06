#!/usr/bin/python3

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

"""Example Python 3.x program using the przetak shared library."""

import ctypes
import platform


class GoString(ctypes.Structure):
    _fields_ = [
        ('p', ctypes.c_char_p),
        ('n', ctypes.c_longlong),
    ]


libprzetak = '../../libprzetak.so'
if platform.system() == 'Darwin':
    libprzetak = '../../libprzetak.dylib'
elif platform.system() == 'Windows':
    libprzetak = '../../przetak.dll'
przetak = ctypes.cdll.LoadLibrary(libprzetak)
przetak.version.argtypes = []
przetak.version.restype = GoString
przetak.evaluate.argtypes = [GoString]
przetak.evaluate.restype = ctypes.c_int


def version():
    pv = przetak.version()
    return pv.p[:pv.n].decode('utf-8')


def evaluate(text):
    b = text if isinstance(text, bytes) else bytes(str(text), 'utf-8')
    pt = ctypes.create_string_buffer(b)
    return przetak.evaluate(GoString(pt.raw, ctypes.sizeof(pt)))


if __name__ == '__main__':
    text = input('Przetak {}> '.format(version()))
    print(evaluate(text))
