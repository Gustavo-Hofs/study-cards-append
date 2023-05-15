# About
I decided to try to solve the challenge presented in https://www.youtube.com/live/PP8NazoBjHU?feature=share to train my knowledge about Golang. My solution (which is in card_append/) uses Golang's "append" function to perform Enqueue and Dequeue operations on a slice.

So I wanted to compare the performance of my solution with the solutions presented by Alexandre Cabral (@[github/o-mago](https://github.com/o-mago)). For that, I modified the https://github.com/o-mago/study codes to not use the "tools" package and use Golang's own benchmark.

## Run benchmark
```
cd cards_append
go test -bench=. -benchmem
cd ..
cd github.com-o-mago-study_modificado
go test main.go main_test.go -bench=. -benchmem
```

## Results
o-mago/study:
![image](https://github.com/Gustavo-Hofs/study-cards-append/assets/54177943/0b542122-449a-4ff8-9bc9-8fcba7acc1c7)
append:
![image](https://github.com/Gustavo-Hofs/study-cards-append/assets/54177943/2acf167a-c167-4751-8de5-62039a50af54)

## Conclusion
Golang's append function seems to be faster and use less mallocs.
