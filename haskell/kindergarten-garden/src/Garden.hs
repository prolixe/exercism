module Garden
  ( Plant(..)
  , garden
  , lookupPlants
  ) where

data Plant
  = Clover
  | Grass
  | Radishes
  | Violets
  deriving (Eq, Show)

type Garden = [(String, [Plant])]

garden :: [String] -> String -> Garden
garden students plants = zip students plantsList
  where
    plantsList = map (map toPlant) $ parsePlants plants

-- Group the flat list of plants into a list of string 4 char long.
parsePlants :: String -> [String]
parsePlants p =
  map
    (\x -> concatMap (take 2 . drop x) (lines p))
    [0,2 .. quot (length p - 1) 2] -- minus 1 for the \n, divide by 2 for the 2 rows

toPlant :: Char -> Plant
toPlant 'C' = Clover
toPlant 'G' = Grass
toPlant 'R' = Radishes
toPlant 'V' = Violets

lookupPlants :: String -> Garden -> [Plant]
lookupPlants student garden = snd . head $ filter ((== student) . fst) garden
