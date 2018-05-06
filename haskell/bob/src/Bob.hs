module Bob (responseFor) where

import Data.Char

responseFor :: String -> String
responseFor xs 
    | isYelling xs && '?' == last xs  = "Calm down, I know what I'm doing!"
    | isYelling xs = "Whoa, chill out!"
    | not ( all isSpace xs) && '?' == last (filter (not . isSpace) xs) = "Sure."
    | not $ any isAlphaNum xs  = "Fine. Be that way!"
    | otherwise = "Whatever."
    where isYelling xs = all (\c -> c `notElem` ['a'..'z'])  xs && any (\c -> c `elem` ['A'..'Z']) xs
            
