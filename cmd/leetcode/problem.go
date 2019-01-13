package leetcode

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

type LeetCode struct {
	UserName        string    `json:"user_name"`
	NumSolved       int       `json:"num_solved"`
	NumTotal        int       `json:"num_total"`
	AcEasy          int       `json:"ac_easy"`
	AcMedium        int       `json:"ac_medium"`
	AcHard          int       `json:"ac_hard"`
	StatStatusPairs []Problem `json:"stat_status_pairs"`
	FrequencyHigh   int       `json:"frequency_high"`
	FrequencyMid    int       `json:"frequency_mid"`
	CategorySlug    string    `json:"category_slug"`
}

//	问题
type Problem struct {
	Stat       Stat       `json:"stat"`
	Status     bool       `json:"status"`
	Difficulty Difficulty `json:"difficulty"`
	PaidOnly   bool       `json:"paid_only"`
	IsFavor    bool       `json:"is_favor"`
	Frequency  int        `json:"frequency"`
	Progress   int        `json:"progress"`
	PathName   string     `json:"path_name"`
}
type Stat struct {
	QuestionID          int    `json:"question_id"`
	QuestionArticleLive bool   `json:"question__article__live"`
	QuestionArticleSlug string `json:"question__article__slug"`
	QuestionTitle       string `json:"question__title"`
	QuestionTitleSlup   string `json:"question__title_slug"`
	QuestionHide        bool   `json:"question__hide"`
	TotalAcs            int    `json:"total_acs"`
	TotalSubmitted      int    `json:"total_submitted"`
	FrontendQuestionID  int    `json:"frontend_question_id"`
	IsNewQuestion       bool   `json:"is_new_question"`
}

type Difficulty struct {
	Level int `json:"level"`
}

func GetAllProblemsPath() []string {
	problems := GetProblemsInstance()

	res := []string{}
	time.Sleep(time.Second)

	for i, _ := range problems {
		res = append(res, problems[i].PathName)
		//fmt.Println(problems[i].Stat.QuestionTitle)
	}
	return res
}

func GetProblemsInstance() []Problem {

	leetcode := new(LeetCode)
	Problemsbuffer := getProblemsBuffer()

	if err := json.Unmarshal(Problemsbuffer, leetcode); err != nil {
		log.Panicf("Json 转换失败: %s\n", err.Error())
	}

	for i := 0; i < len(leetcode.StatStatusPairs); i++ {
		leetcode.StatStatusPairs[i].PathName = string(formatId(leetcode.StatStatusPairs[i].Stat.FrontendQuestionID)) + "." +
			formatName(leetcode.StatStatusPairs[i].Stat.QuestionTitle)
	}

	return leetcode.StatStatusPairs
}

func GetProblemsJosn() string {
	problems := GetProblemsInstance()
	problem_string, err := json.MarshalIndent(problems, " ", " ")
	if err != nil {
		log.Fatalln("Problem Json 序列化失败: ", err.Error())
	}

	return string(problem_string)
}

//	获取题目Buffer
func getProblemsBuffer() []byte {
	request, err := http.Get("https://leetcode.com/api/problems/Algorithms/")

	if err != nil {
		log.Panicln("Lettcode Problem 接口获取失败：", err)
	}

	if request.StatusCode != 200 {
		log.Panicln("Lettcode Problem 接口地址不存在：", err)
	}

	body, _ := ioutil.ReadAll(request.Body)
	return body
}

//	格式化ID 补齐0
func formatId(id int) string {
	strId := strconv.Itoa(id)

	if len(strId) == 1 {
		strId = "000" + strId
	}
	if len(strId) == 2 {
		strId = "00" + strId
	}
	if len(strId) == 3 {
		strId = "0" + strId
	}

	return strId
}

//	格式化提名称
func formatName(name string) string {
	str := ""
	for _, v := range name {
		if v == ' ' {
			str = str + "-"
			continue
		}
		if v == '(' || v == ')' {
			continue
		}
		str = str + string(v)
	}
	return str
}
