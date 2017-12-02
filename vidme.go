package vidme

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
	"vidme/httputils"
)

type VidmeApi struct {
	Status bool `json:"status"`
	Video  struct {
		VideoID               string      `json:"video_id"`
		URL                   string      `json:"url"`
		OriginalSize          string      `json:"original_size"`
		FullURL               string      `json:"full_url"`
		EmbedURL              string      `json:"embed_url"`
		UserID                string      `json:"user_id"`
		State                 string      `json:"state"`
		Title                 string      `json:"title"`
		Description           interface{} `json:"description"`
		Duration              int         `json:"duration"`
		Height                int         `json:"height"`
		Width                 int         `json:"width"`
		DateCreated           string      `json:"date_created"`
		DateStored            string      `json:"date_stored"`
		DateCompleted         string      `json:"date_completed"`
		CommentCount          int         `json:"comment_count"`
		ViewCount             int         `json:"view_count"`
		ShareCount            int         `json:"share_count"`
		RepostCount           int         `json:"repost_count"`
		Version               int         `json:"version"`
		Nsfw                  bool        `json:"nsfw"`
		Moderated             bool        `json:"moderated"`
		Thumbnail             string      `json:"thumbnail"`
		ThumbnailURL          string      `json:"thumbnail_url"`
		ThumbnailGif          interface{} `json:"thumbnail_gif"`
		ThumbnailGifURL       interface{} `json:"thumbnail_gif_url"`
		ThumbnailAi           string      `json:"thumbnail_ai"`
		Storyboard            interface{} `json:"storyboard"`
		Score                 int         `json:"score"`
		LikesCount            int         `json:"likes_count"`
		ChannelID             interface{} `json:"channel_id"`
		Source                interface{} `json:"source"`
		Private               bool        `json:"private"`
		Scheduled             bool        `json:"scheduled"`
		SubscribedOnly        bool        `json:"subscribed_only"`
		DatePublished         string      `json:"date_published"`
		Latitude              int         `json:"latitude"`
		Longitude             int         `json:"longitude"`
		PlaceID               interface{} `json:"place_id"`
		PlaceName             interface{} `json:"place_name"`
		Colors                string      `json:"colors"`
		RedditLink            interface{} `json:"reddit_link"`
		YoutubeOverrideSource interface{} `json:"youtube_override_source"`
		Complete              string      `json:"complete"`
		CompleteURL           string      `json:"complete_url"`
		HotScore              float64     `json:"hot_score"`
		User                  struct {
			UserID         string      `json:"user_id"`
			Username       string      `json:"username"`
			FullURL        string      `json:"full_url"`
			Avatar         string      `json:"avatar"`
			AvatarURL      string      `json:"avatar_url"`
			AvatarAi       string      `json:"avatar_ai"`
			Cover          string      `json:"cover"`
			CoverURL       string      `json:"cover_url"`
			CoverAi        string      `json:"cover_ai"`
			Displayname    string      `json:"displayname"`
			FollowerCount  int         `json:"follower_count"`
			LikesCount     string      `json:"likes_count"`
			VideoCount     int         `json:"video_count"`
			VideoViews     string      `json:"video_views"`
			VideosScores   int         `json:"videos_scores"`
			CommentsScores int         `json:"comments_scores"`
			Bio            string      `json:"bio"`
			Enabled        string      `json:"enabled"`
			GaID           interface{} `json:"ga_id"`
			Vip            bool        `json:"vip"`
		} `json:"user"`
		Formats []interface{} `json:"formats"`
	} `json:"video"`
	Progress struct {
		Progress int `json:"progress"`
	} `json:"progress"`
	Watchers struct {
		Total     int           `json:"total"`
		Countries []interface{} `json:"countries"`
	} `json:"watchers"`
}

func Get(count int) {
	url := "https://api.vid.me/video/" + strconv.Itoa(count)
	fmt.Println(url)
	body := httputils.HttpRequest("GET", url, "", nil)
	// fmt.Printf("%s", body)

	var s = new(VidmeApi)

	err := json.Unmarshal(body, &s)
	if err != nil {
		fmt.Println("whoops:", err)
	}
	if s.Video.CompleteURL != "" {
		b, _ := json.Marshal(s)
		ioutil.WriteFile("videos/"+strconv.Itoa(count)+".json", b, 0644)
		httputils.Download(s.Video.Complete, s.Video.CompleteURL)
	}
}

/*
{
    "status": true,
    "video": {
        "video_id": "40",
        "url": "sqI",
        "original_size": "0",
        "full_url": "https:\/\/vid.me\/sqI",
        "embed_url": "https:\/\/vid.me\/e\/sqI",
        "user_id": "9532",
        "state": "success",
        "title": "This is viddme",
        "description": null,
        "duration": 8,
        "height": 480,
        "width": 360,
        "date_created": "2013-12-12 04:02:12",
        "date_stored": "2013-12-12 04:02:12",
        "date_completed": "2013-12-12 04:02:33",
        "comment_count": 0,
        "view_count": 55,
        "share_count": 0,
        "repost_count": 0,
        "version": 1,
        "nsfw": false,
        "moderated": true,
        "thumbnail": "thumbnails\/40.jpg?v1r1410313240",
        "thumbnail_url": "https:\/\/d1wst0behutosd.cloudfront.net\/thumbnails\/40.jpg?v1r1410313240",
        "thumbnail_gif": null,
        "thumbnail_gif_url": null,
        "thumbnail_ai": "https:\/\/d1ckdm8qo2u5d0.cloudfront.net\/ai\/thumbnails\/40-{}.jpg?v1r1410313240",
        "storyboard": null,
        "score": 2,
        "likes_count": 2,
        "channel_id": null,
        "source": null,
        "private": true,
        "scheduled": false,
        "subscribed_only": false,
        "date_published": "2013-12-12 04:02:33",
        "latitude": 0,
        "longitude": 0,
        "place_id": null,
        "place_name": null,
        "colors": "#494047,#eadfd5,#c6b3a7,#a89990,#8a7a74,#96a1bc",
        "reddit_link": null,
        "youtube_override_source": null,
        "complete": "videos\/40.mp4",
        "complete_url": "https:\/\/d1wst0behutosd.cloudfront.net\/videos\/40.mp4?Expires=1512208095
        &Signature=dTG1MPyi4WuR2BwCK0VhfyLIYvbw3P49E5lAbMKmxx5~BzaWiLCsTPsQ-l4QWM4dRRqb5krCtHj2rgUA78fKueYMNcuqsRisx7o2jNe7APTDwT
        -aZ1vywllSWvRTmj5QJxj945gqL8f9B7deUb8p5F20DnE7Ciu7k6JIYLw-Iw7x8TLF8tx~1bgbSKvwZB~xdWojpSiGKoZfpFKerLl1tUCKY9g5u9IzLI8bh5i
        5AKARjYa6s1P1j4d7dX~N7FIf1Ee711a1jLOaxlLzs61-h2mhZ9oemUilweldsuZt6rXKD7vYx1mN4lFiTZyOsWm6qheEh5GHPsDCcP6oPHsoNg__
        &Key-Pair-Id=APKAJJ6WELAPEP47UKWQ",
        "hot_score": 1423.336875,
        "user": {
            "user_id": "9532",
            "username": "jbboehr",
            "full_url": "https:\/\/vid.me\/jbboehr",
            "avatar": "avatars\/9532.jpg?v4r1452565608",
            "avatar_url": "https:\/\/d1wst0behutosd.cloudfront.net\/avatars\/9532.jpg?v4r1452565608",
            "avatar_ai": "https:\/\/d1ckdm8qo2u5d0.cloudfront.net\/ai\/avatars\/9532-{}.jpg?v4r1452565608",
            "cover": "channel_covers\/9532.jpg?v1r1429289988",
            "cover_url": "https:\/\/d1wst0behutosd.cloudfront.net\/channel_covers\/9532.jpg?v1r1429289988",
            "cover_ai": "https:\/\/d1ckdm8qo2u5d0.cloudfront.net\/ai\/channel_covers\/9532-{}.jpg?v1r1429289988",
            "displayname": "John Boehr\ud83e\udd80",
            "follower_count": 134,
            "likes_count": "575",
            "video_count": 185,
            "video_views": "36957",
            "videos_scores": 271,
            "comments_scores": 83,
            "bio": "\ud83e\udd80",
            "enabled": "1",
            "ga_id": null,
            "vip": true
        },
        "formats": []
    },
    "progress": {
        "progress": 0
    },
    "watchers": {
        "total": 0,
        "countries": []
    }
}

*/
