module LinkedList
  ( LinkedList
  , datum
  , fromList
  , isNil
  , new
  , next
  , nil
  , reverseLinkedList
  , toList
  ) where

data LinkedList a
  = EmptyLinkedList
  | LinkedList a
               (LinkedList a)
  deriving (Eq, Show)

datum :: LinkedList a -> a
datum (LinkedList a _) = a

fromList :: [a] -> LinkedList a
fromList = foldr new EmptyLinkedList

isNil :: LinkedList a -> Bool
isNil EmptyLinkedList = True
isNil _               = False

new :: a -> LinkedList a -> LinkedList a
new = LinkedList

next :: LinkedList a -> LinkedList a
next EmptyLinkedList  = EmptyLinkedList
next (LinkedList a n) = n

nil :: LinkedList a
nil = EmptyLinkedList

reverseLinkedList :: LinkedList a -> LinkedList a
reverseLinkedList = fromList . reverse . toList

toList :: LinkedList a -> [a]
toList EmptyLinkedList  = []
toList (LinkedList a n) = a : toList n
