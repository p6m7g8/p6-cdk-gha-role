import * as cdk from 'aws-cdk-lib'
import { Template } from 'aws-cdk-lib/assertions'
import { P6CDKName } from '../src'

it('p6CDKWebsitePlus snapshot test', () => {
  const app = new cdk.App()
  const stack = new cdk.Stack(app, 'MyStack')

  // WHEN
  new P6CDKName(stack, 'p6-cdk-name', {
  })

  // THEN
  const template = Template.fromStack(stack).toJSON()
  expect(template).toMatchSnapshot()
})
