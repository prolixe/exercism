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
simulate = foldl applyStep

applyStep :: Robot -> Char -> Robot
applyStep (Robot b c) 'R' = mkRobot (turnRight b) c
applyStep (Robot b c) 'L' = mkRobot (turnLeft b) c
applyStep robot 'A'       = advance robot
applyStep robot _         = robot

advance :: Robot -> Robot
advance (Robot b c)
  | b == North = mkRobot b (goNorth c)
  | b == East = mkRobot b (goEast c)
  | b == South = mkRobot b (goSouth c)
  | b == West = mkRobot b (goWest c)
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
