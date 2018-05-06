module Pangram (isPangram) where

import Data.Char

isPangram :: String -> Bool
isPangram text = all (\x -> x `elem` [toLower i | i <- text ]) ['a'..'z'] 
