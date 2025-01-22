#[derive(Debug)]
struct MatrixChainMultiplication {
    num_list: Vec<f32>,
}

impl MatrixChainMultiplication {
    /*
        the following code tries to solve the problem using the traditional recurison which is not optimal nor is space efficient
    */
    pub fn mcm_recursion(&self, idx: usize, prev_idx: usize) -> f32 {
        if idx >= prev_idx {
            return 0.0;
        } 
        let mut val: f32 = f32::MAX;
        for j in idx..prev_idx {
            val = val.min(
                self.mcm_recursion(idx, j) +
                self.mcm_recursion(j+1 , prev_idx) + 
                self.num_list[idx] * 
                self.num_list[j + 1] * 
                self.num_list[prev_idx + 1]
            ); 
        }
        return val;
    }

    /*
        the following code tries to solve the porblem using the Dynamic Programming approach of memoization (top-down) approach which is one of the optimised ways to solve the problem
    */

    pub fn mcm_memoization(&self, idx: usize, prev_idx: usize, dp: &mut Vec<Vec<f32>>) -> f32 {
        if idx >= prev_idx {
            return 0.0;
        }
        if dp[idx][prev_idx] != -1.0 {
            return dp[idx][prev_idx];
        }
        let mut val: f32 = f32::MAX;
        for j in idx..prev_idx {
            val = val.min(
                self.mcm_memoization(idx, j, dp) + 
                self.mcm_memoization(j + 1, prev_idx , dp) + 
                self.num_list[idx] * 
                self.num_list[j + 1] *
                self.num_list[prev_idx + 1]
            );
        }
        dp[idx][prev_idx] = val;
        return val;
    }


    /*
        the following code tries to solve the problem using the Dynamic Programming Approach of Tabulation (bottoms-up) approach which is one of the ways to solve thr problem
    */

    pub fn mcm_tabulation(&self) -> f32 {
        let n = self.num_list.len() - 1;
        let mut dp = vec![vec![0.0; n] ; n];

        for i in 0..n {
            dp[i][i] = 0.0;
        }

        for len in 2..=n{
            for i in 0..=(n-len) {
                let j = i + len - 1;
                dp[i][j] = f32::MAX;
                for k in i..j {
                    let cost = dp[i][k] + dp[k + 1][j] + self.num_list[i] * self.num_list[k + 1] * self.num_list[j + 1];
                    dp[i][j] = dp[i][j].min(cost);
                }
            }
        }
        return dp[0][n-1];
    }
}

fn main() {
    let mcm = MatrixChainMultiplication {
        num_list: vec![30.0,35.0,15.0,5.0,10.0,20.0,25.0]
    };

    println!("recursion: {}", mcm.mcm_recursion(0, mcm.num_list.len() - 2));
    let mut dp = vec![vec![-1.0; mcm.num_list.len()] ; mcm.num_list.len()];
    println!("memoization: {}", mcm.mcm_memoization(0, mcm.num_list.len() - 2,&mut dp));
    println!("tabulation: {}", mcm.mcm_tabulation());
}
