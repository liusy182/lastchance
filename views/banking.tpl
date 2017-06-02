<p>banking</p>

{{range .accounts}}
  <p>{{.ID}} {{.Name}} {{.Number}} {{.Amount}}</p>
{{end}}