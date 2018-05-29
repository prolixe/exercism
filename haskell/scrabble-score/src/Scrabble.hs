module Scrabble
  ( scoreLetter
  , scoreWord
  ) where

import           Data.Char

scoreLetter :: Char -> Integer
scoreLetter letter
  | l `elem` aeioulnrst = 1
  | l `elem` dg = 2
  | l `elem` bcmp = 3
  | l `elem` fhvwy = 4
  | l `elem` k = 5
  | l `elem` jx = 8
  | l `elem` qz = 10
  | otherwise = 0
  where
    aeioulnrst = "AEIOULNRST"
    dg = "DG"
    bcmp = "BCMP"
    fhvwy = "FHVWY"
    k = "K"
    jx = "JX"
    qz = "QZ"
    l = toUpper letter

scoreWord :: String -> Integer
scoreWord = sum . map scoreLetter
