package hw03frequencyanalysis

import (
	"errors"
	"fmt"
	"sort"
	"strings"
)

type wordsFrequency struct {
	Word string
	Freq int
}

type ByFreq []wordsFrequency

func (o ByFreq) Len() int      { return len(o) }
func (o ByFreq) Swap(i, j int) { o[i], o[j] = o[j], o[i] }
func (o ByFreq) Less(i, j int) bool {
	if o[i].Freq == o[j].Freq {
		return o[i].Word < o[j].Word
	}
	return o[i].Freq > o[j].Freq
}

func sortWords(m map[string]int) []wordsFrequency {
	ss := make([]wordsFrequency, 0, len(m))
	for k, v := range m {
		ss = append(ss, wordsFrequency{k, v})
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Freq > ss[j].Freq
	})
	sort.Sort(ByFreq(ss))
	return ss
}

func replaceSymbolsStr(s string) (string, error) {
	if s == "" {
		return "", errors.New("empty name")
	}
	r := strings.NewReplacer(
		"	", " ",
		"\n", " ",
	)
	replacedStr := r.Replace(s)
	return replacedStr, nil
}

func Top10(text string) []string {
	formattedText, err := replaceSymbolsStr(text)
	if err != nil {
		return nil
	}
	words := strings.Fields(formattedText)

	m := make(map[string]int)
	for _, word := range words {
		m[word]++
	}

	s := sortWords(m)[:10]
	output := make([]string, 0, len(s))
	for _, kv := range s {
		output = append(output, kv.Word)
		fmt.Printf("%s, %d\n", kv.Word, kv.Freq)
	}

	return output
}
