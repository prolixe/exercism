module Base
  ( Error(..)
  , rebase
  ) where

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

unDigits :: Integral a => a -> [a] -> a
unDigits base = foldl (\a b -> a * base + b) 0

digits :: Integral a => a -> a -> [a]
digits base = reverse . reversedDigits base

reversedDigits :: Integral a => a -> a -> [a]
reversedDigits _ 0 = []
reversedDigits base number =
  let (rest, lastDigit) = quotRem number base
  in lastDigit : reversedDigits base rest
