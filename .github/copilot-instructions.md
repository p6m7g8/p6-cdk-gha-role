# Copilot Onboarding Instructions

## 1. Summary
This repo defines an AWS CDK v2 Construct (`P6CDKNamer`) that sets an IAM account alias via a Lambda-backed Custom Resource. It uses TypeScript + jsii to publish multi-language packages (JS, Python, Java, .NET, Go), plus Go code under `p6cdkgharole/` for the Go module.

## 2. Tech Stack & Tooling
- Languages: TypeScript (library, Lambda handler), Go (jsii Go bindings)
- Frameworks: AWS CDK v2, jsii, Jest, ESLint
- Runtimes: Node.js ≥24.9.0, Go ≥1.18
- Package Manager: Yarn 4.10.3
- Key files:
  - tsconfig.json / tsconfig.dev.json (compiler settings)
  - jest.config.js (tests)
  - eslint.config.mjs
  - cdk.json (CDK synth entry)
  - package.json (scripts)

## 3. Environment Setup & Bootstrap
Always start in a clean repo clone:
1. Install dependencies:
   ```
   yarn install
   ```
2. Ensure `node` version ≥24.9.0 (`.node-version`) and Go 1.18+ are active.
3. No extra environment variables are needed for local build/test.

## 4. Build & Validation Workflow
Use the npm scripts; all have been verified on a clean checkout.
- Compile & JSII packaging:
  ```
  yarn build
  ```
  (runs compile, lint, jsii:docgen, tests, diagram)
- TypeScript-only compile:
  ```
  yarn compile
  ```
- Tests (Jest + AWS CDK assertions):
  ```
  yarn test
  ```
- Lint:
  ```
  yarn lint
  ```
- CDK synth (generates CloudFormation diagram):
  ```
  yarn synth
  ```
  
Always run `yarn install` if lockfile or dependencies change. If builds fail due to stale artifacts, run:
```
yarn clean
yarn install
yarn build
```

## 5. Project Layout
```
/
├─ package.json          # scripts & dependencies
├─ tsconfig*.json        # TS compiler settings
├─ jest.config.js        # test config
├─ eslint.config.mjs     # lint rules
├─ cdk.json              # CDK entrypoint
├─ bin/visualize.ts      # example synth app
├─ src/                  # TS source
│   ├─ index.ts
│   └─ p6cdkgharole.ts
├─ test/                 # unit tests
│   └─ p6-cdk-namer.test.ts
├─ p6cdkgharole/           # generated Go module (jsii)
│   ├─ go.mod, *.go
│   └─ version, README.md
├─ .github/              # workflows
│   ├─ build.yml
│   ├─ release.yml
│   └─ upgrade-main.yml
└─ README.md             # user‐facing docs
```

## 6. CI Pipelines & Checks
- Build: `.github/workflows/build.yml` (on PR) → runs `cdk-construct-build` (compile + tests)
- Release: `release.yml` (on main) → `cdk-construct-release`
- PR lint: `pull-request-lint.yml` → semantic PR titles
- Auto-merge & auto-approve flows present but not needed locally

## 7. Usage Tips
- Trust these instructions; avoid re-searching unless a command errors.
- When modifying TS code, ensure all tests pass and `yarn lint` is clean.
- After changes to JSII interfaces, bump `p6cdkgharole/version` and regenerate.
- For diagram updates, `yarn diagram:update` syncs assets.

---
_End of onboarding. Only search the repo when instructions lack required detail or if you hit an unexpected error._
