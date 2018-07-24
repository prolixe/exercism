module Queens
  ( boardString
  , canAttack
  ) where

boardString :: Maybe (Int, Int) -> Maybe (Int, Int) -> String
boardString w b =
  unlines $
  map
    (unwords . map ((: []) . decode))
    [[(x, y) | y <- [0 .. 7]] | x <- [0 .. 7]]
  where
    decode x =
      if Just x == w
        then 'W'
        else if Just x == b
               then 'B'
               else '_'

canAttack :: (Int, Int) -> (Int, Int) -> Bool
canAttack (x, y) (x', y') =
  x == x' || y == y' || x - x' == y - y' || x - x' == y' - y
