package kptncook

import (
	"fmt"
	"time"

	"smartmirror.server/utils"
)

type recipeImageType string

const (
	recipeImageFavorite recipeImageType = "favorite"
	recipeImageBlurred  recipeImageType = "blurred"
	recipeImageCover    recipeImageType = "cover"
)

type dailyRecipesResponse []struct {
	Title         string `json:"title"`
	AuthorComment string `json:"authorComment"`
	FavoriteCount int    `json:"favoriteCount"`
	ImageList     []struct {
		URL  string          `json:"url"`
		Type recipeImageType `json:"type"`
	} `json:"imageList"`
}

func getDailyRecipes() ([]dailyRecipesModel, error) {
	if dailyRecipes, found := kptnCookCache.getDailyRecipes(); found {
		return dailyRecipes, nil
	}

	var response dailyRecipesResponse

	// timestamp from 3am of the current day
	now := time.Now()
	threeAM := time.Date(
		now.Year(), now.Month(), now.Day(),
		2, 0, 0, 0, &time.Location{},
	)
	timestamp := threeAM.Unix()

	err := utils.RelaxedHttpRequest(utils.RelaxedHttpRequestOptions{
		URL:      fmt.Sprintf("https://mobile.kptncook.com/recipes/de/%d?kptnkey=6q7QNKy-oIgk-IMuWisJ-jfN7s6&lang=de&recipeFilter=veggie", timestamp),
		Response: &response,
	})

	if err != nil {
		return []dailyRecipesModel{}, fmt.Errorf("Failed to fetch daily recipes from kptn cook: %v", err)
	}

	var dailyRecipes = make([]dailyRecipesModel, 0, len(response))
	for _, recipe := range response {
		newRecipe := dailyRecipesModel{
			Title:         recipe.Title,
			FavoriteCount: recipe.FavoriteCount,
		}

		for _, image := range recipe.ImageList {
			if image.Type == recipeImageFavorite {
				newRecipe.ImageUrl = image.URL
				break
			}
		}

		if newRecipe.ImageUrl != "" {
			dailyRecipes = append(dailyRecipes, newRecipe)
		}

	}

	kptnCookCache.setDailyRecipes(dailyRecipes)

	return dailyRecipes, nil
}
