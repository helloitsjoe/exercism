/// Create a buffer of `count` zeroes.
///
/// Applications often use buffers when serializing data to send over the network.
pub fn create_buffer(count: usize) -> Vec<u8> {
  (0..count).map(|_| 0).collect::<Vec<u8>>()
}

/// Create a vector containing the first five elements of the Fibonacci sequence.
///
/// Fibonacci's sequence is the list of numbers where the next number is a sum of the previous two.
/// Its first five elements are `1, 1, 2, 3, 5`.
pub fn fibonacci() -> Vec<u8> {
  let mut vec = Vec::from([1, 1]);
  while vec.len() < 5 {
    let last = vec[vec.len() - 1];
    let last_last = vec[vec.len() - 2];
    vec.push(last_last + last);
  }
  vec
}

#[test]
fn test_basic() {
  assert_eq!(create_buffer(3), Vec::from([0, 0, 0]));
}

#[test]
fn test_fib() {
  assert_eq!(fibonacci(), Vec::from([1, 1, 2, 3, 5]));
}
