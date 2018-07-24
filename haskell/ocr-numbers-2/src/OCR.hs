module OCR
  ( convert
  ) where

import           Data.List
import           Data.List.Split (chunksOf)

convert :: String -> String
convert xs = intercalate "," $ map (concatMap convertOneDigit) $ tokenize xs

-- format a list of OCR digit into a list of list of digits
-- The outer list is for ',' separated digits
tokenize :: String -> [[[String]]]
tokenize s = map transpose xs
  where
    xs = chunksOf 4 $ map (chunksOf 3) $ lines s

convertOneDigit :: [String] -> String
convertOneDigit xs =
  case elemIndex xs digits of
    Just x  -> show x
    Nothing -> "?"

digits :: [[String]]
digits = concat $ tokenize $ unlines d
  where
    d =
      [ " _     _  _     _  _  _  _  _ "
      , "| |  | _| _||_||_ |_   ||_||_|"
      , "|_|  ||_  _|  | _||_|  ||_| _|"
      , "                              "
      ]
