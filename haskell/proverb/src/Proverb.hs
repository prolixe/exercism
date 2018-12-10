module Proverb
  ( recite
  ) where

recite :: [String] -> String
recite [] = ""
recite list@(want:xs) = concatMap (uncurry forWant) forWantList ++ allFor want
  where
    forWantList = zip list xs

forWant :: String -> String -> String
forWant want lost =
  "For want of a " ++ want ++ " the " ++ lost ++ " was lost.\n"

allFor :: String -> String
allFor want = "And all for the want of a " ++ want ++ "."
