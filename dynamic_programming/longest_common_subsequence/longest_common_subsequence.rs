#[derive(Debug)]
struct LongestCommonSubsequence {
    string1: String, 
    string2: String,
}

impl LongestCommonSubsequence {
    /*
        the following code tries to solve the problem by using the traditional recursion which is not the most optimised solution 
    */
    pub fn lcs_recursion(&self, idx1: i32, idx2: i32) -> i32 {
        if idx1 == 0 || idx2 == 0 {
            return 0;
        } 
        if self.string1.as_bytes()[(idx1 - 1) as usize] == self.string2.as_bytes()[(idx2 - 1) as usize] {
            return 1 + self.lcs_recursion(idx1 - 1, idx2 - 1);
        } else {
            let move_str1 = self.lcs_recursion(idx1 - 1, idx2);
            let move_str2 = self.lcs_recursion(idx1, idx2 - 1);
            return move_str1.max(move_str2);
        }
    }

    /*
        the following code tries to solve the problem using the Dynamic Programming approach of memoization(top-down) approach which is quite optimised
    */
    pub fn lcs_memoization(&self, idx1: i32, idx2: i32) -> i32 {
        let mut dp = vec![vec![-1; self.string2.len() + 1 ]; self.string1.len() + 1]; 
        return self.lcs_memoization_helper(idx1, idx2,&mut dp);
    }


    /*
        Simple helper function to perform the memoization
    */
    pub fn lcs_memoization_helper(&self, idx1: i32, idx2: i32, dp: &mut Vec<Vec<i32>>) -> i32 {
        if idx1 == 0 || idx2 == 0 {
            return 0;
        }
        if dp[idx1 as usize][idx2 as usize] != -1 {
            return dp[idx1 as usize][idx2 as usize];
        }
        if self.string1.as_bytes()[(idx1 - 1) as usize] == self.string2.as_bytes()[(idx2 - 1) as usize] {
            dp[idx1 as usize][idx2 as usize] = 1 + self.lcs_memoization_helper(idx1-1, idx2-1, dp);
        } else {
            let move_str1 = self.lcs_memoization_helper(idx1 - 1, idx2, dp);
            let move_str2 = self.lcs_memoization_helper(idx1, idx2 - 1, dp);
            dp[idx1 as usize][idx2 as usize] = move_str1.max(move_str2);
        }
        return dp[idx1 as usize][idx2 as usize];
    }

    /*
        the following code tries to solve the problem using the Dynamic Programming approach of tabulation (bottom-up) which is quite optimised
    */

    pub fn lcs_tabulation(&self) -> i32 {
        let mut dp = vec![vec![0; self.string2.len() + 1]; self.string1.len() + 1];

        for i in 1..=self.string1.len() {
            for j in 1..=self.string2.len() {
                if self.string1.as_bytes()[i - 1] == self.string2.as_bytes()[j - 1] {
                    dp[i][j] = 1 + dp[i-1][j-1];
                } else {
                    let move_str1 = dp[i-1][j];
                    let move_str2 = dp[i][j - 1];
                    dp[i][j] = move_str1.max(move_str2);
                }
            }
        }
        return dp[self.string1.len()][self.string2.len()];
    }
}

fn main() {
    let lcs1 = LongestCommonSubsequence{
        string1: "MANGROVEFOREST".to_string(),
        string2: "MULTIDIMENSIONAL".to_string(),
    };

    let lcs2 = LongestCommonSubsequence{
        string1: "ABET".to_string(),
        string2: "ACST".to_string(),
    };

    println!("Recursion: {}", lcs1.lcs_recursion(lcs1.string1.len() as i32, lcs1.string2.len() as i32));
    println!("Memoization: {}", lcs2.lcs_memoization(lcs2.string1.len() as i32, lcs2.string2.len() as i32));
    println!("Tabulation: {}", lcs2.lcs_tabulation());
}








