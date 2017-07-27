package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/tokiw/golang-practice/ch07/ex13"
)

var calc = template.Must(template.New("calc").Parse(`
<!DOCTYPE html>
<html>
<head>
	<title>Calc</title>
</head>
<body>
	<form method="POST">
		<input type="text" name="expr" value="{{.Expr}}" />
		<input type="submit" value="calc" />
	</form>
	<div>{{.Result}}</div>
</body>
</html>
`))

type exprAndResult struct {
	Expr   string
	Result string
}

func main() {
	env := make(eval.Env)
	http.HandleFunc("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		exprStr := r.FormValue("expr")

		if exprStr == "" {
			calc.Execute(w, exprAndResult{"", ""})
			return
		}
		expr, err := eval.Parse(exprStr)

		var result string
		if err != nil {
			result = fmt.Sprintf("%v", err)
		} else {
			result = fmt.Sprint(expr.Eval(env))
		}

		calc.Execute(w, exprAndResult{exprStr, result})
	}))

	sockAddress := "localhost:8000"
	fmt.Println("Serving on " + sockAddress)
	log.Fatal(http.ListenAndServe(sockAddress, nil))
}
