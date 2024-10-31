import antfu from '@antfu/eslint-config'

export default antfu({
  ignores: [
    '*.md',
    '.vscode/',
    '.github/',
    '.mergify.yml',
    '.yarn/',
    'assets/',
    'cdk.json',
    'tsconfig.json',
    'tsconfig.dev.json',
    'package.json',
    'yarn.lock',
  ],
  plugins: {
  },
  languageOptions: {
  },
  rules: {
    'no-new': 'off',
    'no-console': 'off',
  },
  settings: {
  },
})
