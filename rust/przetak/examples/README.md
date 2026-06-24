# Rust Example

This example demonstrates how to use the `przetak` crate to evaluate Polish text.

## Running

```bash
cargo run --example example
```

If the shared library is not in the system linker path, set `LD_LIBRARY_PATH`:

```bash
LD_LIBRARY_PATH=$(readlink -f $(pwd)/../..) cargo run --example example
```

The program reads a line from stdin and prints the evaluation result:
- `Normale`  = 0, no toxic speech
- `Abusive` = 1, abusive speech
- `VulgarN` = 2, vulgar with negative connotations
- `VulgarP` = 4, vulgar with positive connotations
