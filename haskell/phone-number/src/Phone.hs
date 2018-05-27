module Phone
  ( number
  ) where

import           Data.Char

-- I tried to mess with the >>= operator (bind) a bit
number :: String -> Maybe String
number xs =
  pure (removeStartingDigit '1' $ filteredNumber xs) >>= checkCorrectDigitCount >>=
  checkValidFirstDigit >>=
  checkValidExchangeCode

filteredNumber :: String -> String
filteredNumber = filter isDigit

removeStartingDigit :: Char -> String -> String
removeStartingDigit c s
  | head s == c = tail s
  | otherwise = s

checkCorrectDigitCount :: String -> Maybe String
checkCorrectDigitCount s
  | 10 == length s = Just s
  | otherwise = Nothing

checkValidFirstDigit :: String -> Maybe String
checkValidFirstDigit s
  | (isNDigit . head) s = Just s
  | otherwise = Nothing

checkValidExchangeCode :: String -> Maybe String
checkValidExchangeCode s
  | isNDigit (s !! 3) = Just s
  | otherwise = Nothing

isNDigit :: Char -> Bool
isNDigit = (`elem` ['2' .. '9'])
