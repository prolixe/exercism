module BST
  ( BST
  , bstLeft
  , bstRight
  , bstValue
  , empty
  , fromList
  , insert
  , singleton
  , toList
  ) where

import           Data.Foldable
import           Data.List     (foldl')

data BST a
  = EmptyTree
  | Node a
         (BST a)
         (BST a)
  deriving (Eq, Show)

instance Ord a => Monoid (BST a) where
  mempty = empty
  mappend x y = fromList (toList x ++ toList y)

instance Foldable BST where
  foldMap f EmptyTree = mempty
  foldMap f (Node x left right) =
    foldMap f left `mappend` f x `mappend` foldMap f right

bstLeft :: BST a -> Maybe (BST a)
bstLeft EmptyTree       = Nothing
bstLeft (Node _ left _) = Just left

bstRight :: BST a -> Maybe (BST a)
bstRight EmptyTree        = Nothing
bstRight (Node _ _ right) = Just right

bstValue :: BST a -> Maybe a
bstValue EmptyTree    = Nothing
bstValue (Node a _ _) = Just a

empty :: BST a
empty = EmptyTree

fromList :: Ord a => [a] -> BST a
fromList = foldl' (flip insert) EmptyTree

insert :: Ord a => a -> BST a -> BST a
insert x EmptyTree = Node x EmptyTree EmptyTree
insert x (Node a left right)
  | x <= a = Node a (insert x left) right
  | x > a = Node a left (insert x right)

singleton :: Ord a => a -> BST a
singleton x = insert x EmptyTree
