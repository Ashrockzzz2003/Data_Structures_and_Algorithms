#[derive(Debug)]
struct LongestIncreasingSubsequence {
    num_list: Vec<i32>,
}

impl LongestIncreasingSubsequence {
    /*
         the following method follows the Dynamic Programming approach of memoization to solve the problem
    */
    pub fn lis_memoization(&self, idx: usize, prev_idx: usize, dp: &mut Vec<Vec<i32>>) -> i32 {
        if idx == self.num_list.len() {
            return 0;
        }
        if dp[idx][prev_idx] != -1 {
            return dp[idx][prev_idx];
        }
        let mut total_length = self.lis_memoization(idx + 1, prev_idx, dp);
        if prev_idx == self.num_list.len() || self.num_list[idx] > self.num_list[prev_idx] {
            total_length = total_length.max(1 + self.lis_memoization(idx + 1, idx, dp));
        }
        dp[idx][prev_idx] = total_length;
        return total_length;

    }

    /*
        the following code tries to solve the problem using the Dynamic Programming approach of Tabulation
    */

    pub fn length_of_lis(&self) -> i32 {
        if self.num_list.is_empty() {
            return 0;
        }
        let mut dp = vec![1; self.num_list.len()];
        let mut max_len = 1;
        for i in 1..self.num_list.len() {
            for j in 0..i {
                if self.num_list[i] > self.num_list[j] {
                    dp[i] = dp[i].max(dp[j] + 1);
                }
            }
            max_len = max_len.max(dp[i]);
        }
        return max_len;
    }

    /*
         the following code tries to solve the problem with the strategy implemented to solve Longest Common Subsequence problem
    */

    pub fn length_of_lis_with_lcs(&self) -> i32 {
        let mut sorted_num_list = self.num_list.clone();
        sorted_num_list.sort();
        return self.lcs(&sorted_num_list);
    }

    pub fn lcs(&self, sorted_num_list: &Vec<i32>) -> i32 {
        let mut dp = vec![vec![0; sorted_num_list.len() + 1]; self.num_list.len() + 1];
        for i in 1..=self.num_list.len() {
            for j in 1..=sorted_num_list.len() {
                if self.num_list[i - 1] == sorted_num_list[j - 1] {
                    dp[i][j] = 1 + dp[i - 1][j - 1];
                } else {
                    dp[i][j] = dp[i - 1][j].max(dp[i][j - 1]);
                }
            }
        }
        return dp[self.num_list.len()][sorted_num_list.len()];
    }
}

fn main() {
    let lis = LongestIncreasingSubsequence {
        num_list: vec![10, 9, 2, 5, 3, 7, 101, 18],
    };

    let mut dp = vec![vec![-1; lis.num_list.len() + 1]; lis.num_list.len()];

    let length_with_memoization = lis.lis_memoization(0, lis.num_list.len(), &mut dp);
    println!("Length of LIS with Memoization: {}", length_with_memoization); 
    println!("Length of LIS: {}", lis.length_of_lis());
    println!("Length of LIS with LCS: {}", lis.length_of_lis_with_lcs());
}
