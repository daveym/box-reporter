package box

// State of returned items
type State string

// Sort - Sort options for returned items
type Sort string

// ItemStatus - Status of retrieved item
type ItemStatus int

// ContentType - Article video or image
type ContentType string

// DetailType - Simple or complete detail
type DetailType string

// FavoriteFilter - Filter by favourite
type FavoriteFilter string

// ItemMediaAttachment - Is Media attached to the Pocket Item.
type ItemMediaAttachment int

// RequestToken - Obtain a code from Pocket
type RequestToken struct {
	Code string `json:"code"`
}

// AuthenticationResponse response from Pocket
type AuthenticationResponse struct {
	Code  string // Code is request Token
	State string
}

// AuthorisationResponse response from Pocket
type AuthorisationResponse struct {
	AccessToken string `json:"access_token"`
	Username    string `json:"username"`
}

// RetrieveRequest - response from Pocket
type RetrieveRequest struct {
	ConsumerKey string         `json:"consumer_key"`
	AccessToken string         `json:"access_token"`
	State       State          `json:"state,omitempty"`
	Favorite    FavoriteFilter `json:"favorite,omitempty"`
	Tag         string         `json:"tag,omitempty"`
	ContentType ContentType    `json:"contentType,omitempty"`
	Sort        Sort           `json:"sort,omitempty"`
	DetailType  DetailType     `json:"detailType,omitempty"`
	Search      string         `json:"search,omitempty"`
	Domain      string         `json:"domain,omitempty"`
	Since       int            `json:"since,omitempty"`
	Count       int            `json:"count,omitempty"`
	Offset      int            `json:"offset,omitempty"`
}

// RetrieveResponse - List of items retrieved from Pocket
type RetrieveResponse struct {
	Status   int
	Complete int
	List     map[string]Item
	Since    int
}

// Item - Individual Pocket Item
type Item struct {
	ItemID        int        `json:"item_id,string"`
	ResolvedID    int        `json:"resolved_id,string"`
	GivenURL      string     `json:"given_url"`
	GivenTitle    string     `json:"given_title"`
	Favorite      int        `json:",string"`
	Status        ItemStatus `json:",string"`
	SortID        int        `json:"sort_id"`
	ResolvedTitle string     `json:"resolved_title"`
	ResolvedURL   string     `json:"resolved_url"`
	Excerpt       string
	IsArticle     int                 `json:"is_article,string"`
	HasImage      ItemMediaAttachment `json:"has_image,string"`
	HasVideo      ItemMediaAttachment `json:"has_video,string"`
	WordCount     int                 `json:"word_count,string"`
}

// ModifyRequest - Send actions to pocket, i.e. add, archive, readd, favourite, unfavourite, delete
type ModifyRequest struct {
	ConsumerKey string    `json:"consumer_key"`
	AccessToken string    `json:"access_token"`
	Actions     []*Action `json:"actions"`
}

// ModifyResponse - Result of modifications against pocket.
type ModifyResponse struct {
	Status        int
	ActionResults []bool `json:"action_results"`
}

// Action represents one action in a bulk modify requests.
type Action struct {
	Action string `json:"action"`
	ItemID int    `json:"item_id,string"`
}
