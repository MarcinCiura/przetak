//! Safe Rust bindings to the Przetak library.
//!
//! Przetak detects Polish vulgar and abusive speech in UTF-8 text. It is resilient to
//! replicating letters, spacing out words, inserting non-letters between letters,
//! and homograph spoofing (replacing letters with similar Unicode characters).
//!
//! # Example
//!
//! ```
//! use przetak::{evaluate, EvaluationResult};
//!
//! assert_eq!(evaluate("przykładowy tekst"), EvaluationResult::Normal);
//! ```

use przetak_sys as sys;

/// Result of evaluating text for toxic speech.
///
/// A bitmask indicating which categories of toxic speech were detected.
#[derive(Clone, Copy, Debug, PartialEq, Eq)]
pub enum EvaluationResult {
    /// No toxic speech detected.
    Normal = 0,
    /// Abusive words detected.
    Abusive = 1,
    /// Vulgar words with negative connotations detected.
    VulgarN = 2,
    /// Vulgar words with positive connotations detected.
    VulgarP = 4,
}

impl std::fmt::Display for EvaluationResult {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        match self {
            EvaluationResult::Normal => write!(f, "Normal"),
            EvaluationResult::Abusive => write!(f, "Abusive"),
            EvaluationResult::VulgarN => write!(f, "VulgarN"),
            EvaluationResult::VulgarP => write!(f, "VulgarP"),
        }
    }
}

impl From<i32> for EvaluationResult {
    fn from(value: i32) -> Self {
        match value {
            0 => EvaluationResult::Normal,
            1 => EvaluationResult::Abusive,
            2 => EvaluationResult::VulgarN,
            4 => EvaluationResult::VulgarP,
            _ => unreachable!(),
        }
    }
}

/// Returns the version string of the underlying Przetak library.
///
/// The version is a date in `YYYY-MM-DD` format representing the last modification
/// of the Przetak package.
pub fn version() -> String {
    unsafe { sys::version().into() }
}

/// Evaluates the given UTF-8 text for Polish vulgar and abusive speech.
///
/// Returns an [`EvaluationResult`] bitmask whose bits are set if the input contains:
/// - `Abusive` (value `1`) — abusive words
/// - `VulgarN` (value `2`) — vulgar words with negative connotations
/// - `VulgarP` (value `4`) — vulgar words with positive connotations
///
/// Multiple bits may be set simultaneously.
pub fn evaluate(s: &str) -> EvaluationResult {
    let result = unsafe { sys::evaluate(s.into()) };
    result.into()
}
