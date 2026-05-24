package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"untisToCalendar/untis"
)

func main() {
	client := &http.Client{}
	uc := untis.NewClient(
		"TorRas",
		"Behördestinkt1",
		"https://gotthard-kuehl-schule.webuntis.com/WebUntis/jsonrpc.do",
		949,
	)

	sess, err := uc.Authenticate(client)
	if err != nil {
		log.Fatalf("authenticate: %v", err)
	}
	fmt.Println("Session:", sess)

	subjects, err := uc.GetSubjects(client, sess)
	if err != nil {
		log.Fatalf("get subjects: %v", err)
	}

	teachers, err := uc.GetTeachers(client, sess)
	if err != nil {
		log.Fatalf("get teachers: %v", err)
	}

	fmt.Println("\nSubjects:")
	for id, s := range subjects {
		name := s.LongName
		if name == "" {
			name = s.Name
		}
		fmt.Printf("%d -> %s\n", id, name)
	}

	fmt.Println("\nTeachers:")
	for id, t := range teachers {
		abbr := t.Abbreviation
		if abbr == "" {
			abbr = t.Name
		}
		fmt.Printf("%d -> %s\n", id, abbr)
	}

	// also print a small timetable sample to correlate IDs
	start := time.Now()
	end := start.AddDate(0, 0, 14)
	lessons, err := uc.GetTimetable(client, sess, start, end)
	if err != nil {
		log.Fatalf("get timetable: %v", err)
	}

	fmt.Println("\nSample lessons (id -> subjectIDs, teacherIDs):")
	for i := 0; i < 10 && i < len(lessons); i++ {
		l := lessons[i]
		fmt.Printf("lesson %d id=%d su=%v te=%v\n", i, l.ID, l.Subjects, l.Teachers)
	}
}
