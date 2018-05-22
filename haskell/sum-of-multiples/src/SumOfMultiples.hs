module SumOfMultiples (sumOfMultiples) where

sumOfMultiples :: [Integer] -> Integer -> Integer
sumOfMultiples factors limit = sum $ filter isMultipleOfAnyFactor [1..(limit-1)]
    where isMultipleOfAnyFactor x = any (divisible x) factors
          divisible x y = x `mod` y == 0 
