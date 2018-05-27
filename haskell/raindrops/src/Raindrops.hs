module Raindrops
  ( convert
  ) where

convert :: Int -> String
convert n =
  if null $ rain n
    then show n
    else rain n
  where
    rain n = pling n ++ plang n ++ plong n

plang :: Int -> String
plang x =
  if x `rem` 5 == 0
    then "Plang"
    else ""

plong :: Int -> String
plong x =
  if x `rem` 7 == 0
    then "Plong"
    else ""

pling :: Int -> String
pling x =
  if x `rem` 3 == 0
    then "Pling"
    else ""
