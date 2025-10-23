package verses

import (
	"fmt"
	"hash/fnv"
	"math/rand"
	"strings"
	"time"
)

type Verse struct {
	Reference string
	Text      string
	Book      string
	Testament string
}

type FilterOptions struct {
	Testament string
	Book      string
}

func GetVerse(daily bool, opts FilterOptions) (*Verse, error) {
	filtered := filterVerses(opts)

	if len(filtered) == 0 {
		return nil, fmt.Errorf("no verses found matching the specified criteria")
	}

	var index int
	if daily {
		index = getDailyIndex(len(filtered))
	} else {
		index = getRandomIndex(len(filtered))
	}

	return &filtered[index], nil
}

func filterVerses(opts FilterOptions) []Verse {
	filtered := make([]Verse, 0, len(allVerses))

	for _, v := range allVerses {
		if opts.Testament != "" && v.Testament != opts.Testament {
			continue
		}

		if opts.Book != "" && !matchesBook(v.Book, opts.Book) {
			continue
		}

		filtered = append(filtered, v)
	}

	return filtered
}

func matchesBook(bookName, filter string) bool {
	if strings.EqualFold(bookName, filter) {
		return true
	}

	bookLower := strings.ToLower(bookName)
	filterLower := strings.ToLower(filter)

	variations := map[string][]string{
		"psalm":  {"psalm", "psalms"},
		"psalms": {"psalm", "psalms"},
	}

	if alts, ok := variations[bookLower]; ok {
		for _, alt := range alts {
			if alt == filterLower {
				return true
			}
		}
	}

	if alts, ok := variations[filterLower]; ok {
		for _, alt := range alts {
			if alt == bookLower {
				return true
			}
		}
	}

	return false
}

func getRandomIndex(count int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(count)
}

func getDailyIndex(count int) int {
	today := time.Now().Format("2006-01-02")

	h := fnv.New32a()
	h.Write([]byte(today))
	seed := int64(h.Sum32())

	r := rand.New(rand.NewSource(seed))
	return r.Intn(count)
}
