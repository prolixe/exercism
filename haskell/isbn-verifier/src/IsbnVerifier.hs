module IsbnVerifier
  ( isbn
  ) where

import           Data.Char

isbn :: String -> Bool
isbn xs
  | null xs = False
  | all ($ filteredIsbn) checks = True
  | otherwise = False
  where
    filteredIsbn = filter isIsbnChar xs

checks :: [String -> Bool]
checks = [isValidFormat, isCorrectLength, isValidFormula . toIsbnInt]

isIsbnChar :: Char -> Bool
isIsbnChar x = isDigit x || 'X' == x

allValidIsbnChar :: String -> Bool
allValidIsbnChar = all isIsbnChar

isCorrectLength :: String -> Bool
isCorrectLength = (10 ==) . length

toIsbnInt :: String -> [Int]
toIsbnInt =
  map
    (\x ->
       if isDigit x
         then digitToInt x
         else 10)

isValidFormula :: [Int] -> Bool
isValidFormula [x1, x2, x3, x4, x5, x6, x7, x8, x9, x10] =
  (x1 * 10 + x2 * 9 + x3 * 8 + x4 * 7 + x5 * 6 + x6 * 5 + x7 * 4 + x8 * 3 +
   x9 * 2 +
   x10) `mod`
  11 ==
  0

isValidFormulaxs = False

isValidFormat :: String -> Bool
isValidFormat xs = all isDigit (init xs) && isIsbnChar (last xs)
