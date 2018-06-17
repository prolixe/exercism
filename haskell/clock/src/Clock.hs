module Clock
  ( clockHour
  , clockMin
  , fromHourMin
  , toString
  , negate
  ) where

import           Text.Printf

newtype Clock =
  Clock Int
  deriving (Show)

clockHour :: Clock -> Int
clockHour (Clock c) = (c `div` 60) `mod` 24

clockMin :: Clock -> Int
clockMin (Clock c) = c `mod` 60

fromHourMin :: Int -> Int -> Clock
fromHourMin hour min = Clock ((hour * 60) + min)

toString :: Clock -> String
toString clock =
  printf "%02d" (clockHour clock) ++ ":" ++ printf "%02d" (clockMin clock)

instance Num Clock where
  (+) (Clock c1) (Clock c2) = Clock (c1 + c2)
  (-) (Clock c1) (Clock c2) = Clock (c1 - c2)
  (*) (Clock c1) (Clock c2) = Clock (c1 * c2)
  negate (Clock c) = Clock (abs (c - 24 * 60))
  abs (Clock c) = Clock (abs c)
  fromInteger i = Clock (fromIntegral i)
  signum (Clock c) = Clock (signum c)

instance Eq Clock where
  (==) c1 c2 = toString c1 == toString c2
