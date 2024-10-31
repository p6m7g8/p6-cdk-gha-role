import * as cdk from 'aws-cdk-lib'
import { Template } from 'aws-cdk-lib/assertions'
import * as iam from 'aws-cdk-lib/aws-iam'
import { P6CDKGHARole } from '../src'

it('p6CDKWebsitePlus snapshot test', () => {
  const app = new cdk.App()
  const stack = new cdk.Stack(app, 'MyStack')

  // WHEN
  new P6CDKGHARole(stack, 'p6-cdk-name', {
    principle: 'arn:aws:iam::346733622331:oidc-provider/token.actions.githubusercontent.com',
    repo: 'p6m7g8/p6m7g8.com',
    policies: [
      iam.ManagedPolicy.fromAwsManagedPolicyName('AmazonS3FullAccess'),
      iam.ManagedPolicy.fromAwsManagedPolicyName('CloudFrontFullAccess'),
    ],
  })

  // THEN
  const template = Template.fromStack(stack).toJSON()
  expect(template).toMatchSnapshot()
})
