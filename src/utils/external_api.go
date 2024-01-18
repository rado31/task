package utils

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type Agify struct {
	Age int `json:"age"`
}

type Genderize struct {
	Gender string `json:"gender"`
}

type Nationalize struct {
	Country []Country `json:"country"`
}

type Country struct {
	Country_id  string  `json:"country_id"`
	Probability float32 `json:"probability"`
}

type Result struct {
	Name        string
	Surname     string
	Patronymic  string
	Age         int
	Gender      string
	Nationality string
}

func Get_data(name, surname, patronymic string) (Result, error) {
	var result Result

	agify_res, agify_err := http.Get("https://api.agify.io/?name=" + name)

	if agify_err != nil {
		return result, agify_err
	}

	gen_res, gen_err := http.Get("https://api.genderize.io/?name=" + name)

	if gen_err != nil {
		return result, gen_err
	}

	nat_res, nat_err := http.Get("https://api.nationalize.io/?name=" + name)

	if gen_err != nil {
		return result, nat_err
	}

	agify_buf, _ := io.ReadAll(agify_res.Body)
	gen_buf, _ := io.ReadAll(gen_res.Body)
	nat_buf, _ := io.ReadAll(nat_res.Body)

	var agify Agify
	var gen Genderize
	var nat Nationalize

	json.Unmarshal(agify_buf, &agify)
	json.Unmarshal(gen_buf, &gen)
	json.Unmarshal(nat_buf, &nat)

	var max_probability float32 = 0.0
	var nationality string

	// the first object in array of response body of
	// "https://api.nationalize.io/?name=" always the max probability,
	// but despite this I still check
	for i := 0; i < len(nat.Country); i++ {
		if nat.Country[i].Probability > max_probability {
			max_probability = nat.Country[i].Probability
			nationality = nat.Country[i].Country_id
		}
	}

	result.Name = name
	result.Surname = surname
	result.Patronymic = patronymic
	result.Age = agify.Age
	result.Gender = gen.Gender
	result.Nationality = nationality

	// when name is random text, like "asdiansidqwe", external api gives
	// null value, that's we must need to check the response
	if result.Age == 0 || result.Gender == "" || result.Nationality == "" {
		return result, errors.New(
			"Couldn't create additional info, cause you've write unusual name",
		)
	}

	return result, nil
}
