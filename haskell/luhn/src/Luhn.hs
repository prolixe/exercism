module Luhn
  ( isValid
  ) where

import           Data.Char (digitToInt, isDigit)

isValid :: String -> Bool
isValid n
  | length filteredInput < 2 = False
  | otherwise =
    ((== 0) . (`mod` 10) . sum . doubleDigit . toDigitList) filteredInput
  where
    filteredInput = filter isDigit n

toDigitList :: String -> [Int]
toDigitList = map digitToInt

doubleDigit :: [Int] -> [Int]
doubleDigit xs =
  reverse $
  map
    (\(d, i) ->
       if even i
         then double d
         else d) $
  zip (reverse xs) [1 ..]
  where
    double d =
      if d * 2 > 9
        then d * 2 - 9
        else d * 2
