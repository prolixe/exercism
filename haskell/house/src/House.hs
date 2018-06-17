module House
  ( rhyme
  ) where

import           Data.List

rhyme :: String
rhyme = intercalate "\n" (map verse [0 .. length list - 1])

verse :: Int -> String
verse v = "This is" ++ verse' v

verse' :: Int -> String
verse' v
  | v < 0 = ""
  | otherwise = list !! v ++ verse' (v - 1)

list =
  reverse
    [ " the horse and the hound and the horn\nthat belonged to"
    , " the farmer sowing his corn\nthat kept"
    , " the rooster that crowed in the morn\nthat woke"
    , " the priest all shaven and shorn\nthat married"
    , " the man all tattered and torn\nthat kissed"
    , " the maiden all forlorn\nthat milked"
    , " the cow with the crumpled horn\nthat tossed"
    , " the dog\nthat worried"
    , " the cat\nthat killed"
    , " the rat\nthat ate"
    , " the malt\nthat lay in"
    , " the house that Jack built.\n"
    ]
