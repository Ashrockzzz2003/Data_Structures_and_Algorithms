#[derive(Debug)]
struct MinDistance {
    string1: String, 
    string2: String,
}

impl MinDistance {
    /*
        the following code tries to solve the problem using Dynamic Programming approach of memoization which is quite optimised to solve the problem at hand
    */
    pub fn minimum_distance(&self) -> i32 {
        let mut dp = vec![vec![-1; self.string2.len() + 1]; self.string1.len() + 1];
        return self.edit_distance(self.string1.len() as i32 - 1, self.string2.len() as i32 - 1, &mut dp);
    }

    /*
        a helper function that perform the memoization
    */

    pub fn edit_distance(&self, idx1: i32, idx2: i32, dp: &mut Vec<Vec<i32>>) -> i32 {
        if idx1 < 0 {
            return idx2 + 1;
        }
        if idx2 < 0 {
            return idx1 + 1;
        }

        if dp[idx1 as usize][idx2 as usize] != -1 {
            return dp[idx1 as usize][idx2 as usize];
        }

        if self.string1.as_bytes()[idx1 as usize] == self.string2.as_bytes()[idx2 as usize] {
            dp[idx1 as usize][idx2 as usize] = self.edit_distance(idx1 - 1, idx2 - 1, dp);
            return dp[idx1 as usize][idx2 as usize];
        } else {
            let insert = 1 + self.edit_distance(idx1, idx2 - 1, dp);
            let delete = 1 + self.edit_distance(idx1 - 1, idx2, dp);
            let replace = 1 + self.edit_distance(idx1 - 1, idx2 - 1,dp);
            dp[idx1 as usize][idx2 as usize] = insert.min(delete.min(replace));
        }
        return dp[idx1 as usize][idx2 as usize];
    }
}

fn main() {
    let min_dist = MinDistance {
        string1: "intention".to_string(),
        string2: "execution".to_string(),
    };

    println!("{}", min_dist.minimum_distance());

    let min_dist2 = MinDistance {
        string1: "horse".to_string(),
        string2: "ros".to_string(),
    };

    println!("{}", min_dist2.minimum_distance());
}

