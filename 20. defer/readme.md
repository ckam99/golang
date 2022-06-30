
In Go language, defer statements delay the execution of the function or method or an anonymous method until the nearby functions returns. In other words, defer function or method call arguments evaluate instantly, but they don’t execute until the nearby functions returns. You can create a deferred method, or function, or anonymous function by using the defer keyword.


### Important Points:

- In Go language, multiple defer statements are allowed in the same program and they are executed in LIFO(Last-In, First-Out) order as shown in Example 2.
- In the defer statements, the arguments are evaluated when the defer statement is executed, not when it is called.
- Defer statements are generally used to ensure that the files are closed when their need is over, or to close the channel, or to catch the panics in the program.