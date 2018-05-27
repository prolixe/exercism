module Accumulate
  ( accumulate
  ) where

-- no collect, map, fmap here! just list comprehension!
-- or the applicative style, I guess.
accumulate :: (a -> b) -> [a] -> [b]
--accumulate f xs = [f x | x <- xs]
accumulate f xs = f <$> xs
