import System.IO
import Control.Monad

score :: String -> Int
score l
    | l == "A X" = 4
    | l == "A Y" = 8
    | l == "A Z" = 3
    | l == "B X" = 1
    | l == "B Y" = 5
    | l == "B Z" = 9
    | l == "C X" = 7
    | l == "C Y" = 2
    | l == "C Z" = 6

main = do
    handle <- openFile "test.txt" ReadMode
    contents <- hGetContents handle
    let singleLines = lines contents
    let scores = map score singleLines
    let total = foldl (+) 0 scores
    print total
    hClose handle