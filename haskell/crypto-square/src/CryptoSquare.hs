module CryptoSquare
  ( encode
  ) where

import           Data.Char       (isAlphaNum, toLower)
import           Data.List       (transpose)
import           Data.List.Split

encode :: String -> String
encode xs = unwords $ transpose $ chunksOf (rect (length normalized)) normalized
  where
    normalized = normalize xs

rect :: Int -> Int
rect = ceiling . (sqrt :: Double -> Double) . fromIntegral

normalize :: String -> String
normalize = map toLower . filter isAlphaNum
