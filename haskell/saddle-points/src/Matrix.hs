module Matrix
  ( saddlePoints
  ) where

import           Data.Array (Array, Ix, assocs)
import           Data.List  (groupBy, intersect, maximumBy, minimumBy)
import           Data.Ord   (compare)

saddlePoints :: (Ord e, Ix i) => Array (i, i) e -> [(i, i)]
saddlePoints matrix =
  map fst $ maxRowList listMatrix `intersect` minColList listMatrix
  where
    listMatrix = assocs matrix

maxRowList :: (Ord e, Eq i) => [((i, i), e)] -> [((i, i), e)]
maxRowList list =
  concatMap (\l -> filter (\l1 -> getElem l1 == maxElem l) l) (rowList list)

minColList :: (Ord e, Eq i) => [((i, i), e)] -> [((i, i), e)]
minColList list =
  concatMap (\l -> filter (\l1 -> getElem l1 == minElem l) l) (colList list)

maxElem :: Ord e => [((i, i), e)] -> e
maxElem = snd . maximumBy (\a b -> compare (snd a) (snd b))

minElem :: Ord e => [((i, i), e)] -> e
minElem = snd . minimumBy (\a b -> compare (snd a) (snd b))

rowList :: Eq i => [((i, i), e)] -> [[((i, i), e)]]
rowList = groupBy (\(r1, c1) (r2, c2) -> fst r1 == fst r2)

colList :: Eq i => [((i, i), e)] -> [[((i, i), e)]]
colList l = map (\c1 -> filter ((== c1) . getCol) l) (getCols l)

getCols :: [((i, i), e)] -> [i]
getCols = map getCol

getCol ((_, c), _) = c

getElem (_, e) = e
