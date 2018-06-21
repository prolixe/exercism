module Atbash
  ( decode
  , encode
  ) where

import           Data.Char
import           Data.List.Split (chunksOf)

decode :: String -> String
decode = filter isAlphaNum . map atbashCipher
  where
    atbashCipher c =
      if c `elem` ['a' .. 'z']
        then reversed !! (ord c - ord 'a')
        else c
    reversed = ['z','y' .. 'a']

encode :: String -> String
encode = unwords . chunksOf 5 . decode . filter isAlphaNum . map toLower
