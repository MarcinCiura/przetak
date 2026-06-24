use std::ffi::c_char;

use crate::GoString;

/// Convert go string to a string slice reference `&str`
impl AsRef<str> for GoString {
    fn as_ref(&self) -> &str {
        if self.n <= 0 {
            return "";
        }

        let len = self.n as usize;
        unsafe {
            let bytes = std::slice::from_raw_parts(self.p as *const u8, len);
            std::str::from_utf8_unchecked(bytes)
        }
    }
}

/// Convert `GoString` into rust dynamic `String`
impl Into<String> for GoString {
    fn into(self) -> String {
        self.as_ref().to_owned()
    }
}

/// Convert string slice to `GoString`
impl From<&str> for GoString {
    fn from(s: &str) -> Self {
        GoString {
            p: s.as_ptr() as *const c_char,
            n: s.len() as isize,
        }
    }
}

/// Convert rust dynamic `String` to `GoString`
impl From<String> for GoString {
    fn from(s: String) -> Self {
        GoString {
            p: s.as_ptr() as *const c_char,
            n: s.len() as isize,
        }
    }
}
