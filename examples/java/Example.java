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

// Example Java program using the przetak shared library.

import static java.nio.charset.StandardCharsets.UTF_8;
import com.sun.jna.Library;
import com.sun.jna.Native;
import com.sun.jna.Structure;
import java.util.Arrays;
import java.util.List;

public class Example {

    public interface Przetak extends Library {
        public class GoString extends Structure {
            public static class ByValue extends GoString implements Structure.ByValue {}
            public String p;
            public long n;
            protected List<String> getFieldOrder() {
                return Arrays.asList(new String[]{"p", "n"});
            }
        }
        public GoString.ByValue version();
        public int evaluate(GoString.ByValue text);
    }

    static private Przetak przetak =
        (Przetak) Native.loadLibrary("przetak", Przetak.class);
        // Or a full system-specific path, e.g. "../../libprzetak.so".

    static public String version() {
        Przetak.GoString.ByValue pv = przetak.version();
        return pv.p.substring(0, (int) pv.n);
    }

    static public long evaluate(String text) {
        Przetak.GoString.ByValue pt = new Przetak.GoString.ByValue();
        pt.p = text;
        pt.n = text.getBytes(UTF_8).length;
        return przetak.evaluate(pt);
    }

    static public void main(String argv[]) {
        System.out.printf("Przetak %s> ", version());
        String text = System.console().readLine();
        System.out.printf("%d\n", evaluate(text));
    }
}
