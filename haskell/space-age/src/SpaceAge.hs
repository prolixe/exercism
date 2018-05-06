module SpaceAge (Planet(..), ageOn) where

data Planet = Mercury
            | Venus
            | Earth
            | Mars
            | Jupiter
            | Saturn
            | Uranus
            | Neptune


earthSecondsPerYear = 31557600 
mercurySecondsPerYear =  earthSecondsPerYear * 0.2408467
venusSecondsPerYear =  earthSecondsPerYear * 0.61519726 
marsSecondsPerYear =  earthSecondsPerYear * 1.8808158 
jupiterSecondsPerYear =  earthSecondsPerYear * 11.862615
saturnSecondsPerYear =  earthSecondsPerYear * 29.447498
uranusSecondsPerYear =  earthSecondsPerYear * 84.016846
netpuneSecondsPerYear =  earthSecondsPerYear * 164.79132


ageOn :: Planet -> Float -> Float
ageOn Mercury seconds = seconds / mercurySecondsPerYear 
ageOn Venus seconds = seconds / venusSecondsPerYear 
ageOn Earth seconds = seconds / earthSecondsPerYear
ageOn Mars seconds = seconds / marsSecondsPerYear
ageOn Jupiter seconds = seconds / jupiterSecondsPerYear 
ageOn Saturn seconds = seconds / saturnSecondsPerYear 
ageOn Uranus seconds = seconds / uranusSecondsPerYear 
ageOn Neptune seconds = seconds / netpuneSecondsPerYear 
