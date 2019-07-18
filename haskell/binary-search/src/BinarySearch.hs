module BinarySearch
  ( find
  ) where

import Data.Array

find :: Ord a => Array Int a -> a -> Maybe Int
find arr key
  | end < start = Nothing
  | m /= key && start == end = Nothing
  | m == key = Just middle
  | m > key = find' arr start (middle - 1) key
  | m < key = find' arr (middle + 1) end key
  where
    (start, end) = bounds arr
    middle = start + ((end - start) `div` 2)
    m = arr ! middle
find _ _ = Nothing

find' :: Ord a => Array Int a -> Int -> Int -> a -> Maybe Int
find' arr start end key
  | end < start = Nothing
  | m /= key && start == end = Nothing
  | m == key = Just middle
  | m > key = find' arr start (middle - 1) key
  | m < key = find' arr (middle + 1) end key
  where
    middle = start + ((end - start) `div` 2)
    m = arr ! middle
find' _ _ _ _ = Nothing
