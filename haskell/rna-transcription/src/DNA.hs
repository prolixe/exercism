module DNA (toRNA) where

import Data.Maybe

toRNA :: String -> Maybe String
toRNA xs = if any isNothing [transcribeRNA x | x <- xs ]
           then Nothing else Just $ map fromJust [transcribeRNA x | x <- xs ] 

transcribeRNA :: Char -> Maybe Char
transcribeRNA c 
    | c == 'G' = Just 'C'
    | c == 'C' = Just 'G'
    | c == 'T' = Just 'A'
    | c == 'A' = Just 'U'
    | otherwise = Nothing

