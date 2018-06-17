module Triangle
  ( rows
  ) where

rows :: Int -> [[Integer]]
rows 0 = []
rows x = map row [1 .. x]

-- We can visualize each row as the sum of the previous row
-- with itself, shifted by 1 and zero padded
row :: Int -> [Integer]
row 1 = [1]
row x = take x $ zipWith (+) zeroPaddedRow shiftedRow
  where
    shiftedRow = 0 : zeroPaddedRow
    zeroPaddedRow = row (x - 1) ++ [0 ..]
