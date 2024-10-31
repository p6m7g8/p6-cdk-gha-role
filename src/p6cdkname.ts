import type { Construct } from 'constructs'
import * as cdk from 'aws-cdk-lib'

export interface IP6CDKNameProps {
}

export class P6CDKName extends cdk.Resource {
  constructor(scope: Construct, id: string, _props: IP6CDKNameProps) {
    super(scope, id)
  }
}
