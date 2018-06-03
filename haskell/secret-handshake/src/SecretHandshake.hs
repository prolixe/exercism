module SecretHandshake
  ( handshake
  ) where

import           Data.Bits

handshake :: Int -> [String]
handshake n
  | popCount (n .&. bit 4) == 1 = reverse (handshake (n - bit 4))
  | popCount (n .&. bit 0) == 1 = "wink" : handshake (n - bit 0)
  | popCount (n .&. bit 1) == 1 = "double blink" : handshake (n - bit 1)
  | popCount (n .&. bit 2) == 1 = "close your eyes" : handshake (n - bit 2)
  | popCount (n .&. bit 3) == 1 = "jump" : handshake (n - bit 3)
  | otherwise = []
