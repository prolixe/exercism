module ProteinTranslation
  ( proteins
  ) where

import           Data.List.Split

proteins :: String -> Maybe [String]
proteins s
  | length s `rem` 3 /= 0 = Nothing
  | otherwise = Just $ takeWhile (/= "STOP") $ map protein $ chunksOf 3 s

protein :: String -> String
protein s
  | s `elem` ["AUG"] = "Methionine"
  | s `elem` ["UUU", "UUC"] = "Phenylalanine"
  | s `elem` ["UUA", "UUG"] = "Leucine"
  | s `elem` ["UCU", "UCC", "UCA", "UCG"] = "Serine"
  | s `elem` ["UAU", "UAC"] = "Tyrosine"
  | s `elem` ["UGU", "UGC"] = "Cysteine"
  | s `elem` ["UGG"] = "Tryptophan"
  | s `elem` ["UAA", "UAG", "UGA"] = "STOP"
  | otherwise = ""
