package main

import "github.com/s-bespalov/packt_go_programming/Chapter9/Exerciese09.01/bookutil/author"

func main() {
	authorInstance := author.NewAuthor("Jane Doe", "jane@example.com")
	chapterTitle := "Introduction to Go modules"
	chapterContent := "Go modules provide a structured way tomanage dependencies and improve code maintainability."
	authorInstance.WriteChapter(chapterTitle, chapterContent)
	authorInstance.ReviewChapter(chapterTitle, "This chapterlooks great, but let's add some more examples.")
	authorInstance.FinalizeChapter(chapterTitle)
}
