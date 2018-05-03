# JSON Overrides

Given a JSON file of base elements, this will apply a set of overrides to that
JSON file from data defined in another file.

## Usage

Basic usage:
```
go run main.go base.json overrides.json
```

Specifying the output file:
```
go run main.go base.json overrides.json overriden.json
```

## Example

Let's say you have two configs, one for your staging environment and one for production.
Some of the elements between environments may be the same, such as the name of your company.
However, other elements like the environment itself will differ. While the configs
below are small, imagine one with many more elements. Over 50% of these elements
might be the same, so it can be a pain having to keep them in sync if an element
changes that is the same in both of them. Instead you could keep a file that is just
the elements that differ from the base config then apply those to the base config
to get your new, final config.

`base.json`
```
{
  "environment": "production",
  "company": "Acme Corp"
}
```

`overrides.json`
```
{
  "environment": "staging"
}
```

`output.json`
```
{
  "company": "Acme Corp",
  "environment": "staging"
}
```

## Testing

```
go test
```
