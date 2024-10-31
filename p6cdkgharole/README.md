DESC

# P6CDKGHARole

## LICENSE

[![License](https://img.shields.io/badge/License-Apache%202.0-yellowgreen.svg)](https://opensource.org/licenses/Apache-2.0)

## Other

![Gitpod ready-to-code](https://img.shields.io/badge/Gitpod-ready--to--code-blue?logo=gitpod) ![Sonarcloud Status](https://sonarcloud.io/api/project_badges/measure?project=p6m7g8_p6-cdk-gha-role&metric=alert_status) ![GitHub commit activity](https://img.shields.io/github/commit-activity/y/p6m7g8/p6-cdk-gha-role) ![GitHub commit activity](https://img.shields.io/github/commit-activity/m/p6m7g8/p6-cdk-gha-role)

## Usage

```go
...
import { P6CDKGHARole } from 'p6-cdk-gha-role';

new P6CDKGHARole(this, 'P6CDKGHARole', {
  principle: arn,
  repo: 'org/repo',
  policies: []
});
```

## Architecture

![./assets/diagram.png](./assets/diagram.png)

## Author

Philip M. Gollucci [pgollucci@p6m7g8.com](mailto:pgollucci@p6m7g8.com)
