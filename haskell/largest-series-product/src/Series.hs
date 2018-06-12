module Series
  ( Error(..)
  , largestProduct
  ) where

import           Data.Char

data Error
  = InvalidSpan
  | InvalidDigit Char
  deriving (Show, Eq)

largestProduct :: Int -> String -> Either Error Integer
largestProduct size xs
  | size > length xs = Left InvalidSpan
  | size < 0 = Left InvalidSpan
  | any (not . isDigit) xs =
    Left $ InvalidDigit ((head . filter (not . isDigit)) xs)
  | otherwise =
    Right $ fromIntegral $ maximum (productList size (map digitToInt xs))

productList :: Int -> [Int] -> [Int]
productList size xs
  | size > length xs = [0]
  | size == 0 && length xs == size = [1]
  | otherwise = product (take size xs) : productList size (drop 1 xs)
