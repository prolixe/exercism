module Sieve
  ( primesUpTo
  ) where

import           Data.List.Ordered

primesUpTo :: Integer -> [Integer]
primesUpTo n
  | n < 2 = []
  | otherwise = eratos [2 .. n]
  where
    eratos []     = []
    eratos (p:xs) = p : eratos (xs `minus` [p * p,p * p + p ..])
