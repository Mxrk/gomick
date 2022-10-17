package gomick

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

const (
	BaseUrl    string = "https://api.comick.fun"
	SearchUrl         = BaseUrl + "/search/title?t=1&q="
	ComicUrl          = BaseUrl + "/comic/"
	ChapterUrl        = BaseUrl + "/chapter"
	ImageUrl          = "https://meo.comick.pictures/"

	OrderHot = "hot"
	OrderNew = "new"
)

var userAgent string

type Gom struct {
	Client http.Client
}

// New creates a gom object.
func New(timeout time.Duration, UserAgent string) Gom {
	var g Gom
	g.Client = http.Client{
		Timeout: time.Second * timeout,
	}

	userAgent = UserAgent

	return g
}

// Search will find results for a given query.
func (g Gom) Search(args SearchArgs) SearchResults {
	v := url.Values{}

	if len(args.Genres) != 0 {
		for _, genre := range args.Genres {
			v.Set("genres", genre)
		}
	}

	if len(args.Exclude) != 0 {
		for _, genre := range args.Exclude {
			v.Set("excludes", genre)
		}
	}

	if len(args.Tags) != 0 {
		for _, tag := range args.Tags {
			v.Set("tags", tag)
		}
	}

	if len(args.Demographic) != 0 {
		for _, d := range args.Demographic {
			v.Set("demographic", strconv.Itoa(d))
		}
	}

	if args.Page != 0 {
		v.Add("page", strconv.Itoa(args.Page))
	}

	if args.Limit != 0 {
		v.Add("limit", strconv.Itoa(args.Limit))
	}

	if args.Time != 0 {
		v.Add("time", strconv.Itoa(args.Time))
	}

	if len(args.Country) != 0 {
		for _, country := range args.Country {
			v.Set("country", country)
		}
	}

	if args.Minimum != 0 {
		v.Add("minimum", strconv.Itoa(args.Minimum))
	}

	if args.From != 0 {
		v.Add("from", strconv.Itoa(args.From))
	}

	if args.To != 0 {
		v.Add("to", strconv.Itoa(args.To))
	}

	if args.Sort != "" {
		v.Add("sort", args.Sort)
	}

	if args.Q != "" {
		v.Add("q", args.Q)
	}

	body := request(g.Client, BaseUrl+"/search"+v.Encode())

	results := SearchResults{}
	jsonErr := json.Unmarshal(body, &results)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	return results
}

// QuerySearch only needs a query as string.
func (g Gom) QuerySearch(query string) SearchResults {
	body := request(g.Client, prepareSearchQuery(query))

	results := SearchResults{}
	jsonErr := json.Unmarshal(body, &results)
	if jsonErr != nil {
		log.Println(jsonErr)
		log.Println(query, prepareSearchQuery(query))
	}

	return results
}

// GetComic returns a ComicResult for a valid slug.
func (g Gom) GetComic(slug string) ComicResult {
	body := request(g.Client, ComicUrl+slug)

	var comic ComicResult
	jsonErr := json.Unmarshal(body, &comic)
	if jsonErr != nil {
		log.Println(jsonErr)
	}

	return comic
}

// GetChapter returns chapters for a given hid.
func (g Gom) GetChapter(hid string) ChapterResult {
	body := request(g.Client, ChapterUrl+"/"+hid)

	var res ChapterResult
	jsonErr := json.Unmarshal(body, &res)
	if jsonErr != nil {
		log.Println(jsonErr)
	}

	return res
}

// GetNew returns the current first page of the new list.
func (g Gom) GetNew() []Chapter {
	body := request(g.Client, ChapterUrl+"?page=1&order=new")

	var res []Chapter
	jsonErr := json.Unmarshal(body, &res)
	if jsonErr != nil {
		log.Println(jsonErr)
	}

	return res
}

// Chapter returns the latest chapters from the main page.
func (g Gom) Chapter(args ChapterArgs) ChapterResult {
	v := url.Values{}

	if args.Lang != "" {
		v.Add("lang", args.Lang)
	}

	if args.Page != 0 {
		v.Add("page", strconv.Itoa(args.Page))
	}

	if args.Order != "" {
		v.Add("order", args.Order)
	}

	body := request(g.Client, ChapterUrl+v.Encode())

	var res ChapterResult
	jsonErr := json.Unmarshal(body, &res)
	if jsonErr != nil {
		log.Println(jsonErr)
	}

	return res
}

// ComicChapter returns chapters of a comic. Only args.ID field is required.
func (g Gom) ComicChapter(args ComicChapterArgs) Chapters {
	v := url.Values{}

	if args.Lang != "" {
		v.Add("lang", args.Lang)
	}

	if args.Page != 0 {
		v.Add("page", strconv.Itoa(args.Page))
	}

	if args.Limit != 0 {
		v.Add("limit", strconv.Itoa(args.Limit))
	}

	if args.ChapOrder != 0 {
		v.Add("chap-order", strconv.Itoa(args.ChapOrder))
	}

	if args.DateOrder != 0 {
		v.Add("date-order", strconv.Itoa(args.DateOrder))
	}

	if args.Chap != "" {
		v.Add("chap", args.Chap)
	}

	if args.ID != 0 {
		v.Add("id", strconv.Itoa(args.ID))
	} else {
		return Chapters{}
	}

	body := request(g.Client, ComicUrl+v.Get("id")+"/chapter"+v.Encode())

	var chapters Chapters
	jsonErr := json.Unmarshal(body, &chapters)
	if jsonErr != nil {
		log.Println(jsonErr)
	}

	return chapters
}

// GetHot returns the current first page of the hot list.
func (g Gom) GetHot() []Chapter {
	body := request(g.Client, ChapterUrl+"?page=1&order=hot")

	var res []Chapter
	jsonErr := json.Unmarshal(body, &res)
	if jsonErr != nil {
		log.Println(jsonErr)
	}

	return res
}

// GetTrending returns the current first page of the trending list. Gender has to be either 1 or 2, day 180, 270, 360, 720.
func (g Gom) GetTrending(gender int, day int, typeQuery string) TopResult {
	url := BaseUrl + "/top"

	if gender == 1 || gender == 2 {
		url += "?gender=" + strconv.Itoa(gender)
	}

	if day == 180 || day == 270 || day == 360 || day == 720 {
		url += "&day=" + strconv.Itoa(day)
	}

	if typeQuery != "" {
		url += "&type=" + typeQuery
	}

	body := request(g.Client, url)

	var res TopResult
	jsonErr := json.Unmarshal(body, &res)
	if jsonErr != nil {
		log.Println(jsonErr)
	}

	return res
}

// GetChapters returns Chapters for a given Comic ID.
func (g Gom) GetChapters(id int) Chapters {
	body := request(g.Client, ComicUrl+strconv.Itoa(id)+"/chapter?lang=en")

	var chapters Chapters
	jsonErr := json.Unmarshal(body, &chapters)
	if jsonErr != nil {
		log.Println(jsonErr)
	}

	return chapters
}

// GetImages returns the images as strings for a given HID.
func (g Gom) GetImages(hid string) []string {
	requestPages := g.GetChapter(hid)

	var pages []string

	for _, image := range requestPages.Chapter.MdImages {
		pages = append(pages, ImageUrl+image.B2Key)
	}

	return pages
}

func prepareSearchQuery(query string) string {
	replacedQuery := strings.ReplaceAll(query, " ", "+")

	return SearchUrl + replacedQuery
}

func request(client http.Client, url string) []byte {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Println(err)
	}

	req.Header.Set("User-Agent", userAgent)

	res, getErr := client.Do(req)
	if getErr != nil {
		log.Println(getErr)
	}

	defer res.Body.Close()

	remainingLimit := res.Header.Get("x-ratelimit-remaining")
	remainingLimitInt, err := strconv.Atoi(remainingLimit)
	if err != nil {
		log.Println(err)
	}

	if remainingLimitInt == 1 {
		tillReset := res.Header.Get("x-ratelimit-reset")
		tillResetInt, err := strconv.Atoi(tillReset)
		if err != nil {
			log.Println(err)
		}

		time.Sleep(time.Duration(tillResetInt) * time.Second)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Println(readErr)
	}

	return body
}
