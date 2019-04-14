# OAuth Credentials

## Development
1. Visit https://console.cloud.google.com/apis/credentials
2. Click "Create credentials"
3. Select OAuth client ID
4. Select Web application
5. Set Authorized JavaScript origins to http://localhost:8080 (and optionally http://localhost:4200)
6. Set Authorized redirect URIs to http://localhost:8080/oauth2redirect (and optionally http://localhost:4200/oauth2redirect)
7. Pick a name (e.g. "Shuffleboard Club Dev")
8. Click Create
9. Click Download JSON
10. `cp ~/Downloads/client_secret_CLIENT-ID.apps.googleusercontent.com.json ./.secrets/oauth2-secret-dev.json`
