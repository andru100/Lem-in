package main 

import (
	"os"
	"io/ioutil"
	"strings"
	"fmt"
	"strconv"
)

var mapnodes = map[string][]string {}
var noRoutesLeft = false
var routefound []string // for adding found route to
var nodeInRoute [] string // used to add nodes used in routes found, so they can be ignore when finding new route
var arrayOfRoutes [][] string
var arrayOfmapRooms [] map[int][]string
var routeswitRooms = map[int][]string{}
var startNode string
var endNode string
var shortestPath int
var numOfAnts int

func wordEnd (a string) int{
	for i, car := range a { // find first space in next line
		if car == 32 {// space found
			return i
		}
	}
	return -1
}

func dashFind (a string) int{
	for i, car := range a { // find first - in next line
		if car == 45 {// "-" found
			return i
		}
	}
	return -1
}

func main () {

	args := os.Args[1] // Take arg from cmd line

	fileIO, err := os.OpenFile(args, os.O_RDWR, 061010)// open file named in args
	if err != nil {
		panic(err)
	}
	defer fileIO.Close()
	rawBytes, err := ioutil.ReadAll(fileIO)// read file
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(rawBytes), "\n") // split text by lines
	
	numOfAnts, err = strconv.Atoi(string(lines[0]))

	for i, str := range lines { 

		if !(strings.Contains (str, "-")) { // check for not a link
			if len(str) >0 && str[0] == 35 { // get start and end nodes chk for #
				if str[2] == 115 { // if line has an "s", so start node
						
					index := wordEnd(lines[i+1])// use word end func to get index to cut string to
					startNode = string(lines[i+1][0:index]) // cut word from beginning to space found
					
				} else if str[2] == 101 { // if line has "e" so end node
					index := wordEnd(lines[i+1])// use word end func to get index to cut to
					endNode = string(lines[i+1][0:index])// get end node from next line char 1
				}
			} else if i > 0{ //standard nodes check first line with num 
				index := wordEnd(str)// use word end func get index to cut to
				if index != -1 { // check for no space found
					mapnodes[string(str)[:index]] = []string{} // add node to map of nodes with empty peers list
				}
			}
		} else { // if its a link message add it to peers
			index := dashFind(str)
			if index != -1 { // check for no space found
				mapnodes[string(str)[:index]] = append(mapnodes[string(str)[:index]], string(str)[index+1:]) 
				mapnodes[string(str)[index+1:]] = append(mapnodes[string(str)[index+1:]], string(str)[:index]) 
			}
		}
	}

	// while loop here with breadth first Search finding routes
	for noRoutesLeft == false {
		a, routesNrooms := bfSearch(mapnodes, startNode, endNode) // run search with start and end nodes
			routesNrooms[1010] = []string{} // add que 
			if noRoutesLeft == false { // check again cos last one returned will be empty so cant index and add to array sort this!! must be better way
				arrayOfRoutes = append(arrayOfRoutes, a) // add found route to list 
				arrayOfmapRooms = append(arrayOfmapRooms, routesNrooms)
				for _, i := range  a[1:len(a)] { // loop through found route ignore start and node and append to node used list to be ignored
						nodeInRoute = append (nodeInRoute, i)
				}
				routefound = []string{} // reset for next route search
				routeswitRooms = map[int][]string{}// reset
			}
	}

	startPath := 0
	neighbour := 0
	// add ants to qoues
	for i:= 1; i<=numOfAnts; i++ { // for each ant
		// find shortest path taking into account ant que length and path length
			CurrentpathTime := len(arrayOfmapRooms[startPath][1010]) + len(arrayOfmapRooms[startPath]) // add route len 2 ants in que
			if startPath == len(arrayOfmapRooms) -1 {
				neighbour = 0
			} else {
				neighbour = startPath +1
			}
			nextPathTime :=  len(arrayOfmapRooms[neighbour][1010]) + len(arrayOfmapRooms[neighbour]) 
			if CurrentpathTime <= nextPathTime { // in case last que is full go to the next one to fill
				arrayOfmapRooms[startPath][1010] = append(arrayOfmapRooms[startPath][1010], "L"+fmt.Sprint(i)+ "-") // put ants in que of shortest path
			} else {
				arrayOfmapRooms[neighbour][1010] = append(arrayOfmapRooms[neighbour][1010], "L"+fmt.Sprint(i)+ "-") // put ants in que of shortest path
				startPath = neighbour//change start path
			}
		
	}
	runtherace(arrayOfmapRooms)// print final array with all routes found
	
}

func runtherace (mapRooms[] map[int][]string){ // send the ants through rooms
	finished := []string{}
	// create que for each route 
	// get lens then find one to send ant to
	// if ant is in room and next room is empty / "." then move to next room make ur room .
	// here handle if lem is at end and next move is the goal node move there and replace ur room with "." / empty!!
	for len(finished) < numOfAnts { // for loop until finish line has all ants
		for _, 	routeInQst := range mapRooms { // loop through different routes with route in question
			if routeInQst[len(routeInQst)-1][1] != "." { // if last box is ocupied pop to finished
				finished = append(finished, routeInQst[len(routeInQst)-1][1])
				routeInQst[len(routeInQst)-1][1] = "." // make it empty
			}
			
			for i:= len(routeInQst)-1; i > 1; i-- { // go backwards from last node 
				if routeInQst[i][1] == "." && routeInQst[i-1][1] != "." { // if room in is empty
					// move to next room and replace ur room with "." 
					index := dashFind(routeInQst[i-1][1])
					if index != -1{
						routeInQst[i][1] = routeInQst[i-1][1][:index+1]  + routeInQst[i][0]// cut the "l(ant num)-" and ad new room to the cut string
						fmt.Printf(routeInQst[i][1]+" ")// print the move to new room
						routeInQst[i-1][1] = "." // set last room to empty
					}
				}
			}

			if len(routeInQst[1010]) >0 && routeInQst[1][1] == "." { // if theres ant in que and 1st box empty
				routeInQst[1][1] = routeInQst[1010][0]  + routeInQst[1][0]// cut the "l(ant num)-" and ad new room to the cut string
				fmt.Printf(routeInQst[1][1]+" ")
				routeInQst[1010] = routeInQst[1010][1:] // cut from que
			}

		} 
		fmt.Printf("\n")
	}
}

// Breadth First Search
func bfSearch (mapd map[string][]string, node string, goal string) ([]string, map[int][]string) {
  visited := []string{}
  popped := [][]string{}
  que := [][]string{}
  que = append(que, []string{node, "None"})
  visited = append(visited, node)
  for len(que) > 0 {
    //pop first in que to m var add topopped for trace back and check if found gaol destination found and break loop
	m := que[0] // pop from que
	que = que[1:]
	popped = append(popped, m) // add popped node to popped list
    if m[0] == goal{
      s, routesNrooms := bfspathtraceback(popped, popped[len(popped)-1]) // send to trace back func with the popped node and its mother so we can find the route that got us to goal node
      for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 { // reverse the route 
        s[i], s[j] = s[j], s[i]
      } 
	  return s, routesNrooms

	} else {
      for _, peers := range mapd[m[0]]{ // loop through poppeds peers
        if xin(peers, visited) == false && xin(peers, nodeInRoute) == false{ // check if node has been visited or used already in other path found
          que = append(que, []string{peers, m[0]})
          visited = append(visited, peers)
		}
	  }
	}
  }

  if len(que) == 0 {
	// for when uve extracted all poss routes and cant find another
	noRoutesLeft = true
  }

  return routefound, routeswitRooms // should never get here if does will be empty []
}

// Takes array with nodes used in breadth first search and trace back to discover route taken to find node
func bfspathtraceback (poppedlist [][]string, chased []string) ([]string, map[int][]string) { // trace back the path found
  if chased[1] == "None" {
	for i:=len(routefound[1:])-1; i>=0; i-- { 
		routeswitRooms[len(routeswitRooms)+1] = []string{routefound[i], "."}  // create list of route but with box for present or empty .
	}
	routeswitRooms[len(routeswitRooms)+1] = []string{endNode, "."}  // create list of route with box for present or empty (.)
	return routefound, routeswitRooms
  } else {
    for i := len(poppedlist)-1; i>=0; i-- {
      if poppedlist[i][0] == chased[1]{
		routefound = append (routefound, poppedlist[i][0])
		bfspathtraceback(poppedlist, poppedlist[i]) 
	  }
	} 
  }
  return routefound, routeswitRooms 
}

func xin(a string, list []string) bool { // finc to find if x is in list
    for _, b := range list {
        if b == a {
            return true
        }
    }
    return false
}