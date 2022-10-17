package gomick

import "time"

type SearchResults []struct {
	Title       string      `json:"title"`
	ID          int         `json:"id"`
	Slug        string      `json:"slug"`
	Rating      string      `json:"rating"`
	RatingCount int         `json:"rating_count"`
	FollowCount int         `json:"follow_count"`
	Demographic interface{} `json:"demographic"`
	MdTitles    []struct {
		Title string `json:"title"`
	} `json:"md_titles"`
	MdCovers []struct {
		URL   string `json:"url"`
		Vol   string `json:"vol"`
		W     int    `json:"w"`
		H     int    `json:"h"`
		Gpurl string `json:"gpurl"`
		B2Key string `json:"b2key"`
	} `json:"md_covers"`
}

type ComicResult struct {
	Comic struct {
		ID                 int     `json:"id"`
		Title              string  `json:"title"`
		Country            string  `json:"country"`
		Status             int     `json:"status"`
		LastChapter        float32 `json:"last_chapter"`
		Demographic        int     `json:"demographic"`
		Mdid               int     `json:"mdid"`
		Hentai             bool    `json:"hentai"`
		VerificationStatus int     `json:"verification_status"`
		UserFollowCount    int     `json:"user_follow_count"`
		FollowRank         int     `json:"follow_rank"`
		CommentCount       int     `json:"comment_count"`
		FollowCount        int     `json:"follow_count"`
		Desc               string  `json:"desc"`
		Parsed             string  `json:"parsed"`
		Slug               string  `json:"slug"`
		Year               int     `json:"year"`
		BayesianRating     string  `json:"bayesian_rating"`
		RatingCount        int     `json:"rating_count"`
		RelateFrom         []struct {
			RelateTo struct {
				Slug  string `json:"slug"`
				Title string `json:"title"`
			} `json:"relate_to"`
			MdRelates struct {
				Name string `json:"name"`
			} `json:"md_relates"`
		} `json:"relate_from"`
		Mies struct {
			MyRelatedsMiesTomyRelatedsRelatedID []struct {
				MiesMiesTomyRelatedsMyID struct {
					Myid   int `json:"myid"`
					Anidbs []struct {
						Slug  string `json:"slug"`
						Title string `json:"title"`
					} `json:"anidbs"`
				} `json:"mies_miesTomy_relateds_my_id"`
			} `json:"my_relateds_miesTomy_relateds_related_id"`
			MyRelatedsMiesTomyRelatedsMyID []struct {
				MiesMiesTomyRelatedsRelatedID struct {
					Myid   int `json:"myid"`
					Anidbs []struct {
						Slug  string `json:"slug"`
						Title string `json:"title"`
					} `json:"anidbs"`
				} `json:"mies_miesTomy_relateds_related_id"`
			} `json:"my_relateds_miesTomy_relateds_my_id"`
		} `json:"mies"`
		MdTitles []struct {
			Title string `json:"title"`
		} `json:"md_titles"`
		MdComicMdGenres []struct {
			MdGenres struct {
				Name  string `json:"name"`
				Type  string `json:"type"`
				Slug  string `json:"slug"`
				Group string `json:"group"`
			} `json:"md_genres"`
		} `json:"md_comic_md_genres"`
		MdCovers []struct {
			URL   string `json:"url"`
			Vol   string `json:"vol"`
			Gpurl string `json:"gpurl"`
			W     int    `json:"w"`
			H     int    `json:"h"`
			B2Key string `json:"b2key"`
		} `json:"md_covers"`
		MuComics struct {
			LicensedInEnglish   bool `json:"licensed_in_english"`
			CompletelyScanlated bool `json:"completely_scanlated"`
			MuComicCategories   []struct {
				MuCategories struct {
					Title string `json:"title"`
					Slug  string `json:"slug"`
				} `json:"mu_categories"`
				PositiveVote int `json:"positive_vote"`
				NegativeVote int `json:"negative_vote"`
			} `json:"mu_comic_categories"`
		} `json:"mu_comics"`
		Iso6391    string `json:"iso639_1"`
		LangName   string `json:"lang_name"`
		LangNative string `json:"lang_native"`
	} `json:"comic"`
	Artists []struct {
		Name string `json:"name"`
		Slug string `json:"slug"`
	} `json:"artists"`
	Authors []struct {
		Name string `json:"name"`
		Slug string `json:"slug"`
	} `json:"authors"`
	LangList    []string    `json:"langList"`
	Demographic string      `json:"demographic"`
	EnglishLink interface{} `json:"englishLink"`
	Genres      []struct {
		Slug string `json:"slug"`
		Name string `json:"name"`
	} `json:"genres"`
	MatureContent bool `json:"matureContent"`
}

type Chapters struct {
	Chapters []struct {
		ID          int         `json:"id"`
		Chap        string      `json:"chap"`
		Title       interface{} `json:"title"`
		Vol         string      `json:"vol"`
		Slug        interface{} `json:"slug"`
		CountryCode string      `json:"country_code"`
		CreatedAt   time.Time   `json:"created_at"`
		UpCount     int         `json:"up_count"`
		DownCount   int         `json:"down_count"`
		Hid         string      `json:"hid"`
		MdGroups    []struct {
			Slug  string `json:"slug"`
			Title string `json:"title"`
		} `json:"md_groups"`
		Iso6391  string `json:"iso639_1"`
		LangName string `json:"langName"`
		//LangName string `json:"lang_name"`
	} `json:"chapters"`
	Total int `json:"total"`
}

type Chapter struct {
	ID           int         `json:"id"`
	Chap         string      `json:"chap"`
	Vol          string      `json:"vol"`
	CountryCode  string      `json:"country_code"`
	Title        interface{} `json:"title"`
	Server       string      `json:"server"`
	Hid          string      `json:"hid"`
	Hash         string      `json:"hash"`
	GroupName    []string    `json:"group_name"`
	ChapterID    interface{} `json:"chapter_id"`
	CreatedAt    time.Time   `json:"created_at"`
	Mdid         string      `json:"mdid"`
	CommentCount int         `json:"comment_count"`
	UpCount      int         `json:"up_count"`
	DownCount    int         `json:"down_count"`
	Status       string      `json:"status"`
	Mseid        interface{} `json:"mseid"`
	MdComics     struct {
		ID          int     `json:"id"`
		Title       string  `json:"title"`
		Country     string  `json:"country"`
		Slug        string  `json:"slug"`
		LastChapter float32 `json:"last_chapter"`
		Links       struct {
			Al    string `json:"al"`
			Ap    string `json:"ap"`
			Kt    string `json:"kt"`
			Mu    string `json:"mu"`
			Nu    string `json:"nu"`
			Mal   string `json:"mal"`
			Raw   string `json:"raw"`
			Engtl string `json:"engtl"`
		} `json:"links"`
		Genres   []int `json:"genres"`
		Hentai   bool  `json:"hentai"`
		MdCovers []struct {
			URL   string `json:"url"`
			W     int    `json:"w"`
			H     int    `json:"h"`
			Gpurl string `json:"gpurl"`
			B2Key string `json:"b2key"`
		} `json:"md_covers"`
	} `json:"md_comics"`
	MdImages []struct {
		Gpurl interface{} `json:"gpurl"`
		H     int         `json:"h"`
		W     int         `json:"w"`
		Name  string      `json:"name"`
		B2Key string      `json:"b2key"`
	} `json:"md_images"`
	MdChaptersGroups []struct {
		MdGroups struct {
			Slug  string `json:"slug"`
			Title string `json:"title"`
		} `json:"md_groups"`
	} `json:"md_chapters_groups"`
	Iso6391  string `json:"iso639_1"`
	LangName string `json:"langName"`
}

type ChapterResult struct {
	Chapter Chapter `json:"chapter"`
	Next    struct {
		ID          int         `json:"id"`
		Hid         string      `json:"hid"`
		Chap        string      `json:"chap"`
		GroupName   interface{} `json:"group_name"`
		CountryCode string      `json:"country_code"`
		Iso6391     string      `json:"iso639_1"`
		LangName    string      `json:"langName"`
		Href        string      `json:"href"`
	} `json:"next"`
	Prev struct {
		ID          int         `json:"id"`
		Hid         string      `json:"hid"`
		Chap        string      `json:"chap"`
		GroupName   interface{} `json:"group_name"`
		CountryCode string      `json:"country_code"`
		Iso6391     string      `json:"iso639_1"`
		LangName    string      `json:"langName"`
		Href        string      `json:"href"`
	} `json:"prev"`
	CoverURL      string `json:"coverURL"`
	MatureContent bool   `json:"matureContent"`
	Chapters      []struct {
		ID          int      `json:"id"`
		Hid         string   `json:"hid"`
		Chap        string   `json:"chap"`
		GroupName   []string `json:"group_name"`
		CountryCode string   `json:"country_code"`
		Iso6391     string   `json:"iso639_1"`
		LangName    string   `json:"langName"`
		Href        string   `json:"href,omitempty"`
	} `json:"chapters"`
	DupGroupChapters []struct {
		ID          int      `json:"id"`
		Hid         string   `json:"hid"`
		Chap        string   `json:"chap"`
		GroupName   []string `json:"group_name"`
		CountryCode string   `json:"country_code"`
		Iso6391     string   `json:"iso639_1"`
		LangName    string   `json:"langName"`
	} `json:"dupGroupChapters"`
	ChapterLangList []struct {
		Iso6391     string `json:"iso639_1"`
		Name        string `json:"name"`
		Native      string `json:"native"`
		CountryCode string `json:"country_code"`
		Hid         string `json:"hid"`
	} `json:"chapterLangList"`
	Canonical      string `json:"canonical"`
	SeoTitle       string `json:"seoTitle"`
	SeoDescription string `json:"seoDescription"`
	ChapTitle      string `json:"chapTitle"`
}

type ChapterArgs struct {
	// Locale code https://developers.google.com/interactive-media-ads/docs/sdks/android/client-side/localization#locale-codes
	Lang string
	// Page as int for pagination. Start with 1.
	Page int
	// Sort by "hot" or "new". Recommend to use OrderHot r OrderNew
	Order string
}

type ComicChapterArgs struct {
	Limit int
	Page  int
	// Default order is 0.
	ChapOrder int
	DateOrder int
	// Locale code https://developers.google.com/interactive-media-ads/docs/sdks/android/client-side/localization#locale-codes
	Lang string
	Chap string
	// ID is a required field which has to be used to get a result.
	ID int
}

type SearchArgs struct {
	// Genres as array. Example: action,romance
	Genres []string
	// genres to exclude.
	Exclude []string
	// Tags as array. Example: family
	Tags []string
	// 1/shounen 2/shoujo 3/seinen 4/josei
	Demographic []int
	Page        int
	// Limit is default set to 50.
	Limit int
	// Time as days. 120 would subtract 120 days.
	Time    int
	Country []string
	// Minimum is the minimum number of chapters.
	Minimum int
	// From year.
	From int
	// To year.
	To int
	// Tachiyomi is required for third party software.
	Tachiyomi bool
	// completed translation
	Completed bool
	// Sort has the following values: "view", "uploaded", "rating", "follow"
	Sort string
	// Q contains the title to search.
	Q string
	// Include alt title for show only.
	T bool
}

type TopResult struct {
	Rank []struct {
		Slug        string  `json:"slug"`
		Title       string  `json:"title"`
		Demographic int     `json:"demographic"`
		Genres      []int   `json:"genres"`
		LastChapter float32 `json:"last_chapter"`
		MdCovers    []struct {
			W     int    `json:"w"`
			H     int    `json:"h"`
			B2Key string `json:"b2key"`
		} `json:"md_covers"`
	} `json:"rank"`
	RecentRank []struct {
		Slug        string      `json:"slug"`
		Title       string      `json:"title"`
		Demographic interface{} `json:"demographic"`
		Genres      []int       `json:"genres"`
		MdCovers    []struct {
			W     int    `json:"w"`
			H     int    `json:"h"`
			B2Key string `json:"b2key"`
		} `json:"md_covers"`
	} `json:"recentRank"`
	Trending struct {
		// trending in the last three days
		Num3 []struct {
			Title    string `json:"title"`
			Slug     string `json:"slug"`
			MdCovers []struct {
				W     int    `json:"w"`
				H     int    `json:"h"`
				B2Key string `json:"b2key"`
			} `json:"md_covers"`
		} `json:"3"`
		// trending in the last seven days
		Num7 []struct {
			Title    string `json:"title"`
			Slug     string `json:"slug"`
			MdCovers []struct {
				W     int    `json:"w"`
				H     int    `json:"h"`
				B2Key string `json:"b2key"`
			} `json:"md_covers"`
		} `json:"7"`
		Num30 []struct {
			Title    string `json:"title"`
			Slug     string `json:"slug"`
			MdCovers []struct {
				W     int    `json:"w"`
				H     int    `json:"h"`
				B2Key string `json:"b2key"`
			} `json:"md_covers"`
		} `json:"30"`
	} `json:"trending"`
	Follows []struct {
		CreatedAt  time.Time `json:"created_at"`
		Identities struct {
			ID     string `json:"id"`
			Traits struct {
				Username string `json:"username"`
				ID       string `json:"id"`
				Gravatar string `json:"gravatar"`
			} `json:"traits"`
		} `json:"identities"`
		MdComics struct {
			ID          int    `json:"id"`
			Title       string `json:"title"`
			Slug        string `json:"slug"`
			FollowCount int    `json:"follow_count"`
			MdCovers    []struct {
				W     int    `json:"w"`
				H     int    `json:"h"`
				B2Key string `json:"b2key"`
			} `json:"md_covers"`
		} `json:"md_comics"`
	} `json:"follows"`
	News []struct {
		Slug        string      `json:"slug"`
		Title       string      `json:"title"`
		Demographic interface{} `json:"demographic"`
		Genres      []int       `json:"genres"`
		CreatedAt   time.Time   `json:"created_at"`
		MdCovers    []struct {
			W     int    `json:"w"`
			H     int    `json:"h"`
			B2Key string `json:"b2key"`
		} `json:"md_covers"`
	} `json:"news"`
	Completions []struct {
		Slug        string      `json:"slug"`
		Title       string      `json:"title"`
		Demographic interface{} `json:"demographic"`
		Genres      []int       `json:"genres"`
		CreatedAt   time.Time   `json:"created_at"`
		UploadedAt  time.Time   `json:"uploaded_at"`
		LastChapter float32     `json:"last_chapter"`
		MdCovers    []struct {
			W     int    `json:"w"`
			H     int    `json:"h"`
			B2Key string `json:"b2key"`
		} `json:"md_covers"`
	} `json:"completions"`
}
