module PrimeFactors
  ( primeFactors
  ) where

import           Data.List.Ordered (minus)

primeFactors :: Integer -> [Integer]
primeFactors n
  | any isPrimeOf primeList = primeFactor : primeFactors (n `div` primeFactor)
  | otherwise = []
  where
    primeList = primesToQ n
    isPrimeOf p = n `rem` p == 0
    primeFactor = head $ dropWhile (not . isPrimeOf) primeList

-- taken from https://wiki.haskell.org/Prime_numbers#Sieve_of_Eratosthenes
primesToQ :: Integer -> [Integer]
primesToQ m = eratos [2 .. m]
  where
    eratos []     = []
    eratos (p:xs) = p : eratos (xs `minus` [p * p,p * p + p .. m])
