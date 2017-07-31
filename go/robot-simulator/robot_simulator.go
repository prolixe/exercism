package robot

import "fmt"

const testVersion = 3

const (
	N Dir = iota
	E
	S
	W
)

const Terminated = Command(255)

func Advance() {
	switch Step1Robot.Dir {
	case N:
		Step1Robot.Y++
	case W:
		Step1Robot.X--
	case S:
		Step1Robot.Y--
	case E:
		Step1Robot.X++
	}
}

func (d Dir) String() string {
	switch d {
	case N:
		return "North"
	case E:
		return "East"
	case S:
		return "South"
	case W:
		return "West"
	default:
		return ""
	}
}

func Right() {
	Step1Robot.Dir++
	if int(Step1Robot.Dir) > 3 {
		Step1Robot.Dir = N
	}
}
func Left() {
	Step1Robot.Dir--
	if int(Step1Robot.Dir) < 0 {
		Step1Robot.Dir = W
	}
}

type Action byte

func StartRobot(cmd chan Command, act chan Action) {
	for c := range cmd {
		act <- Action(c)
	}
	close(act)
}

func Room(rect Rect, startRobot Step2Robot, act chan Action, rep chan Step2Robot) {

	robot := startRobot
	for a := range act {
		switch a {
		case ' ':
			// Initialize
		case 'R':
			robot.Dir++
			if int(robot.Dir) > 3 {
				robot.Dir = N
			}
		case 'L':
			robot.Dir--
			if int(robot.Dir) < 0 {
				robot.Dir = W
			}
		case 'A':
			newPos := robot.Pos
			newPos = advance(newPos, robot.Dir)
			if isPosInsideRect(newPos, rect) {
				robot.Pos = newPos
			}
		}
	}
	rep <- robot
}

func advance(pos Pos, dir Dir) (newPos Pos) {
	switch dir {
	case N:
		pos.Northing++
	case W:
		pos.Easting--
	case S:
		pos.Northing--
	case E:
		pos.Easting++
	}
	newPos = pos
	return
}

func isPosInsideRect(pos Pos, rect Rect) bool {
	if pos.Easting < rect.Min.Easting ||
		pos.Easting > rect.Max.Easting ||
		pos.Northing < rect.Min.Northing ||
		pos.Northing > rect.Max.Northing {

		return false
	}
	return true
}

type Action3 struct {
	name string
	cmd  Command
}

func StartRobot3(name, scr string, act chan Action3, log chan string) {
	//fmt.Printf("Starting robot %s with list of command %s\n", name, scr)
	if len(name) == 0 {
		log <- "Robot has no name"
	}
	for _, c := range []byte(scr) {
		//fmt.Printf("Sending command with byte %d\n\n", c)
		act <- Action3{name, Command(c)}
	}
	act <- Action3{name, Terminated}

}

func findIndex(name string, robots []Step3Robot) (int, bool) {
	for i, robot := range robots {
		if robot.Name == name {
			return i, true
		}
	}
	return -1, false
}

func isPosOccupiedByRobot(pos Pos, currentRobot string, robots []Step3Robot) bool {
	for _, robot := range robots {
		if robot.Step2Robot.Pos == pos && currentRobot != robot.Name {
			return true
		}
	}
	return false
}

func Room3(rect Rect, robots []Step3Robot, action chan Action3, report chan []Step3Robot, log chan string) {
	var doneCount int

	robotNames := make(map[string]bool)
	robotPositions := make(map[Pos]bool)
	for _, robot := range robots {
		if robotNames[robot.Name] {
			log <- fmt.Sprintf("Duplicate name %s", robot.Name)
		}
		robotNames[robot.Name] = true
		if robotPositions[robot.Step2Robot.Pos] {
			log <- fmt.Sprintf("Same position for robot %s", robot.Name)
		}
		robotPositions[robot.Step2Robot.Pos] = true
		if !isPosInsideRect(robot.Step2Robot.Pos, rect) {
			log <- fmt.Sprintf("Robot %s is not in the rectangle at the start", robot.Name)
		}

	}
Loop:
	for a := range action {
		robotName := a.name
		c := a.cmd
		robotIndex, found := findIndex(robotName, robots)
		if !found {
			report <- robots
			log <- fmt.Sprintf("Could not find the robot %s", robotName)
			break
		}
		robot := robots[robotIndex].Step2Robot
		switch c {
		case ' ':
			// Initialize
		case 'R':
			robot.Dir++
			if int(robot.Dir) > 3 {
				robot.Dir = N
			}
		case 'L':
			robot.Dir--
			if int(robot.Dir) < 0 {
				robot.Dir = W
			}
		case 'A':
			newPos := robot.Pos
			newPos = advance(newPos, robot.Dir)
			if isPosInsideRect(newPos, rect) {
				if !isPosOccupiedByRobot(newPos, robotName, robots) {
					robot.Pos = newPos
				} else {
					log <- fmt.Sprintf("Robot %s has bumped into other robot", robotName)
				}
			} else {
				log <- fmt.Sprintf("Robot %s bumped into wall", robotName)
			}
		case Terminated:
			doneCount++
		default:
			// Unknown command
			log <- fmt.Sprintf("Unknown command %d for robot %s", byte(c), robotName)
			report <- robots
			break Loop
		}
		// Update the robot
		robots[robotIndex].Step2Robot = robot
		if doneCount == len(robots) {
			report <- robots
			break
		}
	}

}
