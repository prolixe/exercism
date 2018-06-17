module Triangle
  ( rows
  ) where

rows :: Int -> [[Integer]]
rows 0 = []
rows x = map row [1 .. x]

row :: Int -> [Integer]
row 0 = []
row x = 1 : zipWith (+) (row (x - 1)) shiftedRow
  where
    shiftedRow = drop 1 (row (x - 1)) ++ [0 ..]
-- We can visualize each row as the sum of the previous row
-- with itself, shifted by 1.
