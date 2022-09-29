package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

var firstName string
var lastName string
var stringDate string = "01/02/2003"
var childName1 string
var childName2 string
var childName3 string
var siblingName1 string
var siblingName2 string
var siblingName3 string
var petName1 string
var petName2 string
var petName3 string
var nickName1 string
var nickName2 string
var nickName3 string
var tFrom string
var fromNumber string
var team string
var movie string
var partnerName string
var partnerSurname string
var partnerNickname string

func main() {
	fmt.Println("Welcome to Gonerator! Please start by instering targets first name")

	// Taking input from user
	fmt.Scanln(&firstName)
	fmt.Println(" Target name saved as " + firstName)
	fmt.Println(" Please enter targets last name")
	fmt.Scanln(&lastName)
	fmt.Println(" Target last name saved as " + lastName)
	fmt.Println(" Please enter targets birth date as dd/mm/yyyy (ex. " + stringDate + ")")
	fmt.Scanln(&stringDate)
	fmt.Println(" Target birth date saved as " + stringDate)
	children := askYesNo("Does your target have any child?")
	if children {

		fmt.Println("Please enter child/ren name (Max 3) (Please seperate children name with space)")
		fmt.Scanln(&childName1, &childName2, &childName3)
		fmt.Println(" Target child/ren saved as " + childName1 + " " + childName2 + " " + childName3)
	}
	siblings := askYesNo("Does your target have any sibling?")
	if siblings {

		fmt.Println("Please enter sibling/s name (Max 3) (Please seperate siblings name with space)")
		fmt.Scanln(&siblingName1, &siblingName2, &siblingName3)
		fmt.Println("Target sibling/s saved as " + siblingName1 + " " + siblingName2 + " " + siblingName3)

	}
	pets := askYesNo("Does your target have any pets")
	if pets {

		fmt.Println("Please enter pet/s name (Max 3) ( Please seperate siblings name with space )")
		fmt.Scanln(&petName1, &petName2, &petName3)
		fmt.Println("Target pet/s saved as " + petName1 + " " + petName2 + " " + petName3)
	}
	nickName := askYesNo("Does your target have any nicknames?")
	if nickName {

		fmt.Println("Please enter nickname/s name (Max 3) ( Please seperate siblings name with space )")
		fmt.Scanln(&nickName1, &nickName2, &nickName3)
		fmt.Println("Target nickname/s saved as " + nickName1 + " " + nickName2 + " " + nickName3)

	}
	fmt.Println("Where is your target from (Hometown or Country)")
	fmt.Scanln(&tFrom)
	fmt.Println("Target hometown/country saved as " + tFrom)
	fmt.Println("Is there a known number associated with this hometown/country (ex. Istanbul is known 34 because of plate numbers")
	fmt.Scanln(&fromNumber)
	fmt.Println("Associated number saved as " + fromNumber)
	fmt.Println("What is your targets favorite sports team?")
	fmt.Scanln(&team)
	fmt.Println("Target favorite team saved as " + team)
	fmt.Println("What is your target's favorite movie or TV show?")
	fmt.Scanln(&movie)
	fmt.Println("Target favorite movie/show saved as " + movie)
	partner := askYesNo("Does your target have a partner?")
	if partner {

		fmt.Println("Please enter target's partner name (Please enter as Name LastName) (ex. John Doe)")
		fmt.Scanln(&partnerName, &partnerSurname)
		fmt.Println("Target partner saved as " + partnerName + " " + partnerSurname)
		fmt.Println("Please enter target's partners nickname (If there isn't any leave empty)")
		fmt.Scanln(&partnerNickname)
	} else {
		fmt.Println("LOL")
	}

	fmt.Println("Information gathering finished. Your wordlist is getting ready!")
	writeTXT()

}

func askYesNo(s string) bool {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("%s [y/n]: ", s)

		response, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		response = strings.ToLower(strings.TrimSpace(response))

		if response == "y" || response == "yes" {
			return true
		} else if response == "n" || response == "no" {
			return false
		}
	}
}

func writeTXT() {

	res1 := strings.ToLower(firstName) + string(stringDate[len(stringDate)-2:])
	res2 := strings.ToLower(firstName) + string(stringDate[len(stringDate)-4:])
	res3 := strings.ToLower(lastName) + string(stringDate[len(stringDate)-2:])
	res4 := strings.ToLower(lastName) + string(stringDate[len(stringDate)-4:])
	/* nickCount := len([]rune(nickName1)) */
	res5 := strings.ToLower(nickName1) + string(stringDate[len(stringDate)-2:])
	res6 := strings.ToLower(nickName1) + string(stringDate[len(stringDate)-4:])
	res7 := strings.ToLower(nickName2) + string(stringDate[len(stringDate)-2:])
	res8 := strings.ToLower(nickName2) + string(stringDate[len(stringDate)-4:])
	res9 := strings.ToLower(nickName3) + string(stringDate[len(stringDate)-2:])
	res10 := strings.ToLower(nickName3) + string(stringDate[len(stringDate)-4:])
	res11 := strings.ToLower(petName1) + string(stringDate[len(stringDate)-2:])
	res12 := strings.ToLower(petName1) + string(stringDate[len(stringDate)-4:])
	res13 := strings.ToLower(petName2) + string(stringDate[len(stringDate)-2:])
	res14 := strings.ToLower(petName2) + string(stringDate[len(stringDate)-4:])
	res15 := strings.ToLower(petName3) + string(stringDate[len(stringDate)-2:])
	res16 := strings.ToLower(petName3) + string(stringDate[len(stringDate)-4:])

	currentTime := time.Now()

	f, err := os.Create(firstName + "-" + lastName + "-" + currentTime.Format("01-02-2006") + ".txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	_, err2 := f.WriteString(
		res1 + "\n" +
			res2 + "\n" +
			res3 + "\n" +
			res4 + "\n" +
			res5 + "\n" +
			res6 + "\n" +
			res7 + "\n" +
			res8 + "\n" +
			res9 + "\n" +
			res10 + "\n" +
			res11 + "\n" +
			res12 + "\n" +
			res13 + "\n" +
			res14 + "\n" +
			res15 + "\n" +
			res16 + "\n")

	if err2 != nil {
		log.Fatal(err2)
	}

	fmt.Println("Inputs are succesfully saved!")
}

/* func readFile() {
	currentTime := time.Now()

	// Open file for reading
    file, err := os.Open(firstName + "-" + lastName + "-" + currentTime.Format("01-02-2006") + ".txt")
    if err != nil {
        log.Fatal(err)
    }

    // os.File.Read(), io.ReadFull(), and
    // io.ReadAtLeast() all work with a fixed
    // byte slice that you make before you read

    // ioutil.ReadAll() will read every byte
    // from the reader (in this case a file),
    // and return a slice of unknown slice
    data, err := ioutil.ReadAll(file)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Data as hex: %x\n", data)
    fmt.Printf("Data as string: %s\n", data)
    fmt.Println("Number of bytes read:", len(data))
} */
