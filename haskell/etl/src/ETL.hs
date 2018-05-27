module ETL
  ( transform
  ) where

import           Data.Map (Map)

transform :: Map a String -> Map Char a
transform legacyDatat = error "You need to implement this function."
