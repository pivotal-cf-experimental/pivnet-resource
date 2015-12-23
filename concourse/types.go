package concourse

type Request struct {
	Source  Source  `json:"source"`
	Version Release `json:"version"`
}

type Response []Release

// TODO: Rename to Version
type Release struct {
	ProductVersion string `json:"product_version"`
}

type Source struct {
	APIToken        string `json:"api_token"`
	ProductName     string `json:"product_name"`
	AccessKeyID     string `json:"access_key_id"`
	SecretAccessKey string `json:"secret_access_key"`
}

type InRequest struct {
	Source  Source  `json:"source"`
	Version Release `json:"version"`
}

type InResponse struct {
	Version  Release    `json:"version"`
	Metadata []Metadata `json:"metadata,omitempty"`
}

type Metadata struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

type OutRequest struct {
	Params OutParams `json:"params"`
	Source Source    `json:"source"`
}

type OutParams struct {
	FileGlob        string `json:"file_glob"`
	FilepathPrefix  string `json:"s3_filepath_prefix"`
	VersionFile     string `json:"version_file"`
	ReleaseTypeFile string `json:"release_type_file"`
	ReleaseDateFile string `json:"release_date_file"`
	EulaSlugFile    string `json:"eula_slug_file"`
}

type OutResponse struct {
	Version  Release  `json:"version"`
	Metadata []string `json:"metadata,omitempty"`
}
