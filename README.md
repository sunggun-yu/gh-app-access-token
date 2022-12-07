# gh-app-access-token-cli

Simple Cli tool for operating Github App Installation access token:

- create an access token
- revoke an access token

,with simply wrapping cli functionality on top of <https://github.com/bradleyfalzon/ghinstallation> and <https://github.com/google/go-github>

## Installation

brew:

```bash
brew install sunggun-yu/tap/gh-app-access-token
```

go install:

```bash
go install github.com/sunggun-yu/gh-app-access-token@<version>
```

docker:

```bash
docker pull ghcr.io/sunggun-yu/gh-app-access-token:latest
```

shell script:

```bash
curl -sfL https://raw.githubusercontent.com/sunggun-yu/gh-app-access-token/main/install.sh | sh
```

it place `gh-app-access-token` in `/tmp` directory. please set `INSTALL_PATH` env var to change directory

```bash
export INSTALL_PATH=/<some-dir>
curl -sfL https://raw.githubusercontent.com/sunggun-yu/gh-app-access-token/main/install.sh | sh
```

or,

```bash
curl -sfL https://raw.githubusercontent.com/sunggun-yu/gh-app-access-token/main/install.sh | INSTALL_PATH=/<some-dir> sh
```

## Usage

### Generate a Github App access token

```bash
# generate the Github App access token
gh-app-access-token generate \
  --app-id [app-id] \
  --installation-id [installation-id] \
  --private-key [private-key-file-path]

# generate the Github App access token with file in HOME
gh-app-access-token generate \
  --app-id [app-id] \
  --installation-id [installation-id] \
  --private-key $HOME/private-key.pem

# generate the Github App access token with file in HOME
gh-app-access-token generate \
  --app-id [app-id] \
  --installation-id [installation-id] \
  --private-key ~/private-key.pem

# generate the Github App access token with text in private key file passed into stdin
cat [private-key-file-path] | gh-app-access-token generate \
  --app-id [app-id] \
  --installation-id [installation-id] \
  --private-key -

# generate the Github App access token with private key text passed into stdin
echo "private-key-text" | gh-app-access-token generate \
  --app-id [app-id] \
  --installation-id [installation-id] \
  --private-key -
```

### Generate a Github App access toekn using docker image

```bash
docker run --rm \
  -v <your-github-app-private-key-file-path.pem>:/private-key.pem \
  ghcr.io/sunggun-yu/gh-app-access-token:latest \
  generate --app-id [app-id] --installation-id [installation-id] -f /private-key.pem
```

Use environemnt variable: e.g. run it on circleci or github actions

>⚠️ Note
>
> base64 encoded private key should be set in the environment variable

```bash
docker run --rm --entrypoint="/bin/sh" \
  ghcr.io/sunggun-yu/gh-app-access-token:latest \
  -c \
  "echo -e $GH_APP_PRIVATE_KEY | base64 -d | \
  gh-app-access-token generate \
  --app-id $GH_APP_ID \
  --installation-id $GH_APP_INSTALLATION_ID \
  -f -"
```

>⚠️ Note/Warning
>
> it keeps waiting(hang) if there is no stdin when you pass `-` for arg/value

### Revoke the Github App access token

```bash
# revoke token in argument
gh-app-access-token-cli revoke [access token string]

# revoke the token passed into stdin
cat [access-token-file] | gh-app-access-token-cli revoke -

# revoke the token passed into stdin
echo "access-token-value" | gh-app-access-token-cli revoke -
```

>⚠️ Note/Warning
>
> it keeps waiting(hang) if there is no stdin when you pass `-` for arg/value

### Revoke the Github App access token using docker image

```bash
docker run --rm \
  ghcr.io/sunggun-yu/gh-app-access-token:latest \
  revoke [access token string]
```
