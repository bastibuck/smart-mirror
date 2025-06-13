package garmin

func getStepsToday() (int, error) {
	return 9000, nil
}

// client := garmin.NewClient()
// err := client.Login(os.Getenv("GARMIN_EMAIL"), os.Getenv("GARMIN_PASSWORD"))
// if err != nil {
// 	panic(err)
// }
// api := garmin.NewAPI(client)
// steps, err := api.UserSummary.DailySteps(
// 	time.Date(2025, 6, 2, 0, 0, 0, 0, time.UTC),
// 	time.Date(2025, 6, 8, 0, 0, 0, 0, time.UTC),
// )

// if err != nil {
// 	panic(err)
// }

// fmt.Printf("%+v\n", steps.Aggregations.TotalStepsAverage)
