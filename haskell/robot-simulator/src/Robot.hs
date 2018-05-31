module Robot
  ( Bearing(East, North, South, West)
  , bearing
  , coordinates
  , mkRobot
  , simulate
  , turnLeft
  , turnRight
  ) where

data Bearing
  = North
  | East
  | South
  | West
  deriving (Eq, Show, Enum)

data Robot = Robot
  { bearing     :: Bearing
  , coordinates :: (Integer, Integer)
  }

mkRobot :: Bearing -> (Integer, Integer) -> Robot
mkRobot = Robot

simulate :: Robot -> String -> Robot
simulate = foldl applyInstruction

applyInstruction :: Robot -> Char -> Robot
applyInstruction r 'R' = mkRobot (turnRight (bearing r)) (coordinates r)
applyInstruction r 'L' = mkRobot (turnLeft (bearing r)) (coordinates r)
applyInstruction r 'A' = advance r
applyInstruction r _   = r

advance :: Robot -> Robot
advance r
  | bearing r == North = mkRobot (bearing r) (goNorth $ coordinates r)
  | bearing r == East = mkRobot (bearing r) (goEast $ coordinates r)
  | bearing r == South = mkRobot (bearing r) (goSouth $ coordinates r)
  | bearing r == West = mkRobot (bearing r) (goWest $ coordinates r)
  where
    goNorth c = (fst c, snd c + 1)
    goEast c = (fst c + 1, snd c)
    goSouth c = (fst c, snd c - 1)
    goWest c = (fst c - 1, snd c)

turnLeft :: Bearing -> Bearing
turnLeft direction =
  if direction == North
    then West
    else pred direction

turnRight :: Bearing -> Bearing
turnRight direction =
  if direction == West
    then North
    else succ direction