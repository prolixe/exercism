module Triplet
  ( isPythagorean
  , mkTriplet
  , pythagoreanTriplets
  ) where

import           Data.List (nubBy, sort)

isPythagorean :: (Int, Int, Int) -> Bool
isPythagorean (a, b, c) =
  let [x, y, z] = sort [a, b, c]
  in x ^ 2 + y ^ 2 == z ^ 2

mkTriplet :: Int -> Int -> Int -> (Int, Int, Int)
mkTriplet a b c = (a, b, c)

-- Using List comprehension.
pythagoreanTriplets :: Int -> Int -> [(Int, Int, Int)]
pythagoreanTriplets minFactor maxFactor =
  [ (a, b, c)
  | a <- [minFactor .. maxFactor]
  , b <- [a .. maxFactor]
  , c <- [b .. maxFactor]
  , isPythagorean (a, b, c)
  ]
