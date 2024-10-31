import type { Construct } from 'constructs'
import * as cdk from 'aws-cdk-lib'
import * as iam from 'aws-cdk-lib/aws-iam'
import { P6CDKGHARole } from '../src'

class VisualizeStack extends cdk.Stack {
  constructor(scope: Construct, id: string) {
    super(scope, id)

    new P6CDKGHARole(this, 'MyP6Stack', {
      principle: 'arn:aws:iam::346733622331:oidc-provider/token.actions.githubusercontent.com',
      repo: 'p6m7g8/p6m7g8.com',
      policies: [
        iam.ManagedPolicy.fromAwsManagedPolicyName('AmazonS3FullAccess'),
        iam.ManagedPolicy.fromAwsManagedPolicyName('CloudFrontFullAccess'),
      ],
    })
  }
}

const app = new cdk.App()
new VisualizeStack(app, 'VisualizeStack')
app.synth()
