module Base
  ( Error(..)
  , rebase
  ) where

import           Data.Digits (digits, unDigits)

data Error a
  = InvalidInputBase
  | InvalidOutputBase
  | InvalidDigit a
  deriving (Show, Eq)

rebase :: Integral a => a -> a -> [a] -> Either (Error a) [a]
rebase from to xs
  | from < 2 = Left InvalidInputBase
  | to < 2 = Left InvalidOutputBase
  | any (< 0) xs = Left $ InvalidDigit $ head $ filter (< 0) xs
  | any (>= from) xs = Left $ InvalidDigit $ head $ filter (>= from) xs
  | otherwise = (Right . digits to . unDigits from) xs
