module Triangle
  ( TriangleType(..)
  , triangleType
  ) where

import           Data.List

data TriangleType
  = Equilateral
  | Isosceles
  | Scalene
  | Illegal
  deriving (Eq, Show)

triangleType :: (Num a, Ord a) => a -> a -> a -> TriangleType
triangleType a b c
  | any (<= 0) tList = Illegal
  | sideTooLarge = Illegal
  | and sides = Equilateral
  | 1 == length (filter id sides) = Isosceles
  | otherwise = Scalene
  where
    tList = [a, b, c]
    sides = [a == b, b == c, c == a]
    sideTooLarge = longestSide > sum (tList \\ [longestSide])
    longestSide = maximum tList
