# Description
This project is one of the many possible solutions for the implementation of Walker Alias method.

This one in particular was developed using Go language.

# Alias Method

The alias method is a family of efficient algorithms for sampling from a discrete probability distribution, due to A. J. Walker. That is, it returns integer values 1 ≤ i ≤ n according to some arbitrary probability distribution pi. The algorithms typically use O(n log n) or O(n) preprocessing time, after which random values can be drawn from the distribution in O(1) time.

![Alias Method](https://www.researchgate.net/profile/Mattia_Massone/publication/326356402/figure/fig10/AS:647622192799747@1531416571879/Example-of-alias-method.png)

Internally, the algorithm consults two tables, a probability table Ui and an alias table Ki (for 1 ≤ i ≤ n). To generate a random outcome, a fair diсe is rolled to determine an index into the two tables. Based on the probability stored at that index, a biased coin is then flipped, and the outcome of the flip is used to choose between a result of i and Ki.

Source: [Learn more](https://en.wikipedia.org/wiki/Alias_method)

# How to Run
Go to the project directory and run
```console
go run main.go
```
to test the project. If you want to change the probability weights and the number of random generations you can do that in main file (main.go).