module OCR
  ( convert
  ) where

import           Data.List
import           Data.List.Split (chunksOf)

convert :: String -> String
convert xs = intercalate "," $ map (concatMap convertOneDigit) $ tokenize xs

-- format a list of OCR digit into a list of list of digits
-- formatted in the same way as digits
-- The first list is for ',' separated digits
-- the second list is a list of 'digits' ([String])
-- So we end up with []
tokenize :: String -> [[[String]]]
tokenize s = map transpose xs
  where
    xs = chunksOf 4 $ map (chunksOf 3) $ lines s

convertOneDigit :: [String] -> String
convertOneDigit xs =
  case elemIndex xs digits of
    Just x  -> show x
    Nothing -> "?"

-- Unfortunately formatted this way by hfmt
-- represent 0 to 9
digits :: [[String]]
digits =
  [ [" _ ", "| |", "|_|", "   "]
  , ["   ", "  |", "  |", "   "]
  , [" _ ", " _|", "|_ ", "   "]
  , [" _ ", " _|", " _|", "   "]
  , ["   ", "|_|", "  |", "   "]
  , [" _ ", "|_ ", " _|", "   "]
  , [" _ ", "|_ ", "|_|", "   "]
  , [" _ ", "  |", "  |", "   "]
  , [" _ ", "|_|", "|_|", "   "]
  , [" _ ", "|_|", " _|", "   "]
  ]
