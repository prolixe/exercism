module Matrix
  ( saddlePoints
  ) where

import           Data.Array (Array, Ix, bounds, indices, (!))

saddlePoints :: (Ord e, Ix i, Enum i) => Array (i, i) e -> [(i, i)]
saddlePoints matrix = filter isSaddle $ indices matrix
  where
    isSaddle (r, c) = isMaxRow && isMinCol
      where
        isMaxRow = all (\col -> e >= matrix ! (r, col)) [cmin .. cmax]
        isMinCol = all (\row -> e <= matrix ! (row, c)) [rmin .. rmax]
        e = matrix ! (r, c)
        ((rmin, cmin), (rmax, cmax)) = bounds matrix
