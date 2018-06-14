module Clock
  ( clockHour
  , clockMin
  , fromHourMin
  , toString
  , negate
  ) where

import           Text.Printf

data Clock = Clock
  { clockHour :: Int
  , clockMin  :: Int
  } deriving (Eq, Show)

fromHourMin :: Int -> Int -> Clock
fromHourMin = Clock

toString :: Clock -> String
toString clock =
  printf "%02d" (clockHour clock) ++ ":" ++ printf "%02d" (clockMin clock)
