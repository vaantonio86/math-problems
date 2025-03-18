func GenerateRandomGoCode(n int) []string {
    var results []string
    
    for i := 0; i < n; i++ {
        // Define a random number of variables
        numVars := rand.Intn(10) + 1
        
        // Generate a list of random variable names
        varNames := make([]string, numVars)
        for j := 0; j < numVars; j++ {
            varNames[j] = "x" + strconv.Itoa(rand.Intn(10))
        }
        
        // Define a random number of operations
        numOps := rand.Intn(5) + 1
        
        // Generate a list of random operation strings
        opStrings := make([]string, numOps)
        for j := 0; j < numOps; j++ {
            opStrings[j] = opStr[rand.Intn(len(opStr))]
        }
        
        // Define a random number of parentheses
        numParens := rand.Intn(5) + 1
        
        // Generate a list of random parentheses strings
        parenStrings := make([]string, numParens)
        for j := 0; j < numParens; j++ {
            parenStrings[j] = parenStr[rand.Intn(len(parenStr))]
        }
        
        // Combine the variables, operations, and parentheses into a string
        varExpr := ""
        for _, varName := range varNames {
            varExpr += varName + " "
        }
        for _, opString := range opStrings {
            varExpr += opString + " "
        }
        for _, parenString := range parenStrings {
            varExpr += parenString + " "
        }
        
        // Add the expression to the results slice
        results = append(results, varExpr)
    }
    
    return results
}

// Define a list of operation strings
var opStr = []string{ "+", "-", "*", "/" }

// Define a list of parentheses strings
var parenStr = []string{ "(", ")" }