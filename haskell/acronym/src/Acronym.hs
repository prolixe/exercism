module Acronym (abbreviate) where

import Data.Char

abbreviate :: String -> String
abbreviate = filter isAlpha . map initial . concatMap (words . insertSpaceAtCamelCase) . words . map replaceChar 
    where initial = toUpper . head

replaceChar :: Char -> Char
replaceChar '-' = ' '
replaceChar c = c


insertSpaceAtCamelCase :: String -> String 
insertSpaceAtCamelCase xs  
    | length xs < 2 = xs
    | all isUpper xs = xs
    | otherwise = [head xs] ++ takeWhile isLower (tail xs) ++ " " ++ dropWhile isLower (tail xs)



