module ComplexNumbers
  ( Complex
  , conjugate
  , abs
  , real
  , imaginary
  , mul
  , add
  , sub
  , div
  , complex
  ) where

import           Prelude hiding (abs, div)

-- Data definition -------------------------------------------------------------
type Complex = (Float, Float)

complex :: (Float, Float) -> Complex
complex (a, b) = (a, b)

-- unary operators -------------------------------------------------------------
conjugate :: Complex -> Complex
conjugate (r, i) = (r, negate i)

abs :: Complex -> Float
abs (r, i) = sqrt (r ** 2 + i ** 2)

real :: Complex -> Float
real (r, i) = r

imaginary :: Complex -> Float
imaginary (r, i) = i

-- binary operators ------------------------------------------------------------
mul :: Complex -> Complex -> Complex
mul (r1, i1) (r2, i2) = (r1 * r2 - i1 * i2, r1 * i2 + i1 * r2)

add :: Complex -> Complex -> Complex
add (r1, i1) (r2, i2) = (r1 + r2, i1 + i2)

sub :: Complex -> Complex -> Complex
sub (r1, i1) (r2, i2) = (r1 - r2, i1 - i2)

div :: Complex -> Complex -> Complex
div (r1, i1) (r2, i2) = (real, imaginary)
  where
    real = (r1 * r2 + i1 * i2) / (r2 ** 2 + i2 ** 2)
    imaginary = (i1 * r2 - r1 * i2) / (r2 ** 2 + i2 ** 2)
