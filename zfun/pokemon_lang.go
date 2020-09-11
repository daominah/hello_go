package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/agnivade/levenshtein"
	"github.com/mywrap/gofast"
	"github.com/mywrap/log"
	"github.com/mywrap/textproc"
	"golang.org/x/net/html"
)

func main() {
	langs, err := getAllProgLangs()
	if err != nil {
		log.Fatal(err)
	}
	pokemons, err := getAllPokemons()
	if err != nil {
		log.Fatal(err)
	}
	for _, lang := range langs {
		for _, pokemon := range pokemons {
			dist := levenshtein.ComputeDistance(lang, pokemon)
			if dist > 2 {
				continue
			}
			if float64(dist) / float64(gofast.MinInts(len(lang), len(pokemon))) >= 0.5 {
				continue
			}
			log.Debugf("dist: %v, lang: %v, pokemon: %v",
				dist, lang, pokemon)
		}
	}
}

func getAllProgLangs() ([]string, error) {
	url0 := "https://en.wikipedia.org/wiki/List_of_programming_languages"
	tree, err := getPageTree(url0)
	if err != nil {
		return nil, err
	}
	langElems, err := textproc.HTMLXPath(tree, `/html/body/div[3]/div[3]/div[5]/div[1]/div/ul/li`)
	if err != nil {
		return nil, err
	}
	langs := make([]string, 0)
	for _, elem := range langElems[:len(langElems)-8] {
		//log.Debugf("i: %v", i)
		lang := textproc.HTMLGetText(elem)
		//log.Debugf(lang)
		langs = append(langs, lang)
	}
	return langs, nil
}

func getAllPokemons() ([]string, error) {
	url0 := "https://pokemondb.net/pokedex/national"
	tree, err := getPageTree(url0)
	if err != nil {
		return nil, err
	}
	pokeElems, err := textproc.HTMLXPath(tree, `//*[@id="main"]//span/a[contains(@class, 'ent-name')]`)
	if err != nil {
		return nil, err
	}
	pokemons := make([]string, 0)
	for _, elem := range pokeElems {
		//log.Debugf("i: %v", i)
		pokemon := textproc.HTMLGetText(elem)
		//log.Debugf(pokemon)
		pokemons = append(pokemons, pokemon)
	}
	return pokemons, nil
}

func getPageTree(pageURL string) (*html.Node, error) {
	resp, err := http.Get(pageURL)
	if err != nil {
		return nil, fmt.Errorf("http Get: %v", err)
	}
	defer resp.Body.Close()
	bodyB, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("ioutil ReadAll respBody: %v", err)
	}
	htmlTree, err := html.Parse(bytes.NewReader(bodyB))
	if err != nil {
		return nil, fmt.Errorf("html Parse: %v", err)
	}
	return htmlTree, nil
}
