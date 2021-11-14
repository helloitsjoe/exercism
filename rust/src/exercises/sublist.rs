#[derive(Debug, PartialEq)]
pub enum Comparison {
  Equal,
  Sublist,
  Superlist,
  Unequal,
}

pub fn compare<T: PartialEq>(first: &[T], second: &[T], comparison: Comparison) -> Comparison {
  for (i, el) in second.iter().enumerate() {
    if &first[0] != el {
      continue;
    }
    let end = i + first.len();
    if end > second.len() {
      return Comparison::Unequal;
    }
    let slice = &second[i..i + first.len()];
    if first == slice {
      return comparison;
    }
  }

  return Comparison::Unequal;
}

pub fn sublist<T: PartialEq>(first: &[T], second: &[T]) -> Comparison {
  if first.len() == second.len() {
    if first == second {
      return Comparison::Equal;
    }
  }

  if first.len() == 0 {
    return Comparison::Sublist;
  }

  if second.len() == 0 {
    return Comparison::Superlist;
  }

  if first.len() < second.len() {
    return compare(first, second, Comparison::Sublist);
  } else {
    return compare(second, first, Comparison::Superlist);
  }

  return Comparison::Unequal;
}

#[test]
fn empty_equals_empty() {
  let v: &[u32] = &[];
  assert_eq!(Comparison::Equal, sublist(v, v));
}
#[test]
fn test_empty_is_a_sublist_of_anything() {
  assert_eq!(Comparison::Sublist, sublist(&[], &['a', 's', 'd', 'f']));
}
#[test]
fn test_anything_is_a_superlist_of_empty() {
  assert_eq!(Comparison::Superlist, sublist(&['a', 's', 'd', 'f'], &[]));
}
#[test]
fn test_1_is_not_2() {
  assert_eq!(Comparison::Unequal, sublist(&[1], &[2]));
}
#[test]
fn test_compare_larger_equal_lists() {
  use std::iter::repeat;
  let v: Vec<char> = repeat('x').take(1000).collect();
  assert_eq!(Comparison::Equal, sublist(&v, &v));
}
#[test]
fn test_sublist_at_start() {
  assert_eq!(Comparison::Sublist, sublist(&[1, 2, 3], &[1, 2, 3, 4, 5]));
}
#[test]
fn sublist_in_middle() {
  assert_eq!(Comparison::Sublist, sublist(&[4, 3, 2], &[5, 4, 3, 2, 1]));
}
#[test]
fn sublist_at_end() {
  assert_eq!(Comparison::Sublist, sublist(&[3, 4, 5], &[1, 2, 3, 4, 5]));
}
#[test]
fn partially_matching_sublist_at_start() {
  assert_eq!(Comparison::Sublist, sublist(&[1, 1, 2], &[1, 1, 1, 2]));
}
#[test]
fn sublist_early_in_huge_list() {
  let huge: Vec<u32> = (1..1_000_000).collect();
  assert_eq!(Comparison::Sublist, sublist(&[3, 4, 5], &huge));
}
#[test]
fn huge_sublist_not_in_huge_list() {
  let v1: Vec<u64> = (10..1_000_001).collect();
  let v2: Vec<u64> = (1..1_000_000).collect();
  assert_eq!(Comparison::Unequal, sublist(&v1, &v2));
}
#[test]
fn superlist_at_start() {
  assert_eq!(Comparison::Superlist, sublist(&[1, 2, 3, 4, 5], &[1, 2, 3]));
}
#[test]
fn superlist_in_middle() {
  assert_eq!(Comparison::Superlist, sublist(&[5, 4, 3, 2, 1], &[4, 3, 2]));
}
#[test]
fn superlist_at_end() {
  assert_eq!(Comparison::Superlist, sublist(&[1, 2, 3, 4, 5], &[3, 4, 5]));
}
#[test]
fn superlist_early_in_huge_list() {
  let huge: Vec<u32> = (1..1_000_000).collect();
  assert_eq!(Comparison::Superlist, sublist(&huge, &[3, 4, 5]));
}
#[test]
fn recurring_values_sublist() {
  assert_eq!(
    Comparison::Sublist,
    sublist(&[1, 2, 1, 2, 3], &[1, 2, 3, 1, 2, 1, 2, 3, 2, 1])
  );
}
#[test]
fn recurring_values_unequal() {
  assert_eq!(
    Comparison::Unequal,
    sublist(&[1, 2, 1, 2, 3], &[1, 2, 3, 1, 2, 3, 2, 3, 2, 1])
  );
}
