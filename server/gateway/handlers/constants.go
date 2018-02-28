package handlers

const (
	headerContentType = "Content-Type"
	contentTypeJSON = "application/json"
	headerAccessControlAllowOrigin = "Access-Control-Allow-Origin"
	headerAccept      = "Accept"
)

const (
	//githubCurrentUserAPI is the URL for GitHub's current user API
	githubCurrentUserAPI = "https://api.github.com/user"

	//acceptGitHubV3JSON is the value you should include in
	//the Accept header when making requests to the GitHub API
	acceptGitHubV3JSON = "application/vnd.github.v3+json"
)
