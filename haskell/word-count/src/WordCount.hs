module WordCount
  ( wordCount
  ) where

import           Data.Char       (isAlphaNum, toLower)
import           Data.List       (nub)
import           Data.List.Extra
import           Text.Regex

wordCount :: String -> [(String, Int)]
wordCount xs = nub $ map (\w -> (w, length $ filter (== w) ws)) ws
  where
    ws =
      map (dropSuffix "'" . dropPrefix "'" . map toLower) $
      filter (not . null) $
      map (filter isAlphaNumPlusAp) $ splitRegex (mkRegex "[, \n]") xs
      where
        isAlphaNumPlusAp x = isAlphaNum x || x == '\''
