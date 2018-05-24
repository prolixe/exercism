module School (School, add, empty, grade, sorted) where

import Data.List

type School = [(Int, String)]

add :: Int -> String -> School -> School
add gradeNum student school = (gradeNum, student):school

empty :: School
empty = []

grade :: Int -> School -> [String]
grade gradeNum school = sort [snd x | x <- school, fst x == gradeNum]

sorted :: School -> [(Int, [String])]
sorted school = sort $ nub [ (g,  grade g school)| (g,_) <- school ]
