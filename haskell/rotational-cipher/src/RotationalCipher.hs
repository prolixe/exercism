module RotationalCipher
  ( rotate
  ) where

import           Data.Char

rotate :: Int -> String -> String
rotate offset = map (rotate' offset)

rotate' :: Int -> Char -> Char
rotate' offset c
  | isLetter c =
    toEnum $ charOffset + (fromEnum c + offset - charOffset) `rem` 26
  | otherwise = c
  where
    charOffset
      | isUpper c = fromEnum 'A'
      | otherwise = fromEnum 'a'
