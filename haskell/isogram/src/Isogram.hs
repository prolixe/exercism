module Isogram (isIsogram) where

import Data.Char

isIsogram :: String -> Bool
isIsogram xs = all isUniqueLetter loweredString
    where loweredString = map toLower xs
          isUniqueLetter x = (not . isLetter) x 
                             || 1 == length (filter (x==) loweredString)
