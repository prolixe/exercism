module ListOps
  ( length
  , reverse
  , map
  , filter
  , foldr
  , foldl'
  , (++)
  , concat
  ) where

import           Prelude hiding (concat, filter, foldr, length, map, reverse,
                          (++))

foldl' :: (b -> a -> b) -> b -> [a] -> b
foldl' f z [] = z
foldl' f z (x:xs) =
  let z' = f z x
  in seq z' $ foldl' f z' xs

foldr :: (a -> b -> b) -> b -> [a] -> b
foldr f z []     = z
foldr f z (x:xs) = f x $ foldr f z xs

length :: [a] -> Int
length xs = sum [1 | x <- xs]

reverse :: [a] -> [a]
reverse []     = []
reverse (x:xs) = reverse xs ++ [x]

map :: (a -> b) -> [a] -> [b]
map f xs = [f x | x <- xs]

filter :: (a -> Bool) -> [a] -> [a]
filter _ [] = []
filter p (x:xs)
  | p x = x : filter p xs
  | otherwise = filter p xs

(++) :: [a] -> [a] -> [a]
(++) [] ys     = ys
(++) (x:xs) ys = x : xs ++ ys

concat :: [[a]] -> [a]
concat = foldr (++) []
