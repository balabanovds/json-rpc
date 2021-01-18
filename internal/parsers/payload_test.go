package parsers_test

var smallInput = []byte(`{
	"st": 1,
	"sid": 486,
	"tt": "active",
	"gr": 0,
	"uuid": "de305d54-75b4-431b-adb2-eb6b9e546014",
	"ip": "127.0.0.1",
	"ua": "user_agent",
	"tz": -6,
	"v": 1
}`)

type smallPayload struct {
	St   int    `json:"st"`
	Sid  int    `json:"-"`
	Tt   string `json:"-"`
	Gr   int    `json:"-"`
	Uuid string `json:"uuid"`
	Ip   string `json:"-"`
	Ua   string `json:"ua"`
	Tz   int    `json:"tz"`
	V    int    `json:"-"`
}

var mediumInput = []byte(`{
	"person": {
    "id": "d50887ca-a6ce-4e59-b89f-14f0b5d03b03",
    "name": {
      "fullName": "Leonid Bugaev",
      "givenName": "Leonid",
      "familyName": "Bugaev"
    },
    "email": "leonsbox@gmail.com",
    "gender": "male",
    "location": "Saint Petersburg, Saint Petersburg, RU",
    "geo": {
      "city": "Saint Petersburg",
      "state": "Saint Petersburg",
      "country": "Russia",
      "lat": 59.9342802,
      "lng": 30.3350986
    },
    "bio": "Senior engineer at Granify.com",
    "site": "http://flickfaver.com",
    "avatar": "https://d1ts43dypk8bqh.cloudfront.net/v1/avatars/d50887ca-a6ce-4e59-b89f-14f0b5d03b03",
    "employment": {
      "name": "www.latera.ru",
      "title": "Software Engineer",
      "domain": "gmail.com"
    },
    "facebook": {
      "handle": "leonid.bugaev"
    },
    "github": {
      "handle": "buger",
      "id": 14009,
      "avatar": "https://avatars.githubusercontent.com/u/14009?v=3",
      "company": "Granify",
      "blog": "http://leonsbox.com",
      "followers": 95,
      "following": 10
    },
    "twitter": {
      "handle": "flickfaver",
      "id": 77004410,
      "bio": null,
      "followers": 2,
      "following": 1,
      "statuses": 5,
      "favorites": 0,
      "location": "",
      "site": "http://flickfaver.com",
      "avatar": null
    },
    "linkedin": {
      "handle": "in/leonidbugaev"
    },
    "googleplus": {
      "handle": null
    },
    "angellist": {
      "handle": "leonid-bugaev",
      "id": 61541,
      "bio": "Senior engineer at Granify.com",
      "blog": "http://buger.github.com",
      "site": "http://buger.github.com",
      "followers": 41,
      "avatar": "https://d1qb2nb5cznatu.cloudfront.net/users/61541-medium_jpg?1405474390"
    },
    "klout": {
      "handle": null,
      "score": null
    },
    "foursquare": {
      "handle": null
    },
    "aboutme": {
      "handle": "leonid.bugaev",
      "bio": null,
      "avatar": null
    },
    "gravatar": {
      "handle": "buger",
      "urls": [
      ],
      "avatar": "http://1.gravatar.com/avatar/f7c8edd577d13b8930d5522f28123510",
      "avatars": [
        {
          "url": "http://1.gravatar.com/avatar/f7c8edd577d13b8930d5522f28123510",
          "type": "thumbnail"
        }
      ]
    },
    "fuzzy": false
  },
  "company": null
}`)

type mediumPayload struct {
	Person  *_person `json:"person"`
	Company string   `json:"company"`
}

type _person struct {
	ID         string      `json:"id,omitempty"`
	Name       *name       `json:"name,omitempty"`
	Email      string      `json:"email,omitempty"`
	Gender     string      `json:"gender,omitempty"`
	Location   string      `json:"location,omitempty"`
	Geo        *geo        `json:"geo,omitempty"`
	Bio        string      `json:"bio,omitempty"`
	Site       string      `json:"site,omitempty"`
	Avatar     string      `json:"avatar,omitempty"`
	Employment *employment `json:"employment,omitempty"`
	Facebook   *facebook   `json:"facebook,omitempty"`
	Github     *github     `json:"github,omitempty"`
	Twitter    *twitter    `json:"twitter,omitempty"`
	Linkedin   *linkedin   `json:"linkedin,omitempty"`
	Googleplus *googleplus `json:"googleplus,omitempty"`
	Angellist  *angellist  `json:"angellist,omitempty"`
	Klout      *klout      `json:"klout,omitempty"`
	Foursquare *foursquare `json:"foursquare,omitempty"`
	Aboutme    *aboutme    `json:"aboutme,omitempty"`
	Gravatar   *gravatar   `json:"gravatar,omitempty"`
	Fuzzy      bool        `json:"fuzzy,omitempty"`
}

type name struct {
	FullName   string `json:"fullName,omitempty"`
	GivenName  string `json:"givenName,omitempty"`
	FamilyName string `json:"familyName,omitempty"`
}

type geo struct {
	City    string  `json:"city,omitempty"`
	State   string  `json:"state,omitempty"`
	Country string  `json:"country,omitempty"`
	Lat     float64 `json:"lat,omitempty"`
	Lng     float64 `json:"lng,omitempty"`
}

type employment struct {
	Name   string `json:"name,omitempty"`
	Title  string `json:"title,omitempty"`
	Domain string `json:"domain,omitempty"`
}

type facebook struct {
	Handle string `json:"handle,omitempty"`
}

type github struct {
	Handle    string `json:"handle,omitempty"`
	ID        int64  `json:"id,omitempty"`
	Avatar    string `json:"avatar,omitempty"`
	Company   string `json:"company,omitempty"`
	Blog      string `json:"blog,omitempty"`
	Followers int    `json:"followers,omitempty"`
	Following int    `json:"following,omitempty"`
}

type twitter struct {
	Handle    string `json:"handle,omitempty"`
	ID        int64  `json:"id,omitempty"`
	Bio       string `json:"bio,omitempty"`
	Followers int    `json:"followers,omitempty"`
	Following int    `json:"following,omitempty"`
	Statuses  int    `json:"statuses,omitempty"`
	Favorites int    `json:"favorites,omitempty"`
	Location  string `json:"location,omitempty"`
	Site      string `json:"site,omitempty"`
	Avatar    string `json:"avatar,omitempty"`
}
type linkedin struct {
	Handle string `json:"handle,omitempty"`
}

type googleplus struct {
	Handle string `json:"handle,omitempty"`
}

type angellist struct {
	Handle    string `json:"handle,omitempty"`
	ID        int64  `json:"id,omitempty"`
	Bio       string `json:"bio,omitempty"`
	Blog      string `json:"blog,omitempty"`
	Site      string `json:"site,omitempty"`
	Followers int    `json:"followers,omitempty"`
	Avatar    string `json:"avatar,omitempty"`
}

type klout struct {
	Handle string `json:"handle,omitempty"`
	Score  int    `json:"score,omitempty"`
}

type foursquare struct {
	Handle string `json:"handle,omitempty"`
}

type aboutme struct {
	Handle string `json:"handle,omitempty"`
	Bio    string `json:"bio,omitempty"`
	Avatar string `json:"avatar,omitempty"`
}

type gravatar struct {
	Handle  string   `json:"handle,omitempty"`
	Urls    []string `json:"urls,omitempty"`
	Avatar  string   `json:"avatar,omitempty"`
	Avatars []avatar `json:"avatars,omitempty"`
}

type avatar struct {
	URL  string `json:"url,omitempty"`
	Type string `json:"type,omitempty"`
}
