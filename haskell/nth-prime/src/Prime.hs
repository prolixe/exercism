module Prime
  ( nth
  ) where

import           Data.List  (unfoldr)
import           Data.Maybe (listToMaybe)

nth :: Int -> Maybe Integer
nth n
  | n > 0 = Just $ (last . take n) $ filter isPrime (2 : [3,5 ..])
  | otherwise = Nothing

factors :: Integer -> [Integer]
factors n =
    unfoldr (\(d,n) -> listToMaybe [(x, (x, div n x)) -- (factor, (factor, n divided by factor))
                                    | x <- takeWhile ((<=n).(^2)) [d..] ++ [n|n>1] -- take all x starting from d while under sqrt n + n
                                    , mod n x==0]) -- if x is a factor of n
                                    (2,n) -- start with 2


isPrime :: Integer -> Bool
isPrime n = n > 1 && head (factors n) == n
