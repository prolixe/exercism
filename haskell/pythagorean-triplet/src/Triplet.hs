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

pythagoreanTriplets :: Int -> Int -> [(Int, Int, Int)]
pythagoreanTriplets minFactor maxFactor =
  nubBy isSameTriangle $
  filter isPythagorean $ addToTupple <$> range <*> ((,) <$> range <*> range)
  where
    addToTupple a (b, c) = (a, b, c)
    range = [minFactor .. maxFactor]
    isSameTriangle (a, b, c) (d, e, f) = sort [a, b, c] == sort [d, e, f]
