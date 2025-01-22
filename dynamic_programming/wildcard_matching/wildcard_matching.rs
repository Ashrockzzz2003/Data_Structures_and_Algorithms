#[derive(Debug)]
struct WildCardMatch {
    string1: String,
    string2: String
}

impl WildCardMatch {

    /*
        The function follows the dynamic programming priniciple: memoization
    */

    pub fn is_match(&self) -> bool {
        let mut dp = vec![vec![None; self.string2.len() + 1]; self.string1.len() + 1];
        return self.wild_card_match(self.string1.len() - 1, self.string2.len() - 1,&mut dp);
    }

    /*
        Simple helper fucntion to perform the memoization
    */
    pub fn wild_card_match(&self, idx1: usize, idx2: usize, dp: &mut Vec<Vec<Option<bool>>>) -> bool {
        if idx1 == 0 && idx2 == 0 {
            return true; 
        }

        if idx2 == 0 && idx1 > 0 {
            return false;
        }

        if idx1 == 0 {
            for k in 1..=idx2 {
                if self.string2.chars().nth(idx2 - k) != Some('*') {
                    return false;
                }
            }
            return true;
        }

        if let Some(result) = dp[idx1][idx2] {
            return result;
        }

        let mut result = false;
        
        if self.string1.chars().nth(idx1 -1) == self.string2.chars().nth(idx2 - 1) ||
        self.string2.chars().nth(idx2 - 1) == Some('?') {
            result = self.wild_card_match(idx1 - 1, idx2 - 1, dp);
        }

        if self.string2.chars().nth(idx2 - 1) == Some('*') {
            result = self.wild_card_match(idx1, idx2 - 1, dp) || self.wild_card_match(idx1 - 1, idx2, dp);
        }

        dp[idx1][idx2] = Some(result);
        return result;
    }
}

fn main() {
    let wcm = WildCardMatch {
        string1: "acdcb".to_string(),
        string2: "a*c?b".to_string(),
    };

    println!("{}", wcm.is_match());
}
