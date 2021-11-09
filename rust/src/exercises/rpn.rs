#[derive(Debug)]
pub enum CalculatorInput {
  Add,
  Subtract,
  Multiply,
  Divide,
  Value(i32),
}

use CalculatorInput::{Add, Divide, Multiply, Subtract, Value};

pub fn evaluate(inputs: &[CalculatorInput]) -> Option<i32> {
  println!("Inputs: {:?}", inputs);
  if inputs.len() == 0 {
    return None;
  }

  let mut stack: Vec<i32> = Vec::new();

  for input in inputs {
    if let Value(val) = input {
      stack.push(*val);
    } else {
      if stack.len() < 2 {
        return None;
      }

      let last = stack.pop().unwrap();
      let prev = stack.pop().unwrap();

      match input {
        Add => stack.push(prev + last),
        Subtract => stack.push(prev - last),
        Multiply => stack.push(prev * last),
        Divide => stack.push(prev / last),
        _ => {}
      }
    }
  }

  if stack.len() > 1 {
    return None;
  }

  return stack.pop();
}

#[test]
fn test_simple_addition() {
  assert_eq!(evaluate(&[Value(4), Value(8), Add]), Some(12));
}

#[test]
fn test_simple_subtraction() {
  assert_eq!(evaluate(&[Value(8), Value(5), Subtract]), Some(3));
}

#[test]
fn test_simple_multiplication() {
  assert_eq!(evaluate(&[Value(8), Value(5), Multiply]), Some(40));
}

#[test]
fn test_simple_division() {
  assert_eq!(evaluate(&[Value(20), Value(5), Divide]), Some(4));
}

#[test]
fn test_complex_operation() {
  assert_eq!(
    evaluate(&[
      Value(4),
      Value(8),
      Add,
      Value(7),
      Value(5),
      Subtract,
      Divide
    ]),
    Some(6)
  );
}

#[test]
fn test_empty_input_returns_none() {
  assert_eq!(evaluate(&[]), None);
}

#[test]
fn test_too_few_operands_returns_none() {
  assert_eq!(evaluate(&[Value(2), Add]), None);
}

#[test]
fn test_too_many_operands_returns_none() {
  assert_eq!(evaluate(&[Value(2), Value(2), Value(3), Subtract]), None);
}
