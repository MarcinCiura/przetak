# `przetak`

Safe Rust bindings to the [Przetak](https://github.com/piotreknow02/przetak) library for detecting Polish vulgar and abusive speech.

## Installation

```toml
[dependencies]
przetak = "0.1"
```

## Usage

```rust
use przetak::{evaluate, EvaluationResult};

assert_eq!(evaluate("przykładowy tekst"), EvaluationResult::Normal);
```

## License

Licensed under the [Apache License, Version 2.0](https://www.apache.org/licenses/LICENSE-2.0).
