pub struct Allergies {
    score: u32,
}

#[derive(Debug, PartialEq, Eq)]
pub enum Allergen {
    Eggs,
    Peanuts,
    Shellfish,
    Strawberries,
    Tomatoes,
    Chocolate,
    Pollen,
    Cats,
}

fn get_value(allergen: &Allergen) -> u32 {
    match allergen {
        Allergen::Eggs => 1,
        Allergen::Peanuts => 2,
        Allergen::Shellfish => 4,
        Allergen::Strawberries => 8,
        Allergen::Tomatoes => 16,
        Allergen::Chocolate => 32,
        Allergen::Pollen => 64,
        Allergen::Cats => 128,
    }
}

fn get_allergen(factor: u32) -> Allergen {
    match factor {
        1 => Allergen::Eggs,
        2 => Allergen::Peanuts,
        4 => Allergen::Shellfish,
        8 => Allergen::Strawberries,
        16 => Allergen::Tomatoes,
        32 => Allergen::Chocolate,
        64 => Allergen::Pollen,
        128 => Allergen::Cats,
        _ => panic!("oh no!"),
    }
}

impl Allergies {
    pub fn new(score: u32) -> Self {
        Allergies { score }
    }

    pub fn is_allergic_to(&self, allergen: &Allergen) -> bool {
        // 00001001 & 11111111
        // 00001000 & 11111111
        // println!("{:#010b}", self.score & 255);
        // println!("{:#010b}", get_value(allergen) & 255);
        if self.score & get_value(allergen) == get_value(allergen) & 255 {
            return true;
        }

        false
    }

    pub fn allergies(&self) -> Vec<Allergen> {
        let mut factor = 1;
        let mut allergens = vec![];

        for i in 0..8 {
            if self.score >> 1 == 1 {
                allergens.push(get_allergen(factor));
            }
            factor *= 2;
        }

        allergens
    }
}

// use allergies::*;
fn compare_allergy_vectors(expected: &[Allergen], actual: &[Allergen]) {
    for element in expected {
        if !actual.contains(element) {
            panic!("Allergen missing\n  {element:?} should be in {actual:?}");
        }
    }
    if actual.len() != expected.len() {
        panic!(
            "Allergy vectors are of different lengths\n  expected {expected:?}\n  got {actual:?}"
        );
    }
}
#[test]
fn is_not_allergic_to_anything() {
    let allergies = Allergies::new(0);
    assert!(!allergies.is_allergic_to(&Allergen::Peanuts));
    assert!(!allergies.is_allergic_to(&Allergen::Cats));
    assert!(!allergies.is_allergic_to(&Allergen::Strawberries));
}

#[test]
fn is_allergic_to_eggs() {
    assert!(Allergies::new(1).is_allergic_to(&Allergen::Eggs));
}
#[test]
fn is_allergic_to_eggs_and_shellfish_but_not_strawberries() {
    let allergies = Allergies::new(5);
    assert!(allergies.is_allergic_to(&Allergen::Eggs));
    assert!(allergies.is_allergic_to(&Allergen::Shellfish));
    assert!(!allergies.is_allergic_to(&Allergen::Strawberries));
}
#[test]
fn no_allergies_at_all() {
    let expected = &[];
    let allergies = Allergies::new(0).allergies();
    compare_allergy_vectors(expected, &allergies);
}
#[test]
fn allergic_to_just_eggs() {
    let expected = &[Allergen::Eggs];
    let allergies = Allergies::new(1).allergies();
    compare_allergy_vectors(expected, &allergies);
}
#[test]
#[ignore]
fn allergic_to_just_peanuts() {
    let expected = &[Allergen::Peanuts];
    let allergies = Allergies::new(2).allergies();
    compare_allergy_vectors(expected, &allergies);
}
#[test]
#[ignore]
fn allergic_to_just_strawberries() {
    let expected = &[Allergen::Strawberries];
    let allergies = Allergies::new(8).allergies();
    compare_allergy_vectors(expected, &allergies);
}
#[test]
#[ignore]
fn allergic_to_eggs_and_peanuts() {
    let expected = &[Allergen::Eggs, Allergen::Peanuts];
    let allergies = Allergies::new(3).allergies();
    compare_allergy_vectors(expected, &allergies);
}
#[test]
#[ignore]
fn allergic_to_eggs_and_shellfish() {
    let expected = &[Allergen::Eggs, Allergen::Shellfish];
    let allergies = Allergies::new(5).allergies();
    compare_allergy_vectors(expected, &allergies);
}
#[test]
#[ignore]
fn allergic_to_many_things() {
    let expected = &[
        Allergen::Strawberries,
        Allergen::Tomatoes,
        Allergen::Chocolate,
        Allergen::Pollen,
        Allergen::Cats,
    ];
    let allergies = Allergies::new(248).allergies();
    compare_allergy_vectors(expected, &allergies);
}
#[test]
#[ignore]
fn allergic_to_everything() {
    let expected = &[
        Allergen::Eggs,
        Allergen::Peanuts,
        Allergen::Shellfish,
        Allergen::Strawberries,
        Allergen::Tomatoes,
        Allergen::Chocolate,
        Allergen::Pollen,
        Allergen::Cats,
    ];
    let allergies = Allergies::new(255).allergies();
    compare_allergy_vectors(expected, &allergies);
}
#[test]
#[ignore]
fn scores_over_255_do_not_trigger_false_positives() {
    let expected = &[
        Allergen::Eggs,
        Allergen::Shellfish,
        Allergen::Strawberries,
        Allergen::Tomatoes,
        Allergen::Chocolate,
        Allergen::Pollen,
        Allergen::Cats,
    ];
    let allergies = Allergies::new(509).allergies();
    compare_allergy_vectors(expected, &allergies);
}
