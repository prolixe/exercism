module Matrix
  ( Matrix
  , cols
  , column
  , flatten
  , fromList
  , fromString
  , reshape
  , row
  , rows
  , shape
  , transpose
  ) where

import           Data.Char
import qualified Data.List       as L
import           Data.List.Split (chunksOf)
import qualified Data.Vector     as V

newtype Matrix a =
  Matrix [[a]]
  deriving (Eq, Show)

cols :: Matrix a -> Int
cols (Matrix m) = length (head m)

column :: Int -> Matrix a -> V.Vector a
column i (Matrix m) = V.fromList (map (!! i) m)

flatten :: Matrix a -> V.Vector a
flatten (Matrix m) = V.fromList (concat m)

fromList :: [[a]] -> Matrix a
fromList = Matrix

fromString :: Read a => String -> Matrix a
fromString xs = fromList (map (map read . words) (lines xs))

reshape :: (Int, Int) -> Matrix a -> Matrix a
reshape dimensions (Matrix m) = fromList (chunksOf (fst dimensions) (concat m))

row :: Int -> Matrix a -> V.Vector a
row x (Matrix m) = V.fromList (m !! x)

rows :: Matrix a -> Int
rows (Matrix m) = length m

shape :: Matrix a -> (Int, Int)
shape (Matrix m)
  | null m = (0, 0)
  | otherwise = (length m, length (head m))

transpose :: Matrix a -> Matrix a
transpose (Matrix m) = fromList (L.transpose m)
