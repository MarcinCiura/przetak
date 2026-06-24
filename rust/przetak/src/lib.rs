use przetak_sys as sys;

/// Bit mask from evaluator function represented as an enum type
pub enum EvaluationResult {
    Normal = 0,
    Abusive = 1,
    VulgarN = 2,
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

/// Returns the library version as a `String`
pub fn version() -> String {
    unsafe { sys::version().into() }
}

/// Evaluates given text for occurrence of polish vulgar words
pub fn evaluate(s: &str) -> EvaluationResult {
    let result = unsafe { sys::evaluate(s.into()) };
    result.into()
}
