// This stub file contains items which aren't used yet; feel free to remove this module attribute
// to enable stricter warnings.
#![allow(unused)]

use std::collections::HashMap;

pub fn construct_map(coll: &[&str]) -> HashMap<String, u8> {
  let mut map = HashMap::new();
  for word in coll {
    let count = *map.entry(word.to_string()).or_insert(0);
    map.insert(word.to_string(), count + 1);
  }
  return map;
}

pub fn can_construct_note(magazine: &[&str], note: &[&str]) -> bool {
  let magazine_map = construct_map(magazine);
  let note_map = construct_map(note);

  for (word, count) in note_map {
    if magazine_map.get(&word).unwrap_or(&0) < &count {
      return false;
    }
  }
  return true;
}

#[test]
fn test_false() {
  let magazine = "two times three is not four"
    .split_whitespace()
    .collect::<Vec<&str>>();
  let note = "two times two is four"
    .split_whitespace()
    .collect::<Vec<&str>>();

  assert!(!can_construct_note(&magazine, &note));
}

#[test]
fn test_true() {
  let magazine = "two times two is four"
    .split_whitespace()
    .collect::<Vec<&str>>();
  let note = "two times two is four"
    .split_whitespace()
    .collect::<Vec<&str>>();

  assert!(can_construct_note(&magazine, &note));
}

#[test]
fn test_capital_false() {
  let magazine = "Two times two is four"
    .split_whitespace()
    .collect::<Vec<&str>>();
  let note = "two times two is four"
    .split_whitespace()
    .collect::<Vec<&str>>();

  assert!(!can_construct_note(&magazine, &note));
}

#[test]
fn test_capital_true() {
  let magazine = "Two times two is four"
    .split_whitespace()
    .collect::<Vec<&str>>();
  let note = "two times Two is four"
    .split_whitespace()
    .collect::<Vec<&str>>();

  assert!(can_construct_note(&magazine, &note));
}

#[test]
fn test_more_words_than_needed_true() {
  let magazine = "two times two is four and one times two is two"
    .split_whitespace()
    .collect::<Vec<&str>>();
  let note = "two times two is four"
    .split_whitespace()
    .collect::<Vec<&str>>();

  assert!(can_construct_note(&magazine, &note));
}
