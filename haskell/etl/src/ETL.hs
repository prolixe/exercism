module ETL
  ( transform
  ) where

import           Data.Char
import           Data.Map  (Map, fromList, toList, union)

transform :: Map a String -> Map Char a
transform legacyData =
  fromList (concatMap (uncurry rearrangeList) (toList legacyData))
  where
    rearrangeList k v = [(toLower c, k) | c <- v]
