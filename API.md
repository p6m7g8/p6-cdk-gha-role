# API Reference <a name="API Reference" id="api-reference"></a>

## Constructs <a name="Constructs" id="Constructs"></a>

### P6CDKName <a name="P6CDKName" id="p6-cdk-name.P6CDKName"></a>

#### Initializers <a name="Initializers" id="p6-cdk-name.P6CDKName.Initializer"></a>

```typescript
import { P6CDKName } from 'p6-cdk-name'

new P6CDKName(scope: Construct, id: string, _props: IP6CDKNameProps)
```

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#p6-cdk-name.P6CDKName.Initializer.parameter.scope">scope</a></code> | <code>constructs.Construct</code> | *No description.* |
| <code><a href="#p6-cdk-name.P6CDKName.Initializer.parameter.id">id</a></code> | <code>string</code> | *No description.* |
| <code><a href="#p6-cdk-name.P6CDKName.Initializer.parameter._props">_props</a></code> | <code><a href="#p6-cdk-name.IP6CDKNameProps">IP6CDKNameProps</a></code> | *No description.* |

---

##### `scope`<sup>Required</sup> <a name="scope" id="p6-cdk-name.P6CDKName.Initializer.parameter.scope"></a>

- *Type:* constructs.Construct

---

##### `id`<sup>Required</sup> <a name="id" id="p6-cdk-name.P6CDKName.Initializer.parameter.id"></a>

- *Type:* string

---

##### `_props`<sup>Required</sup> <a name="_props" id="p6-cdk-name.P6CDKName.Initializer.parameter._props"></a>

- *Type:* <a href="#p6-cdk-name.IP6CDKNameProps">IP6CDKNameProps</a>

---

#### Methods <a name="Methods" id="Methods"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#p6-cdk-name.P6CDKName.toString">toString</a></code> | Returns a string representation of this construct. |
| <code><a href="#p6-cdk-name.P6CDKName.applyRemovalPolicy">applyRemovalPolicy</a></code> | Apply the given removal policy to this resource. |

---

##### `toString` <a name="toString" id="p6-cdk-name.P6CDKName.toString"></a>

```typescript
public toString(): string
```

Returns a string representation of this construct.

##### `applyRemovalPolicy` <a name="applyRemovalPolicy" id="p6-cdk-name.P6CDKName.applyRemovalPolicy"></a>

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

###### `policy`<sup>Required</sup> <a name="policy" id="p6-cdk-name.P6CDKName.applyRemovalPolicy.parameter.policy"></a>

- *Type:* aws-cdk-lib.RemovalPolicy

---

#### Static Functions <a name="Static Functions" id="Static Functions"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#p6-cdk-name.P6CDKName.isConstruct">isConstruct</a></code> | Checks if `x` is a construct. |
| <code><a href="#p6-cdk-name.P6CDKName.isOwnedResource">isOwnedResource</a></code> | Returns true if the construct was created by CDK, and false otherwise. |
| <code><a href="#p6-cdk-name.P6CDKName.isResource">isResource</a></code> | Check whether the given construct is a Resource. |

---

##### `isConstruct` <a name="isConstruct" id="p6-cdk-name.P6CDKName.isConstruct"></a>

```typescript
import { P6CDKName } from 'p6-cdk-name'

P6CDKName.isConstruct(x: any)
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

###### `x`<sup>Required</sup> <a name="x" id="p6-cdk-name.P6CDKName.isConstruct.parameter.x"></a>

- *Type:* any

Any object.

---

##### `isOwnedResource` <a name="isOwnedResource" id="p6-cdk-name.P6CDKName.isOwnedResource"></a>

```typescript
import { P6CDKName } from 'p6-cdk-name'

P6CDKName.isOwnedResource(construct: IConstruct)
```

Returns true if the construct was created by CDK, and false otherwise.

###### `construct`<sup>Required</sup> <a name="construct" id="p6-cdk-name.P6CDKName.isOwnedResource.parameter.construct"></a>

- *Type:* constructs.IConstruct

---

##### `isResource` <a name="isResource" id="p6-cdk-name.P6CDKName.isResource"></a>

```typescript
import { P6CDKName } from 'p6-cdk-name'

P6CDKName.isResource(construct: IConstruct)
```

Check whether the given construct is a Resource.

###### `construct`<sup>Required</sup> <a name="construct" id="p6-cdk-name.P6CDKName.isResource.parameter.construct"></a>

- *Type:* constructs.IConstruct

---

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#p6-cdk-name.P6CDKName.property.node">node</a></code> | <code>constructs.Node</code> | The tree node. |
| <code><a href="#p6-cdk-name.P6CDKName.property.env">env</a></code> | <code>aws-cdk-lib.ResourceEnvironment</code> | The environment this resource belongs to. |
| <code><a href="#p6-cdk-name.P6CDKName.property.stack">stack</a></code> | <code>aws-cdk-lib.Stack</code> | The stack in which this resource is defined. |

---

##### `node`<sup>Required</sup> <a name="node" id="p6-cdk-name.P6CDKName.property.node"></a>

```typescript
public readonly node: Node;
```

- *Type:* constructs.Node

The tree node.

---

##### `env`<sup>Required</sup> <a name="env" id="p6-cdk-name.P6CDKName.property.env"></a>

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

##### `stack`<sup>Required</sup> <a name="stack" id="p6-cdk-name.P6CDKName.property.stack"></a>

```typescript
public readonly stack: Stack;
```

- *Type:* aws-cdk-lib.Stack

The stack in which this resource is defined.

---




## Protocols <a name="Protocols" id="Protocols"></a>

### IP6CDKNameProps <a name="IP6CDKNameProps" id="p6-cdk-name.IP6CDKNameProps"></a>

- *Implemented By:* <a href="#p6-cdk-name.IP6CDKNameProps">IP6CDKNameProps</a>



