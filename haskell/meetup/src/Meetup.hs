module Meetup
  ( Weekday(..)
  , Schedule(..)
  , meetupDay
  ) where

import qualified Data.Time.Calendar          as D

import           Data.List                   (nub)
import           Data.Time.Calendar.WeekDate (toWeekDate)

data Weekday
  = Monday
  | Tuesday
  | Wednesday
  | Thursday
  | Friday
  | Saturday
  | Sunday
  deriving (Enum)

data Schedule
  = First
  | Second
  | Third
  | Fourth
  | Last
  | Teenth
  deriving (Eq)

meetupDay :: Schedule -> Weekday -> Integer -> Int -> D.Day
meetupDay schedule weekday year month =
  head $ filter (matchWeekDay . toWeekDate) (candidates schedule year month)
  where
    matchWeekDay (_, _, w) = (fromEnum weekday + 1) == w

candidates :: Schedule -> Integer -> Int -> [D.Day]
candidates schedule year month
  | schedule == First = dayList [1 .. 7]
  | schedule == Second = dayList [8 .. 14]
  | schedule == Third = dayList [15 .. 21]
  | schedule == Fourth = dayList [22 .. 28]
  | schedule == Last = take 7 $ nub $ dayList [31,30 .. 22]
  | schedule == Teenth = dayList [13 .. 19]
  where
    dayList = map (D.fromGregorian year month)
