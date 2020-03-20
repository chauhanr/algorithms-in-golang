{{with "output"}}{{printf "pipeline - %q" .}}{{end}}

{{with $x := "output" | printf "%q"}}{{$x}}{{end}}
