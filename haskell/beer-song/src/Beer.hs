module Beer
  ( song
  ) where

verse :: Int -> String
verse x =
  show x ++
  " bottle" ++
  plural x ++
  " of beer on the wall, " ++
  show x ++
  " bottle" ++
  plural x ++
  " of beer.\n\
         \Take one down and pass it around, " ++
  show (x - 1) ++ " bottle" ++ plural (x - 1) ++ " of beer on the wall.\n\n"

plural :: Int -> String
plural 1 = ""
plural x = "s"

-- I give up for the last 2 "loop". It's not worth making more special case.
song :: String
song =
  concatMap verse [99,98 .. 2] ++
  "1 bottle of beer on the wall, 1 bottle of beer.\n\
         \Take it down and pass it around, no more bottles of beer on the wall.\n\
         \\n\
         \No more bottles of beer on the wall, no more bottles of beer.\n\
         \Go to the store and buy some more, 99 bottles of beer on the wall.\n"
