module RunLength
  ( decode
  , encode
  ) where

import           Data.Char (isDigit)
import           Data.List

decode :: String -> String
decode = concat . decode' . groupBy bothAreDigit
  where
    bothAreDigit x y = isDigit x && isDigit y

decode' :: [String] -> [String]
decode' [] = []
decode' [x] = [x]
decode' (x:y:xs) =
  if all isDigit x
    then replicate (read x :: Int) y ++ decode' xs
    else x : decode' (y : xs)

encode :: String -> String
encode = concatMap (uncurry tupleTostring . countChar) . group

tupleTostring :: Int -> Char -> String
tupleTostring 0 c = ""
tupleTostring 1 c = [c]
tupleTostring x c = show x ++ [c]

countChar :: String -> (Int, Char)
countChar xs = (length xs, head xs)
