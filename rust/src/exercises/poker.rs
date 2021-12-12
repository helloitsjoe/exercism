/// Given a list of poker hands, return a list of those hands which win.
///
/// Note the type signature: this function should return _the same_ reference to
/// the winning hand(s) as were passed in, not reconstructed strings which happen to be equal.
use std::cmp::Ordering;
use std::collections::{HashMap, HashSet};
use std::fmt;

type Hand = Vec<(u32, char)>;
type Counts = HashMap<u32, u32>;

struct PokerHand {
  cards: Hand,
  counts: Counts,
}

impl PokerHand {
  fn new(hand: &str) -> Self {
    let cards = hand_as_vec(hand);
    let counts = get_counts(&cards);

    PokerHand { cards, counts }
  }
}

impl std::fmt::Debug for PokerHand {
  fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
    f.debug_struct("PokerHand")
      .field("cards", &self.cards)
      .finish()
  }
}

impl PartialOrd for PokerHand {
  fn partial_cmp(&self, other: &Self) -> Option<Ordering> {
    Some(compare_hands(&self, &other))
  }
}

impl PartialEq for PokerHand {
  fn eq(&self, other: &Self) -> bool {
    compare_hands(&self, &other) == Ordering::Equal
  }
}

pub fn winning_hands<'a>(hands: &[&'a str]) -> Vec<&'a str> {
  // unimplemented!("Out of {:?}, which hand wins?", hands)
  // Sort all hands
  let hands_vec = Vec::from(hands);
  let mut poker_hands_vec = hands_vec
    .iter()
    .map(|h| PokerHand::new(h))
    .collect::<Vec<PokerHand>>();

  poker_hands_vec.sort_by(|a, b| a.partial_cmp(&b).unwrap_or(Ordering::Less));
  let top_hand = &poker_hands_vec[0];

  let mut winners = Vec::new();

  for &hand in hands {
    // TODO: Better way to do this?
    let poker_hand = PokerHand::new(hand);
    if compare_hands(&poker_hand, &top_hand) == Ordering::Equal {
      winners.push(hand);
    }
  }

  winners
}

fn hand_as_vec<'a>(hand: &'a str) -> Hand {
  let mut hand: Hand = hand
    .split(' ')
    .map(|card: &str| {
      let mut ch = card.chars();
      let count = ch.clone().count();

      if count == 3 {
        let suit = ch.skip(2).next().unwrap();
        return (10, suit);
      }

      let val = ch.next().unwrap();
      let val = val.to_digit(10).unwrap_or_else(|| match val {
        'J' => 11,
        'Q' => 12,
        'K' => 13,
        'A' => 14,
        _ => 0,
      });

      let suit = ch.next().unwrap();

      return (val, suit);
    })
    .collect();

  hand.sort_by(|a, b| a.partial_cmp(b).unwrap());

  hand
}

fn is_straight(sorted_hand: &Hand) -> bool {
  println!("{:?}", sorted_hand);
  let straight = sorted_hand.iter().fold(0, |prev, &(curr, _)| {
    if prev == 0 || curr == prev + 1 {
      curr
    } else if curr == 14 && sorted_hand[0].0 == 2 {
      make_ace_low_straight(sorted_hand);
      curr
    } else {
      100
    }
  });

  straight != 100
}

fn is_flush(hand: &Hand) -> bool {
  let flush = hand.iter().fold('X', |prev, &(_, curr)| {
    if prev == 'X' || prev == curr {
      curr
    } else {
      'N'
    }
  });

  flush != 'N'
}

fn get_counts(hand: &Hand) -> Counts {
  let mut hm: Counts = HashMap::new();
  for (val, _ch) in hand {
    hm.insert(*val, *hm.get(val).unwrap_or(&0) + 1);
  }
  hm
}

fn has_count(num: u32, counts: &Counts) -> bool {
  return counts
    .iter()
    .find_map(|(_k, &v)| if v == num { Some(true) } else { None })
    .unwrap_or(false);
}

fn has_two_pair(counts: &Counts) -> bool {
  let mut pairs = 0;
  for (_, &val) in counts {
    if val == 2 {
      pairs += 1;
    }
  }
  pairs == 2
}

fn get_hand_rank<'a>(hand: &PokerHand) -> u32 {
  let straight = is_straight(&hand.cards);
  let flush = is_flush(&hand.cards);
  let two_pair = has_two_pair(&hand.counts);
  let two = has_count(2, &hand.counts);
  let three = has_count(3, &hand.counts);
  let four = has_count(4, &hand.counts);

  if straight && flush {
    1
  } else if four {
    2
  } else if three && two {
    3
  } else if flush {
    4
  } else if straight {
    5
  } else if three {
    6
  } else if two_pair {
    7
  } else if two {
    8
  } else {
    9
  }
}

fn compare_hands(a: &PokerHand, b: &PokerHand) -> Ordering {
  let rank_a = get_hand_rank(a);
  let rank_b = get_hand_rank(b);

  let rough = rank_a.cmp(&rank_b);

  if rough != Ordering::Equal {
    return rough;
  }

  // This is very awkward, there must be a
  // better way to compare hands generically
  if rank_a == 7 {
    compare_equal_two_pair(&a.counts, &b.counts)
  } else if rank_a == 5 {
    compare_straights(&a.cards, &b.cards)
  } else {
    compare_equal_hands(&a.cards, &b.cards)
  }
}

fn compare_equal_two_pair(counts_a: &Counts, counts_b: &Counts) -> Ordering {
  let mut counts_a_vec: Vec<_> = counts_a.iter().collect();
  let mut counts_b_vec: Vec<_> = counts_b.iter().collect();

  // So much sorting/reversing! Fix this.
  counts_a_vec.sort_by(|a, b| a.0.cmp(b.0).reverse());
  counts_a_vec.sort_by(|a, b| a.1.cmp(b.1).reverse());
  counts_b_vec.sort_by(|a, b| a.0.cmp(b.0).reverse());
  counts_b_vec.sort_by(|a, b| a.1.cmp(b.1).reverse());

  return counts_a_vec.cmp(&counts_b_vec).reverse();
}

fn make_ace_low_straight(hand: &Hand) -> Hand {
  let mut new_hand = hand.clone();
  if hand.first().unwrap().0 == 2 && hand.last().unwrap().0 == 14 {
    new_hand.rotate_right(1);
    new_hand[0].0 = 1;
  }
  return new_hand;
}

fn compare_straights(a: &Hand, b: &Hand) -> Ordering {
  let sanitized_a = make_ace_low_straight(a);
  let sanitized_b = make_ace_low_straight(b);
  compare_equal_hands(&sanitized_a, &sanitized_b)
}

fn compare_equal_hands(a: &Hand, b: &Hand) -> Ordering {
  for (i, card) in a.iter().enumerate() {
    if card.0 > b[i].0 {
      return Ordering::Less;
    } else if card.0 < b[i].0 {
      return Ordering::Greater;
    }
  }
  Ordering::Equal
}

// ======= TESTS =======

fn hs_from<'a>(input: &[&'a str]) -> HashSet<&'a str> {
  let mut hs = HashSet::new();
  for item in input.iter() {
    hs.insert(*item);
  }
  hs
}
/// Test that the expected output is produced from the given input
/// using the `winning_hands` function.
///
/// Note that the output can be in any order. Here, we use a HashSet to
/// abstract away the order of outputs.
fn test<'a, 'b>(input: &[&'a str], expected: &[&'b str]) {
  assert_eq!(hs_from(&winning_hands(input)), hs_from(expected))
}
#[test]
fn test_single_hand_always_wins() {
  test(&["4S 5S 7H 8D JC"], &["4S 5S 7H 8D JC"])
}

#[test]
fn test_get_hand_as_vec() {
  assert_eq!(
    hand_as_vec("4S 5S 7H 10D JC"),
    vec![(4, 'S'), (5, 'S'), (7, 'H'), (10, 'D'), (11, 'C')]
  );
}

#[test]
fn test_is_straight() {
  assert_eq!(
    is_straight(&vec![(4, 'S'), (5, 'S'), (7, 'H'), (10, 'D'), (11, 'C')]),
    false
  );
  assert_eq!(
    is_straight(&vec![(4, 'S'), (5, 'S'), (6, 'H'), (7, 'D'), (8, 'C')]),
    true
  );
}

#[test]
fn test_is_straight_aces() {
  assert_eq!(
    is_straight(&vec![(2, 'S'), (3, 'H'), (4, 'D'), (5, 'C'), (14, 'S')]),
    true
  );
  assert_eq!(
    is_straight(&vec![(10, 'S'), (11, 'S'), (12, 'H'), (13, 'D'), (14, 'C')]),
    true
  );
}

#[test]
fn test_is_flush() {
  assert_eq!(
    is_flush(&vec![(4, 'S'), (5, 'S'), (7, 'H'), (10, 'D'), (11, 'C')]),
    false
  );
  assert_eq!(
    is_flush(&vec![(4, 'S'), (5, 'S'), (7, 'S'), (10, 'S'), (11, 'S')]),
    true
  );
}

#[test]
fn test_pair() {
  assert_eq!(
    has_count(
      2,
      &get_counts(&vec![(4, 'S'), (5, 'S'), (7, 'H'), (10, 'D'), (11, 'C')])
    ),
    false
  );
  assert_eq!(
    has_count(
      2,
      &get_counts(&vec![(4, 'S'), (4, 'C'), (7, 'H'), (10, 'D'), (11, 'C')])
    ),
    true
  );
}

#[test]
fn test_duplicate_hands_always_tie() {
  let input = &["3S 4S 5D 6H JH", "3S 4S 5D 6H JH", "3S 4S 5D 6H JH"];
  assert_eq!(&winning_hands(input), input)
}
#[test]
fn test_highest_card_of_all_hands_wins() {
  test(
    &["4D 5S 6S 8D 3C", "2S 4C 7S 9H 10H", "3S 4S 5D 6H JH"],
    &["3S 4S 5D 6H JH"],
  )
}
#[test]
fn test_a_tie_has_multiple_winners() {
  test(
    &[
      "4D 5S 6S 8D 3C",
      "2S 4C 7S 9H 10H",
      "3S 4S 5D 6H JH",
      "3H 4H 5C 6C JD",
    ],
    &["3S 4S 5D 6H JH", "3H 4H 5C 6C JD"],
  )
}
#[test]
fn test_high_card_can_be_low_card_in_an_otherwise_tie() {
  // multiple hands with the same high cards, tie compares next highest ranked,
  // down to last card
  test(&["3S 5H 6S 8D 7H", "2S 5D 6D 8C 7S"], &["3S 5H 6S 8D 7H"])
}
#[test]
fn test_one_pair_beats_high_card() {
  test(&["4S 5H 6C 8D KH", "2S 4H 6S 4D JH"], &["2S 4H 6S 4D JH"])
}
#[test]
fn test_highest_pair_wins() {
  test(&["4S 2H 6S 2D JH", "2S 4H 6C 4D JD"], &["2S 4H 6C 4D JD"])
}
#[test]
fn test_two_pairs_beats_one_pair() {
  test(&["2S 8H 6S 8D JH", "4S 5H 4C 8C 5C"], &["4S 5H 4C 8C 5C"])
}
#[test]
fn test_two_pair_ranks() {
  // both hands have two pairs, highest ranked pair wins
  test(&["2S 8H 2D 8D 3H", "4S 5H 4C 8S 5D"], &["2S 8H 2D 8D 3H"])
}
#[test]
fn test_two_pairs_second_pair_cascade() {
  // both hands have two pairs, with the same highest ranked pair,
  // tie goes to low pair
  test(&["2S QS 2C QD JH", "JD QH JS 8D QC"], &["JD QH JS 8D QC"])
}
#[test]
fn test_two_pairs_last_card_cascade() {
  // both hands have two identically ranked pairs,
  // tie goes to remaining card (kicker)
  test(&["JD QH JS 8D QC", "JS QS JC 2D QD"], &["JD QH JS 8D QC"])
}
#[test]
fn test_three_of_a_kind_beats_two_pair() {
  test(&["2S 8H 2H 8D JH", "4S 5H 4C 8S 4H"], &["4S 5H 4C 8S 4H"])
}
#[test]
fn test_three_of_a_kind_ranks() {
  //both hands have three of a kind, tie goes to highest ranked triplet
  test(&["2S 2H 2C 8D JH", "4S AH AS 8C AD"], &["4S AH AS 8C AD"])
}
#[test]
fn test_three_of_a_kind_cascade_ranks() {
  // with multiple decks, two players can have same three of a kind,
  // ties go to highest remaining cards
  test(&["4S AH AS 7C AD", "4S AH AS 8C AD"], &["4S AH AS 8C AD"])
}
#[test]
fn test_straight_beats_three_of_a_kind() {
  test(&["4S 5H 4C 8D 4H", "3S 4D 2S 6D 5C"], &["3S 4D 2S 6D 5C"])
}
#[test]
fn test_aces_can_end_a_straight_high() {
  // aces can end a straight (10 J Q K A)
  test(&["4S 5H 4C 8D 4H", "10D JH QS KD AC"], &["10D JH QS KD AC"])
}
#[test]
fn test_aces_can_end_a_straight_low() {
  // aces can start a straight (A 2 3 4 5)
  test(&["4S 5H 4C 8D 4H", "4D AH 3S 2D 5C"], &["4D AH 3S 2D 5C"])
}
#[test]
fn test_straight_cascade() {
  // both hands with a straight, tie goes to highest ranked card
  test(&["4S 6C 7S 8D 5H", "5S 7H 8S 9D 6H"], &["5S 7H 8S 9D 6H"])
}
#[test]
fn test_straight_scoring() {
  // even though an ace is usually high, a 5-high straight is the lowest-scoring straight
  test(&["2H 3C 4D 5D 6H", "4S AH 3S 2D 5H"], &["2H 3C 4D 5D 6H"])
}
#[test]
fn test_flush_beats_a_straight() {
  test(&["4C 6H 7D 8D 5H", "2S 4S 5S 6S 7S"], &["2S 4S 5S 6S 7S"])
}
#[test]
fn test_flush_cascade() {
  // both hands have a flush, tie goes to high card, down to the last one if necessary
  test(&["4H 7H 8H 9H 6H", "2S 4S 5S 6S 7S"], &["4H 7H 8H 9H 6H"])
}
#[test]
fn test_full_house_beats_a_flush() {
  test(&["3H 6H 7H 8H 5H", "4S 5C 4C 5D 4H"], &["4S 5C 4C 5D 4H"])
}
#[test]
fn test_full_house_ranks() {
  // both hands have a full house, tie goes to highest-ranked triplet
  test(&["4H 4S 4D 9S 9D", "5H 5S 5D 8S 8D"], &["5H 5S 5D 8S 8D"])
}
#[test]
fn test_full_house_cascade() {
  // with multiple decks, both hands have a full house with the same triplet, tie goes to the pair
  test(&["5H 5S 5D 9S 9D", "5H 5S 5D 8S 8D"], &["5H 5S 5D 9S 9D"])
}
#[test]
fn test_four_of_a_kind_beats_full_house() {
  test(&["4S 5H 4D 5D 4H", "3S 3H 2S 3D 3C"], &["3S 3H 2S 3D 3C"])
}
#[test]
fn test_four_of_a_kind_ranks() {
  // both hands have four of a kind, tie goes to high quad
  test(&["2S 2H 2C 8D 2D", "4S 5H 5S 5D 5C"], &["4S 5H 5S 5D 5C"])
}
#[test]
fn test_four_of_a_kind_cascade() {
  // with multiple decks, both hands with identical four of a kind, tie determined by kicker
  test(&["3S 3H 2S 3D 3C", "3S 3H 4S 3D 3C"], &["3S 3H 4S 3D 3C"])
}
#[test]
fn test_straight_flush_beats_four_of_a_kind() {
  test(&["4S 5H 5S 5D 5C", "7S 8S 9S 6S 10S"], &["7S 8S 9S 6S 10S"])
}
#[test]
fn test_straight_flush_ranks() {
  // both hands have straight flush, tie goes to highest-ranked card
  test(&["4H 6H 7H 8H 5H", "5S 7S 8S 9S 6S"], &["5S 7S 8S 9S 6S"])
}
