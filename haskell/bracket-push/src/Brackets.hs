module Brackets
  ( arePaired
  ) where

import           Data.Maybe

import           Data.List  (foldl')

data Bracket
  = Round
  | Curly
  | Square -- () or {} or []
  deriving (Eq)

data Side
  = Open
  | Close
  deriving (Eq)

toBracket :: Char -> Maybe Bracket
toBracket x
  | x `elem` "()" = Just Round
  | x `elem` "{}" = Just Curly
  | x `elem` "[]" = Just Square
  | otherwise = Nothing

toBraketSide :: Char -> Maybe Side
toBraketSide x
  | x `elem` "({[" = Just Open
  | x `elem` ")}]" = Just Close
  | otherwise = Nothing

arePaired :: String -> Bool
arePaired xs =
  null $ foldl' matchBracket [] $ zip (map toBracket xs) (map toBraketSide xs)

matchBracket ::
     [Maybe Bracket] -> (Maybe Bracket, Maybe Side) -> [Maybe Bracket]
matchBracket brackets (b, side)
  | side == Just Close && not (any isJust brackets) = [Nothing]
  | side == Just Close && head brackets /= b = [Nothing]
  | side == Just Close && head brackets == b = tail brackets
  | side == Just Open = b : brackets
  | otherwise = brackets
