package main

import (
	"github.com/kylesliu/awesome-golang-leetcode/cmd/leetcode"
)

func main() {
	problems := leetcode.GetSortedProblemsInstance()
	//problems := leetcode.GetProblemsJson()
	//fmt.Println(problems)
	//for _, v := range problems {
	//	fmt.Println(v.Stat.QuestionID, v.Stat.QuestionTitle, v.PathName,v.Stat.QuestionTitleSlup)
	//err := leetcode.UrlPath("https://leetcode.com/problems/" + v.Stat.QuestionTitleSlup )
	//if err != nil {
	//	fmt.Println(v.Stat.QuestionID, v.Stat.QuestionTitle, v.PathName,v.Stat.QuestionTitleSlup)
	//	fmt.Println(err.Error())
	//}
	//}

	//	生成Problem 目录

	//leetcode.MakeDir(problems)

	//leetcode.GetReadmeTemplateBuffer()

	//	GitBook
	leetcode.MakeGitbookSummary(problems)

}
