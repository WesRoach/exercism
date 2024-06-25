use std::collections::HashSet;

pub fn anagrams_for<'a>(word: &str, possible_anagrams: &[&'a str]) -> HashSet<&'a str> {
    // todo!("For the '{word}' word find anagrams among the following words: {possible_anagrams:?}");
    // For the `word` word find anagrams among the following word candidates `possible_anagrams`
    // Return a HashSet with the anagrams
    // The anagrams must use all letters exactly once
    // Words are not anagrams of themselves
    // Words are not anagrams of themselves even if letter case is partially different
    // Words are not anagrams of themselves even if letter case is completely different
    // Words other than themselves can be anagrams
    // Handles case of Greek letters
    let word_lower = word.to_lowercase();
    let mut word_sorted: Vec<char> = word_lower.chars().collect();
    word_sorted.sort();

    let mut anagrams: HashSet<&'a str> = HashSet::new();
    for candidate in possible_anagrams {
        let candidate_lower = candidate.to_lowercase();
        if candidate_lower == word_lower {
            continue;
        }
        let mut candidate_sorted: Vec<char> = candidate_lower.chars().collect();
        candidate_sorted.sort();

        if candidate_sorted.len() != word_sorted.len() {
            continue;
        } else if candidate_sorted == word_sorted {
            anagrams.insert(candidate);
        }
    }
    return anagrams;
}
