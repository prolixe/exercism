module DNA (nucleotideCounts) where

import qualified Data.Map as Map  

nucleotideCounts :: String -> Either String (Map.Map Char Int)
nucleotideCounts xs 
    | any notANucleotide xs = Left xs
    | otherwise = Right $ Map.fromList $  map countNucleotide ['A', 'C', 'G', 'T'] 
    where notANucleotide x =  x `notElem` ['A', 'C', 'G', 'T']
          countNucleotide x= (x, fromIntegral (length (filter (x==) xs)))



