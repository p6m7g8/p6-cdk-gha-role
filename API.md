# API Reference <a name="API Reference" id="api-reference"></a>

## Constructs <a name="Constructs" id="Constructs"></a>

### P6CDKGHARole <a name="P6CDKGHARole" id="p6-cdk-gha-role.P6CDKGHARole"></a>

#### Initializers <a name="Initializers" id="p6-cdk-gha-role.P6CDKGHARole.Initializer"></a>

```typescript
import { P6CDKGHARole } from 'p6-cdk-gha-role'

new P6CDKGHARole(scope: Construct, id: string, props: IP6CDKGHARoleProps)
```

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#p6-cdk-gha-role.P6CDKGHARole.Initializer.parameter.scope">scope</a></code> | <code>constructs.Construct</code> | *No description.* |
| <code><a href="#p6-cdk-gha-role.P6CDKGHARole.Initializer.parameter.id">id</a></code> | <code>string</code> | *No description.* |
| <code><a href="#p6-cdk-gha-role.P6CDKGHARole.Initializer.parameter.props">props</a></code> | <code><a href="#p6-cdk-gha-role.IP6CDKGHARoleProps">IP6CDKGHARoleProps</a></code> | *No description.* |

---

##### `scope`<sup>Required</sup> <a name="scope" id="p6-cdk-gha-role.P6CDKGHARole.Initializer.parameter.scope"></a>

- *Type:* constructs.Construct

---

##### `id`<sup>Required</sup> <a name="id" id="p6-cdk-gha-role.P6CDKGHARole.Initializer.parameter.id"></a>

- *Type:* string

---

##### `props`<sup>Required</sup> <a name="props" id="p6-cdk-gha-role.P6CDKGHARole.Initializer.parameter.props"></a>

- *Type:* <a href="#p6-cdk-gha-role.IP6CDKGHARoleProps">IP6CDKGHARoleProps</a>

---

#### Methods <a name="Methods" id="Methods"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#p6-cdk-gha-role.P6CDKGHARole.toString">toString</a></code> | Returns a string representation of this construct. |
| <code><a href="#p6-cdk-gha-role.P6CDKGHARole.applyRemovalPolicy">applyRemovalPolicy</a></code> | Apply the given removal policy to this resource. |

---

##### `toString` <a name="toString" id="p6-cdk-gha-role.P6CDKGHARole.toString"></a>

```typescript
public toString(): string
```

Returns a string representation of this construct.

##### `applyRemovalPolicy` <a name="applyRemovalPolicy" id="p6-cdk-gha-role.P6CDKGHARole.applyRemovalPolicy"></a>

```typescript
public applyRemovalPolicy(policy: RemovalPolicy): void
```

Apply the given removal policy to this resource.

The Removal Policy controls what happens to this resource when it stops
being managed by CloudFormation, either because you've removed it from the
CDK application or because you've made a change that requires the resource
to be replaced.

The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).

###### `policy`<sup>Required</sup> <a name="policy" id="p6-cdk-gha-role.P6CDKGHARole.applyRemovalPolicy.parameter.policy"></a>

- *Type:* aws-cdk-lib.RemovalPolicy

---

#### Static Functions <a name="Static Functions" id="Static Functions"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#p6-cdk-gha-role.P6CDKGHARole.isConstruct">isConstruct</a></code> | Checks if `x` is a construct. |
| <code><a href="#p6-cdk-gha-role.P6CDKGHARole.isOwnedResource">isOwnedResource</a></code> | Returns true if the construct was created by CDK, and false otherwise. |
| <code><a href="#p6-cdk-gha-role.P6CDKGHARole.isResource">isResource</a></code> | Check whether the given construct is a Resource. |

---

##### `isConstruct` <a name="isConstruct" id="p6-cdk-gha-role.P6CDKGHARole.isConstruct"></a>

```typescript
import { P6CDKGHARole } from 'p6-cdk-gha-role'

P6CDKGHARole.isConstruct(x: any)
```

Checks if `x` is a construct.

Use this method instead of `instanceof` to properly detect `Construct`
instances, even when the construct library is symlinked.

Explanation: in JavaScript, multiple copies of the `constructs` library on
disk are seen as independent, completely different libraries. As a
consequence, the class `Construct` in each copy of the `constructs` library
is seen as a different class, and an instance of one class will not test as
`instanceof` the other class. `npm install` will not create installations
like this, but users may manually symlink construct libraries together or
use a monorepo tool: in those cases, multiple copies of the `constructs`
library can be accidentally installed, and `instanceof` will behave
unpredictably. It is safest to avoid using `instanceof`, and using
this type-testing method instead.

###### `x`<sup>Required</sup> <a name="x" id="p6-cdk-gha-role.P6CDKGHARole.isConstruct.parameter.x"></a>

- *Type:* any

Any object.

---

##### `isOwnedResource` <a name="isOwnedResource" id="p6-cdk-gha-role.P6CDKGHARole.isOwnedResource"></a>

```typescript
import { P6CDKGHARole } from 'p6-cdk-gha-role'

P6CDKGHARole.isOwnedResource(construct: IConstruct)
```

Returns true if the construct was created by CDK, and false otherwise.

###### `construct`<sup>Required</sup> <a name="construct" id="p6-cdk-gha-role.P6CDKGHARole.isOwnedResource.parameter.construct"></a>

- *Type:* constructs.IConstruct

---

##### `isResource` <a name="isResource" id="p6-cdk-gha-role.P6CDKGHARole.isResource"></a>

```typescript
import { P6CDKGHARole } from 'p6-cdk-gha-role'

P6CDKGHARole.isResource(construct: IConstruct)
```

Check whether the given construct is a Resource.

###### `construct`<sup>Required</sup> <a name="construct" id="p6-cdk-gha-role.P6CDKGHARole.isResource.parameter.construct"></a>

- *Type:* constructs.IConstruct

---

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#p6-cdk-gha-role.P6CDKGHARole.property.node">node</a></code> | <code>constructs.Node</code> | The tree node. |
| <code><a href="#p6-cdk-gha-role.P6CDKGHARole.property.env">env</a></code> | <code>aws-cdk-lib.ResourceEnvironment</code> | The environment this resource belongs to. |
| <code><a href="#p6-cdk-gha-role.P6CDKGHARole.property.stack">stack</a></code> | <code>aws-cdk-lib.Stack</code> | The stack in which this resource is defined. |

---

##### `node`<sup>Required</sup> <a name="node" id="p6-cdk-gha-role.P6CDKGHARole.property.node"></a>

```typescript
public readonly node: Node;
```

- *Type:* constructs.Node

The tree node.

---

##### `env`<sup>Required</sup> <a name="env" id="p6-cdk-gha-role.P6CDKGHARole.property.env"></a>

```typescript
public readonly env: ResourceEnvironment;
```

- *Type:* aws-cdk-lib.ResourceEnvironment

The environment this resource belongs to.

For resources that are created and managed by the CDK
(generally, those created by creating new class instances like Role, Bucket, etc.),
this is always the same as the environment of the stack they belong to;
however, for imported resources
(those obtained from static methods like fromRoleArn, fromBucketName, etc.),
that might be different than the stack they were imported into.

---

##### `stack`<sup>Required</sup> <a name="stack" id="p6-cdk-gha-role.P6CDKGHARole.property.stack"></a>

```typescript
public readonly stack: Stack;
```

- *Type:* aws-cdk-lib.Stack

The stack in which this resource is defined.

---




## Protocols <a name="Protocols" id="Protocols"></a>

### IP6CDKGHARoleProps <a name="IP6CDKGHARoleProps" id="p6-cdk-gha-role.IP6CDKGHARoleProps"></a>

- *Implemented By:* <a href="#p6-cdk-gha-role.IP6CDKGHARoleProps">IP6CDKGHARoleProps</a>


#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#p6-cdk-gha-role.IP6CDKGHARoleProps.property.principle">principle</a></code> | <code>aws-cdk-lib.Arn</code> | *No description.* |
| <code><a href="#p6-cdk-gha-role.IP6CDKGHARoleProps.property.repo">repo</a></code> | <code>string</code> | *No description.* |
| <code><a href="#p6-cdk-gha-role.IP6CDKGHARoleProps.property.policies">policies</a></code> | <code>aws-cdk-lib.aws_iam.IManagedPolicy[]</code> | *No description.* |

---

##### `principle`<sup>Required</sup> <a name="principle" id="p6-cdk-gha-role.IP6CDKGHARoleProps.property.principle"></a>

```typescript
public readonly principle: Arn;
```

- *Type:* aws-cdk-lib.Arn

---

##### `repo`<sup>Required</sup> <a name="repo" id="p6-cdk-gha-role.IP6CDKGHARoleProps.property.repo"></a>

```typescript
public readonly repo: string;
```

- *Type:* string

---

##### `policies`<sup>Optional</sup> <a name="policies" id="p6-cdk-gha-role.IP6CDKGHARoleProps.property.policies"></a>

```typescript
public readonly policies: IManagedPolicy[];
```

- *Type:* aws-cdk-lib.aws_iam.IManagedPolicy[]

---

