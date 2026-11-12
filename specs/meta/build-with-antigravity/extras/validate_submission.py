#!/usr/bin/env python3
"""
validate_submission.py - Validator for Build with Antigravity App Manifests (Schema 1.0)
"""

import sys
import os
import re
import argparse

try:
    import yaml
except ImportError:
    yaml = None

def validate_dict(data, file_path):
    errors = []
    
    # 1. Check schema_version
    ver = str(data.get("schema_version", ""))
    if ver != "1.0":
        errors.append(f"Invalid or missing schema_version: expected '1.0', got '{ver}'")

    # 2. Check name and slug
    name = data.get("name", "")
    if not isinstance(name, str) or len(name.strip()) < 3 or len(name) > 64:
        errors.append(f"Field 'name' must be between 3 and 64 characters (got '{name}')")

    slug = data.get("slug", "")
    if not re.match(r"^[a-z0-9-]+$", slug) or len(slug) > 48:
        errors.append(f"Field 'slug' must be lowercase alphanumeric and hyphens (max 48 chars), got '{slug}'")

    # 3. Check headline
    headline = data.get("headline", "")
    if not isinstance(headline, str) or len(headline.strip()) < 10 or len(headline) > 140:
        errors.append(f"Field 'headline' must be between 10 and 140 characters")

    # 4. Check author
    author = data.get("author", {})
    if not isinstance(author, dict):
        errors.append("Field 'author' must be an object")
    else:
        if not author.get("name"):
            errors.append("Field 'author.name' is required")
        linkedin = author.get("linkedin", "")
        if not re.match(r"^https:\/\/(www\.)?linkedin\.com\/in\/[a-zA-Z0-9_-]+\/?$", linkedin):
            errors.append(f"Field 'author.linkedin' must be a valid LinkedIn profile URL, got '{linkedin}'")

    # 5. Check repository
    repo = data.get("repository", {})
    if not isinstance(repo, dict):
        errors.append("Field 'repository' must be an object")
    else:
        repo_url = repo.get("url", "")
        if not repo_url.startswith("http://") and not repo_url.startswith("https://") and not repo_url.startswith("git@"):
            errors.append(f"Field 'repository.url' must be a valid URL, got '{repo_url}'")
        license_str = repo.get("license", "")
        if not license_str:
            errors.append("Field 'repository.license' is required (e.g. 'MIT', 'Apache-2.0', 'Internal')")

    # 6. Check showcase
    showcase = data.get("showcase", {})
    if not isinstance(showcase, dict) or not showcase.get("hero_image"):
        errors.append("Field 'showcase.hero_image' is required")

    return errors

def main():
    parser = argparse.ArgumentParser(description="Validate Build with Antigravity app manifest against Schema 1.0")
    parser.add_argument("file", help="Path to YAML or JSON manifest file")
    args = parser.parse_args()

    if not os.path.exists(args.file):
        print(f"Error: File not found at {args.file}", file=sys.stderr)
        sys.exit(1)

    import json
    data = None
    with open(args.file, "r", encoding="utf-8") as f:
        content = f.read()
        if args.file.endswith(".json"):
            data = json.loads(content)
        else:
            if yaml is None:
                print("Error: PyYAML not installed for YAML parsing", file=sys.stderr)
                sys.exit(1)
            data = yaml.safe_load(content)

    errors = validate_dict(data, args.file)
    if errors:
        print(f"FAILED: {len(errors)} schema validation errors found in {args.file}:")
        for err in errors:
            print(f"  - {err}")
        sys.exit(1)
    else:
        print(f"SUCCESS: {args.file} conforms to Build with Antigravity Schema 1.0!")
        sys.exit(0)

if __name__ == "__main__":
    main()
