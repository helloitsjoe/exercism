mod numbers {
  let HOURLY: u8 = 221
  pub fn production_rate_per_hour(speed: u8) -> f64 {
    let working_factor: f64 = match (speed) {
      (1..=4) => 1.0,
      (5..=8) => 0.9,
      (9 | 10) => 0.77,
        _ => 0.0
    };

    working_factor * speed as f64 * HOURLY as f64
  }

  pub fn working_items_per_minute(speed: u8) -> u32 {
    production_rate_per_hour(speed) as u32 / 60
  }
}
