/*
Given a set of cards each with a value, two players take turns to pick a card from the set. The player with the highest value wins the game. Find the maximum possible score for the first player if both players play optimally and first player is the one to start the game.

The idea is this:
    1. Maximize the score of the first player
    2. Minimize the score of the second player

left -> i
right -> j

Cases:
    1. Player 1 picks left card, player 2 picks left card       (i, j) -> value[i] + (i + 2, j)
    2. Player 1 picks left card, player 2 picks right card      (i, j) -> value[i] + (i + 1, j - 1)

    3. Player 1 picks right card, player 2 picks left card      (i, j) -> value[j] + (i + 1, j - 1)
    4. Player 1 picks right card, player 2 picks right card     (i, j) -> value[j] + (i, j - 2)


Logic: max(left + min(1, 2), right + min(3, 4)) -> From the definition of the subproblem, we can see that we are trying to maximize the score of the first player and minimize the score of the second playerq
*/


#[derive(Debug)]
struct RemovalGame {
    cards: Vec<i32>,
}

impl RemovalGame {

    /*
        the following code tries to solve the problem in the traditional recurisve way which is not space nor time optimised
    */
    pub fn removal_game_recursion(&self, idx1: usize, idx2: usize) -> i32 {
        if idx1 == idx2 {
            return 0;
        }
        if idx1 == idx2 - 1 {
            return self.cards[idx1].max(self.cards[idx2]);  
        }
        
        return self.cards[idx1].max(
            self.removal_game_recursion(idx1 + 2, idx2).min(self.removal_game_recursion(idx1 + 1, idx2 - 1))
        )
        .max(
            self.cards[idx2].max(
                self.removal_game_recursion(idx1 + 1, idx2 - 1).min(self.removal_game_recursion(idx1, idx2 - 2))
            )
        );
    }

    /*
        the following code tries to solve the problem using the Dynamic Programming approach of memoization (top-down) where we space and time optimise the code 
    */
    pub fn removal_game_memoization(&self, idx1: usize, idx2: usize, dp: &mut Vec<Vec<i32>>) -> i32 {
        if idx1 == idx2 {
            return 0;
        }

        if idx1 == idx2 - 1 {
            return self.cards[idx1].max(self.cards[idx2]);
        }

        if dp[idx1][idx2] != -1 {
            return dp[idx1][idx2];
        }

        dp[idx1][idx2] = self.cards[idx1].max(
            self.removal_game_memoization(idx1 + 2, idx2, dp).min(self.removal_game_memoization(idx1 + 1, idx2 - 1, dp))
        )
        .max(
            self.cards[idx2].max(
                self.removal_game_memoization(idx1 + 1, idx2 - 1, dp).min(self.removal_game_memoization(idx1, idx2 - 2, dp))
            )
        );

        return dp[idx1][idx2];
    }

    /*
        the following code tries to the solve the problem using the Dynamic programming approahc of memoization (bottoms-up) where space and time optimize the code
    */
    pub fn removal_game_tabulation(&self) -> i32 {
        let n = self.cards.len();
        let mut dp = vec![vec![0; n]; n];
        
        for i in 0..n {
            dp[i][i] = self.cards[i];
        }
        
        for i in 0..n-1 {
            dp[i][i+1] = self.cards[i].max(self.cards[i+1]);
        }

        for len in 3..=n {
            for i in 0..=n-len {
                let j = i + len - 1;
                if j < n {  
                    let take_first = if i + 2 < n {
                        self.cards[i].max(
                            dp[i + 2][j].min(dp[i + 1][j - 1])
                        )
                    } else {
                        self.cards[i]
                    };
                    
                    let take_last = if j >= 2 {
                        self.cards[j].max(
                            dp[i + 1][j - 1].min(dp[i][j - 2])
                        )
                    } else {
                        self.cards[j]
                    };
                    
                    dp[i][j] = take_first.max(take_last);
                }
            }
        }
        
        return dp[0][n-1];
    }
}

fn main() {
    let rmg = RemovalGame {
        cards: vec![4, 5, 1, 3],
    };
    println!("Recursion: {}", rmg.removal_game_recursion(0, rmg.cards.len() - 1));
    let mut dp = vec![vec![-1; rmg.cards.len()]; rmg.cards.len()];
    println!("Memoization: {}", rmg.removal_game_memoization(0, rmg.cards.len() - 1, &mut dp));
    println!("Tabulation: {}", rmg.removal_game_tabulation());
}
