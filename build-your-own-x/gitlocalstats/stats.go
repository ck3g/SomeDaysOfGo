package main

// stats calcualtes and prints the stats
func stats(email string) {
	commits := processRepositories(email)
	printCommitsStats(commits)
}

func processRepositories(email string) map[int]int {
	return make(map[int]int)
}

func printCommitsStats(commits map[int]int) {

}
