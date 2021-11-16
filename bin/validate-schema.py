#!/usr/bin/env python
import json
import sys
import jsonschema_rs


def main():
    with open(sys.argv[1], 'r') as schema_file:
        schema = json.loads(schema_file.read())
        obj = json.loads(sys.stdin.read())
        validator = jsonschema_rs.JSONSchema(schema)


if __name__ == "__main__":
    main()
