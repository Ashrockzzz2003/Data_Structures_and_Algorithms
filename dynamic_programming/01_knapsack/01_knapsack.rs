#[derive(Debug)]
struct KnapSack01 {
    weights: Vec<i32>,
    value: Vec<i32>,
    capacity: i32
}

impl KnapSack01 {
    /*
        the following code tries to solve the problem at hand using the traditional recursion approach
    */
    pub fn knapsack_recursive(&self, idx: usize, current_capacity: i32) -> i32 {
        if idx == 0 || current_capacity == 0 {
            return 0;
        }
        if self.weights[idx - 1] > current_capacity {
            return self.knapsack_recursive(idx - 1, current_capacity);
        } else {
            let not_pick = self.knapsack_recursive(idx - 1 , current_capacity);
            let pick = self.knapsack_recursive(idx - 1 , current_capacity - self.weights[idx - 1]) + self.value[idx - 1];
            return not_pick.max(pick);
        }
    }

    /*
        the following code tries to solve the problem using the Dynamic Programming approach of Memoization which is quite optimised 
    */
    pub fn knapsack_memoization(&self, idx: usize, current_capacity: i32, memo: &mut Vec<Vec<i32>>) -> i32 {
        if idx == 0 || current_capacity == 0 {
            return 0; 
        }
        if memo[idx][current_capacity as usize] != -1 {
            return memo[idx][current_capacity as usize];
        }
        if self.weights[idx - 1] > current_capacity {
            memo[idx][current_capacity as usize] =  self.knapsack_memoization(idx-1, current_capacity, memo);
        } else {
            let not_pick = self.knapsack_memoization(idx - 1 , current_capacity, memo);
            let pick = self.knapsack_memoization(idx - 1 , current_capacity - self.weights[idx - 1] , memo) + self.value[idx - 1];
            memo[idx][current_capacity as usize] = not_pick.max(pick);
        }
        return memo[idx][current_capacity as usize];
    }

    /*
        the following code tries to solve the problem using the Dynamic Programming approach of Tabulation which is quite optimised
    */
    pub fn knapsack_tabulation(&self) -> i32 {
        let weight_sz = self.weights.len();
        let mut dp = vec![vec![0; (self.capacity + 1) as usize]; weight_sz + 1];

        for i in 1..=weight_sz {
            for j in 1..=(self.capacity as usize) {
                if self.weights[i - 1] > j as i32 {
                    dp[i][j] = dp[i - 1][j];
                } else {
                    let not_pick = dp[i - 1][j];
                    let pick = dp[i - 1][j - self.weights[i - 1] as usize] + self.value[i - 1];
                    dp[i][j] = not_pick.max(pick);
                }
            }
        }
        return dp[weight_sz][self.capacity as usize];
    }
}

fn main() {
    let knapsack = KnapSack01 {
        weights: vec![10,5,5],
        value: vec![60,50,50],
        capacity: 10
    };

    let result_recursive = knapsack.knapsack_recursive(knapsack.weights.len(), knapsack.capacity);
    println!("Recusive: {}", result_recursive);

    let mut memo = vec![vec![-1; (knapsack.capacity + 1) as usize]; knapsack.weights.len() + 1];
    let result_memoization = knapsack.knapsack_memoization(knapsack.weights.len(), knapsack.capacity, &mut memo);
    println!("Memoized: {}", result_memoization);

    let result_dp = knapsack.knapsack_tabulation();
    println!("tabulation: {}", result_dp);
}
