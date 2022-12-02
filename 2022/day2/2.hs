import System.IO
import Control.Monad

score :: String -> Int
score l
    | l == "A X" = 0 + 3
    | l == "A Y" = 3 + 1
    | l == "A Z" = 6 + 2
    | l == "B X" = 0 + 1
    | l == "B Y" = 3 + 2
    | l == "B Z" = 6 + 3
    | l == "C X" = 0 + 2
    | l == "C Y" = 3 + 3
    | l == "C Z" = 6 + 1

main = do
    handle <- openFile "test.txt" ReadMode
    contents <- hGetContents handle
    let singleLines = lines contents
    let scores = map score singleLines
    let total = foldl (+) 0 scores
    print total
    hClose handle