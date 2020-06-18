# robot
![](https://github.com/street-bot/robot/workflows/release/badge.svg)

The application that runs on the robot. Automatically built.


### Pulling latest release file from GitHub
```bash
GITHUB_TOKEN=<a personal access token>
OWNER="street-bot"
REPO="robot"
RELEASE_ID=$(curl -H "Authorization: token $GITHUB_TOKEN" -sL https://api.github.com/repos/street-bot/robot/releases/latest | jq -r ".assets[] | select(.name | contains(\"robot_linux_amd64.tar.gz\")) | .id")
curl -H 'Accept: application/octet-stream' -H "Authorization: token $GITHUB_TOKEN" -LJO "https://api.github.com/repos/$OWNER/$REPO/releases/assets/$RELEASE_ID"
```