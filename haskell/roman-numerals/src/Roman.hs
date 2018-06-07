module Roman
  ( numerals
  ) where

import qualified Data.Map as M

numerals :: Integer -> Maybe String
numerals n = Just $ singleNumeral n "" mapArabicToRoman

mapArabicToRoman :: [(String, Integer)]
mapArabicToRoman =
  [ ("M", 1000)
  , ("CM", 900)
  , ("D", 500)
  , ("CD", 400)
  , ("C", 100)
  , ("XC", 90)
  , ("L", 50)
  , ("XL", 40)
  , ("X", 10)
  , ("IX", 9)
  , ("V", 5)
  , ("IV", 4)
  , ("I", 1)
  ]

singleNumeral :: Integer -> String -> [(String, Integer)] -> String
singleNumeral numeral roman list
  | null list = roman
  | numeral > currentArabic =
    singleNumeral (numeral - currentArabic) (roman ++ currentRoman) list
  | numeral == currentArabic =
    singleNumeral (numeral - currentArabic) (roman ++ currentRoman) (tail list)
  | otherwise = singleNumeral numeral roman (tail list)
  where
    currentArabic = snd $ head list
    currentRoman = fst $ head list
