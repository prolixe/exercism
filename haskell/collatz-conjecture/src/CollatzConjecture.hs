module CollatzConjecture (collatz) where

collatz :: Integer -> Maybe Integer
collatz x 
    | x <= 0 = Nothing
    | otherwise = Just $ fromIntegral (length $ collatzList x)-1

collatzList :: Integer -> [Integer]
collatzList n
    | n <= 0 = []
    | n == 1 = [n]
    | otherwise = n:collatzList (collatzNext n)
    where collatzNext n = if even n then  n `div` 2 else 3*n+1

