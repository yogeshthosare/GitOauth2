package main
  
import (
        "net/http"
	"net/url"
	"io/ioutil"
	"fmt"
	"os"
	"WebApp/config"
	helper "GitOauth2/helper"
	"gopkg.in/src-d/go-git.v4"
        ghttp "gopkg.in/src-d/go-git.v4/plumbing/transport/http"
)

func CallbackListener(w http.ResponseWriter, r *http.Request) {
        code := r.URL.Query().Get("code")
	fmt.Println("Code retrieved : ", code)
        fmt.Println("Retrieving access token now...")
	token, err := helper.RetrieveBearerToken(code)
  	if err != nil {
    		fmt.Println("Filed to retrieve token")
  	}

//	Initiate actual git clone using access token
	cwd, err := os.Getwd(); 
	if err != nil {
        	fmt.Println("Failed to get current working directory")
    	}

	dir, err := ioutil.TempDir(cwd, "cloned-repo")
	if err != nil {
        	fmt.Println("Failure while creating temp directory")
	}

//	For deleting cloned repo, while programme testing we can uncomment 
//	defer os.RemoveAll(dir)
	
	gr, err := git.PlainClone(dir, false, &git.CloneOptions{
		URL:    config.OAuthConfig.RepoEndpoint,
		Auth: &ghttp.BasicAuth{
			Username: "yogesh123", // anything except an empty string
        		Password: token.AccessToken,
    		},
		Progress: os.Stdout,
	}) 
	if err != nil{
		fmt.Println("#Failed to clone with err: ", err)
		fmt.Fprintln(w, "**Failed to clone with err: "+err.Error())
		return
	}
	
	ref, err := gr.Head() 
	if err != nil {
		fmt.Println("Failed to get with ref: ", err)
		fmt.Fprintln(w, "Failed to get with ref: "+err.Error())
		return
	}

        // ... retrieving the commit object
        commit, err := gr.CommitObject(ref.Hash()) 
	if err != nil {
		fmt.Println("Failed to retrieve commit object: ", err)
		fmt.Fprintln(w, "Failed to retrieve commit object: "+err.Error())
		return
	}else{
        	fmt.Println("Clone successful, check at location: ", dir)
        	fmt.Fprintln(w, "Clone successful, check at location: "+ dir)
	}

        fmt.Println(commit)
	fmt.Fprintln(w, commit)
}

func InitiateCloneRepo(w http.ResponseWriter, r *http.Request) {
	var Url *url.URL
        authorizeEndpoint := config.OAuthConfig.AuthorizeEndpoint
        Url, err := url.Parse(authorizeEndpoint)
        if err != nil {
                fmt.Println("Error parsing url :", err.Error())
        }

        parameters := url.Values{}
        parameters.Add("client_id", config.OAuthConfig.ClientId)
        parameters.Add("redirect_uri", config.OAuthConfig.RedirectUri)
        Url.RawQuery = parameters.Encode()

	http.Redirect(w, r, Url.String(), http.StatusTemporaryRedirect)
}

func main() {
	http.HandleFunc("/clone-repo", InitiateCloneRepo)
	http.HandleFunc("/oauth/redirect", CallbackListener)
	fmt.Println("App listening on localhost port:", config.OAuthConfig.Port)
	http.ListenAndServe(config.OAuthConfig.Port, nil)
}
