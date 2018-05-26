module Isogram (isIsogram) where

import Data.List (nub)
import Data.Char (toLower, isAlpha)

-- This way is really the cleanest!
isIsogram :: String -> Bool
isIsogram = (\x -> nub x == x) . map toLower . filter isAlpha
