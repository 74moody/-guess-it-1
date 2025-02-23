
# guess-it-1!

Math skills to guess numbers. 

### Instructions

That is right, you must build a program that given a number as standard input, prints out a range in which the next number provided should be.

The data received by the program, as always, will be presented as the following example:

```console
189
113
121
114
145
110
...
```

This data represents a graph in which the values of the x axis are the number of the lines (0, 1, 2, 3, 4, 5 ...) and the values of the y axis are the actual numbers (189, 113, 121, 114, 145, 110...).

Each of the numbers will be your standard input and the purpose of your program is for you to find the range in which the next number will be in.
This range should have a space separating the lower limit from the upper one like in the example:

```console
>$ ./your_program
189 --> the standard input
120 200    --> the range for the next input, in this case for the number 113
113 --> the standard input
160 230    --> the range for the next input, in this case for the number 121
121 --> the standard input
110 140    --> the range for the next input, in this case for the number 114
114 --> the standard input
100 200    --> the range for the next input, in this case for the number 145
145 --> the standard input
1 99      --> the range for the next input, in this case for the number 110
110 --> the standard input
100 150    --> the range for the next input, in this case for the number
145
...
```



