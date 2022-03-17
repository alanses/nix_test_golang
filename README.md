How to run
go build -height -width -count_iterations -matrix

-height int from 1 to infinity flag for height of matrix
-width int from 1 to infinity flag for height of matrix
-count_iterations int from 1 to infinity flag for count of iterations
-matrix string value types 
- rand mean matrix will create randomly by height, width
- "" empty string will takes mocked value
- [0,1,0][1,1,0][1,0,1] example string, by this struct will create matrix

examples of use

./main -height=3 -width=3 -count_iterations=3 -matrix=[0,1,0][1,1,0][1,0,1] 

./main -height=25 -width=25 -count_iterations=10 -matrix=rand

./main -height=10 -width=10 -count_iterations=10 -matrix=default