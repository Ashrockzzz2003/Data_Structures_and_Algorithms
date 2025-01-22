import Data.List (sortOn)
import Data.Maybe (isJust, fromMaybe)
import Control.Applicative ((<|>))

-- Type synonyms for clarity
type Position = (Int, Int)
type BoardSize = (Int, Int)

-- All possible knight moves
knightMoves :: [Position]
knightMoves = [(-2, -1), (-1, -2), (1, -2), (2, -1), (2, 1), (1, 2), (-1, 2), (-2, 1)]

-- Function to check if a position is within board boundaries
isValid :: Position -> BoardSize -> Bool
isValid (x, y) (rows, cols) = x >= 0 && y >= 0 && x < rows && y < cols

-- Function to calculate valid next moves
nextMoves :: Position -> [Position] -> BoardSize -> [Position]
nextMoves (x, y) visited boardSize =
  filter (`notElem` visited) . filter (`isValid` boardSize) $
  map (\(dx, dy) -> (x + dx, y + dy)) knightMoves

-- Warnsdorff's heuristic: sort moves by the number of onward moves
warnsdorffSort :: Position -> [Position] -> BoardSize -> [Position]
warnsdorffSort pos visited boardSize =
  sortOn (\move -> length (nextMoves move visited boardSize)) (nextMoves pos visited boardSize)

-- Recursive function to solve the Knight's Tour
knightTour :: [Position] -> Int -> BoardSize -> Maybe [Position]
knightTour visited moveCount boardSize
  | moveCount == totalCells = Just visited -- Tour complete
  | otherwise = case sortedMoves of
      [] -> Nothing -- Dead end
      _  -> foldl (\acc move -> acc <|> knightTour (move : visited) (moveCount + 1) boardSize) Nothing sortedMoves
  where
    sortedMoves = warnsdorffSort (head visited) visited boardSize
    totalCells = uncurry (*) boardSize

-- Function to solve the Knight's Tour with user-defined settings
solveKnightTour :: Position -> BoardSize -> Maybe [Position]
solveKnightTour start boardSize = knightTour [start] 1 boardSize

-- Function to display the board with the knight's path
printBoard :: Maybe [Position] -> BoardSize -> IO ()
printBoard Nothing _ = putStrLn "No solution exists."
printBoard (Just path) (rows, cols) = do
  let board = [[findMove (x, y) | y <- [0 .. cols - 1]] | x <- [0 .. rows - 1]]
  putStrLn "\nKnight's Tour Solution:\n"
  mapM_ printRow board
  where
    findMove pos = case lookup pos (zip (reverse path) [1..]) of
      Just step -> show step
      Nothing -> "."
    printRow row = putStrLn $ unwords (map (\cell -> pad cell) row)
    pad str = replicate (3 - length str) ' ' ++ str -- Pad cells for alignment

-- Interactive user interface
main :: IO ()
main = do
  putStrLn "Welcome to the Knight's Tour Solver!"
  putStrLn "Enter the number of rows (M):"
  rows <- readLn
  putStrLn "Enter the number of columns (N):"
  cols <- readLn
  putStrLn "Enter the starting position as two integers (row col):"
  [startRow, startCol] <- fmap (map read . words) getLine
  let start = (startRow, startCol)
      boardSize = (rows, cols)
  if isValid start boardSize
    then do
      let solution = solveKnightTour start boardSize
      printBoard solution boardSize
    else putStrLn "Invalid starting position!"