// Copyright (c) 2023 Proton AG
//
// This file is part of Proton Mail Bridge.
//
// Proton Mail Bridge is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// Proton Mail Bridge is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with Proton Mail Bridge. If not, see <https://www.gnu.org/licenses/>.

package kb

import (
	_ "embed"
	"encoding/json"
	"regexp"
	"strings"

	"github.com/bradenaw/juniper/xslices"
	"golang.org/x/exp/slices"
)

//go:embed kbArticleList.json
var articleListString []byte

// Article is a struct that holds information about a knowledge-base article.
type Article struct {
	Index    uint64   `json:"index"`
	URL      string   `json:"url"`
	Title    string   `json:"title"`
	Keywords []string `json:"keywords"`
	Score    int
}

type ArticleList []*Article

// GetArticleList returns the list of KB articles.
func GetArticleList() (ArticleList, error) {
	var articles ArticleList
	err := json.Unmarshal(articleListString, &articles)

	return articles, err
}

// GetSuggestions return a list of up to 3 suggestions for KB articles matching the given user input.
func GetSuggestions(userInput string) (ArticleList, error) {
	userInput = strings.ToUpper(userInput)
	articles, err := GetArticleList()
	if err != nil {
		return ArticleList{}, err
	}

	for _, article := range articles {
		for _, keyword := range article.Keywords {
			if strings.Contains(userInput, strings.ToUpper(keyword)) {
				article.Score++
			}
		}
	}

	articles = xslices.Filter(articles, func(article *Article) bool { return article.Score > 0 })
	slices.SortFunc(articles, func(lhs, rhs *Article) bool { return lhs.Score > rhs.Score })

	if len(articles) > 3 {
		return articles[:3], nil
	}

	return articles, nil
}

func simplifyUserInput(input string) string {
	// replace any sequence not matching of the following with a single space:
	// - letters in any language (accentuated or not)
	// - numbers
	// - the apostrophe character '
	return strings.TrimSpace(regexp.MustCompile(`[^\p{L}\p{N}']+`).ReplaceAllString(input, " "))
}