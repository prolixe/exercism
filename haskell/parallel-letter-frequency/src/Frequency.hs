module Frequency
  ( frequency
  ) where

import           Control.Parallel.Strategies
import           Data.Char                   (isLetter)
import qualified Data.Map                    as Map
import qualified Data.Text                   as Text

frequency :: Int -> [Text.Text] -> Map.Map Char Int
frequency nWorkers texts =
  Map.unionsWith (+) $ map fst $ parMap parFrequency getMap texts
  where
    parFrequency :: Strategy (Map.Map Char Int, Text.Text)
    parFrequency (m, t) = do
      (m, t) <- rpar ((frequency' . Text.filter isLetter . Text.toLower) t, t)
      return (m, t)
    getMap :: Text.Text -> (Map.Map Char Int, Text.Text)
    getMap text = (Map.empty, text)

frequency' :: Text.Text -> Map.Map Char Int
frequency' = Text.foldl' countChar Map.empty
  where
    countChar m c =
      if Map.member c m
        then Map.update (\x -> Just (x + 1)) c m
        else Map.insert c 1 m
