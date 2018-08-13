module CustomSet
  ( delete
  , difference
  , empty
  , fromList
  , insert
  , intersection
  , isDisjointFrom
  , isSubsetOf
  , member
  , null
  , size
  , toList
  , union
  ) where

import           Data.List (nub, sort, (\\))
import           Prelude   hiding (null)

type CustomSet a = [a]

delete :: Eq a => a -> CustomSet a -> CustomSet a
delete x set = set \\ [x]

difference :: Eq a => CustomSet a -> CustomSet a -> CustomSet a
difference setA setB = setA \\ setB

empty :: Eq a => CustomSet a
empty = []

fromList :: (Ord a, Eq a) => [a] -> CustomSet a
fromList = sort . nub

insert :: (Ord a, Eq a) => a -> CustomSet a -> CustomSet a
insert x set
  | x `elem` set = set
  | otherwise = (sort . nub) (x : (toList set))

intersection :: Eq a => CustomSet a -> CustomSet a -> CustomSet a
intersection setA setB = [x | x <- setA, x `elem` setB]

isDisjointFrom :: Eq a => CustomSet a -> CustomSet a -> Bool
isDisjointFrom setA setB = all (`notElem` toList setB) $ toList setA

isSubsetOf :: Eq a => CustomSet a -> CustomSet a -> Bool
isSubsetOf setA setB = all (`elem` toList setB) setA

member :: Eq a => a -> CustomSet a -> Bool
member x set = x `elem` (toList set)

null :: CustomSet a -> Bool
null set = length set == 0

size :: Eq a => CustomSet a -> Int
size set = length $ toList set

toList :: CustomSet a -> [a]
toList set = set

union :: (Ord a, Eq a) => CustomSet a -> CustomSet a -> CustomSet a
union setA setB = sort $ nub (setA ++ setB)
