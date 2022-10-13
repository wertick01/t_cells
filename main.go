package main

import (
	"allignment-t-cells/models"
	"encoding/json"
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var a = []int{10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
var groups = []string{
	"TRAC (1)",
	"TRAJ",
	"TRAV",
	"TRBC",
	"TRBD",
	"TRBJ",
	"TRBV",
	"TRGC",
	"TRGJ",
	"TRGV",
	"TRDC (1)",
	"TRDD",
	"TRDJ",
	"TRDV",
}

func main() {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var primers = []models.Primers{}
	var keyPrimers = []models.KeyPrimers{}

	var mass = []models.KeySequence{}
	var data = []models.Data{}

	var prefinal = []models.PreFinal{}

	groupDir, err := ioutil.ReadDir(pwd + "/Groups/")
	if err != nil {
		log.Fatal(err)
	}
	mass = GroupSorter(groupDir, mass, data, pwd+"/Groups/")

	primerDir, err := ioutil.ReadDir(pwd + "/Primers/")
	if err != nil {
		log.Fatal(err)
	}
	keyPrimers = PrimerSorter(primerDir, keyPrimers, primers, pwd+"/Primers/")

	for _, val := range keyPrimers {
		for _, seq := range mass {
			// if val.Key == seq.Key {
			for _, primer := range val.Primers {
				for _, sequence := range seq.Sequences {
					if strings.Contains(strings.ToUpper(sequence.Value), primer.Primer) {
						sequence.Value = strings.ToUpper(sequence.Value)
						prefinal = append(prefinal, models.PreFinal{Type: seq.Key, Primer: primer, Sequence: sequence})
					}
				}
				// }
			}
		}
	}
	fmt.Println(len(prefinal))

	f, _ := os.Create(pwd + "/Results/all_to_all.json")
	defer f.Close()

	as_json, err := json.Marshal(prefinal)
	if err != nil {
		log.Fatal(err)
	}
	f.Write(as_json)

	// var mass []Data

	// jsonFile, err := os.Open(pwd + "/data.json")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// jsonParser := json.NewDecoder(jsonFile)
	// if err = jsonParser.Decode(&mass); err != nil {
	// 	log.Fatal(err)
	// }

	// if err = Sorter(pwd, mass, groups); err != nil {
	// 	log.Fatal(err)
	// }
}

func GroupSorter(groupDir []fs.FileInfo, mass []models.KeySequence, data []models.Data, path string) []models.KeySequence {
	for _, f := range groupDir {
		jsonFile, err := os.Open(path + f.Name())
		if err != nil {
			log.Fatal(err)
		}
		jsonParser := json.NewDecoder(jsonFile)
		if err = jsonParser.Decode(&data); err != nil {
			log.Fatal(err)
		}
		mass = append(mass, models.KeySequence{Key: f.Name(), Sequences: data})
	}
	return mass
}

func PrimerSorter(primerDir []fs.FileInfo, keyPrimers []models.KeyPrimers, primers []models.Primers, path string) []models.KeyPrimers {
	for _, f := range primerDir {
		jsonFile, err := os.Open(path + f.Name())
		if err != nil {
			log.Fatal(err)
		}
		jsonParser := json.NewDecoder(jsonFile)
		if err = jsonParser.Decode(&primers); err != nil {
			log.Fatal(err)
		}
		keyPrimers = append(keyPrimers, models.KeyPrimers{Key: f.Name(), Primers: primers})
	}
	return keyPrimers
}

func Sorter(pwd string, data []models.Data, groups []string) error {
	for _, group := range groups {
		mass := []models.Data{}
		for _, val := range data {
			if strings.Contains(val.Key, group) {
				mass = append(mass, models.Data{
					Key:   val.Key,
					Value: val.Value,
				},
				)
			}
		}
		f, _ := os.Create(pwd + "/Groups/" + group + ".json")
		defer f.Close()

		as_json, err := json.Marshal(mass)
		if err != nil {
			return err
		}
		f.Write(as_json)
	}
	return nil
}

func Splitter(data []models.Data) []models.Substrings {
	var mass []models.Substrings
	for _, count := range a {
		for _, val := range data {
			if len(val.Value) > 20 {
				sub := &models.Substrings{}
				sbstr := []string{}
				for i := 0; i < len(val.Value)-count; i++ {
					sbstr = append(sbstr, val.Value[i:i+count])
				}
				sub.Key = val.Key
				sub.Substrings = sbstr
				sub.Length = count
				mass = append(mass, *sub)
			}
		}
	}
	return mass
}

func Writer(path string, res []models.Substrings) {
	f, _ := os.Create(path)
	defer f.Close()

	as_json, err := json.Marshal(res)
	if err != nil {
		log.Fatal(err)
	}
	f.Write(as_json)
}

func Alligner(data []models.Data, substrings []models.Substrings) string {
	aa := 0
	bb := 0
	// var rt int
	// var wg sync.WaitGroup
	var ct int
	// var result []AllignerSeq
	var res models.AllignerSeq
	f, _ := os.Create("/home/mrred/Рабочий стол/Работа/t_cells/Statistics.json")
	defer f.Close()
	// for i := 0; i < len(data)/10; i++ {
	// 	wg.Add(10)
	for _, value := range substrings {

		aa++
		// wg.Add(1)
		// go func(value Substrings) {
		// 	defer wg.Done()
		for _, substring := range value.Substrings {
			res = AllignHelper(value, data, ct, substring)
			if res.Count > 10 {
				as_json, _ := json.Marshal(res)
				// if err != nil {
				// 	log.Fatal(err)
				// }
				f.Write(as_json)
				// if result[aa-1].Count > 5 {
				// 	fmt.Println(result[aa-1])
				// }
			}
		}
		fmt.Println(aa, res)
		if aa%871 == 0 {
			bb++
			fmt.Println(bb)
		}
		// }(value)
		// fmt.Println(len(result))
		// }(value)
	}
	// time.Sleep(5 * time.Second)
	// wg.Wait()
	// }
	// wg.Wait()
	// as_json, err := json.Marshal(result)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// f.Write(as_json)
	// result = append(result, res)
	return "+"
}

func AllignHelper(value models.Substrings, data []models.Data, ct int, substring string) models.AllignerSeq {
	// for _, substring := range value.Substrings {
	stat := &models.AllignerSeq{}
	ms := []string{}
	ct = 0
	for _, sequence := range data {
		if strings.Contains(sequence.Value, substring) {
			ms = append(ms, sequence.Key)
			ct++
		}
	}
	stat.Primer = substring
	stat.Count = ct
	stat.Sequences = ms
	if ct < 10 {
		return models.AllignerSeq{Count: 0}
	}
	// if ct > 10 {
	// 	result = append(result, *stat)
	// }
	// }
	return *stat
}
