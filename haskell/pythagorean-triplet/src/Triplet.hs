module Triplet
  ( isPythagorean
  , mkTriplet
  , pythagoreanTriplets
  ) where

import           Control.Monad (guard)
import           Data.List     (sort)

isPythagorean :: (Int, Int, Int) -> Bool
isPythagorean (a, b, c) =
  let [x, y, z] = sort [a, b, c]
  in x ^ 2 + y ^ 2 == z ^ 2

mkTriplet :: Int -> Int -> Int -> (Int, Int, Int)
mkTriplet a b c = (a, b, c)

pythagoreanTriplets :: Int -> Int -> [(Int, Int, Int)]
pythagoreanTriplets min max = do
  a <- [min .. max]
  b <- [a .. max]
  c <- [b .. max]
  guard $ isPythagorean (a, b, c)
  return (a, b, c)
