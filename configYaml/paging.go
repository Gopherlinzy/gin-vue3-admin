package configYaml

type Paging struct {
	PerPage         int    `mapstructure:"perpage" json:"perpage" yaml:"perpage"`
	UrlQueryPage    string `mapstructure:"url_query_page" json:"url_query_page" yaml:"url_query_page"`
	UrlQuerySort    string `mapstructure:"url_query_sort" json:"url_query_sort" yaml:"url_query_sort"`
	UrlQueryOrder   string `mapstructure:"url_query_order" json:"url_query_order" yaml:"url_query_order"`
	UrlQueryPerPage string `mapstructure:"url_query_per_page" json:"url_query_per_page" yaml:"url_query_per_page"`
}
