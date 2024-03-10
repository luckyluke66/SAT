module Dpll(
    dpll,
    Cnf,
    Var,
    Clause
) where

import Data.List

type Cnf = [Clause]
type Var = Int
type Clause = [Int]

findUnit :: Cnf -> Maybe Var
findUnit cnf = 
    let unit = find (\cl -> length cl == 1) cnf
    in case unit of 
        Just [x] -> Just x
        Nothing -> Nothing

chooseVar :: Cnf -> Var
chooseVar cnf = 
    case findUnit cnf of 
        Just x -> x
        Nothing -> head $ head cnf

dpll :: Cnf ->  Bool
dpll []  = True
dpll cnf 
    | [] `elem` cnf = False
    | otherwise = 
        let x = chooseVar cnf 
        in  dpll (simplify x cnf) || dpll (simplify (-x) cnf)
                
simplify :: Var -> Cnf -> Cnf
simplify x cnf = map (delete (-x)) clauses
    where clauses = filter (notElem x) cnf