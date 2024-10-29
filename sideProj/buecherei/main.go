package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type Buch struct {
    Titel       string `json:"titel"`
    Autor       string `json:"autor"`
    Genre       string `json:"genre"`
    Isbn        uint64 `json:"isbn"`
    Regal       string `json:"regal"`
    Ausgeliehen bool   `json:"ausgeliehen"`
}

type Verleih struct {
    Buch           *Buch     `json:"buch"`
    Kunde          *Kunde    `json:"kunde"`
    AusleihDatum   string    `json:"ausleihDatum"`
    RueckgabeDatum string    `json:"rueckgabeDatum"`
}


type Person struct {
	Vorname  string
	Nachname string
	Adresse  string
	Telefon  string
}

type Mitarbeiter struct {
	Person
	MitarbeiterNummer string
}

type Kunde struct {
	Person
	KundeNummer string
}

type Regal struct {
	Bezeichnung string
	Buecher     []*Buch
}

type Bibliothek struct {
	Buecher      []*Buch
	RegalListe   []*Regal
	VerleihListe []*Verleih
}

func (r *Regal) AddBuch(b *Buch) {
	if len(r.Buecher) < 4 {
		r.Buecher = append(r.Buecher, b)
	} else {
		fmt.Println("Das Regal is voll!")
	}
}

func (r *Regal) RemoveBuch(buch *Buch) {
	for i, b := range r.Buecher {
		if b == buch {
			r.Buecher = append(r.Buecher[:i], r.Buecher[i+1:]...)
			break
		}
	}
}

func (bib *Bibliothek) AddBuch(b *Buch) {
	bib.Buecher = append(bib.Buecher, b)
}

func (bib *Bibliothek) AddRegal(r *Regal) {
	bib.RegalListe = append(bib.RegalListe, r)
}

func (bib *Bibliothek) VerleihBuch(b *Buch, k *Kunde) {
	if !b.Ausgeliehen {
		ausleihDatum := time.Now()
		rueckgabeDatum := ausleihDatum.AddDate(0, 0, 14)
		bib.VerleihListe = append(bib.VerleihListe, &Verleih{b, k, ausleihDatum.Format("02-01-2006"), rueckgabeDatum.Format("02-01-2006")})
		b.Ausgeliehen = true
		fmt.Println("Buch erfolgreich ausgeliehen. Rückgabedatum: " + rueckgabeDatum.Format("02-01-2006"))
	} else {
		fmt.Println("Buch ist bereits ausgeliehen.")
	}
}

func (bib *Bibliothek) ReturnBuch(b *Buch)  {
	for i, verleih := range bib.VerleihListe {
		if (verleih.Buch == b) {
			bib.VerleihListe = append(bib.VerleihListe[:i], bib.VerleihListe[i+1:]...)
			b.Ausgeliehen = false
			fmt.Println("Buch erfolgreich zurückgegeben.")
			return
		}
	}
	fmt.Println("Buch nicht gefunden oder nicht ausgeliehen.") 
}

func readInput(prompt string, reader *bufio.Reader) string {
    fmt.Print(prompt)
    input, _ := reader.ReadString('\n')
    return strings.TrimSpace(input)
}

func saveToFile(bib *Bibliothek, filename string) {
    data, err := json.MarshalIndent(bib, "", "  ")
    if err != nil {
        log.Fatalf("Fehler beim Serialisieren: %v", err)
    }

    err = os.WriteFile(filename, data, 0644)
    if err != nil {
        log.Fatalf("Fehler beim Schreiben in die Datei: %v", err)
    }
}

func loadFromFile(filename string) (*Bibliothek, error) {
    data, err := os.ReadFile(filename)
    if err != nil {
        return nil, fmt.Errorf("Fehler beim Lesen der Datei: %v", err)
    }

    var bib Bibliothek
    err = json.Unmarshal(data, &bib)
    if err != nil {
        return nil, fmt.Errorf("Fehler beim Deserialisieren: %v", err)
    }

    fmt.Println("Daten erfolgreich geladen.")
    return &bib, nil
}

func main() {
	var bibliothek *Bibliothek
    filename := "bibliothek.json"

	loadedBib, err := loadFromFile(filename)
    if err != nil {
        fmt.Println("Leere Bibliothek wird gestartet.")
        bibliothek = &Bibliothek{}
    } else {
        bibliothek = loadedBib
    }

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\nMenü:")
		fmt.Println("1. Regal hinzufügen")
		fmt.Println("2. Buch hinzufügen");
		fmt.Println("3. Buch ausleihen");
		fmt.Println("4. Buch zurückgeben");
		fmt.Println("5. Alle Bücher anzeigen");
		fmt.Println("6. Beenden");
		var choice int
		fmt.Print("\nWählen Sie eine Option (1-6): ");
		fmt.Scanln(&choice)	
		switch choice {
			case 1:
				addRegal(reader, bibliothek)
				break
			case 2:
				addBook(reader, bibliothek)
				break
			case 3:
				rentBook(reader, bibliothek)
				break
			case 4: 
				returnBook(reader, bibliothek)
				break
			case 5:
				viewBooks(bibliothek)
				break
			case 6:
				fmt.Println("Programm beendet.")
				os.Exit(0)
			default:
				fmt.Println("Fehler: Bitte geben Sie eine gültige Zahl ein.")
				os.Exit(1)
		}
		saveToFile(bibliothek, filename)
	}
}

func addRegal(reader *bufio.Reader, bib *Bibliothek) *Regal {
	bezeichnung := readInput("Bezeichnung des Regals eingeben: ", reader)

    for _, r := range bib.RegalListe {
        if strings.EqualFold(r.Bezeichnung, bezeichnung) {
            fmt.Println("Fehler: Ein Regal mit diesem Namen existiert bereits.")
            return nil
        }
    }

    regal := Regal{Bezeichnung: bezeichnung}
    bib.AddRegal(&regal)
    fmt.Println("Regal erfolgreich hinzugefügt.")
    return &regal
}

func addBook(reader *bufio.Reader, bib *Bibliothek) {
	titel := readInput("Titel des Buches eingeben: ", reader)
    autor := readInput("Autor des Buches eingeben: ", reader)
    genre := readInput("Genre des Buches eingeben: ", reader)
	isbnInp := readInput("ISBN des Buches eingeben: ", reader)

	for _, buch := range bib.Buecher {
		if strings.EqualFold(titel, buch.Titel) {
			fmt.Println("Dieses Buch ist bereits in der Bibliothek.")
			return
		}
	}
	
	isbnInp = strings.ReplaceAll(isbnInp, "-", "")
	isbn, err := strconv.ParseUint(isbnInp, 10, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "buecherei: %v", err)
		os.Exit(1)
	}

	regalName := readInput("\nWählen Sie ein Regal für das Buch (geben Sie den Namen ein): ", reader)

	var regal *Regal
	for _, r := range bib.RegalListe {
		if strings.EqualFold(strings.TrimSpace(r.Bezeichnung), regalName) {
			regal = r
			break
		}
	}

	if regal != nil {
		buch := Buch{titel, autor, genre, isbn, regal.Bezeichnung, false}
		regal.AddBuch(&buch)
		bib.AddBuch(&buch)
		fmt.Println("Buch erfolgreich hinzugefügt.")
	} else {
		antwort := readInput("Regal nicht gefunden. Möchten Sie ein neues Regal hinzufügen? (ja/nein): ", reader)
		if strings.EqualFold(antwort, "ja") {
			neuesRegal := addRegal(reader, bib)
			buch := Buch{titel, autor, genre, isbn, neuesRegal.Bezeichnung, false}
			neuesRegal.AddBuch(&buch)
			bib.AddBuch(&buch)
			fmt.Println("Buch erfolgreich hinzugefügt.")
		} else {
			fmt.Println("Buch konnte nicht hinzugefügt werden.")
		}
	}
}


func rentBook(reader *bufio.Reader, bib *Bibliothek) {
	titel := readInput("Titel des Buches eingeben, das Sie ausleihen möchten: ", reader)
	vorname := readInput("Vorname des Kunden: ", reader)
	nachname := readInput("Nachname des Kunden: ", reader)

	kunde := Kunde{Person{vorname, nachname, "", ""}, "K_001"}

	for _, buch := range bib.Buecher {
		if strings.EqualFold(buch.Titel, titel) {
			bib.VerleihBuch(buch, &kunde)
			return
		}
	}
	
	fmt.Println("Buch nicht gefunden.")
}

func returnBook(reader *bufio.Reader, bib *Bibliothek) {
	titel := readInput("Titel des Buches eingeben, das Sie zurückgeben möchten: ", reader)

	for _, buch := range bib.Buecher {
		if strings.EqualFold(buch.Titel, titel) && buch.Ausgeliehen {
			bib.ReturnBuch(buch)
			return
		}
	}

	fmt.Println("Buch nicht gefunden oder nicht ausgeliehen.")
}

func viewBooks(bib *Bibliothek)  {
	if len(bib.Buecher) != 0 {
		fmt.Println("Alle Bücher in der Bibliothek:")
		for _, buch := range bib.Buecher {
			status := "Verfügbar"
			if buch.Ausgeliehen {
				rueckgabeDatum := ""
				for _, verleih := range bib.VerleihListe {
					if buch == verleih.Buch {
						rueckgabeDatum = verleih.RueckgabeDatum
						break
					}
				}
				status = "Ausgeliehen, muss bis zum " + rueckgabeDatum + " zurückgegeben werden"
			}

			fmt.Printf("%s - %s\n", buch.Titel, status)
		}
	} else {
		fmt.Println("Die Bibliothek ist leer.")
	}
}