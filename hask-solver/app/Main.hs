module Main where

import Data.List
import ParseDimacs
import Dpll

main :: IO ()
main = do
    input <- readFile "input.cnf"
    let cnf = parseDimacs input
    print $ dpll cnf

