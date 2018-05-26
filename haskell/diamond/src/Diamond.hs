module Diamond (diamond) where

import Data.Char

diamond :: Char -> Maybe [String]
diamond c 
    | isLetter c = Just (diamond' c 0 ++ (tail . reverse) (diamond' c 0))
    | otherwise = Nothing

diamond' :: Char -> Int -> [String]
diamond' 'A' x = [createRow 'A' x]
diamond' c x = diamond' (pred c) (x+1) ++ [createRow c x] 

createRow :: Char -> Int -> String
createRow 'A' l = addSpace l ++ ['A'] ++ addSpace l
createRow c l = addSpace l ++ [c] ++ innerSpace c ++ [c] ++ addSpace l
    
addSpace :: Int-> String
addSpace x = replicate x ' '

innerSpace :: Char -> String
innerSpace 'A' = ""
innerSpace c = addSpace (innerSpaceMap c) 
    where innerSpaceMap ch = snd . head $ filter ((==ch) . fst ) $ zip ['B'..'Z'] [1, 3 ..]




