package kptncook

import (
	"fmt"

	"smartmirror.server/utils"
)

type dailyRecipesResponse []struct {
	Title string `json:"title"`
}

func getDailyRecipes() (dailyRecipesModel, error) {
	if dailyRecipes, found := kptnCookCache.getDailyRecipes(); found {
		return dailyRecipes, nil
	}

	var response dailyRecipesResponse

	err := utils.RelaxedHttpRequest(utils.RelaxedHttpRequestOptions{
		URL:      "https://mobile.kptncook.com/recipes/de/1?kptnkey=6q7QNKy-oIgk-IMuWisJ-jfN7s6&lang=de&recipeFilter=veggie",
		Response: &response,
	})

	if err != nil {
		return dailyRecipesModel{}, fmt.Errorf("Failed to fetch daily recipes from kptn cook: %v", err)
	}

	var dailyRecipes = make(dailyRecipesModel, 0, len(response))
	for _, recipe := range response {
		dailyRecipes = append(dailyRecipes, struct {
			Title string `json:"title"`
		}{
			Title: recipe.Title,
		})
	}

	kptnCookCache.setDailyRecipes(dailyRecipes)

	return dailyRecipes, nil
}
