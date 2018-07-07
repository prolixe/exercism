module Queens
  ( boardString
  , canAttack
  ) where

import           Data.List
import           Data.Maybe

boardString :: Maybe (Int, Int) -> Maybe (Int, Int) -> String
boardString white black
  | isNothing white && isNothing black = emptyBoard
  | isNothing white = placeInRank 'B' (fromJust black) emptyBoard
  | isNothing black = placeInRank 'W' (fromJust white) emptyBoard
  | otherwise =
    placeInRank 'B' (fromJust black) $
    placeInRank 'W' (fromJust white) emptyBoard
  where
    emptyBoard = unlines $ replicate 8 $ intersperse ' ' $ replicate 8 '_'
    placeInRank c queen board =
      unlines $
      map
        (\(i, r) ->
           if i == fst queen
             then placeQueen c queen
             else r) $
      zip [0 ..] (lines board)
    placeQueen c queen =
      intersperse ' ' $
      map
        (\i ->
           if snd queen == i
             then c
             else '_')
        [0 .. 7]

canAttack :: (Int, Int) -> (Int, Int) -> Bool
canAttack (rankA, fileA) (rankB, fileB) =
  rankA == rankB || fileA == fileB || abs (rankA - rankB) == abs (fileA - fileB)
