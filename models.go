package patreon

import "time"

type (
	// Response contains one type of requested information.
	Response struct {
		Data     *Data  `json:"data"`
		Included *Data  `json:"included"`
		Links    *Links `json:"links"`
	}

	// ResponseList contains arrays of many types of requested information.
	ResponseList struct {
		Data     []Data `json:"data"`
		Included []Data `json:"included"`
		Links    *Links `json:"links"`
	}

	// Data contains information of requested type of data.
	Data struct {
		ID            int            `json:"id,string"`
		Type          string         `json:"type"`
		Attributes    *Attributes    `json:"attributes"`
		Relationships *Relationships `json:"relationships"`
	}

	// Attributes contains meta-information about data.
	Attributes struct {
		About                         string             `json:"about,omitempty"`
		AmountCents                   int                `json:"amount_cents,omitempty"`
		CommentCount                  int                `json:"comment_count,omitempty"`
		Created                       *DateTime          `json:"created,string,omitempty"`
		CreatedAt                     *DateTime          `json:"created_at,string,omitempty"`
		CreationCount                 int                `json:"creation_count,omitempty"`
		CreationName                  string             `json:"creation_name,omitempty"`
		DeclinedSince                 *DateTime          `json:"declined_since,string,omitempty"`
		DiscordID                     int64              `json:"discord_id,string,omitempty"`
		Email                         string             `json:"email,omitempty"`
		Facebook                      string             `json:"facebook,omitempty"`
		FacebookID                    int64              `json:"facebook_id,string,omitempty"`
		FirstName                     string             `json:"first_name,omitempty"`
		FullName                      string             `json:"full_name,omitempty"`
		HasShippingAddress            bool               `json:"has_shipping_address,omitempty"`
		Gender                        int                `json:"gender,omitempty"`
		IsDeleted                     bool               `json:"is_deleted,omitempty"`
		IsEmailVerified               bool               `json:"is_email_verified,omitempty"`
		IsNuked                       bool               `json:"is_nuked,omitempty"`
		IsSuspended                   bool               `json:"is_suspended,omitempty"`
		SocialConnections             *SocialConnections `json:"social_connections,omitempty"`
		HasPassword                   bool               `json:"has_password,omitempty"`
		ImageSmallURL                 string             `json:"image_small_url,omitempty"`
		ImageURL                      string             `json:"image_url,omitempty"`
		IsMonthly                     bool               `json:"is_monthly,omitempty"`
		IsNSFW                        bool               `json:"is_nsfw,omitempty"`
		IsPaused                      bool               `json:"is_paused,omitempty"`
		LastName                      string             `json:"last_name,omitempty"`
		LikeCount                     int                `json:"like_count,omitempty"`
		MainVideoEmbed                string             `json:"main_video_embed,omitempty"`
		MainVideoURL                  string             `json:"main_video_url,omitempty"`
		OneLiner                      string             `json:"one_liner,omitempty"`
		OutstandingPaymentAmountCents int                `json:"outstanding_payment_amount_cents,omitempty"`
		PatronCount                   int                `json:"patron_count,omitempty"`
		PatronPaysFees                bool               `json:"patron_pays_fees,omitempty"`
		PayPerName                    string             `json:"pay_per_name,omitempty"`
		PledgeCapCents                int                `json:"pledge_cap_cents,omitempty"`
		PledgeSum                     int                `json:"pledge_sum,omitempty"`
		PledgeURL                     string             `json:"pledge_url,omitempty"`
		PublishedAt                   *DateTime          `json:"published_at,string,omitempty"`
		Summary                       string             `json:"summary,omitempty"`
		ThanksEmbed                   string             `json:"thanks_embed,omitempty"`
		ThanksMsg                     string             `json:"thanks_msg,omitempty"`
		ThanksVideoURL                string             `json:"thanks_video_url,omitempty"`
		ThumbURL                      string             `json:"thumb_url,omitempty"`
		TotalHistoricalAmountCents    int                `json:"total_historical_amount_cents,omitempty"`
		Twitch                        string             `json:"twitch,omitempty"`
		Twitter                       string             `json:"twitter,omitempty"`
		URL                           string             `json:"url,omitempty"`
		Vanity                        string             `json:"vanity,omitempty"`
		YouTube                       string             `json:"youtube,omitempty"`
	}

	// Links contains data about current
	Links struct {
		Self    string `json:"self"`
		Related string `json:"related"`
	}

	// DateTime contains date/time data.
	DateTime struct {
		time.Time
	}

	// Error contains information about reasons of invalid request.
	Error struct {
		ID       string `json:"id"`
		Code     int    `json:"code"`
		CodeName string `json:"code_name"`
		Status   int    `json:"status,string"`
		Title    string `json:"title"`
		Detail   string `json:"detail"`
	}

	// Relationships contains other types of data which relevant for current
	// response.
	Relationships struct {
		Campaign *Response     `json:"campaign"`
		Creator  *Response     `json:"creator"`
		Goals    *ResponseList `json:"goals"`
		Pledges  *ResponseList `json:"pledges"`
		Rewards  *ResponseList `json:"rewards"`
	}

	// SocialConnections contains data about connections to other platforms.
	SocialConnections struct {
		DeviantArt *SocialConnection `json:"deviantart"`
		Discord    *SocialConnection `json:"discord"`
		Facebook   *SocialConnection `json:"facebook"`
		Spotify    *SocialConnection `json:"spotify"`
		Twitch     *SocialConnection `json:"twitch"`
		Twitter    *SocialConnection `json:"twitter"`
		YouTube    *SocialConnection `json:"youtube"`
	}

	// SocialConnection contains data about single social connection: user ID and
	// his scopes.
	SocialConnection struct {
		UserID string   `json:"user_id"`
		Scopes []string `json:"scopes"`
	}
)

const (
	IncludeRewards = "rewards"
	IncludeCreator = "creator"
	IncludeGoals   = "goals"
	IncludePledge  = "pledge"
	IncludePledges = "pledges"
)
