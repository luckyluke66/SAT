package main

import (
	"fmt"
	"os"
	"strings"
    "strconv"
    "slices"
    "time"
)

type Cnf []Clause
type Clause []Var
type Var int


func parseDimacs(xs string) Cnf {
    lines := strings.Split(xs, "\n")
    return RemoveZeroes(RemoveComments(lines))
}

func RemoveComments(xs []string) []string {
    var result []string
    for _, x := range xs {
        if len(x) > 0 && x[0] != 'c' && x[0] != 'p' {
            result = append(result, x)
        }
    }
    return result
}

func RemoveZeroes(xs []string) Cnf {
    var result Cnf
    for _, x := range xs {
        var clause []Var
        for _, word := range strings.Fields(x) {
            num, err := strconv.Atoi(word) 
            if err == nil {
                clause = append(clause, Var(num))
            }
        }
        if len(clause) > 0 {
            result = append(result, clause[:len(clause)-1])
        }
    }
    return result
}

func getCnf(filename string) (Cnf, error) {
    file, err := os.ReadFile(filename)
    
    if err != nil {
        return nil, err
    }

    return parseDimacs(string(file)), nil
}


func dpll(cnf Cnf, literals []Var, unitCount int, nodes int) ([]Var, bool, int, int) {
    if containsEmpty(cnf) {
        return literals, false, unitCount, nodes
    }

    if len(cnf) == 0 {
        return literals, true, unitCount, nodes
    }

    x := findUnit(cnf)
    if x == 0 {
        x = cnf[0][0]  
        unitCount++
    }

    if newLiterals, ok, unitCount, nodes:= dpll(simplify(x, cnf), append(literals, x), unitCount, nodes); ok {
        nodes++
        return newLiterals, true, unitCount, nodes
    } else {
        nodes++
        return dpll(simplify(-x, cnf), append(literals, -x), unitCount, nodes)
    }
}
func Dpll(cnf Cnf) ([]Var, bool, int, int) {
    return dpll(cnf, []Var{}, 0, 0)
}

func pickVar(formula Cnf) Var {
    if findUnit(formula) != 0 {
        return findUnit(formula)
    } else {
        return formula[0][0]
    }
}

func findUnit(formula Cnf) Var {
    for _, clause := range formula {
        if len(clause) == 1 {
            return clause[0]
        }
    }
    return 0
}

func containsEmpty(formula Cnf, ) bool {
    for _, clause := range formula {
        if len(clause) == 0 {
            return true
        }
    }
    return false
}

func simplify(x Var, cnf Cnf) Cnf {
    var result Cnf
    for _, clause := range cnf {
        if !contains(clause, x) {
            result = append(result, deleteVar(-x, clause))
        }
    }
    return result
}

func deleteVar(x Var, clause []Var) []Var {
    var result []Var
    for _, v := range clause {
        if v != x {
            result = append(result, v)
        }
    }
    return result
}

func contains(slice []Var, item Var) bool {
    for _, a := range slice {
        if a == item {
            return true
        }
    }
    return false
}

func printData(vars []int, sat bool, initTime time.Duration, time time.Duration, unit int, treeNodes int) {
    if sat {
        fmt.Println("SAT")
        fmt.Println(vars)
    } else {
        fmt.Println("UNSAT")
        fmt.Println()
    }
    fmt.Println(initTime)
    fmt.Println(time)
    fmt.Println(unit)
    fmt.Println(treeNodes)
}


func main() {
    args := os.Args
    filepath := args[1]

    startParse := time.Now()
    cnf, err := getCnf(filepath)
    parseTime := time.Since(startParse)

    if err != nil {
        fmt.Println(err)
        return
    }

    startDpll := time.Now()
    vars, sat, unitCount, nodes := Dpll(cnf)
    intvars := make([]int, len(vars))
    dpllTime := time.Since(startDpll)
    
    for i, v := range vars {
        intvars[i] = int(v)
    }

    slices.Sort(intvars)
    printData(intvars, sat, parseTime, dpllTime, unitCount, nodes)
}