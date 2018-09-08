module RailFenceCipher
  ( encode
  , decode
  ) where

import Data.List (nub, sort, sortBy)

encode :: Int -> [a] -> [a]
encode rail message = concatMap encodeRow [0 .. rail - 1]
  where
    indexedMessage = zip [0 ..] message
    step = (rail - 1) * 2
    encodeRow row = map snd $ filter (isElemInRow step row . fst) indexedMessage

-- isElemInRow use the regular pattern of the rail fence to 
-- decide if a particular position in the list is on the given row
isElemInRow :: Int -> Int -> Int -> Bool
isElemInRow step row pos =
  (pos + row) `rem` step == 0 || (pos - row) `rem` step == 0

--decode undo the encoding by mapping an index to every element and sort them
decode :: Int -> [a] -> [a]
decode rail message =
  map snd $ sortBy (\(a, _) (b, _) -> compare a b) mappedMessage
  where
    mapIndex = mapPos rail (length message)
    mappedMessage = zip mapIndex message

-- Map pos generate a tuple of encoded position
mapPos :: Int -> Int -> [Int]
mapPos rail len = encode rail [0 .. len - 1]
