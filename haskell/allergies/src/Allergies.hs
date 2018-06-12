module Allergies
  ( Allergen(..)
  , allergies
  , isAllergicTo
  ) where

import           Data.Bits

data Allergen
  = Eggs
  | Peanuts
  | Shellfish
  | Strawberries
  | Tomatoes
  | Chocolate
  | Pollen
  | Cats
  deriving (Eq, Enum)

allergies :: Int -> [Allergen]
allergies score
  | popCount (score .&. bit 0) == 1 = Eggs : allergies (score - bit 0)
  | popCount (score .&. bit 1) == 1 = Peanuts : allergies (score - bit 1)
  | popCount (score .&. bit 2) == 1 = Shellfish : allergies (score - bit 2)
  | popCount (score .&. bit 3) == 1 = Strawberries : allergies (score - bit 3)
  | popCount (score .&. bit 4) == 1 = Tomatoes : allergies (score - bit 4)
  | popCount (score .&. bit 5) == 1 = Chocolate : allergies (score - bit 5)
  | popCount (score .&. bit 6) == 1 = Pollen : allergies (score - bit 6)
  | popCount (score .&. bit 7) == 1 = Cats : allergies (score - bit 7)
  | otherwise = []

isAllergicTo :: Allergen -> Int -> Bool
isAllergicTo allergen score = allergen `elem` allergies score
