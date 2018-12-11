module Cipher
  ( caesarDecode
  , caesarEncode
  , caesarEncodeRandom
  ) where

import qualified Data.Map as Map
import Data.Maybe
import Data.Tuple (swap)
import System.Random

caesarDecode :: String -> String -> String
caesarDecode key = zipWith decode (cycle key)

caesarEncode :: String -> String -> String
caesarEncode key = zipWith encode (cycle key)

caesarEncodeRandom :: String -> IO (String, String)
caesarEncodeRandom text = do
  g <- newStdGen
  let randomKey = take (length text) $ randoms g :: String
  return (randomKey, caesarEncode randomKey text)

encode :: Char -> Char -> Char
encode key c = newChar
  where
    offset = (fromEnum key - fromEnum 'a') `mod` 26
    newChar =
      toEnum $ ((fromEnum c - fromEnum 'a' + offset) `mod` 26) + fromEnum 'a' :: Char

decode :: Char -> Char -> Char
decode key c = newChar
  where
    offset = (fromEnum 'a' - fromEnum key) `mod` 26
    newChar =
      toEnum $ ((fromEnum c - fromEnum 'a' + offset) `mod` 26) + fromEnum 'a' :: Char
