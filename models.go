package patreon

import (
	"net/url"
	"time"
)

type (
	Webhook struct {
		Data     *Data  `json:"data,omitempty"`
		Included []Data `json:"included,omitempty"`
	}

	User struct {
		Data     *Data  `json:"data,omitempty"`
		Included *Data  `json:"included,omitempty"`
		Links    *Links `json:"links,omitempty"`
	}

	Campaign struct {
		Data     *Data  `json:"data,omitempty"`
		Included []Data `json:"included,omitempty"`
		Links    *Links `json:"links,omitempty"`
	}

	Pledge struct {
		Data     []Data `json:"data,omitempty"`
		Included []Data `json:"included,omitempty"`
		Links    *Links `json:"links,omitempty"`
		Meta     *Meta  `json:"meta,omitempty"`
	}

	// Data contains information of requested type of data.
	Data struct {
		ID            int            `json:"id,string,omitempty"`
		Type          string         `json:"type,omitempty"`
		Attributes    *Attributes    `json:"attributes,omitempty"`
		Relationships *Relationships `json:"relationships,omitempty"`
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
		Facebook                      *URL               `json:"facebook,omitempty"`
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
		ImageSmallURL                 *URL               `json:"image_small_url,string,omitempty"`
		ImageURL                      *URL               `json:"image_url,string,omitempty"`
		IsMonthly                     bool               `json:"is_monthly,omitempty"`
		IsNSFW                        bool               `json:"is_nsfw,omitempty"`
		IsPaused                      bool               `json:"is_paused,omitempty"`
		LastName                      string             `json:"last_name,omitempty"`
		LikeCount                     int                `json:"like_count,omitempty"`
		MainVideoEmbed                string             `json:"main_video_embed,omitempty"`
		MainVideoURL                  *URL               `json:"main_video_url,string,omitempty"`
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
		ThanksVideoURL                *URL               `json:"thanks_video_url,string,omitempty"`
		ThumbURL                      *URL               `json:"thumb_url,omitempty"`
		TotalHistoricalAmountCents    int                `json:"total_historical_amount_cents,omitempty"`
		Twitch                        *URL               `json:"twitch,omitempty"`
		Twitter                       string             `json:"twitter,omitempty"`
		URL                           *URL               `json:"url,omitempty"`
		Vanity                        string             `json:"vanity,omitempty"`
		YouTube                       *URL               `json:"youtube,omitempty"`
	}

	// Links contains data about current
	Links struct {
		Self    *URL `json:"self,omitempty"`
		Related *URL `json:"related,omitempty"`
	}

	// DateTime contains date/time data.
	DateTime struct{ time.Time }

	// DateTime contains url data.
	URL struct{ *url.URL }

	Errors struct {
		Errors []Error `json:"errors"`
	}

	// Error contains information about reasons of invalid request.
	Error struct {
		ID       string `json:"id,omitempty"`
		Code     int    `json:"code,omitempty"`
		CodeName string `json:"code_name,omitempty"`
		Status   int    `json:"status,string,omitempty"`
		Title    string `json:"title,omitempty"`
		Detail   string `json:"detail,omitempty"`
	}

	// Relationships contains other types of data which relevant for current response.
	Relationships struct {
		Campaign *User   `json:"campaign,omitempty"`
		Creator  *User   `json:"creator,omitempty"`
		Goals    *Pledge `json:"goals,omitempty"`
		Pledges  *Pledge `json:"pledges,omitempty"`
		Rewards  *Pledge `json:"rewards,omitempty"`
	}

	// SocialConnections contains data about connections to other platforms.
	SocialConnections struct {
		DeviantArt *SocialConnection `json:"deviantart,omitempty"`
		Discord    *SocialConnection `json:"discord,omitempty"`
		Facebook   *SocialConnection `json:"facebook,omitempty"`
		Spotify    *SocialConnection `json:"spotify,omitempty"`
		Twitch     *SocialConnection `json:"twitch,omitempty"`
		Twitter    *SocialConnection `json:"twitter,omitempty"`
		YouTube    *SocialConnection `json:"youtube,omitempty"`
	}

	// SocialConnection contains data about single social connection: user ID and his scopes.
	SocialConnection struct {
		UserID interface{} `json:"user_id,omitempty"`
		Scopes []string    `json:"scopes,omitempty"`
	}

	Meta struct {
		Count int `json:"count,omitempty"`
	}
)
