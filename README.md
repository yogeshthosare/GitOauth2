# GitOauth2
OAuth 2.0 based implementation to authenticate the user with Github


Objective - To write a simple web application that will connect to Github and clone specified repository.
App uses Oauth2.0 based implementation to authenticate user with Github, App then clones the repository on behalf of the user.


Application flow - 

1. App requests Github for Read-access on the repositories.

2. Github asks for user credentials. 

3. User enters credentials.

4. Github authorizes and gives token to App. 

5. App clones the repository on behalf of the user.



Read more about building GitHub Oauth Apps from below links.
https://docs.github.com/en/developers/apps/authorizing-oauth-apps
