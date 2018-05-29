module Beer
  ( song
  ) where

import           Data.List.Utils

template =
  "$BOTTLES bottle$PLURAL of beer on the wall, $BOTTLES bottle$PLURAL of beer.\n\
         \Take one down and pass it around, $MINUS1 bottle$PLURAL of beer on the wall.\n\n"

-- I give up for the last 3 "loop". It's not worth making more special case.
song :: String
song =
  concatMap songRefrain [99,98 .. 3] ++
  "2 bottles of beer on the wall, 2 bottles of beer.\n\
         \Take one down and pass it around, 1 bottle of beer on the wall.\n\
         \\n\
         \1 bottle of beer on the wall, 1 bottle of beer.\n\
         \Take it down and pass it around, no more bottles of beer on the wall.\n\
         \\n\
         \No more bottles of beer on the wall, no more bottles of beer.\n\
         \Go to the store and buy some more, 99 bottles of beer on the wall.\n"

songRefrain :: Int -> String
songRefrain x = songPluralized
  where
    songBottles = replaceAll "$BOTTLES" (show x) template
    songBottles2 = replaceAll "$MINUS1" (show (x - 1)) songBottles
    songPluralized =
      last . take 3 $
      iterate
        (replace
           "$PLURAL"
           (if x /= 1
              then "s"
              else ""))
        songBottles2

replaceAll :: String -> String -> String -> String
replaceAll toFind toReplace toParse =
  last . take 10 $ iterate (replace toFind toReplace) toParse
