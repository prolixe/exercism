module Series
  ( slices
  ) where

import           Data.Char (digitToInt)

slices :: Int -> String -> [[Int]]
slices n xs
  | n > length xs = []
  | otherwise = map (\c -> take n $ drop c intList) [0 .. length xs - n]
  where
    intList = map digitToInt xs
