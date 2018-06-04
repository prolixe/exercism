module Anagram
  ( anagramsFor
  ) where

import           Data.Char
import           Data.List (sort)

anagramsFor :: String -> [String] -> [String]
anagramsFor xs = filter (anagramOf lxs . map toLower)
  where
    lxs = map toLower xs

anagramOf :: String -> String -> Bool
anagramOf subject candidate =
  sort subject == sort candidate && subject /= candidate
