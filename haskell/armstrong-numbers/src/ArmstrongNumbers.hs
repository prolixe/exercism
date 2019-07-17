module ArmstrongNumbers
  ( armstrong
  ) where

import Data.Char (digitToInt)

armstrong :: (Show a, Integral a) => a -> Bool
armstrong x = x == fromIntegral (sum (map (^ l) digits))
  where
    digits = map digitToInt sX
    l = length sX
    sX = show x
