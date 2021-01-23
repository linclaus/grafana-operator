package gografana

import "time"

// Board represents Grafana dashboard.
type Board struct {
	ID           uint     `json:"id,omitempty"`
	UID          string   `json:"uid"`
	Title        string   `json:"title"`
	Tags         []string `json:"tags"`
	Editable     bool     `json:"editable"`
	HideControls bool     `json:"hideControls"`
	Style        string   `json:"style"`
	Timezone     string   `json:"timezone"`
	Version      uint     `json:"version"`
	IsStarred    bool     `json:"isStarred"`
	FolderId     uint     `json:"folderId"`
	FolderUid    string   `json:"folderUid"`
	FolderTitle  string   `json:"folderTitle"`
	FolderUrl    string   `json:"folderUrl"`
	Url          string   `json:"url,omitempty"` //TODO: ??
	Rows         []*Row   `json:"rows"`
}

type CreateDashboardRequest struct {
	Board Board `json:"dashboard"`
	//The id of the folder to save the dashboard in.
	FolderId uint `json:"folderId"`
	//Set to true if you want to overwrite existing dashboard with newer version, same dashboard title in folder or same dashboard uid.
	Overwrite bool `json:"overwrite,omitempty"`
	//Set a commit message for the version history.
	Message string `json:"message"`
}

type CreateFolderRequest struct {
	UID   string `json:"uid"`
	Title string `json:"title"`
}

type CreateFolderResponse struct {
	ID        int       `json:"id"`
	UID       string    `json:"uid"`
	Title     string    `json:"title"`
	URL       string    `json:"url"`
	HasAcl    bool      `json:"hasAcl"`
	CanSave   bool      `json:"canSave"`
	CanEdit   bool      `json:"canEdit"`
	CanAdmin  bool      `json:"canAdmin"`
	CreatedBy string    `json:"createdBy"`
	Created   time.Time `json:"created"`
	UpdatedBy string    `json:"updatedBy"`
	Updated   time.Time `json:"updated"`
	Version   int       `json:"version"`
}

type CreateAPIKeyResponse struct {
	Name string `json:"name"`
	Key  string `json:"key"`
}

/*
Status Codes:
---------------
200 – Created
400 – Errors (invalid json, missing or invalid fields, etc)
401 – Unauthorized
403 – Access denied
412 – Precondition failed

The 412 status code is used for explaing that you cannot create the dashboard and why. There can be different reasons for this:

The dashboard has been changed by someone else, status=version-mismatch
A dashboard with the same name in the folder already exists, status=name-exists
A dashboard with the same uid already exists, status=name-exists
The dashboard belongs to plugin <plugin title>, status=plugin-dashboard
*/
type CreateDashboardResponse struct {
	ID      uint   `json:"id,omitempty"`
	UID     string `json:"uid"`
	Url     string `json:"url"`
	Status  string `json:"status"`
	Version uint   `json:"version"`
	Message string `json:"message,omitempty"`
}

type Panel_5_0 struct {
	AliasColors struct {
	} `json:"aliasColors"`
	Alert *struct {
		AlertRuleTags struct {
		} `json:"alertRuleTags"`
		Conditions []struct {
			Evaluator struct {
				Params []int64 `json:"params"`
				Type   string  `json:"type"`
			} `json:"evaluator"`
			Operator struct {
				Type string `json:"type"`
			} `json:"operator"`
			Query struct {
				Params []string `json:"params"`
			} `json:"query"`
			Reducer struct {
				Params []interface{} `json:"params"`
				Type   string        `json:"type"`
			} `json:"reducer"`
			Type string `json:"type"`
		} `json:"conditions"`
		ExecutionErrorState string `json:"executionErrorState"`
		For                 string `json:"for"`
		Frequency           string `json:"frequency"`
		Handler             int    `json:"handler"`
		Message             string `json:"message"`
		Name                string `json:"name"`
		NoDataState         string `json:"noDataState"`
		Notifications       []struct {
			ID int `json:"id"`
		} `json:"notifications"`
	} `json:"alert,omitempty"`
	Bars       bool   `json:"bars"`
	DashLength int    `json:"dashLength"`
	Dashes     bool   `json:"dashes"`
	Datasource string `json:"datasource"`
	Fill       int    `json:"fill"`
	GridPos    struct {
		H int `json:"h"`
		W int `json:"w"`
		X int `json:"x"`
		Y int `json:"y"`
	} `json:"gridPos,omitempty"`
	ID      int `json:"id"`
	Options struct {
		Orientation   string `json:"orientation"`
		ReduceOptions struct {
			Calcs  []string `json:"calcs"`
			Fields string   `json:"fields"`
			Values bool     `json:"values"`
		} `json:"reduceOptions"`
		ShowThresholdLabels  bool `json:"showThresholdLabels"`
		ShowThresholdMarkers bool `json:"showThresholdMarkers"`
	} `json:"options"`
	FieldConfig struct {
		Defaults struct {
			Unit     string `json:"unit"`
			Decimals int    `json:"decimals"`
		} `json:"defaults"`
	} `json:"fieldConfig"`
	Legend struct {
		Avg          bool `json:"avg"`
		Current      bool `json:"current"`
		Max          bool `json:"max"`
		Min          bool `json:"min"`
		Show         bool `json:"show"`
		Total        bool `json:"total"`
		Values       bool `json:"values"`
		AlignAsTable bool `json:"alignAsTable"`
	} `json:"legend,omitempty"`
	Lines           bool          `json:"lines"`
	Linewidth       int           `json:"linewidth"`
	Links           []interface{} `json:"links"`
	NullPointMode   string        `json:"nullPointMode"`
	Percentage      bool          `json:"percentage"`
	Pointradius     int           `json:"pointradius"`
	Points          bool          `json:"points"`
	Renderer        string        `json:"renderer"`
	SeriesOverrides []interface{} `json:"seriesOverrides"`
	SpaceLength     int           `json:"spaceLength"`
	Stack           bool          `json:"stack"`
	SteppedLine     bool          `json:"steppedLine"`
	Targets         []struct {
		Expr           string `json:"expr"`
		Format         string `json:"format"`
		Instant        bool   `json:"instant"`
		IntervalFactor int    `json:"intervalFactor"`
		LegendFormat   string `json:"legendFormat"`
		RefID          string `json:"refId"`
	} `json:"targets"`
	Thresholds []interface{} `json:"thresholds"`
	TimeFrom   interface{}   `json:"timeFrom"`
	TimeShift  interface{}   `json:"timeShift"`
	Title      string        `json:"title"`
	DataFormat string        `json:"dataFormat"`
	Tooltip    struct {
		Shared    bool   `json:"shared"`
		Sort      int    `json:"sort"`
		ValueType string `json:"value_type"`
	} `json:"tooltip,omitempty"`
	Transparent bool   `json:"transparent"`
	Type        string `json:"type"`
	Xaxis       struct {
		Buckets interface{}   `json:"buckets"`
		Mode    string        `json:"mode"`
		Name    interface{}   `json:"name"`
		Show    bool          `json:"show"`
		Values  []interface{} `json:"values"`
	} `json:"xaxis,omitempty"`
	Yaxes []struct {
		Format   string      `json:"format"`
		Label    interface{} `json:"label"`
		LogBase  int         `json:"logBase"`
		Max      interface{} `json:"max"`
		Min      interface{} `json:"min"`
		Show     bool        `json:"show"`
		Decimals int         `json:"decimals"`
	} `json:"yaxes,omitempty"`
	Color struct {
		CardColor   string  `json:"cardColor"`
		ColorScale  string  `json:"colorScale"`
		ColorScheme string  `json:"colorScheme"`
		Exponent    float64 `json:"exponent"`
		Mode        string  `json:"mode"`
	} `json:"color"`
}

type Row struct {
	Title     string      `json:"title"`
	ShowTitle bool        `json:"showTitle"`
	Collapse  bool        `json:"collapse"`
	Editable  bool        `json:"editable"`
	Height    string      `json:"height"`
	Panels    []Panel_5_0 `json:"panels"`
}

type GetDashboardByUIdResponse struct {
	Meta struct {
		Type        string    `json:"type"`
		CanSave     bool      `json:"canSave"`
		CanEdit     bool      `json:"canEdit"`
		CanAdmin    bool      `json:"canAdmin"`
		CanStar     bool      `json:"canStar"`
		Slug        string    `json:"slug"`
		URL         string    `json:"url"`
		Expires     time.Time `json:"expires"`
		Created     time.Time `json:"created"`
		Updated     time.Time `json:"updated"`
		UpdatedBy   string    `json:"updatedBy"`
		CreatedBy   string    `json:"createdBy"`
		Version     int       `json:"version"`
		HasAcl      bool      `json:"hasAcl"`
		IsFolder    bool      `json:"isFolder"`
		FolderID    int       `json:"folderId"`
		FolderTitle string    `json:"folderTitle"`
		FolderURL   string    `json:"folderUrl"`
	} `json:"meta"`
	Dashboard Board `json:"dashboard"`
}

type DataSource struct {
	ID          int                    `json:"id"`
	OrgID       int                    `json:"orgId"`
	Name        string                 `json:"name"`
	Type        string                 `json:"type"`
	TypeLogoURL string                 `json:"typeLogoUrl"`
	Access      string                 `json:"access"`
	URL         string                 `json:"url"`
	Password    string                 `json:"password"`
	User        string                 `json:"user"`
	Database    string                 `json:"database"`
	BasicAuth   bool                   `json:"basicAuth"`
	IsDefault   bool                   `json:"isDefault"`
	JSONData    map[string]interface{} `json:"jsonData"`
	ReadOnly    bool                   `json:"readOnly"`
}

type CreateDataSourceResponse struct {
	DataSource *DataSource `json:"datasource"`
	ID         int         `json:"id"`
	Message    string      `json:"message"`
	Name       string      `json:"name"`
}

type CreateNotificationChannelResponse struct {
	ID           int    `json:"id"`
	UID          string `json:"uid"`
	Name         string `json:"name"`
	Type         string `json:"type"`
	IsDefault    bool   `json:"isDefault"`
	SendReminder bool   `json:"sendReminder"`
	Settings     struct {
		Addresses string `json:"addresses"`
	} `json:"settings"`
	Created time.Time `json:"created"`
	Updated time.Time `json:"updated"`
}

type APIKey struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Role string `json:"role"`
}

type Folder struct {
	ID        int       `json:"id"`
	UID       string    `json:"uid"`
	Title     string    `json:"title"`
	URL       string    `json:"url"`
	HasACL    bool      `json:"hasAcl"`
	CanSave   bool      `json:"canSave"`
	CanEdit   bool      `json:"canEdit"`
	CanAdmin  bool      `json:"canAdmin"`
	CreatedBy string    `json:"createdBy"`
	Created   time.Time `json:"created"`
	UpdatedBy string    `json:"updatedBy"`
	Updated   time.Time `json:"updated"`
	Version   int       `json:"version"`
}

type NotificationChannel struct {
	ID                    int       `json:"id"`
	Name                  string    `json:"name"`
	Type                  string    `json:"type"`
	IsDefault             bool      `json:"isDefault"`
	SendReminder          bool      `json:"sendReminder"`
	DisableResolveMessage bool      `json:"disableResolveMessage"`
	Frequency             string    `json:"frequency"`
	Created               time.Time `json:"created"`
	Updated               time.Time `json:"updated"`
	Settings              struct {
		Addresses   string `json:"addresses"`
		AutoResolve bool   `json:"autoResolve"`
		HTTPMethod  string `json:"httpMethod"`
		UploadImage bool   `json:"uploadImage"`
	} `json:"settings"`
}
