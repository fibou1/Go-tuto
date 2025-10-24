package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Question struct {
	Prompt  string
	Choices []string
	Answer  int
}

type Category struct {
	Name      string
	Questions []Question
}

type ScoreEntry struct {
	Player    string    `json:"player"`
	Category  string    `json:"category"`
	Score     int       `json:"score"`
	Total     int       `json:"total"`
	Title     string    `json:"title"`
	Timestamp time.Time `json:"timestamp"`
}

const scoreboardFile = "scores.json"

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("============================")
	fmt.Println("        GoQuiz v1.0")
	fmt.Println("Apprends et révise en t'amusant !")
	fmt.Println("============================\n")

	fmt.Print("Quel est ton prénom ou pseudo ? ")
	player, _ := reader.ReadString('\n')
	player = strings.TrimSpace(player)
	if player == "" {
		player = "Joueur Mystère"
	}

	categories := loadCategories()

	fmt.Println("Choisis un thème :")
	for i, category := range categories {
		fmt.Printf("%d. %s\n", i+1, category.Name)
	}
	fmt.Println("Tape 'help' pour obtenir de l'aide.")

	category, err := selectCategory(reader, categories)
	if err != nil {
		fmt.Println("Aucune catégorie sélectionnée, fin du jeu. À bientôt !")
		return
	}

	score := askQuestions(reader, category)
	title := rankTitle(score, len(category.Questions))

	fmt.Printf("\nBravo %s ! Tu as %d/%d bonnes réponses — %s !\n", player, score, len(category.Questions), title)

	if err := saveScore(ScoreEntry{
		Player:    player,
		Category:  category.Name,
		Score:     score,
		Total:     len(category.Questions),
		Title:     title,
		Timestamp: time.Now(),
	}); err != nil {
		fmt.Printf("Impossible d'enregistrer le score : %v\n", err)
	} else {
		fmt.Println("Ton score a été enregistré dans le tableau des meilleurs scores !")
	}

	if err := displayBestScores(); err != nil {
		fmt.Printf("\nImpossible d'afficher les meilleurs scores : %v\n", err)
	}
}

func loadCategories() []Category {
	return []Category{
		{
			Name: "Informatique générale",
			Questions: []Question{
				{
					Prompt:  "Quel système d'exploitation est open source ?",
					Choices: []string{"Windows", "macOS", "Linux", "ChromeOS"},
					Answer:  2,
				},
				{
					Prompt:  "Quel est le protocole utilisé pour récupérer des pages web ?",
					Choices: []string{"FTP", "HTTP", "SMTP", "SSH"},
					Answer:  1,
				},
				{
					Prompt:  "Quelle structure de données fonctionne en FIFO ?",
					Choices: []string{"Pile", "File", "Arbre", "Table de hachage"},
					Answer:  1,
				},
				{
					Prompt:  "Quel algorithme est utilisé pour chiffrer les mots de passe ?",
					Choices: []string{"SHA-256", "Bubble Sort", "Dijkstra", "Quick Sort"},
					Answer:  0,
				},
			},
		},
		{
			Name: "Cybersécurité",
			Questions: []Question{
				{
					Prompt:  "Quel est le but d'un pare-feu ?",
					Choices: []string{"Sauvegarder les données", "Empêcher l'accès non autorisé", "Accélérer le réseau", "Analyser les virus"},
					Answer:  1,
				},
				{
					Prompt:  "Quelle est la meilleure pratique pour créer un mot de passe ?",
					Choices: []string{"Utiliser uniquement des chiffres", "Utiliser son prénom", "Utiliser une phrase secrète longue", "Réutiliser le même mot de passe"},
					Answer:  2,
				},
				{
					Prompt:  "Quel type d'attaque consiste à submerger un service de requêtes ?",
					Choices: []string{"Phishing", "Man-in-the-middle", "DoS/DDoS", "Injection SQL"},
					Answer:  2,
				},
				{
					Prompt:  "Quel outil détecte les logiciels malveillants ?",
					Choices: []string{"Proxy", "Antivirus", "Routeur", "Switch"},
					Answer:  1,
				},
			},
		},
		{
			Name: "Data & IA",
			Questions: []Question{
				{
					Prompt:  "Quel langage est le plus utilisé pour l'analyse de données ?",
					Choices: []string{"Go", "R", "HTML", "CSS"},
					Answer:  1,
				},
				{
					Prompt:  "Quel type de base de données stocke des documents JSON ?",
					Choices: []string{"SQL", "Graph", "Clé-valeur", "NoSQL"},
					Answer:  3,
				},
				{
					Prompt:  "Quelle étape consiste à nettoyer les données ?",
					Choices: []string{"Data Cleaning", "Data Mining", "Data Warehousing", "Data Streaming"},
					Answer:  0,
				},
				{
					Prompt:  "Quel algorithme est typiquement utilisé pour la classification supervisée ?",
					Choices: []string{"K-means", "Réseau de neurones", "Apriori", "PageRank"},
					Answer:  1,
				},
			},
		},
	}
}

func selectCategory(reader *bufio.Reader, categories []Category) (Category, error) {
	for attempts := 0; attempts < 3; attempts++ {
		fmt.Print("\nEntre le numéro de la catégorie : ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(strings.ToLower(input))

		switch input {
		case "help":
			fmt.Println("Entrer le numéro du thème (1, 2 ou 3). Bonus : essaye 'sudo start quiz' 😉")
			continue
		case "sudo start quiz":
			fmt.Println("Haha ! Même avec sudo, on ne contourne pas les menus 🤖")
			continue
		}

		index, err := parseCategoryIndex(input, len(categories))
		if err != nil {
			fmt.Printf("Commande inconnue : %s. Astuce geek : %s\n", input, geekJoke())
			continue
		}

		return categories[index], nil
	}

	return Category{}, errors.New("aucune catégorie valide choisie")
}

func parseCategoryIndex(input string, total int) (int, error) {
	index, err := strconv.Atoi(input)
	if err != nil {
		return -1, errors.New("invalid index")
	}

	index--
	if index < 0 || index >= total {
		return -1, errors.New("invalid index")
	}

	return index, nil
}

func geekJoke() string {
	jokes := []string{
		"Il y a 10 types de personnes : celles qui comprennent le binaire et les autres.",
		"Pourquoi les programmeurs confondent Halloween et Noël ? Parce que OCT 31 == DEC 25.",
		"Un paquet TCP entre dans un bar et demande : 'Je peux garder vos numéros ?'",
	}
	now := time.Now().UnixNano()
	return jokes[int(now)%len(jokes)]
}

func askQuestions(reader *bufio.Reader, category Category) int {
	fmt.Printf("\n--- %s ---\n", strings.ToUpper(category.Name))
	score := 0

	for i, question := range category.Questions {
		fmt.Printf("\nQuestion %d: %s\n", i+1, question.Prompt)
		for idx, choice := range question.Choices {
			fmt.Printf("  %c) %s\n", 'A'+idx, choice)
		}

		answer := promptAnswer(reader, len(question.Choices))
		if answer == question.Answer {
			fmt.Println("✅ Correct !")
			score++
		} else {
			fmt.Printf("❌ Mauvaise réponse. La bonne réponse était %c) %s\n", 'A'+question.Answer, question.Choices[question.Answer])
		}
	}

	return score
}

func promptAnswer(reader *bufio.Reader, choiceCount int) int {
	for {
		fmt.Print("Ta réponse (A, B, C...) : ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(strings.ToLower(input))

		if len(input) == 1 {
			idx := int(input[0] - 'a')
			if idx >= 0 && idx < choiceCount {
				return idx
			}
		}

		fmt.Println("Réponse invalide. Essaie encore !")
	}
}

func rankTitle(score, total int) string {
	ratio := float64(score) / float64(total)

	switch {
	case ratio <= 0.3:
		return "Apprenti·e codeur·se"
	case ratio <= 0.6:
		return "Développeur·se en progression"
	case ratio < 1.0:
		return "Expert·e sécurité"
	default:
		return "Cyber Mastermind"
	}
}

func saveScore(entry ScoreEntry) error {
	entries, err := loadScores()
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return err
	}

	entries = append(entries, entry)
	sort.SliceStable(entries, func(i, j int) bool {
		if entries[i].Score == entries[j].Score {
			return entries[i].Timestamp.After(entries[j].Timestamp)
		}
		return entries[i].Score > entries[j].Score
	})

	file, err := os.Create(scoreboardFile)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(entries)
}

func loadScores() ([]ScoreEntry, error) {
	file, err := os.Open(scoreboardFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var entries []ScoreEntry
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&entries); err != nil {
		return nil, err
	}

	return entries, nil
}

func displayBestScores() error {
	entries, err := loadScores()
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			fmt.Println("\nPas encore de meilleurs scores, sois le premier !")
			return nil
		}
		return err
	}

	fmt.Println("\n=== Tableau des meilleurs scores ===")
	limit := 5
	if len(entries) < limit {
		limit = len(entries)
	}

	for i := 0; i < limit; i++ {
		entry := entries[i]
		fmt.Printf("%d. %s — %d/%d (%s) sur %s\n", i+1, entry.Player, entry.Score, entry.Total, entry.Title, entry.Category)
	}

	return nil
}
