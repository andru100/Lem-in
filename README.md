# Lem-in
Solution to the Lem-in problem where you are given n amount of ants, rooms and connections between rooms/nodes in a txt file. The program uses Breadth First Search to locate all possible routes by shortest first. Then takes into account route length and assigns ants to ques for each route. So that they make it to the end node in shortest time possible with the least steps.
To see the solution type 'go run . ' and add your chosen example.txt eg 'go run . example01.txt'

# Here are the original instructions given for the task


Instructions
You need to create tunnels and rooms.
A room will never start with the letter L or with # and must have no spaces.
You join the rooms together with as many tunnels as you need.
A tunnel joins only two rooms together never more than that.
A room can be linked to an infinite number of rooms and by as many tunnels as deemed necessary.
Each room can only contain one ant at a time (except at ##start and ##end which can contain as many ants as necessary).
To be the first to arrive, ants will need to take the shortest path or paths. They will also need to avoid traffic jams as well as walking all over their fellow ants.
You will only display the ants that moved at each turn, and you can move each ant only once and through a tunnel (the room at the receiving end must be empty).
The rooms names will not necessarily be numbers, and in order.
Any unknown command will be ignored.
The program must handle errors carefully. In no way can it quit in an unexpected manner.
The coordinates of the rooms will always be int.
Your project must be written in Go.
The code must respect the good practices.
It is recommended to have test files for unit testing.
Allowed packages
Only the standard Go packages are allowed.
Usage
Example 1 :

$ go run . test0.txt
3
##start
1 23 3
2 16 7
3 16 3
4 16 5
5 9 3
6 1 5
7 4 8
##end
0 9 5
0-4
0-6
1-3
4-3
5-2
3-5
4-2
2-1
7-6
7-2
7-4
6-5

L1-3 L2-2
L1-4 L2-5 L3-3
L1-0 L2-6 L3-4
L2-0 L3-0
$
