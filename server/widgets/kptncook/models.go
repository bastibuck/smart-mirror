package kptncook

type dailyRecipesModel struct {
	Title         string `json:"title"`
	ImageUrl      string `json:"imageUrl"`
	FavoriteCount int    `json:"favoriteCount"`
}
