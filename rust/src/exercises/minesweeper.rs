fn create_vec(minefield: &[&str]) -> Vec<Vec<String>> {
  return (0..minefield.len())
    .map(|_r| {
      let len = minefield[0].chars().count();
      (0..len).map(|_c| String::from("")).collect()
    })
    .collect();
}

fn check_for_mines(r_idx: usize, c_idx: usize, minefield: &[&str]) -> String {
  let orig = minefield[r_idx].as_bytes()[c_idx] as char;
  if orig == '*' {
    return String::from(orig);
  }

  let mut mines = 0;

  for offset_r in 0..=2 {
    for offset_c in 0..=2 {
      // checked_sub is useful for checking potentially out-of-bounds indices
      let row_offset = (r_idx + offset_r).checked_sub(1).unwrap_or(1000);
      let col_offset = (c_idx + offset_c).checked_sub(1).unwrap_or(1000);

      if row_offset >= minefield.len() || col_offset >= minefield[0].len() {
        println!("out of bounds");
        continue;
      }

      if minefield[row_offset].as_bytes()[col_offset] as char == '*' {
        mines += 1;
      }
    }
  }
  if mines == 0 {
    return String::from(" ");
  }
  mines.to_string()
}

pub fn annotate(minefield: &[&str]) -> Vec<String> {
  let mut output = create_vec(minefield);

  for (r_idx, row) in minefield.iter().enumerate() {
    for (c_idx, _char) in row.chars().enumerate() {
      output[r_idx][c_idx] = check_for_mines(r_idx, c_idx, minefield);
    }
  }

  return output.iter().map(|row| row.join("")).collect();
}

fn remove_annotations(board: &[&str]) -> Vec<String> {
  board.iter().map(|r| remove_annotations_in_row(r)).collect()
}
fn remove_annotations_in_row(row: &str) -> String {
  row
    .chars()
    .map(|ch| match ch {
      '*' => '*',
      _ => ' ',
    })
    .collect()
}
fn run_test(test_case: &[&str]) {
  let cleaned = remove_annotations(test_case);
  let cleaned_strs = cleaned.iter().map(|r| &r[..]).collect::<Vec<_>>();
  let expected = test_case.iter().map(|&r| r.to_string()).collect::<Vec<_>>();
  assert_eq!(expected, annotate(&cleaned_strs));
}
#[test]
fn no_rows() {
  #[rustfmt::skip]
    run_test(&[
    ]);
}
#[test]
fn no_columns() {
  #[rustfmt::skip]
    run_test(&[
        "",
    ]);
}
#[test]
fn no_mines() {
  #[rustfmt::skip]
    run_test(&[
        "   ",
        "   ",
        "   ",
    ]);
}
#[test]
fn board_with_only_mines() {
  #[rustfmt::skip]
    run_test(&[
        "***",
        "***",
        "***",
    ]);
}
#[test]
fn mine_surrounded_by_spaces() {
  #[rustfmt::skip]
    run_test(&[
        "111",
        "1*1",
        "111",
    ]);
}
#[test]
fn space_surrounded_by_mines() {
  #[rustfmt::skip]
    run_test(&[
        "***",
        "*8*",
        "***",
    ]);
}
#[test]
fn horizontal_line() {
  #[rustfmt::skip]
    run_test(&[
        "1*2*1",
    ]);
}
#[test]
fn horizontal_line_mines_at_edges() {
  #[rustfmt::skip]
    run_test(&[
        "*1 1*",
    ]);
}
#[test]
fn vertical_line() {
  #[rustfmt::skip]
    run_test(&[
        "1",
        "*",
        "2",
        "*",
        "1",
    ]);
}
#[test]
fn vertical_line_mines_at_edges() {
  #[rustfmt::skip]
    run_test(&[
        "*",
        "1",
        " ",
        "1",
        "*",
    ]);
}
#[test]
fn cross() {
  #[rustfmt::skip]
    run_test(&[
        " 2*2 ",
        "25*52",
        "*****",
        "25*52",
        " 2*2 ",
    ]);
}
#[test]
fn large_board() {
  #[rustfmt::skip]
    run_test(&[
        "1*22*1",
        "12*322",
        " 123*2",
        "112*4*",
        "1*22*2",
        "111111",
    ]);
}
