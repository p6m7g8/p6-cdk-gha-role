import type { Construct } from 'constructs'
import * as cdk from 'aws-cdk-lib'
import * as iam from 'aws-cdk-lib/aws-iam'

export interface IP6CDKGHARoleProps {
  readonly principle: cdk.Arn
  readonly repo: string
  readonly policies?: iam.IManagedPolicy[]
}

export class P6CDKGHARole extends cdk.Resource {
  constructor(scope: Construct, id: string, props: IP6CDKGHARoleProps) {
    super(scope, id)

    const role = new iam.Role(this, 'gha-p6m7g8-p6-roles', {
      assumedBy: new iam.FederatedPrincipal(
        props.principle as unknown as string,
        {
          StringEquals: {
            'token.actions.githubusercontent.com:aud': 'sts.amazonaws.com',
          },
          StringLike: {
            [`token.actions.githubusercontent.com:sub`]: `repo:${props.repo}:*`,
          },
        },
        'sts:AssumeRoleWithWebIdentity',
      ),
    })
    if (props.policies) {
      props.policies.forEach(policy => role.addManagedPolicy(policy))
    }
  }
}
