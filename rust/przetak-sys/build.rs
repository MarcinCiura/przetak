use std::env;
use std::path::PathBuf;
use std::process::Command;

fn main() {
    // Tell cargo to look for the shared library in the parent directory
    let manifest_dir = env::var("CARGO_MANIFEST_DIR").unwrap();
    let manifest_path = PathBuf::from(&manifest_dir);
    let parent_dir = manifest_path.parent().unwrap();
    let grandparent_dir = parent_dir.parent().unwrap();

    // Build the Go static library via the Makefile in the parent directory
    let make_status = Command::new("make")
        .arg("static")
        .arg("-C")
        .arg(grandparent_dir)
        .status()
        .expect("Failed to invoke make to build libprzetak.a");

    if !make_status.success() {
        panic!("make static failed to build libprzetak.a");
    }

    println!("cargo:rustc-link-search={}", grandparent_dir.display());

    // Tell cargo to link the static library
    println!("cargo:rustc-link-lib=static=przetak");

    // Tell cargo to re-run when the header file changes
    println!(
        "cargo:rerun-if-changed={}/libprzetak.h",
        grandparent_dir.display()
    );

    // The bindgen::Builder is the main entry point
    let header_path = grandparent_dir.join("libprzetak.h");
    let header_path_str = header_path.to_str().unwrap();

    let bindings = bindgen::Builder::default()
        // The input header we would like to generate bindings for
        .header(header_path_str)
        // Tell cargo to invalidate the built crate whenever any of the
        // included header files changed
        .parse_callbacks(Box::new(bindgen::CargoCallbacks::new()))
        // Allowlist only the public API functions
        .allowlist_function("version")
        .allowlist_function("evaluate")
        // Generate Rust-style enums and derive traits
        .default_enum_style(bindgen::EnumVariation::Rust {
            non_exhaustive: false,
        })
        // Generate partial_eq and eq traits
        .derive_partialeq(true)
        .derive_eq(true)
        .derive_debug(true)
        .derive_copy(true)
        // Disable layout tests in release builds for better performance
        .disable_header_comment()
        .size_t_is_usize(true);

    let bindings = bindings.generate().expect("Unable to generate bindings");

    // Write the bindings to the $OUT_DIR/bindings.rs file
    let out_path = PathBuf::from(env::var("OUT_DIR").unwrap());
    // Add attributes to allow unsafe_code in extern blocks
    bindings
        .write_to_file(out_path.join("bindings.rs"))
        .expect("Couldn't write bindings!");
}
