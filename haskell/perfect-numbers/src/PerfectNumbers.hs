module PerfectNumbers (classify, Classification(..)) where

data Classification = Deficient | Perfect | Abundant deriving (Eq, Show)

classify :: Int -> Maybe Classification
classify x
    | x < 1 = Nothing
    | facSum < x = Just Deficient
    | facSum == x = Just Perfect 
    | facSum > x = Just Abundant
    where facSum = sum $ factorList x


factorList :: Int -> [Int]
factorList n = [x | x <- [1..n-1] , n `mod` x == 0 ]    


