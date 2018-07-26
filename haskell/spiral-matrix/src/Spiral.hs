module Spiral
  ( spiral
  ) where

import           Data.List (transpose)

-- Algo taken from Rosetta Stone.
spiral' :: Int -> Int -> Int -> [[Int]]
spiral' 0 _ _ = [[]]
spiral' h w s = [s .. s + w - 1] : rot90 (spiral' w (h - 1) (s + w))
  where
    rot90 = map reverse . transpose

spiral :: Int -> [[Int]]
spiral 0    = []
spiral size = spiral' size size 1
--
-- How it works
-- spiral' 3 3 1
-- gets expanded into the following
-- [1,2,3]:rot90 ([4,5]:rot90( [6,7]:rot90([8]:rot90 ([9]:[[]]))))
-- so we can conceptually think of this as a spiral of dimension 1 1 starting with 9 (spiral' 1 1 9),
-- appended by a spiral' 2 1 8 etc etc
--
-- Each rot90 rotates the whole tail of the list!
--
