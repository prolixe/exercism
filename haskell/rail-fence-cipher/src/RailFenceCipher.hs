module RailFenceCipher
  ( encode
  , decode
  ) where

import Data.List (nub, sort, sortBy)

encode :: Int -> [a] -> [a]
encode rail message = concatMap encodeRow [0 .. rail - 1]
  where
    indexedMessage = zip [0 ..] message
    encodeRow row = map snd $ filter (isElemInRow rail row . fst) indexedMessage

-- isElemInRow use the regular pattern of the rail fence to 
-- decide if a particular position in the list is on the given row
isElemInRow :: Int -> Int -> Int -> Bool
isElemInRow rail row pos =
  (pos + row) `rem` step == 0 || (pos - row) `rem` step == 0
  where
    step = (rail - 1) * 2

--decode undo the encoding by mapping an index to every element and sort them
decode :: Ord a => Int -> [a] -> [a]
decode rail message = map snd $ sort mappedMessage
  where
    index = encodedPos rail (length message)
    mappedMessage = zip index message

-- encodedPos generate a list of encoded position
-- such as the value as index x is the original position 
-- of the xth element 
-- (e.g. [0,4,1,3,2] means that the 2nd element was originally at the 4th )
encodedPos :: Int -> Int -> [Int]
encodedPos rail len = encode rail [0 .. len - 1]
