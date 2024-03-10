{-
dpll :: [[Int]] -> Bool
dpll [] = True
dpll ls@(x:xs)
    | [] `elem` ls = False
    | otherwise = dpll xTrue || dpll xFalse
        where 
            xTrue           = removeVars var $ removeClauses var ls
            xFalse          = removeVars (-var) $ removeClauses (-var) ls
            removeVars v    = map (delete v)
            removeClauses v = filter (var `elem`)
            var             = head x
-}

dpll :: [[Int]] -> Bool
dpll [] = True
dpll cnf
    | [] `elem` cnf = False
    | otherwise = let
        unit = findUnit cnf
        pure = findPure cnf
        in case unit of
            Just x -> dpll $ simplify x cnf
            Nothing -> case pure of
                Just x -> dpll $ simplify x cnf
                Nothing -> let
                    (x:xs) = head cnf
                    in dpll (simplify x cnf) || dpll (simplify (-x) cnf)

findUnit :: [[Int]] -> Maybe Int
findUnit cnf = 
    let unit = find (\cl -> length cl == 1) cnf
    in case unit of 
        Just [x] -> Just x
        Nothing -> Nothing