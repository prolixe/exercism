module ResistorColors
  ( Color(..)
  , value
  ) where

data Color
  = Black
  | Brown
  | Red
  | Orange
  | Yellow
  | Green
  | Blue
  | Violet
  | Grey
  | White
  deriving (Enum, Eq, Show, Read)

value :: [Color] -> Int
value = foldl addColor 0
  where
    addColor num c = 10 * num + fromEnum c
