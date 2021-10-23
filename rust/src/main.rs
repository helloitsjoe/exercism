mod exercises;

use chrono::Utc;
use exercises::{gigasecond, reverse_string};

fn main() {
    reverse_string::reverse("foo");
    gigasecond::after(Utc::now());
}
