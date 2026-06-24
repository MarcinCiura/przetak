use libloading::{Library, Symbol};
use std::ffi::CString;
use std::os::raw::c_char;

#[repr(C)]
struct GoString {
    p: *const c_char,
    n: isize,
}

fn load_library() -> Library {
    let lib_path = if cfg!(target_os = "windows") {
        "../../przetak.dll"
    } else if cfg!(target_os = "macos") {
        "../../libprzetak.dylib"
    } else {
        "../../libprzetak.so"
    };
    unsafe { Library::new(lib_path).expect("Failed to load przetak library") }
}

fn version(lib: &Library) -> String {
    unsafe {
        let func: Symbol<unsafe extern "C" fn() -> GoString> = lib.get(b"version").unwrap();
        let gs = func();
        let slice = std::slice::from_raw_parts(gs.p as *const u8, gs.n as usize);
        std::str::from_utf8(slice).unwrap().to_string()
    }
}

fn evaluate(lib: &Library, text: &str) -> i32 {
    unsafe {
        let func: Symbol<unsafe extern "C" fn(GoString) -> i32> = lib.get(b"evaluate").unwrap();
        let c_text = CString::new(text).unwrap();
        let gs = GoString {
            p: c_text.as_ptr(),
            n: c_text.as_bytes().len() as isize,
        };
        func(gs)
    }
}

fn main() {
    let lib = load_library();
    let v = version(&lib);
    print!("Przetak {}> ", v);
    let mut input = String::new();
    std::io::stdin().read_line(&mut input).unwrap();
    println!("{}", evaluate(&lib, input.trim()));
}
