module ParseDimacs(
    parseDimacs
) where

import Dpll(Cnf)

parseDimacs :: String -> Cnf
parseDimacs xs = parseDimacs' $ lines xs

parseDimacs' :: [String] -> Cnf
parseDimacs' xs = map (map read . words) $ removeZeroes $ removeComments xs

removeComments :: [String] -> [String]
removeComments = filter (\(x:xs) -> x /= 'c' && x /= 'p')

removeZeroes :: [String] -> [String]
removeZeroes = map (init . init)