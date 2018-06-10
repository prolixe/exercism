module PrimeFactors
  ( primeFactors
  ) where

primeFactors :: Integer -> [Integer]
primeFactors n = factors n (2 : [3,5 ..])

factors n (p:ps)
  | n `rem` p == 0 = p : factors (n `div` p) (p : ps)
  | n == 1 = []
  | otherwise = factors n ps
