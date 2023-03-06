# Go-ClientServer
A Go project with communication between clients and a server.\
The data is processed concurrently using go routines.\
The communcation is done via messages between the client and the server and there is a config file with the port and the number of clients.\
In the #3 the client sends an array of integers to the server and the server responds with the inverse numbers array and their sum. Ex:  12,13,14 => 21,31,41 with sum 93\
In the #5 the client sends an array of strings and the server and the server responds by transforming the strings that represent a binary number into base 10, deleting others. Ex: 2dasdas,12,dasdas,1010,101 => 10,3\
In the #12 the client sends an array of natural numbers and the server responds by doubling the first digit of each element and adding them. Ex: 23,43,26,74 => (223+443+226+774) 
