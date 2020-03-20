Input: {{printf "%q" .}} 
Output 0: {{title .}} 
Output 1: {{title . | printf "%q"}} 
Output 2: {{printf "%q" . | title}} 
